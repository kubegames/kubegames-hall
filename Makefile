NAME=ALL
PUSH=false
DEPLOY=true

install:
	cd ./cmd && ./build.sh $(NAME) $(PUSH) $(DEPLOY)
	
delete:
	cd ./cmd && ./clean.sh $(NAME)

proto:
	./app/build.sh
	
clean:
	./app/clean.sh

tools:
	go install github.com/kubegames/kubegames-ctl/kubegames@v1.0.0
	go install github.com/kubegames/kubegames-ctl/grpc-gateway/protoc-gen-openapiv2@v1.0.0
	go install github.com/kubegames/kubegames-ctl/protobuf/protoc-gen-gofast@v1.0.0
	go install github.com/kubegames/kubegames-ctl/protobuf/protoc-gen-gin@v1.0.0