package debugcode

import (
	"github.com/gostaticanalysis/builtinprint"
	"golang.org/x/tools/go/analysis"
)

// Analyzers returns all analyzers of debugcode.
func Analyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		builtinprint.Analyzer,
	}
}
