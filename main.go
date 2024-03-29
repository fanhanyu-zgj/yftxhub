package main

import (
	"fmt"
	"os"
	"yftxhub/app/cmd"
	"yftxhub/app/cmd/make"
	"yftxhub/bootstrap"
	btsConfig "yftxhub/config"
	"yftxhub/pkg/config"
	"yftxhub/pkg/console"

	"github.com/spf13/cobra"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {
	// 配置初始化，依赖命令行 --env 参数
	// var env string
	// flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .envtesting 文件")
	// flag.Parse()
	// config.InitConfig(env)

	// // 初始化 Logger
	// bootstrap.SetupLogger()
	// // 设置 gin 的运行模式，支持 debug, release, test
	// // release 会屏蔽调试信息，官方建议生产环境中使用
	// // 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// // 故此设置为 release，有特殊情况手动改为 debug 即可

	// gin.SetMode(gin.ReleaseMode)

	// // 初始化 Gin 实例
	// r := gin.New()

	// // 初始化DB
	// bootstrap.SetupDB()

	// // 初始化 Redis
	// bootstrap.SetupRedis()

	// // 初始化路由绑定
	// bootstrap.SetupRoute(r)

	// // 运行服务
	// err := r.Run(":" + config.Get("app.port"))
	// if err != nil {
	// 	// 错误处理，端口被占用了或者其他错误
	// 	fmt.Println(err.Error())
	// }

	var rootCmd = &cobra.Command{
		Use:   "Yftxhub",
		Short: "A simple forum project",
		Long:  `Default will run "serve" command,you can see use "-h" flag to see all subcommands`,
		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {
			// 配置初始化，依赖命令行 --env 参数
			config.InitConfig(cmd.Env)

			// 初始化 Logger
			bootstrap.SetupLogger()

			// 初始化 数据库
			bootstrap.SetupDB()

			// 初始化 redis
			bootstrap.SetupRedis()

			// 初始化 缓存
		},
	}

	//  注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
		make.CmdMake,
		cmd.CmdMigrate,
	)
	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}

}
