package clients

import (
	"flag"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/selector"
	"github.com/micro/go-plugins/broker/nsq"
	"github.com/micro/go-plugins/registry/etcdv3"
	"strings"
	"time"
)

var (
	etcdAddr = flag.String("etcd", "", "register etcd address")
	etcdUser = flag.String("user","","etcd username")
	etcdPass = flag.String("pass","","etcd password")
	nsqdAddr = flag.String("nsqd", "", "nsqd address")
)

func init()  {
	flag.Parse()
	// 初始化注册中心
	InitRegistry(strings.Split(*etcdAddr,","),*etcdUser,*etcdPass)
	// 初始化消息中间件
	InitBroker(strings.Split(*nsqdAddr,","))
	// 初始化rpc客户端
	InitClient()
}

// 定义注册中心
var Registry registry.Registry

// 初始化注册中心
func InitRegistry(etcdAddrs []string,etcdUser string, etcdPass string) registry.Registry {
	if len(etcdAddrs) < 1 || strings.TrimSpace(etcdAddrs[0]) == "" {
		return nil
	}
	if strings.TrimSpace(etcdUser) != "" && strings.TrimSpace(etcdPass) !="" {
		Registry = etcdv3.NewRegistry(func(options *registry.Options) {
			options.Addrs = etcdAddrs
			etcdv3.Auth(etcdUser,etcdPass)
		})
	}else {
		Registry = etcdv3.NewRegistry(func(options *registry.Options) {
			options.Addrs = etcdAddrs
		})
	}
	return Registry
}

// 定义一个nsqBroker
var NsqBroker broker.Broker

// 初始化 broker
func InitBroker(nsqdAddrs []string) {
	if len(nsqdAddrs) < 1 || strings.TrimSpace(nsqdAddrs[0]) == "" {
		NsqBroker = nil
	}else{
		NsqBroker = nsq.NewBroker(func(options *broker.Options) {
			options.Addrs = nsqdAddrs
		})
	}
}

// 定义 client 对象
var Client client.Client

// 初始还一个默认的rpc客户端
func InitClient() {
	s := selector.NewSelector(selector.Registry(Registry))
	options := []client.Option{client.Registry(Registry),
		client.RequestTimeout(15*time.Second),
		client.Retries(3),
		client.Selector(s)}
	if NsqBroker != nil {
		options = append(options, client.Broker(NsqBroker))
	}
	Client = client.NewClient(options...)
}

// 指定参数创建rpc客户端
func GetClient(retries int ,requestTimeout time.Duration) client.Client {
	if retries < 1 {
		retries = 1
	}
	if requestTimeout < time.Second {
		requestTimeout = 3 * time.Second
	}
	// new 负载均衡选择器
	s := selector.NewSelector(selector.Registry(Registry))
	options := []client.Option{client.Registry(Registry),
		client.RequestTimeout(requestTimeout),
		client.Retries(retries),
		client.Selector(s)}
	if NsqBroker != nil {
		options = append(options, client.Broker(NsqBroker))
	}
	return client.NewClient(options...)
}

