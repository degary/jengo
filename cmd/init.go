package cmd

import (
	"bytes"
	"fmt"
	"github.com/degary/jengo/internal/dockeroper"
	"github.com/degary/jengo/internal/initpro"
	"github.com/spf13/cobra"
	"os/exec"
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
		for _, containerName := range containerList {
			if containerName == iniName {
				fmt.Println("docker container已存在")
				return
			}
		}
		fmt.Println(containerList, err)
		//写start.sh
		bashPath := proPath + "/start.sh"
		content := fmt.Sprintf(initpro.StartFile, profile, jvm)
		err = initpro.CopyFile(bashPath, content)
		if err != nil {
			fmt.Println(err)
		}
		command := fmt.Sprintf("docker run -itd --name=%s nginx", iniName)
		cmdOs := exec.Command("/bin/bash", "-c", command)
		buf := bytes.Buffer{}
		//将cmd2的标准输出设置为buf
		cmdOs.Stdout = &buf
		//运行命令，阻塞直到完成
		cmdOs.Run()
		fmt.Println(buf.String())
	},
}

func init() {
	initCmd.Flags().StringVarP(&iniName, "name", "n", "", "项目名称")
	initCmd.Flags().StringVarP(&other, "other", "o", "", desc)
	initCmd.Flags().StringVarP(&profile, "profile", "p", "test", "java profile设置")
	initCmd.Flags().IntVarP(&jvm, "jvm", "j", 256, "xms和xmx参数")
}
