# Workspace set-up & Go project basics

The goal of this exercise is to set-up your Go development environment and practice basic commands using the `go` executable.

## Steps

### Installation

1. Download Go.
    - Consult your package manager. If it provides a up-to-date version of Go install it using it.
        - MacOS [brew]: `brew install go`
        - Ubuntu [apt]: `sudo apt install golang-go`
        - Arch [pacman]: `sudo pacman -S go`
        - Windows [cocolatey]: `choco install golang`
    - If your package manager does not provide you with Go package or if it is not up-to-date, install Go manually using [the official guide](https://go.dev/doc/install).
2. Verify your installation.
    - `go version`
3. Choose your IDE.
    - If you are fan of JetBrains tools you can use their dedicated IDE for Go, [GoLand](https://www.jetbrains.com/go/).
    - If you prefer VS Code just install the [VS Code Go](https://code.visualstudio.com/docs/languages/go) plugin.
    - If you are old school or you just like staying in your terminal using vim/neovim with [vim-go](https://github.com/fatih/vim-go) is also an excelent choice.

### Workspace

The go command and the tools it invokes consult environment variables for configuration. We will just take a quick glimpse at them.

1. Go downloads and installs all packages to a directory specified by the `GOPATH` environment variable.
    - To learn more about `GOPATH` execute `go env GOPATH`.
    - By default:
        - `$HOME/go` on Unix,
        - `%USERPROFILE%\go` on Windows
    - I personally recommend hiding the `go` directory by setting the `GOPATH` explicitly.
        - Add `export GOPATH=$HOME/.go` to your bashrc or zshrc.
2. Go SDK can be specified by the `GOROOT` variable.
    - You do not need to set this variable unless you plan to use different Go version.
3. To view the current state of all of yours Go environment variables execute: `go env`.
4. To view all configurable variables and their meaning execute: `go help environment`.

### Go

In this last section, we will create a new Go project and practice basic `go` commands.

#### Modules

1. Create a new directory for our code and move into it.
2. The `go.mod` file represents a go module. It is similar to files like:
    - `*.csproj` in C#
    - `pom.xml` in Java Maven project
    - `Cargo.toml` in Rust
    - `package.json` in Javascript
3. To create a new module execute: `go mod init <name-of-module>`.
    - You can name your module whatever you would like, e.g. "my-project". However, it is a common practice to name your module the same way as the repository in which it will be hosted. If I wanted to host this module on GitHub, I would call it "github.com/course-go/my-project".
    - The `go.mod` file specifies basic information about your project like it's name, the Go version it uses and list of it's dependencies.
    - To learn more about the `go.mod` execute: `go help go.mod`
    - To learn more about go modules execute: `go help modules`

#### Building and executing Go files

1. Create a new file named `main.go` which will serve as the projects entrypoint and insert the following code into it.
    - This is just a basic hello world program.

```
package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}
```

2. To execute the program run: `go run .`. 
    - This compiles your program and executes it immediately. 
    - The current directory has to contain the `package main` and the `main` function (the file it-self does not have to be called `main.go` but it is a good convention).
3. Try chaging the package name from `main` to `main2` and run the `go run .` again.
    - Compiler should throw an error: `package my-project is not a main package`
4. Changing the name of the `main()` function results in a similar error.
    - `runtime.main_main·f: function main is undeclared in the main package`
5. Change the source code back to the functioning state as specified above.
5. To build your program execute: `go build .`.
    - This command compiles your program and generates an executable binary.
    - The binary generates into the current directory and is named after the project's name, as specified by `go.mod`.
    - The change the output name or directory you can use the `-o` switch like so: `go build -o bin/my-binary .`.

#### Formatting

1. Change the `main.go` source code to the following snippet:
    - This is definitely not a well-formatted code. Luckily, Go ships with a default formatter.

```
package main
import"fmt"
func main(){
fmt.Println("Hello world!")}
```

2. To format all go files in the current directory execute: `go fmt .`.
    - This will change the source code to follow the uniform Go style format.
    - Note that most IDEs and editors automatically execute the `fmt` for you.
3. To format all go files in the source tree you can use the wildcard pattern: `go fmt ./...`.
    - The `./...` can be used with many other Go commands like `go list`, `go get`, `go test` etc.

#### Dependencies

1. Let's try to add some dependencies to our project. 
    - We will replace our "Hello world!" string with a constant from "github.com/course-go/code" repository.
    - To add the `course-go code` repository to the project's dependencies execute: `go get github.com/course-go/code`
    - Note that `course-go/code` was added to the `go.mod` file with a check-sum of it in `go.sum`.
    - The source code of the repository got downloaded to the `$GOPATH/pkg` directory.
        - Feel free to check it out, to get a better sense of how it works.
2. We can now use the exported code from `course-go/code` like this:

```
package main

import (
	"fmt"

	"github.com/course-go/code/pkg/hello"
)

func main() {
	fmt.Println(hello.World)
}
```

#### External tools

1. We will install the `go present` tool which is used in the course for presenting slides.
    - To install it, simply run `go install golang.org/x/tools/cmd/present@latest`.
        - Notice that we have to specify the version (in this case `latest`) when we are not in a Go module.
    – This downloads the source code, builds it and saves the executable binary into the *$GOPATH/bin* directory.
        - Do not forget to add this directory to your path for convenience like so: `export PATH="$GOPATH/bin:$PATH"`
2. Clone the [lectures repository](https://github.com/course-go/lectures).
3. Move into it and execute the present binary: `present`.
4. This should start a web server where you can check-out the slides.
