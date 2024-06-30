package config

import (
	"github.com/foodi-org/foodi-user-proxy/internal/handler/pkg/pkgconsul"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

var servConf = Config{}

type (
	Config struct {
		rest.RestConf
		Auth struct {
			AccessSecret string
			AccessExpire int64
		}

		UserRpc zrpc.RpcClientConf

		//AccountRpc zrpc.RpcClientConf

		//ArticleRpc zrpc.RpcClientConf

		Consul consul.Conf
	}

	ProxyConsulJson struct {
		ServiceName string `json:"serviceName"`

		Host string `json:"host"`

		Port int `json:"port"`

		Consul struct {
			Host []string `json:"host"`
			Key  string   `json:"key"`
		} `json:"consul"`

		AccountConf struct {
			Target string `json:"target"`
		} `json:"accountConf"`

		UserRpcConf struct {
			Target string `json:"target"`
		} `json:"UserRpcConf"`

		Auth struct {
			AccessSecret string `json:"accessSecret"`
			AccessExpire int64  `json:"accessExpire"`
		} `json:"auth"`
	}

	ServYaml struct {
		UnmarshalKey string     `yaml:"UnmarshalKey"`
		ConsulYaml   ConsulYaml `yaml:"Consul"`
	}

	ConsulYaml struct {
		Key   string   `yaml:"Key"`
		Meta  Meta     `yaml:"Meta"`
		Tag   []string `yaml:"Tag"`
		Host  string   `yaml:"Host"`
		Token string   `yaml:"Token"`
		TTL   int      `yaml:"TTL"`
	}

	Meta struct {
		Protocol string `yaml:"Protocol"`
	}
)

func ServConf() *Config {
	return &servConf
}

func InitServConf(path string, filename string) error {
	var (
		servYaml  ServYaml
		proxyJson ProxyConsulJson
	)

	// 解析yaml配置文件
	file, err := os.ReadFile(filepath.Join(path, "etc", filename))
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(file, &servYaml); err != nil {
		return err
	}

	servConf.Consul = consul.Conf{
		Host:  servYaml.ConsulYaml.Host,
		Key:   servYaml.ConsulYaml.Key,
		Token: servYaml.ConsulYaml.Token,
		Tag:   servYaml.ConsulYaml.Tag,
		TTL:   servYaml.ConsulYaml.TTL,
	}
	if err = servConf.Consul.Validate(); err != nil {
		return err
	}
	consulClient := pkgconsul.NewConsulDO()
	consulClient.SetConfig(servConf.Consul)

	// consul connect
	if err = consulClient.Connect(); err != nil {
		panic(err)
	}

	// load service config from consul with key
	if err = consulClient.LoadJsonConfig(servYaml.UnmarshalKey, &proxyJson); err != nil {
		panic(err)
	}

	// set proxy config
	servConf.Name = proxyJson.ServiceName
	servConf.Host = proxyJson.Host
	servConf.Port = proxyJson.Port
	//servConf.AccountRpc.Target = proxyJson.UserConf.Target
	servConf.UserRpc.Target = proxyJson.UserRpcConf.Target
	//servConf.ArticleRpc.Target = proxyJson.UserConf.Target
	servConf.Auth.AccessSecret = proxyJson.Auth.AccessSecret
	servConf.Auth.AccessExpire = proxyJson.Auth.AccessExpire

	return nil
}
