package make

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeSeeder = &cobra.Command{
	Use:   "seeder",
	Short: "Create seeder file,exapmle: make seeder user",
	Run:   runMakeSeed,
	Args:  cobra.ExactArgs(1),
}

func runMakeSeed(cmd *cobra.Command, args []string) {
	// 格式化模型名称,返回一个 Model 对象
	model := makeModelFromString(args[0])
	// 拼接目标文件路径
	filepath := fmt.Sprintf("database/seeders/%s_seeder.go", model.PackageName)
	// 基于模版创建文件(做好变量替换)
	createFileFromStub(filepath, "seeder", model)
}
