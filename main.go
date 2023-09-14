package main

import (
	"fmt"

	"github.com/crnvl96/monads-in-golang/monads"
	"github.com/crnvl96/monads-in-golang/operations"
)

func main() {
	firstOp := operations.AddOne(2)
	secondOp := monads.RunWithLogs(firstOp, operations.Square)
	result := monads.RunWithLogs(secondOp, operations.MultiplyByThree)

	fmt.Println(result)
}
