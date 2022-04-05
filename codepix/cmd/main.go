package main

import (
	"github/bonfimjustino7/codepix-go/application/grpc"
	"github/bonfimjustino7/codepix-go/infrastructure/db"
	"os"

	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {

	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)

}
