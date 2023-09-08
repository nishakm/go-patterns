u Go compile time plugins

Typically, plugins are invoked at run time. Since go's standard library plugin has [several warnings](https://pkg.go.dev/plugin#hdr-Warnings) associated with it, we will instead try to create plugins that will be compiled into the binary.

## Why would we want to do this?

This allows users of the source code to compile in their own special plugin for a binary they are distributing. Other than that, I don't see any other reason to implement compile-time plugins in go.

## Example

The typical example of using compile time plugins is go's `database/sql` module which has the concept of "drivers". I used [this blog post](https://eli.thegreenplace.net/2019/design-patterns-in-gos-databasesql-package/).

This is a toy example that imitates the thing above. I'm using [this plugin example](https://github.com/vladimirvivien/go-plugin-example) where the main function prints a greeting and each plugin implements the greeting in a different language.

But rather than implementing them using shared libraries, I'll implement them using the `database/sql` technique.

We want:

- A user facing Go Type
- A "greeter" interface that the user facing Go Type uses

For a user to use a specific "greeter" they need to do the following:

```
// import the main package
import "plugins/greet"

// import the specific greeter to use
import _ "plugins/english"

// invoke the greeter
func main() {
  greeting, err := greet.In("english")
  ...
}

```

## Functionality

We want to be able to:

- discover what plugins are available to use
- register the plugin with the main package
- wire up the plugin to the application
- expose some options from the main package to the plugin
- expose some options from the plugin to the main package

## Implementation

1. To simulate two people implementing two different things, I'm creating one directory for the "greet" library and another one for the "english" plugin.
1. Create the required interface in `greet.go`. I'm calling this interface `Driver` for now.
2. Create the plugin `english.go`. In order to implement an interface, you have to create a struct which then has a method that implements the interface.
3. Create the `main.go` app which imports the `plugins/greet` and the `plugins/english` modules and creates the `greet.Driver` interface by instantiating the `EnglishDriver` implementation.

The code should look like this:

`greet.go`
```
package greet

type Driver interface {
	In(language string) (string, error)
}
```

`english.go`
```
package english

import "errors"

type EnglishDriver struct {
	Greeting string
}

func (ed EnglishDriver) In(lang string) (string, error) {
	ed.Greeting = "Hello all!"
	if lang == "english" {
		return ed.Greeting, nil
	}
	return "", errors.New("unknown driver")
}
```

`main.go`
```
package main

import (
	"fmt"
	"plugins/english"
	"plugins/greet"
)

func main() {
	var driver greet.Driver = english.EnglishDriver{}
	output, _ := driver.In("english")
	fmt.Println(output)
}

```
Now No. 3 doesn't follow the implicit instantiation where I don't know before hand what kind of concrete implementation I am instantiating. The `database/sql` pattern of doing this is registering the plugin with the core library.

4. Add some functionality to `greet.go` to "register" the a Plugin object

Now the code looks like this:
```
package greet

import "errors"

type Plugin interface {
	Language(name string) (string, error)
}

var plugins = make(map[string]Plugin)

func Register(name string, plugin Plugin) {
	if plugin == nil {
		panic("A Plugin object is required, but is nil")
	}
	if _, dup := plugins[name]; dup {
		panic("Register called twice for plugin name " + name)
	}
	plugins[name] = plugin
}

func In(name string) (string, error) {
	plugin := plugins[name]
	if plugin == nil {
		return "", errors.New("No plugin registered with name " + name)
	} else {
		return plugin.Language(name)
	}
}
```
I changed the interface to `Plugin` and its method `Language` to make it clearer what's happening. I added a global map where the keys are the plugin names and the values are the particular Plugin object, a function to add the plugin to the map, and a function to return the required string.

I would normally just return the plugin object and call whatever method, but I'm trying to keep the code small.

5. Add `init()` function to `english.go`. This is how the plugin makes itself known to the greet module by adding a key to the `plugins` map.

The code in `english.go` now looks like this:
```
package english

import (
	"errors"
	"plugins/greet"
)

var pluginName = "english"

type EnglishPlugin struct{}

func (ep EnglishPlugin) Language(lang string) (string, error) {
	if lang == "english" {
		return "Hello all!", nil
	}
	return "", errors.New("unknown driver")
}

func init() {
	if pluginName != "" {
		greet.Register(pluginName, &EnglishPlugin{})
	}
}
```

So now, the main.go code can do this:
```
package main

import (
	"fmt"
	_ "plugins/english"
	"plugins/greet"
)

func main() {
	greeting, err := greet.In("english")
	fmt.Println(greeting)
	greeting, err = greet.In("hindi")
	if err != nil {
		fmt.Println(err)
	}
}
```

And running this code gives:
```
$ go run main.go 
Hello all!
No plugin registered with name hindi
```
