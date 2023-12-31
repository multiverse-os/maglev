package maglev

import (
	config "github.com/multiverse-os/maglev/config"
)

func DefaultConfig(name string) config.Settings { return config.DefaultConfig(name) }

func LoadConfig(path string) (config.Settings, error) {
	return config.Load(path)
}
