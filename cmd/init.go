package cmd

import (
	"fmt"
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
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化项目",
	Long:  "该子命令用于初始化项目",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init")
	},
}

func init() {
	initCmd.Flags().StringVarP(&iniName, "name", "n", "", "项目名称")
	initCmd.Flags().StringVarP(&other, "other", "o", "", desc)
}
