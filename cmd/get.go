package cmd

import (
	db "cc/DB"
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var GET = &cobra.Command{
	Use:   "get",
	Short: "get the information about the objects",
	Long:  "get the information about the objects such as pods deployments etc",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Plese Provide a valid subcommand ")
	},
}

var Name string

var Deployments = &cobra.Command{
	Use:   "deployments",
	Short: "get the information about the deployments",
	Long:  "get the information about the deployments",
	Run: func(cmd *cobra.Command, args []string) {
		object := "deployments"
		year := strconv.Itoa(time.Now().Year())
		month := strconv.Itoa(int(time.Now().Month()))
		day := strconv.Itoa(time.Now().Day())
		hour := strconv.Itoa(time.Now().Hour())
		min := strconv.Itoa(time.Now().Minute())

		DBCLIENT := db.CreateClient()
		DBCLIENT.ReadData(year, month, day, hour, min, object, Name)

	},
}

var Pods = &cobra.Command{
	Use:   "pods",
	Short: "get the information about the pods",
	Long:  "get the information about the pods",
	Run: func(cmd *cobra.Command, args []string) {
		object := "pods"
		year := strconv.Itoa(time.Now().Year())
		month := strconv.Itoa(int(time.Now().Month()))
		day := strconv.Itoa(time.Now().Day())
		hour := strconv.Itoa(time.Now().Hour())
		min := strconv.Itoa(time.Now().Minute())

		DBCLIENT := db.Client{
			ReadAccess:  true,
			WriteAccess: false,
			HashId:      "1234567890",
		}
		DBCLIENT.ReadData(year, month, day, hour, min, object, Name)

	},
}
