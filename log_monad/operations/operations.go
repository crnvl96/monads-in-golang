package operations

import (
	"fmt"

	"github.com/crnvl96/monads-in-golang/log_monad/monads"
)

// These must return the monadic WrapperType
func Square(i int) monads.NumberWithLogs {
	return monads.NumberWithLogs{
		Result: i * i,
		Logs:   []string{fmt.Sprintf("Squared %d to get %d", i, i*i)},
	}
}

func AddOne(i int) monads.NumberWithLogs {
	return monads.NumberWithLogs{
		Result: i + 1,
		Logs:   []string{fmt.Sprintf("Added one to %d to get %d", i, i+1)},
	}
}

func MultiplyByThree(i int) monads.NumberWithLogs {
	return monads.NumberWithLogs{
		Result: i * 3,
		Logs:   []string{fmt.Sprintf("Multiplied three to %d to get %d", i, i*3)},
	}
}
