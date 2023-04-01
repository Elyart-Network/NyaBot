package plugin

var cqPlugins = make(map[string]GoCqPlugin)

func CqPlugRegister(plugin GoCqPlugin) {
	cqPlugins[plugin.Info().Name] = plugin
	return
}
