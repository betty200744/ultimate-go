package register

import (
	"fmt"

	"ultimate-go/utils"

	"github.com/hashicorp/consul/api"
)

type Client interface {
	Service(service, tag string) (addrs []string, err error)
	Register(name string, port int) (err error)
	DeRegister(id string) (err error)
}

type client struct {
	consul *api.Client
}

func NewConsulClient(addr string) (Client, error) {
	config := api.DefaultConfig()
	config.Address = addr
	c, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &client{consul: c}, nil
}

func (c *client) Service(service, tag string) (addrs []string, err error) {
	items, _, err := c.consul.Health().Service(service, tag, false, nil)
	for _, i2 := range items {
		add := fmt.Sprintf("%s:%d", i2.Service.Address, i2.Service.Port)
		addrs = append(addrs, add)
	}
	return
}
func (c *client) Register(name string, port int) (err error) {
	ip := utils.GetIp()
	service := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", name, ip, port),
		Name:    name,
		Port:    port,
		Address: ip,
	}
	check := &api.AgentServiceCheck{
		TCP:      fmt.Sprintf("%s:%d", ip, port),
		Interval: "5s",
		Timeout:  "15s",
	}
	service.Check = check
	err = c.consul.Agent().ServiceRegister(service)
	if err != nil {
		return
	}
	return
}
func (c *client) DeRegister(id string) (err error) {
	c.consul.Agent().ServiceDeregister(id)
	return
}
