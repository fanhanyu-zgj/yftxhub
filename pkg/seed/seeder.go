// Package seed 处理数据库填充相关逻辑
package seed

import (
	"yftxhub/pkg/console"
	"yftxhub/pkg/database"

	"gorm.io/gorm"
)

// 存放所有 Seeder
var seeders []Seeder

// 按顺序执行的 Seeder 数组
// 支持一些必须按顺序执行的 seeder，例如 topic 创建的
// 时候必须依赖于 User，所以 TopicSeeder 应该在 UserSeeder 后执行
var orderedSeederNames []string

type SeederFunc func(*gorm.DB)

// Seeder 对应每一个 database/seeders 目录下的 Seeder 文件
type Seeder struct {
	Func SeederFunc
	Name string
}

// Add 注册到 Seeders 数组中
func Add(name string, fn SeederFunc) {
	seeders = append(seeders, Seeder{
		Name: name,
		Func: fn,
	})
}

// SetRunOrder 设置[按照顺序执行的 Sseder 数组]
func SetRunOrder(names []string) {
	orderedSeederNames = names
}

// GetSeeder 通过名称来获取 Seeder 对象
func GetSeeder(name string) Seeder {
	for _, sdr := range seeders {
		if name == sdr.Name {
			return sdr
		}
	}
	return Seeder{}
}

// RunAll 运行所有 Seeder
func RunAll() {
	// 先运行 ordered 的
	execeuted := make(map[string]string)
	for _, name := range orderedSeederNames {
		sdr := GetSeeder(name)
		if len(sdr.Name) > 0 {
			console.Warning("Runing Orderded Seeder：" + sdr.Name)
			sdr.Func(database.DB)
			execeuted[name] = name
		}
	}
	// 再运行剩下的
	for _, sdr := range seeders {
		// 过滤已执行
		if _, ok := execeuted[sdr.Name]; !ok {
			console.Warning("Runing Seeder: " + sdr.Name)
			sdr.Func(database.DB)
		}
	}
}

// RunSeeder 运行单个 Seeder
func RunSeeder(name string) {
	for _, sdr := range seeders {
		if sdr.Name == name {
			sdr.Func(database.DB)
			break
		}
	}
}
