package make

import (
	"fmt"
	"strings"
	"yftxhub/pkg/console"

	"github.com/spf13/cobra"
)

var CmdMakeApiController = &cobra.Command{
	Use:   "apicontroller",
	Short: "Create api controller, example: make controller v1/user",
	Run:   runMakeAPIController,
	Args:  cobra.ExactArgs(1), // 只允许且必须传一个参数
}

func runMakeAPIController(cmd *cobra.Command, args []string) {
	// 参数处理 要求附带 API 版本（v1 或者 v2）
	array := strings.Split(args[0], "/")
	if len(array) != 2 {
		console.Exit("api controler name format: v1/user")
	}

	// apiVersion 用来拼接目标路径
	// name 用来生成 cmd.Model 实例
	// 格式化模型名称,返回一个 Model 对象
	apiVersion, name := array[0], array[1]
	model := makeModelFromString(name)

	// 组建目标目录
	filePath := fmt.Sprintf("app/http/controllers/api/%s/%s_controller.go", apiVersion, model.TableName)
	// 从模版中创建文件（做好变量替换）
	createFileFromStub(filePath, "apicontroller", model)

}
