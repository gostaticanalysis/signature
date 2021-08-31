package main

import (
	"github.com/gostaticanalysis/signature"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(signature.Analyzer) }
