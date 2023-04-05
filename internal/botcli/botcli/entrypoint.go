package botcli

func Entry(callback interface{}, cmd []string, protocol string) {
	switch protocol {
	case "GoCqHttp":
		cqProcessor(callback, cmd)
	case "Discord":
		discordProcessor(callback, cmd)
	case "Telegram":
		telegramProcessor(callback, cmd)
	case "Slack":
		slackProcessor(callback, cmd)
	}
}
