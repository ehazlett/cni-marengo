package api

//go:generate protoc -I.:../vendor:../vendor/github.com/gogo/protobuf --gogo_out=plugins=grpc,import_path=github.com/ehazlett/marengo/api,Mgogoproto/gogo.proto=github.com/gogo/protobuf/gogoproto:. api.proto ipam.proto
