package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	depName string
	depBnum int
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "jenkins任务发布",
	Long:  "jenkins任务发布",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("deploy: name=%s,num=%d,", depName, depBnum)
	},
}

func init() {
	deployCmd.Flags().StringVarP(&depName, "name", "n", "", "项目名称")
	deployCmd.Flags().IntVarP(&depBnum, "bnum", "b", 0, "备份编号")
}
