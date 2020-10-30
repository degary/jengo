package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	rollbackProNum  int
	rollbackProName string
)

var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Long:  "回滚项目到对应的项目编号",
	Short: "回滚项目",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("回滚项目%s到%d\n", rollbackProName, rollbackProNum)
	},
}

func init() {
	rollbackCmd.Flags().StringVarP(&rollbackProName, "rbname", "n", "default", "需要回滚的项目名称")
	rollbackCmd.Flags().IntVarP(&rollbackProNum, "rbnum", "r", 0, "回滚到项目编号")
}
