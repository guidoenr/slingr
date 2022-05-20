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

import "fmt" // a library to import debuggin and printing funcs

func main() {
	fmt.Println("Hi there!")
}

