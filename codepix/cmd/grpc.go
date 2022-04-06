package cmd

import (
	"github/bonfimjustino7/codepix-go/application/grpc"
	"github/bonfimjustino7/codepix-go/infrastructure/db"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
)

var portNumber int
var database *gorm.DB

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Start gRPC server",

	Run: func(cmd *cobra.Command, args []string) {
		database = db.ConnectDB(os.Getenv("env"))
		grpc.StartGrpcServer(database, portNumber)
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
	grpcCmd.Flags().IntVarP(&portNumber, "port", "p", 5051, "gRPC Server port")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
