package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/connect"
)

const (
	name = "web-svr"
	port = 30003
)

// todo
func handleResponse(resp *http.Response) {
	/*	b, _ := io.ReadAll(resp.Body)
		fmt.Println(string(b[:]))*/
}

func HTTPClient(client *api.Client, _url string) {
	svr, _ := connect.NewService("web-svr", client)
	defer svr.Close()
	httpClient := svr.HTTPClient()
	resp, _ := httpClient.Get(_url)
	defer resp.Body.Close()
	handleResponse(resp)
}

func ServiceRegister(client *api.Client) {
	agent := client.Agent()
	s := &api.AgentServiceRegistration{
		Name: name,
		Port: port,
	}
	agent.ServiceRegister(s)
}

func main() {
	client, _ := api.NewClient(api.DefaultConfig())
	ServiceRegister(client)
	HTTPClient(client, "http://localhost:8500/v1/catalog/service/web-svr")

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)

}
