package config

type config struct {
	Server   server   `yaml:"server"`
	GoCqHttp goCqHttp `yaml:"gocqhttp"`
	Database database `yaml:"database"`
	Queue    queue    `yaml:"queue"`
	Cache    cache    `yaml:"cache"`
	Search   search   `yaml:"search"`
}

type server struct {
	HttpPort   string `yaml:"http_port"`
	WsPort     string `yaml:"ws_port"`
	EnableWs   bool   `yaml:"enable_ws"`
	DebugMode  bool   `yaml:"debug_mode"`
	FileLogger bool   `yaml:"file_logger"`
}

type goCqHttp struct {
	Enable   bool   `yaml:"enable"`
	HttpHost string `yaml:"http_host"`
	WsListen bool   `yaml:"ws_listen"`
}

type database struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type queue struct {
	Type      string `yaml:"type"`
	Host      string `yaml:"host"`
	IndexName string `yaml:"index_name"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
}

type cache struct {
	Type      string `yaml:"type"`
	Host      string `yaml:"host"`
	IndexName string `yaml:"index_name"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
}

type search struct {
	Type      string `yaml:"type"`
	Host      string `yaml:"host"`
	IndexName string `yaml:"index_name"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
}

var serverDef = server{
	HttpPort:   "3000",
	WsPort:     "4000",
	EnableWs:   false,
	DebugMode:  true,
	FileLogger: false,
}

var goCqHttpDef = goCqHttp{
	Enable:   true,
	HttpHost: "http://127.0.0.1:5700",
	WsListen: false,
}

var databaseDef = database{
	Type: "internal",
}

var queueDef = queue{
	Type: "internal",
}

var cacheDef = cache{
	Type: "internal",
}

var searchDef = search{
	Type: "internal",
}
