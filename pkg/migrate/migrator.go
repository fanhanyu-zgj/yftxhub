// Package migrate 处理数据库迁移
package migrate

import (
	"os"
	"yftxhub/pkg/console"
	"yftxhub/pkg/database"
	"yftxhub/pkg/file"

	"gorm.io/gorm"
)

// Migrator 数据迁移操作类
type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

// Migration 对应数据的 migrations 表里的一条数据
type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique"`
	Batch     int
}

// NewMigrator 创建 Migrator 实例，用以之前迁移操作
func NewMigrator() *Migrator {
	// 初始化必要属性
	migrator := &Migrator{
		Folder:   "database/migrations/",
		DB:       database.DB,
		Migrator: database.DB.Migrator(),
	}

	// migrations 不存在的话就创建它
	migrator.CreateMigrationsTable()

	return migrator
}

func (migrator *Migrator) CreateMigrationsTable() {
	migration := Migration{}
	if !migrator.Migrator.HasTable(&migration) {
		migrator.Migrator.CreateTable(&migration)
	}
}

// 执行所有未迁移过的文件
func (migrator *Migrator) Up() {

	// 读取多有迁移文件,确保按照时间排序
	migrationFiles := migrator.readAllMigrationFiles()

	// 当前的批次
	batch := migrator.getBatch()

	// 获取所有迁移数据
	migrations := []Migration{}
	migrator.DB.Find(&migrations)

	// 可以通过此值来判断数据库是否已是最新
	runed := false

	// 对迁移文件进行遍历，如果没有执行过，就执行 up 回调
	for _, mfile := range migrationFiles {
		// 对比文件名称，看是否已经运行过
		if mfile.isNotMigrated(migrations) {
			migrator.runUpMigration(mfile, batch)
			runed = true
		}
	}
	if !runed {
		console.Success("database is up to date.")
	}
}

// Rollback 回滚上一个操作
func (migrator *Migrator) Rollback() {
	// 获取最后一批次次的迁移数据
	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)
	migrations := []Migration{}
	migrator.DB.Where("batch = ?", lastMigration.Batch).Order("id DESC").Find(&migrations)

	// 回滚最后一批次的迁移
	if !migrator.rollbackMigrations(migrations) {
		console.Success("[migrations] table is empty, nothing to rollback.")
	}

}

func (migrator *Migrator) rollbackMigrations(migrations []Migration) bool {
	// 标记是否真的有执行了迁移回退的操作
	runed := false
	for _, _migration := range migrations {
		//友好提示
		console.Warning("rollback " + _migration.Migration)

		//执行迁移文件的 down 方法
		mfile := getMigrationFile(_migration.Migration)
		if mfile.Down != nil {
			mfile.Down(database.DB.Migrator(), database.SQLDB)
		}
		runed = true
		// 回退成功了就删除掉这条记录
		migrator.DB.Delete(&_migration)
		// 打印运行状态
		console.Success("finish " + mfile.FileName)
	}
	return runed
}

// Rollback 回滚所有迁移文件
func (migrator *Migrator) Reset() {

	migrations := []Migration{}
	// 按照倒序读取所有文件
	migrator.DB.Order("id DESC").Find(&migrations)

	// 回滚所有迁移
	if !migrator.rollbackMigrations(migrations) {
		console.Success("[migrations] table is empty, nothing to rollback.")
	}

}

// Refresh 回滚所有迁移,并运行所有迁移
func (migrator *Migrator) Refresh() {

	// 回滚所有迁移
	migrator.Rollback()
	// 再次执行所有迁移
	migrator.Up()

}

// Fresh Drop 所有的表并重新运行所有迁移
func (migrator *Migrator) Fresh() {
	// 获取数据库名称，用以提示
	dbname := database.CurrentDatabase()

	//删除所有表
	err := database.DeleteAllTables()
	if err != nil {
		console.ExitIf(err)
		console.Success("clearup database " + dbname)
	}

	// 重新创建 migrations 表
	migrator.CreateMigrationsTable()
	console.Success("[migrations] table created.")

	// 重新调用 up 命令
	migrator.Up()

}

// 获取当前这个批次的值
func (migrator *Migrator) getBatch() int {
	// 默认为 1
	batch := 1

	// 取最后执行的一条迁移数据
	lastMigration := Migration{}
	migrator.DB.Order("id Desc").First(&lastMigration)
	// 如果有值的话 ，加一
	if lastMigration.ID > 0 {
		batch = lastMigration.Batch + 1
	}
	return batch
}

// 从文件目录读取文件，保证正确的时间排序
func (migrator *Migrator) readAllMigrationFiles() []MigrationFile {
	// 读取 database/migrations/ 目录下的所有文件
	// 默认是会按照文件名称进行排序
	files, err := os.ReadDir(migrator.Folder)
	console.ExitIf(err)

	var migrationFiles []MigrationFile
	for _, f := range files {
		// 去除文件后缀 .go
		fileName := file.FileNameWithOutExtension(f.Name())
		// 通过迁移文件的名称获取 [MigrationFile] 对象
		mfile := getMigrationFile(fileName)
		// 加个判断，确保迁移文件可用再放进 migrationFiles 数组中
		if len(mfile.FileName) > 0 {
			migrationFiles = append(migrationFiles, mfile)
		}
	}

	// 返回排序好的[MigrationFile]数组
	return migrationFiles
}

// 执行迁移，执行迁移的 up 方法
func (migrator *Migrator) runUpMigration(mfile MigrationFile, batch int) {
	// 执行 up 区块的 SQL
	if mfile.Up != nil {
		// 友好提示
		console.Warning("migrating " + mfile.FileName)
		// 执行 up 方法
		mfile.Up(database.DB.Migrator(), database.SQLDB)
		// 提示已迁移了哪个文件
		console.Success("migrated " + mfile.FileName)
	}
	// 入库
	err := migrator.DB.Create(&Migration{Migration: mfile.FileName, Batch: batch}).Error
	console.ExitIf(err)
}
