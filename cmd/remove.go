package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	projectName string
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Long:  "删除对应的docker服务",
	Short: "删除服务",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("删除docker项目: %s\n", projectName)
	},
}

func init() {
	removeCmd.Flags().StringVarP(&projectName, "name", "n", "default", "需要删除的docker项目名称")
}
