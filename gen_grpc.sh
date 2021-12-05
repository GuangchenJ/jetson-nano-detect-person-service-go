#!/bin/bash

# 生成 Detect 的代码
protoc-3.18.1.0 --proto_path=proto/ --go_out=./ --go-grpc_out=./ ./proto/detect.proto

printf "Generate 'detect.proto' finished\n"
