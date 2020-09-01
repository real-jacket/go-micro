build:
	protoc --proto_path=. -I. --go_out=plugins=micro:. proto/vessel/vessel.proto
	docker build -t vessel-service .

run:
	docker run -p 50052:5001\
	-e MICRO_SERVER_ADDRESS=:50051\
	-e MICRO_REGISTRY=mdns vessel-service