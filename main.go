package main

import (
	"fmt"

	"github.com/crnvl96/monads-in-golang/log_monad/monads"
	"github.com/crnvl96/monads-in-golang/log_monad/operations"
)

// Log Monad
func main() {
	wrapped := monads.WrapNumberWithLogs(2)
	firstOp := monads.RunNumberWithLogs(wrapped, operations.AddOne)
	secondOp := monads.RunNumberWithLogs(firstOp, operations.Square)
	result := monads.RunNumberWithLogs(secondOp, operations.MultiplyByThree)
	fmt.Println(result)
}
