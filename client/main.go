package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	"go_server/configs/db"
	"go_server/datasource/database"
	Detect "go_server/detect_service"
	"go_server/internel/biz/service"
	"go_server/internel/data/model"
)

func main() {
	config := db.InitConfig()

	// conn, err := grpc.Dial(conf.Basic.RpcNetwork()+":"+conf.Basic.RpcPort(), grpc.WithTransportCredentials(Detect.GetClientCreds()))
	conn, err := grpc.Dial(config.Grpc.Host+":"+config.Grpc.Port, grpc.WithTransportCredentials(Detect.GetClientCreds()))
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

		// r := defs.RpcResponse{DetectResponse: resp}.ToResponse()
		// 		bytes, err := json.Marshal(r.WithoutImg())
		// 		if err != nil {
		// 			log.Println("WithoutImg: ", err.Error())
		// 			return
		// 		}
		// 		fmt.Println(string(bytes))
		if resp.Status {
			for _, newPedesFlowInfo := range resp.GetCameraRect() {
				err = pedesFlowService.NewPedesFlowInfo(
					model.PedesFlowInfo{CameraID: uint(newPedesFlowInfo.CameraId), Time: time.Now().Format("20060102150405"), PersonNum: uint(len(newPedesFlowInfo.BoxRect))})
			}
		}

		// err = pedesFlowService.NewPedesFlowInfo(
		//     model.PedesFlowInfo{CameraID: resp.CameraRect.CameraId, Time: time.Now().Format("20060102150405"), PersonNum: uint(len(resp.CameraRect))})
		//
		if nil != err {
			break
		}
		time.Sleep(time.Millisecond * 1000)
	}
}
