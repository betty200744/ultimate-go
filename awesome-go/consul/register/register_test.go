package register

import (
	"fmt"
	"testing"

	"ultimate-go/utils"
)

var (
	c Client
)

func TestMain(m *testing.M) {
	c, _ = NewConsulClient("127.0.0.1:8500")
	m.Run()
}
func Test_client_Register(t *testing.T) {
	type args struct {
		name string
		port int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Register",
			args: args{
				name: "web",
				port: 8080,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.Register(tt.args.name, tt.args.port); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_DeRegister(t *testing.T) {
	name := "web_d"
	ip := utils.GetIp()
	port := 8080
	c.Register(name, port)
	id := fmt.Sprintf("%s-%s-%d", name, ip, port)
	fmt.Println(id)
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "DeRegister",
			args: args{id: "greeter-192.168.31.77-50051"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.DeRegister(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeRegister() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_Service(t *testing.T) {
	name := "web"
	ip := utils.GetIp()
	port := 8080
	c.Register(name, port)
	addr := fmt.Sprintf("%s:%d", ip, port)
	type args struct {
		service string
		tag     string
	}
	tests := []struct {
		name      string
		args      args
		wantAddrs []string
		wantErr   bool
	}{
		{
			name: "Service",
			args: args{
				service: "greeer",
			},
			wantAddrs: []string{addr},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAddrs, err := c.Service(tt.args.service, tt.args.tag)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("Service() gotAddrs = %v, want %v", gotAddrs, tt.wantAddrs)
		})
	}
}
