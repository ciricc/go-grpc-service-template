genproto:
	protoc --go_out=. --go_opt=module=github.com/ciricc/go-grpc-service-template --go-grpc_out=. \
	--go-grpc_opt=module=github.com/ciricc/go-grpc-service-template -Iprotos protos/*.proto && \
	protoc --go_out=. --go_opt=module=github.com/ciricc/go-grpc-service-template --go-grpc_out=. \
	--go-grpc_opt=module=github.com/ciricc/go-grpc-service-template -Iprotos protos/types/*.proto
build:
	docker build -f Dockerfile --build-arg GITHUB_TOKEN=${GITHUB_TOKEN} -t ciricc/printer .