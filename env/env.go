package env

// Type env type
type Type string

const (
	Local   = "local"   // 本地环境
	Dev     = "dev"     // 调试环境
	Qa      = "qa"      // 测试环境
	Test    = "test"    // 测试环境
	Prod    = "prod"    // 正式环境
	Onl     = "onl"     // 正式环境
	RunTime = "runtime" // RunTime 环境变量
)
