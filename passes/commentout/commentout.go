package commentout

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/ast/astutil"
)

const doc = "commentout finds a commented out debug code without reason"

var Analyzer = &analysis.Analyzer{
	Name: "commentout",
	Doc:  doc,
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		cmap := ast.NewCommentMap(pass.Fset, file, file.Comments)
		for _, cg := range file.Comments {

			path, _ := astutil.PathEnclosingInterval(file, cg.Pos(), cg.Pos())
			if len(path) != 1 || isMapped(pass.Fset, cmap, cg) {
				continue
			}

			if ok, n := isDebugComment(cg); ok && n == len(cg.List) {
				pass.Reportf(cg.Pos(), "do not leave a commented out debug code without reason")
			}
		}
	}
	return nil, nil
}

func isMapped(fset *token.FileSet, cmap ast.CommentMap, cg *ast.CommentGroup) bool {
	for n, cgs := range cmap {
		if !isSameLine(fset, cg.Pos(), n.Pos()) {
			continue
		}
		for i := range cgs {
			if cg == cgs[i] {
				return true
			}
		}
	}
	return false
}

func isSameLine(fset *token.FileSet, x, y token.Pos) bool {
	return x != token.NoPos && y != token.NoPos &&
		fset.Position(x).Line == fset.Position(y).Line
}

func isDebugComment(cg *ast.CommentGroup) (bool, int) {
	if cg == nil || len(cg.List) <= 1 || strings.TrimSpace(cg.Text()) == "" {
		return false, 0
	}

	if isDebugCommentInTop(cg) {
		return true, len(cg.List)
	}

	return isDebugComment(&ast.CommentGroup{List: cg.List[1:]})
}

func isDebugCommentInTop(cg *ast.CommentGroup) bool {
	src := fmt.Sprintf("package a\n%s", cg.Text())
	_, err := parser.ParseFile(token.NewFileSet(), "a.go", src, 0)
	return err == nil
}
