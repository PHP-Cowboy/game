package config

var Conf *Config

type Config struct {
	ServersConf ServersConf `json:"serversConf"`
}

type ServersConf struct {
	Nats           NatsConfig      `json:"nats" `
	Center         []*CenterConfig `json:"center" `
	ServerList     []*Server       `json:"serverList" `
	ServerTypeList map[string][]*Server
}

type NatsConfig struct {
	Url string `json:"url"`
}

type CenterConfig struct {
	ID         string `json:"id" `
	Host       string `json:"host" `
	ClientPort int    `json:"clientPort" `
	ServerType string `json:"serverType" `
}

type Server struct {
	Id               string `json:"id" `
	ServerType       string `json:"serverType" `
	HandleTimeOut    int    `json:"handleTimeOut" `
	RPCTimeOut       int    `json:"rpcTimeOut" `
	MaxRunRoutineNum int    `json:"maxRunRoutineNum" `
}
