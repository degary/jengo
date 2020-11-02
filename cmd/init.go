package cmd

import (
	"fmt"
	"github.com/degary/jengo/internal/dockeroper"
	"github.com/degary/jengo/internal/initpro"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

var desc = strings.Join([]string{
	"该命令用于初始化项目时的其他选项,例如: ",
	"-v 挂载选项等",
}, "\n")
var (
	iniName string
	other   string
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化项目",
	Long:  "该子命令用于初始化项目",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("initproject")
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
		viper.AddConfigPath(".")
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Printf("read config failed:%v", err)
		}
		projectPath := viper.GetString("projectpath")
		backupPath := viper.GetString("backuppath")
		proPath := filepath.Join(projectPath, iniName)
		backPath := filepath.Join(backupPath, iniName)
		//创建项目目录
		if err := initpro.Makedir(proPath); err != nil {
			fmt.Errorf("创建项目目录报错了: %s\n", err)
		}
		//创建备份目录
		if err := initpro.Makedir(backPath); err != nil {
			fmt.Errorf("创建项目目录报错了: %s\n", err)
		}
		initpro.StartDocker()
		containerList, err := dockeroper.ListContainer()
		fmt.Println(containerList, err)
	},
}

func init() {
	initCmd.Flags().StringVarP(&iniName, "name", "n", "", "项目名称")
	initCmd.Flags().StringVarP(&other, "other", "o", "", desc)
}
