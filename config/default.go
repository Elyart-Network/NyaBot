package config

var serverDef = server{
	ListenPort: "3000",
	RpcPort:    "3001",
	DebugMode:  true,
	FileLogger: false,
}

var miraiDef = mirai{
	Enable: false,
}

var goCqHttpDef = goCqHttp{
	Enable:   false,
	HostUrl:  "http://127.0.0.1:5700",
	Delay:    3000,
	EnableWs: false,
}

var databaseDef = database{
	Type: "sqlite",
}

var loggingDef = logging{
	External:    false,
	MongoUri:    "mongodb://",
	CacheNum:    80,
	InternalLog: false,
}

var cacheDef = cache{
	External: false,
}

var searchDef = search{
	Enable: false,
}

var pluginDef = plugin{
	LuaEnable:    true,
	LuaScriptDir: "scripts",
	LuaSandbox:   false,
}
