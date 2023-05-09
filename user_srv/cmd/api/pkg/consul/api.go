package consul

import (
	"fmt"

	"github.com/hashicorp/consul/api"

	"github.com/sjxiang/ddshop/user_srv/cmd/api/pkg/conf"
)

func Register(addr string, port int, name string, tags []string, id string) error {
	
	cfg := api.DefaultConfig()
	cfg.Address = conf.Conf.Consul.Addr
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// 生成对应的检查对象
	check := &api.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s/v1/health", conf.Conf.App.Addr),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",  // 故障检查失败后，consul 自动将注册服务删除
	}

	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = addr
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	return nil 
}


// 根据 ID 获取服务
func QueryServiceByID(id string) {
	cfg := api.DefaultConfig()
	cfg.Address = conf.Conf.Consul.Addr
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == %s", id))
	if err != nil {
		panic(err)
	}


	var (
		svrHost string
		svrPort int
	)

	for k, v := range data {
		svrHost = v.Address
		svrPort = v.Port
		fmt.Println("获取到服务：" + k)
		fmt.Println("IP:PORT", svrHost, svrPort)
	}
}