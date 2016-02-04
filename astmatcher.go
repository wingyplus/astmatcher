package astmatcher

import (
	"go/ast"
	"go/parser"
	"go/token"
	"regexp"
)

var file *ast.File

type MatcherFunc func(*ast.FuncDecl) bool

func (matcher MatcherFunc) Match(fnDecl *ast.FuncDecl) bool {
	return matcher(fnDecl)
}

func HasName(pattern string) MatcherFunc {
	return func(fnDecl *ast.FuncDecl) bool {
		matched, err := regexp.MatchString(pattern, fnDecl.Name.Name)
		return (err == nil) && matched
	}
}

func FuncDecl(matchers ...MatcherFunc) []*ast.FuncDecl {
	fnDecls := []*ast.FuncDecl{}

	for _, decl := range file.Decls {
		switch fnDecl := decl.(type) {
		case *ast.FuncDecl:
			if match(fnDecl, matchers) {
				fnDecls = append(fnDecls, fnDecl)
			}
		}
	}

	return fnDecls
}

func match(fnDecl *ast.FuncDecl, matchers []MatcherFunc) bool {
	return len(matchers) == 0 || matchers[0].Match(fnDecl)
}

func ParseSrc(src string) {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "src.go", src, 0)

	file = f
}
