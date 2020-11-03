package cmd

import (
	"fmt"
	"github.com/degary/jengo/internal/dockeroper"
	"github.com/degary/jengo/internal/initpro"
	"github.com/spf13/cobra"
	"strings"
)

var desc = strings.Join([]string{
	"该命令用于初始化项目时的其他选项,例如: ",
	"-v 挂载选项等",
}, "\n")
var (
	iniName string
	other   string
	jvm     int
	profile string
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化项目",
	Long:  "该子命令用于初始化项目",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("initproject")
		proPath, backPath := initpro.InitViper(iniName, "config", "toml", ".")
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
		//写start.sh
		bashPath := proPath + "/start.sh"
		content := fmt.Sprintf(initpro.StartFile, profile, jvm)
		err = initpro.CopyFile(bashPath, content)

	},
}

func init() {
	initCmd.Flags().StringVarP(&iniName, "name", "n", "", "项目名称")
	initCmd.Flags().StringVarP(&other, "other", "o", "", desc)
	initCmd.Flags().StringVarP(&profile, "profile", "p", "test", "java profile设置")
	initCmd.Flags().IntVarP(&jvm, "jvm", "j", 256, "xms和xmx参数")
}
