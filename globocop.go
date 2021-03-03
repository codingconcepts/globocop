package main

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/analysis/singlechecker"
	"golang.org/x/tools/go/ast/inspector"
)

func main() {
	doc := `check for global variables

This checker reports global variables (i.e exported or unexported variables).`

	a := analysis.Analyzer{
		Name:     "globocop",
		Doc:      doc,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Run:      run,
	}

	singlechecker.Main(&a)
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}

	inspect.WithStack(nodeFilter, func(n ast.Node, push bool, stack []ast.Node) bool {
		ident := n.(*ast.Ident)

		// Bail if we don't have enough to figure out what this identifier is.
		if ident.Obj == nil || ident.Obj.Data == nil {
			return false
		}

		// Bail if the identifier isn't at the global level.
		if len(stack) > 4 {
			return false
		}

		pass.Report(analysis.Diagnostic{
			Pos:     ident.Pos(),
			Message: fmt.Sprintf("global %s %q", ident.Obj.Kind, ident.Name),
		})

		return false
	})

	return nil, nil
}
