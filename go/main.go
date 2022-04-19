// package = project = workspace
/*
a package can have a lot of files inside, like main.go, helper.go and support.go
but every file should has the 'package main' at the beginning,
why the 'main' name?
in go there are two types of packages

Executable: generates a file that we can run (like this one)
Reusable: code used as 'helpers' good place to put reusable logic, if we
are coding some kind of libraries or maybe a code to share with our friends
we gonna use the 'reusable' packages, not using the 'package main' at the top
something like 'package guido-tools' or whatelse 

*/
package main //so if the package is exectuable, there will be a 'package main' at the top of the file

import "fmt"

func main() {
	fmt.Println("Hi there!")
}

/*

go build -> compiles a bunch of go source code files, and creates a executable 'main' file in the dir
and it can be runned with $ ./main

go run -> compiles and executes on or two fiels

go fmt -> fomrats all thhe code in each file in the current directory

go install -> compiles and installs a package [like maven]

go get -> downloads the raw source code [like maven]

*/
