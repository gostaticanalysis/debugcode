package main

import (
	"github.com/gostaticanalysis/builtinprint"
	"github.com/gostaticanalysis/debugcode/passes/commentout"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(
		builtinprint.Analyzer,
		commentout.Analyzer,
	)
}
