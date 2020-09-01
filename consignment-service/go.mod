module github.com/microke/consignment-service

go 1.13

require (
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro v1.18.0
	github.com/microke/vessel-service v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20200822124328-c89045814202
)

replace github.com/microke/vessel-service => ../vessel-service
