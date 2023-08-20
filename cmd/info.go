package cmd

import (
	db "cc/DB"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var clusterInfoCmd = &cobra.Command{
	Use:   "clusterinfo",
	Short: "Print cluster information",
	Run: func(cmd *cobra.Command, args []string) {

		clusterinfo, err := db.GetPastData(db.CALANDER, db.YEAR, db.MONTH, db.DAY, db.HOUR, db.MINUTE)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(color.YellowString(clusterinfo.ClusterInfo.ClusterName))
	},
}

var clusterInfoCmdhistory = &cobra.Command{
	Use:   "clusterinfohistory",
	Short: "Print cluster information",
	Run: func(cmd *cobra.Command, args []string) {

		time.Sleep(1 * time.Second)
		color.Yellow("Fetching cluster information")
		time.Sleep(1 * time.Second)
		fmt.Printf(".")
		time.Sleep(1 * time.Second)
		fmt.Printf(".")

		clusterinfo, err := db.GetPastData(db.CALANDER, year, month, day, hour, min)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(clusterinfo.ClusterInfo)
	},
}
