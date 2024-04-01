package v1

import (
	"yftxhub/app/models/topic"
	"yftxhub/app/policies"
	"yftxhub/app/requests"
	"yftxhub/pkg/auth"
	"yftxhub/pkg/response"

	"github.com/gin-gonic/gin"
)

type TopicsController struct {
	BaseAPIController
}

func (ctrl *TopicsController) Store(c *gin.Context) {

	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	topicModel := topic.Topic{
		Title:      request.Title,
		Body:       request.Body,
		CategoryID: request.CategoryID,
		UserID:     auth.CurrentUID(c),
	}
	topicModel.Create()
	if topicModel.ID > 0 {
		response.Created(c, topicModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *TopicsController) Update(c *gin.Context) {
	// 验证 url 参数是否正确
	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		response.Abort404(c)
		return
	}
	if ok := policies.CanModifyTopic(c, topicModel); !ok {
		response.Abort403(c)
		return
	}
	// 表单验证
	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}
	// 保存数据
	topicModel.Title = request.Title
	topicModel.Body = request.Body
	topicModel.CategoryID = request.CategoryID
	rowsAaffected := topicModel.Save()
	if rowsAaffected > 0 {
		response.Data(c, topicModel)
	} else {
		response.Abort500(c, "更新失败，请重新尝试")
	}
}
