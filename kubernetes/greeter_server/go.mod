module greeter-server

go 1.15

require (
	google.golang.org/grpc v1.39.0
	proto v0.0.0
)

replace proto => ./../proto
