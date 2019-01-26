## Go workshop - Devconf.cz 2019

Welcome and thank you for coming here!
At the end of this workshop you will understand reasons behind Go's popularity.
We will walk you through the language's syntax and idioms, so that you can start using the language quickly end effectively.
Let's begin by installing the compiler and tools.

#### Installing Go
#### Linux
On *rpm* systems:
```
sudo yum install golang

$ mkdir -p $HOME/go
$ echo 'export GOPATH=$HOME/go' >> $HOME/.bashrc
$ source $HOME/.bashrc
```

On *deb* systems:
[github.com/golang](https://github.com/golang/go/wiki/Ubuntu) manual

#### Mac & Windows
[golang.org](https://golang.org/doc/install#macos) manual

#### Verifying installation
Before we start everyone should be able to get the source code and run it:

(all *.go* files in the repository are for simplicity in the *main* package and
they all have *main()* function, so that you can easily run them with `go run`)

```
go get github.com/alexxa/go_workshop
go run 01_hello_world.go
```

#### Labs
Use go's *test* command to execute the tests in `labs/lab*`:
```
go test ./labs/lab1
```
