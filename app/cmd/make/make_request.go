package make

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeRequest = &cobra.Command{
	Use:   "request",
	Short: "Create request file, example: make request user",
	Run:   runMakeRequest,
	Args:  cobra.ExactArgs(1), // 只允许且必须传一个参数
}

func runMakeRequest(cmd *cobra.Command, args []string) {
	// 格式化模型名称,返回一个 Model 对象
	model := makeModelFromString(args[0])

	// 组建目标目录
	filePath := fmt.Sprintf("app/requests/%s_request.go", model.PackageName)
	// 从模版中创建文件（做好变量替换）
	createFileFromStub(filePath, "request", model)

}
