package viperx

import (
	"fmt"

	"os"
	"sync"

	"github.com/betacat-co/pkgx/env"

	"github.com/spf13/viper"
)

type Viper struct {
	C *viper.Viper
}

var multipleViper sync.Map
var C = viper.New()

func NewConfig(filePath string) *viper.Viper {

	C.SetConfigName("default")              // name of config file (without extension)
	C.SetConfigName(os.Getenv(env.RunTime)) // name of config file (without extension)
	C.AddConfigPath(filePath)               // path to look for the config file in
	C.WatchConfig()                         // 动态加载配置文件
	err := viper.ReadInConfig()             // Find and read the config file
	if err != nil {                         // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	return C
}
