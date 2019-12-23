package a

import (
	"go.uber.org/zap"
)

func test() {
	l, _ := zap.NewProduction()

	l.Error("testing 123")

	l.Error("Testing 123") //want `log messages should not be capitalized`

}
