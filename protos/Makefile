all:

ping:
	protoc -I=. pingpong.proto --js_out=import_style=commonjs:../web-server/ui/src/protos --grpc-web_out=import_style=commonjs,mode=grpcwebtext:../web-server/ui/src/protos
	protoc -I=. pingpong.proto --go_out=plugins=grpc:../ping-server/protos --go_opt=paths=source_relative pingpong.proto

cping:
	rm ../web-server/ui/src/protos/pingpong*

.PHONY: clean
clean: cping