package main

// import a plugin by simply add a blank import.
import (
	"github.com/Elyart-Network/NyaBot/config"
	_ "github.com/Elyart-Network/NyaBot/internal/logging" // Logging Plugin
	"github.com/Elyart-Network/NyaBot/server"
	// _ "github.com/Elyart-Network/NyaBot/examples/plugin" // Example Plugin
	_ "github.com/Elyart-Network/NyaBot/internal/lua" // Lua Plugin
	// _ "github.com/Elyart-Network/NyaBot/internal/grpc"        // gRPC Plugin
	// _ "github.com/Elyart-Network/NyaBot/internal/javascript"  // JavaScript Plugin
)

func main() {
	config.EnvInit() // Initialize environment variables
	server.Start()   // Start server (framework)
}
