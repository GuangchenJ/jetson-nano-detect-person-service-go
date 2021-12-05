package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	"go_server/client/defs"
	"go_server/configs/db"
	"go_server/datasource/database"
	Detect "go_server/detect_service"
	"go_server/internel/biz/service"
	"go_server/internel/data/model"
)

func main() {
	config := db.InitConfig()

	// conn, err := grpc.Dial(conf.Basic.RpcNetwork()+":"+conf.Basic.RpcPort(), grpc.WithTransportCredentials(Detect.GetClientCreds()))
	conn, err := grpc.Dial("10.100.214.20:50005", grpc.WithTransportCredentials(Detect.GetClientCreds()))
	if err != nil {
		log.Panicln(err.Error())
	}
	detectResultServiceClient := Detect.NewDetectResultServiceClient(conn)

	ctx := context.Background()

	mysqlDb := database.NewMySqlDB(config.DataBase)

	pedesFlowService := service.NewPedesFlowService(mysqlDb)

	for {
		resp, err := detectResultServiceClient.DetectedRect(ctx, &Detect.DetectRequest{Status: true})
		if err != nil {
			log.Println(err.Error())
			continue
		}

		r := defs.RpcResponse{DetectResponse: resp}.ToResponse()
		// 		bytes, err := json.Marshal(r.WithoutImg())
		// 		if err != nil {
		// 			log.Println("WithoutImg: ", err.Error())
		// 			return
		// 		}
		// 		fmt.Println(string(bytes))

		err = pedesFlowService.NewPedesFlowInfo(
			model.PedesFlowInfo{CameraID: 1, Time: time.Now().Format("20060102150405"), PersonNum: uint(len(r.Rect))})
		if nil != err {
			break
		}
		time.Sleep(time.Millisecond * 1000)
	}
}