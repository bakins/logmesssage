package main

import (
	"github.com/bakins/logmessage/internal/logmessage"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(logmessage.Analyzer)
}