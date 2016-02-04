package astmatcher

import (
	"go/ast"
	"testing"
)

func TestFuncDecl(t *testing.T) {
	src := `package main

func hello() {}
`
	ParseSrc(src)
	if len(FuncDecl()) != 1 {
		t.Error("Expect 1 function")
	}
}

func TestFuncDecl_HasName(t *testing.T) {
	src := `package main

func hello() {}

func gopher() {}
`

	ParseSrc(src)
	fnDecls := FuncDecl()
	if len(fnDecls) != 2 {
		t.Error("Expect 2 functions")
	}
	nameEqual(t, fnDecls[0], "hello")
	nameEqual(t, fnDecls[1], "gopher")

	fnDecls = FuncDecl(HasName("gopher"))
	if len(fnDecls) != 1 {
		t.Error("Expect 1 function")
	}
	nameEqual(t, fnDecls[0], "gopher")
}

func TestFuncDecl_HasName_MatchPartial(t *testing.T) {
	src := `package main

func TestSomething() string {
        return "hello"
}
`
	ParseSrc(src)
	fnDecls := FuncDecl(HasName("Test"))
	if len(fnDecls) != 1 {
		t.Error("Expect 1 function")
		return
	}
	nameEqual(t, fnDecls[0], "TestSomething")
}

func nameEqual(t *testing.T, fnDecl *ast.FuncDecl, name string) {
	if fnDecl.Name.Name != name {
		t.Errorf("Expect name %s but got %s", fnDecl.Name.Name, name)
	}
}
