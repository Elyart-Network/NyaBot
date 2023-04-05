package main

// import a plugin by simply add a blank import.
import (
	// === Import plugins here ===
	_ "github.com/Elyart-Network/NyaBot/plugin/Test"
	// _ "github.com/Elyart-Network/NyaBot/plugin/Example"        // Example Plugin
	// _ "github.com/Elyart-Network/NyaBot/plugin/Internal_ExLua" // Lua Script Support
	"github.com/Elyart-Network/NyaBot/server"
)

func main() {
	server.Start()
}
