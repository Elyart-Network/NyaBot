package main

// import a plugin by simply add a blank import.
import (
	"github.com/Elyart-Network/NyaBot/internal/server"
	// === Import plugins here ===
	_ "github.com/Elyart-Network/NyaBot/plugin/flarum_group" // Help Plugin
	// _ "github.com/Elyart-Network/NyaBot/examples/plugin"        // Example Plugin
	// === Internal Plugins ===
	// _ "github.com/Elyart-Network/NyaBot/internal/botcli" // Bot CLI
	// _ "github.com/Elyart-Network/NyaBot/internal/lua"    // Lua Plugin
)

func main() {
	server.Start()
}
