proto-compile:
	protoc -I proto proto/*.proto --gofast_out=plugins=grpc:proto