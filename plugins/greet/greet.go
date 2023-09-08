package greet

import "errors"

// Plugin is the interface that plugins need to implement
// Language is the method that takes the plugin name and returns the
// greeting with the string and error object
type Plugin interface {
	Language(name string) (string, error)
}

// a map storing a mapping of the plugin name to its relevant
// Plugin interface object
var plugins = make(map[string]Plugin)

// Register registers a particular language plugin by name
func Register(name string, plugin Plugin) {
	if plugin == nil {
		panic("A Plugin object is required, but is nil")
	}
	if _, dup := plugins[name]; dup {
		panic("Register called twice for plugin name " + name)
	}
	plugins[name] = plugin
}

// In is the main function that looks for available plugins
// If the plugin is there, it calls the plugin's Language implementation
func In(name string) (string, error) {
	plugin := plugins[name]
	if plugin == nil {
		return "", errors.New("No plugin registered with name " + name)
	} else {
		return plugin.Language(name)
	}
}
