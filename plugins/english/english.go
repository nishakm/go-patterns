// package english implements the greet.Plugin interface
package english

import (
	"errors"
	"plugins/greet"
)

// pluginName is used to register this plugin with the greet module
var pluginName = "english"

// EnglishPlugin is just used to implement the greet.Plugin interface
type EnglishPlugin struct{}

// The EnglishPlugin implements the Language method in greet.Plugin
func (ep EnglishPlugin) Language(lang string) (string, error) {
	if lang == "english" {
		return "Hello all!", nil
	}
	return "", errors.New("unknown driver")
}

// func init() registers this plugin with the greet module
func init() {
	if pluginName != "" {
		greet.Register(pluginName, &EnglishPlugin{})
	}
}
