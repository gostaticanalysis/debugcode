package main

import (
	"github.com/gostaticanalysis/debugcode/passes/commentout"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(commentout.Analyzer) }
