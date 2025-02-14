package core

import "os"

type configtype struct {
	Path     string
	Theme    string
	Autosave bool
}

var usrHome, _ = os.UserHomeDir()

// Config :: Access the config
var Config = configtype{
	usrHome,
	"dark",
	false,
}

// NewConfig :: Initiates user config
func NewConfig() {
	Config.Path = "/Users/devallabharath/Library/Mobile Documents/com~apple~CloudDocs/Notes"
	Config.Autosave = true
	Config.Theme = "dark"
}
