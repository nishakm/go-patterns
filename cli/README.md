# Go CLI creation

Most people use cobra to create a cli. The directory construction looks like this:

```
▾ goproject/
    ▾ cmd/
        subcommand1.go
        subcommand2.go
      main.go
    ▾ pkg/
```

`cmd` is where the CLI commands are created. `pkg` is where libraries are located.

## Using Cobra

Cobra has a generator that will create all the boilerplate for you.

```
$ go install github.com/spf13/cobra-cli@latest
$ mkdir cli
$ cd cli
$ go mod init cli
$ cobra-cli init
```
There should now be a `cmd` directory and a `main.go` file.

```
package main

import "cli/cmd"

func main() {
  cmd.Execute()
}
```

Package `cmd` is in the `cmd` folder. The main file is called `root.go`. The file creates a `cobra.Command` object and the function `Execute` calls the Command's `Execute` method.

All flags are defined in the `init()` function.

## Application

Let's make a greeting app where the default greeting is in English and we have subcommands for Hindi and Tamil. Our UX should look like this:
```
$ cli
Hello
$ cli hindi
नमस्ते
$ cli tamil
வணக்கம்
```

## Default Run
In the `cobra.Command object there is a `Run` property which is a user defined function. We'll define that function like this:

```
Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Hello")
},
```

Now we can build the project using `go build` which will give the binary executable called `cli`. Running just `./cli` will produce "Hello" which is the UX we want.

## Adding a subcommand

One can add a subcommand boilerplate by running:

```
$ cobra-cli add hindi 
```

When running the command, a file called `hindi.go` is created in the `cmd` directory. This creates a new `cobra.Command` object which stores a function to run something. Here we can put the "hindi" greeting.

We can similarly add another one for `tamil`.

## Adding command line options

`cobra-cli` cannot add optional flags. But we can use either `PersistentFlags()` or `Flags()` methods of the `Command` object.

For example, we will add a flag `-e, --english` to the hindi command which will just spell "hello" using the hindi script. To do this, go to the `init()` function in `hindi.go` and add the following:

```
hindiCmd.Flags().BoolP("english", "e", false, "Write hello in hindi")
```

The flag called `english` should now be accessable to the Run function:

```
if cmd.Flags().Changed("english) {...
```

Unfortunately, the full set of methods accessable to the `FlagSet` object is not documented in the Cobra user guide. So to do more sophisticated checks, look at the godoc for Cobra.
