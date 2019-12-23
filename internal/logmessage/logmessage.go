package logmessage

import (
	"go/ast"
	"go/token"
	"go/types"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyzer ...
var Analyzer = &analysis.Analyzer{
	Name:     "log",
	Doc:      "reports issues about log messages",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

// based on https://github.com/golang/lint/blob/master/lint.go
// and https://arslan.io/2019/06/13/using-go-analysis-to-write-a-custom-linter/

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		ce := n.(*ast.CallExpr)
		if len(ce.Args) < 1 {
			return
		}

		if !isFunc(ce.Fun, "Debug", "Error", "Fatal", "Info", "Panic", "Warn") {
			return
		}

		// Check if first argument is a string
		str, ok := ce.Args[0].(*ast.BasicLit)
		if !ok || str.Kind != token.STRING {
			return
		}

		t := pass.TypesInfo.TypeOf(ce.Fun)
		s, ok := t.(*types.Signature)
		if !ok {
			return
		}

		// I got deep into signature tuples and gave up for now and used String
		// Check if function accepts zapcore.Field as arguments
		if !strings.Contains(s.String(), "go.uber.org/zap/zapcore.Field") {
			return
		}

		st, _ := strconv.Unquote(str.Value)
		if st == "" {
			return
		}

		first, _ := utf8.DecodeRuneInString(st)
		if unicode.IsUpper(first) {
			pass.Reportf(ce.Lparen, "log messages should not be capitalized")
		}
	})
	return nil, nil
}

// based on https://github.com/golang/lint/blob/master/lint.go

func isFunc(expr ast.Expr, names ...string) bool {
	sel, ok := expr.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	found := false
	for _, n := range names {
		if isIdent(sel.Sel, n) {
			found = true
			break
		}
	}
	return found
}

func isIdent(expr ast.Expr, ident string) bool {
	id, ok := expr.(*ast.Ident)
	return ok && id.Name == ident
}
