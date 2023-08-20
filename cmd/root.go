package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "terrago",
	Short: "A sample Cobra CLI tool",
}

var (
	year  int
	month int
	day   int
	hour  int
	min   int
)

func init() {
	rootCmd.AddCommand(clusterInfoCmd)
	// clusterinfo history will take year, month, day, hour, min as input
	rootCmd.AddCommand(clusterInfoCmdhistory)
	rootCmd.AddCommand(GET)

	// clusterinfo history will take year, month, day, hour, min as input
	clusterInfoCmdhistory.Flags().IntVar(&year, "year", 0, "Year")
	clusterInfoCmdhistory.Flags().IntVar(&month, "month", 0, "Month")
	clusterInfoCmdhistory.Flags().IntVar(&day, "day", 0, "Day")
	clusterInfoCmdhistory.Flags().IntVar(&hour, "hour", 0, "Hour")
	clusterInfoCmdhistory.Flags().IntVar(&min, "min", 0, "Minute")

	GET.AddCommand(Deployments)
	Deployments.Flags().StringVar(&Name, "name", "", "Name of the deployment")
	GET.AddCommand(Pods)
	Pods.Flags().StringVar(&Name, "name", "", "Name of the pod")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
