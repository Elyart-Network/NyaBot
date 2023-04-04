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
	ListenPort string `yaml:"listen_port"`
	DebugMode  bool   `yaml:"debug_mode"`
	FileLogger bool   `yaml:"file_logger"`
}

type goCqHttp struct {
	Enable    bool   `yaml:"enable"`
	HostUrl   string `yaml:"host_url"`
	Delay     int    `yaml:"delay"`
	EnableWs  bool   `yaml:"enable_ws"`
	WsForward bool   `yaml:"ws_forward"`
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
	ListenPort: "3000",
	DebugMode:  true,
	FileLogger: false,
}

var goCqHttpDef = goCqHttp{
	Enable:    true,
	HostUrl:   "http://127.0.0.1:5700",
	Delay:     3000,
	EnableWs:  false,
	WsForward: false,
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
