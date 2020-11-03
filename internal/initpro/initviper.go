package initpro

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
)

func InitViper(iniName, configname, configtype, configpath string) (string, string) {
	viper.SetConfigName(configname)
	viper.SetConfigType(configtype)
	viper.AddConfigPath(configpath)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("read config failed:%v", err)
	}
	projectPath := viper.GetString("projectpath")
	backupPath := viper.GetString("backuppath")
	proPath := filepath.Join(projectPath, iniName)
	backPath := filepath.Join(backupPath, iniName)
	return proPath, backPath
}
