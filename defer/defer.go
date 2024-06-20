package main

import (
	"fmt"
	"os"
)

// defer is used to execute after all the surrounding code has been executed

func createFile(p string) *os.File {
	f , err := os.Create(p)
	fmt.Println("file created....")
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing to file...")
	fmt.Fprintln(f , "hello I am writing this programmatically");
	fmt.Println("wrote to the file")
}

// os.Stderr Stderr is capital because it is exported by the os package to use by other packages

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr , "error : %v\n" , err);
		os.Exit(1);
	}
}


func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f);
	writeFile(f);
}