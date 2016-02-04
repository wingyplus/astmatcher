package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/wingyplus/astmatcher"
)

func findTests(src string) []string {
	astmatcher.ParseSrc(src)

	fnDecls := astmatcher.FuncDecl(astmatcher.HasName("Test"))
	funcs := make([]string, 0, len(fnDecls))

	for _, fnDecl := range fnDecls {
		funcs = append(funcs, fnDecl.Name.Name)
	}

	return funcs
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: findtest [file]")
		os.Exit(1)
	}

	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, fn := range findTests(string(b)) {
		fmt.Println(fn)
	}
}
