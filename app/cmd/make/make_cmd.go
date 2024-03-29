package make

import (
	"fmt"
	"yftxhub/pkg/console"

	"github.com/spf13/cobra"
)

var CmdMakeCMD = &cobra.Command{
	Use:   "cmd",
	Short: "Create a command ,should be snake_case ,example: make cmd backup_database",
	Run:   runMakeCMD,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个 参数
}

func runMakeCMD(cmd *cobra.Command, args []string) {
	// 格式化模型名称,返回一个 Model 对象
	model := makeModelFromString(args[0])

	// 拼接目标文件路径
	filePath := fmt.Sprintf("app/cmd/%s.go", model.PackageName)

	// 从模版中创建文件（做好变量替换）
	createFileFromStub(filePath, "cmd", model)

	// 友好提示
	console.Success("command name:" + model.PackageName)
	console.Success("command variable mame: cmd.Cmd" + model.StructName)
	console.Success("please edit main.go's app.Command slice to register command")
}
