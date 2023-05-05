package config

type config struct {
	Server   server   `yaml:"server"`
	Mirai    mirai    `yaml:"mirai"`
	GoCqHttp goCqHttp `yaml:"gocqhttp"`
	Database database `yaml:"database"`
	Logging  logging  `yaml:"logging"`
	Cache    cache    `yaml:"cache"`
	Search   search   `yaml:"search"`
	Plugin   plugin   `yaml:"plugin"`
}

type server struct {
	ListenPort string `yaml:"listen_port"`
	DebugMode  bool   `yaml:"debug_mode"`
	FileLogger bool   `yaml:"file_logger"`
}

type mirai struct {
	Enable bool `yaml:"enable"`
}

type goCqHttp struct {
	Enable   bool   `yaml:"enable"`
	HostUrl  string `yaml:"host_url"`
	Delay    int    `yaml:"delay"`
	EnableWs bool   `yaml:"enable_ws"`
}

// SQLite or MySQL or PostgreSQL
type database struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// MongoDB
type logging struct {
	External    bool   `yaml:"external"`
	MongoUri    string `yaml:"mongo_uri"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	CacheNum    int    `yaml:"cache_num"`
	InternalLog bool   `yaml:"internal_log"`
}

// InMem or Redis
type cache struct {
	External  bool   `yaml:"external"`
	Host      string `yaml:"host"`
	IndexName string `yaml:"index_name"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
}

// Elasticsearch
type search struct {
	Enable    bool   `yaml:"enable"`
	Host      string `yaml:"host"`
	IndexName string `yaml:"index_name"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
}

// Plugins
type plugin struct {
	LuaEnable    bool   `yaml:"lua_enable"`
	LuaScriptDir string `yaml:"lua_script_dir"`
	LuaSandbox   bool   `yaml:"lua_sandbox"`
}
