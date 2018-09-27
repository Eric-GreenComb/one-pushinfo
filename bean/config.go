package bean

// ServerConfig ServerConfig Struct
type ServerConfig struct {
	Port []string
	Mode string
}

// NsqConfig NsqConfig
type NsqConfig struct {
	Host     string
	Topic    string
	TopicNum int
}
