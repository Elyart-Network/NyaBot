package config

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
