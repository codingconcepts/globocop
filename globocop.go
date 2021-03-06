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
			return true
		}

		// Bail if the identifier isn't at the global level (and recurs
		// into nested nodes).
		if len(stack) > 4 {
			return true
		}

		// Bail if it's not a variable (and don't recurs into nested nodes,
		// as we're outside of functions and types etc).
		if ident.Obj.Kind != ast.Var {
			return false
		}

		pass.Report(analysis.Diagnostic{
			Pos:     ident.Pos(),
			Message: fmt.Sprintf("global var %q", ident.Name),
		})

		return true
	})

	return nil, nil
}
