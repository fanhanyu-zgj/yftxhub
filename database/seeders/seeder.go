//Package seeders  存放数据填充文件

func Initialize() {
	// 触发本目录下的其他文件中的 init 方法
	// 指定优先于同目录下的其他文件运行
	seed.SetRunOrder([]string{
		"SeedUsersTable",
	})
}

