package main

import (
	"os"

	"github.com/jandersn01/microservices/shipping/config"
	"github.com/jandersn01/microservices/shipping/internal/adapters/grpc"
	"github.com/jandersn01/microservices/shipping/internal/application/core/api"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{FieldMap: log.FieldMap{
		"msg": "message",
	}})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main(){
	application := api.NewApplication()
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()

}