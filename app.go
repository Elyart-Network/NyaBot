package main

// import a plugin by simply add a blank import.
import (
	// === Import plugins here ===
	_ "github.com/Elyart-Network/NyaBot/plugin/Example"
	"github.com/Elyart-Network/NyaBot/server"
)

func main() {
	server.Start()
}
