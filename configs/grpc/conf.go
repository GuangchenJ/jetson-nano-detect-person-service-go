package grpc

import (
	"errors"
	"log"
	"os"
)

var (
	Basic = &basic{}

	env = &envDef{
		basicRpcNetwork: "JETSON_NANO_DETECT_SERVER_BASIC_RPC_NETWORK",
		basicRpcPort:    "JETSON_NANO_DETECT_SERVER_BASIC_RPC_PORT",
	}
)

func init() {
	// typeOfEnv := reflect.TypeOf(*env)
	// valueOfEnv := reflect.ValueOf(*env)
	//
	// for i := 0; i < typeOfEnv.NumField(); i++ {
	//	if value := os.Getenv(valueOfEnv.Field(i).String()); value != "" {
	//		SetEnv(typeOfEnv.Field(i).Name, value)
	//	} else {
	//		EnvError(errors.New("env should not be empty"))
	//	}
	// }
}

func SetEnv(key, value string) {
	switch key {
	case "basicRpcNetwork":
		Basic.rpcNetwork = value
	case "basicRpcPort":
		Basic.rpcPort = value
	default:
		EnvError(errors.New("unknown env name"))
	}
}

func EnvError(err error) {
	log.Println(err)
	os.Exit(-1)
}

type envDef struct {
	basicRpcNetwork string
	basicRpcPort    string
}

type basic struct {
	rpcNetwork string
	rpcPort    string
}

func (b basic) RpcNetwork() string {
	return b.rpcNetwork
}

func (b basic) RpcPort() string {
	return b.rpcPort
}
