module greeter_gateway

go 1.15
require (
	proto v0.0.0
	google.golang.org/grpc v1.39.0
)

replace (
	proto => ./../proto
)