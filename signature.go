package signature

import (
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "signature finds low readability functions"

var Analyzer = &analysis.Analyzer{
	Name: "signature",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var (
	flagNumParams  int
	flagNumResults int
)

func init() {
	Analyzer.Flags.IntVar(&flagNumParams, "params", 5, "max number of params")
	Analyzer.Flags.IntVar(&flagNumResults, "results", 3, "max number of results")
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.FuncLit)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			typ := pass.TypesInfo.TypeOf(n.Name)
			checkSignature(pass, n.Pos(), typ)
		case *ast.FuncLit:
			typ := pass.TypesInfo.TypeOf(n)
			checkSignature(pass, n.Pos(), typ)
		}
	})

	return nil, nil
}

func checkSignature(pass *analysis.Pass, pos token.Pos, typ types.Type) {
	sig, _ := typ.(*types.Signature)
	if sig == nil {
		return
	}
	checkNumParams(pass, pos, sig)
	checkNumResults(pass, pos, sig)
}

func checkNumParams(pass *analysis.Pass, pos token.Pos, sig *types.Signature) {
	if params := sig.Params(); params != nil && params.Len() > flagNumParams {
		pass.Reportf(pos, "too many params")
	}
}

func checkNumResults(pass *analysis.Pass, pos token.Pos, sig *types.Signature) {
	if rets := sig.Results(); rets != nil && rets.Len() > flagNumResults {
		pass.Reportf(pos, "too many results")
	}
}
