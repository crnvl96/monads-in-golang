package main

import "fmt"

type numberWithLogs struct {
	logs   []string
	result int
}

func main() {
	fmt.Println(addOne(square(wrapWithLogs(2))))
}

func wrapWithLogs(i int) numberWithLogs {
	return numberWithLogs{
		result: i,
		logs:   []string{},
	}
}

func square(i numberWithLogs) numberWithLogs {
	return numberWithLogs{
		result: i.result * i.result,
		logs:   append(i.logs, fmt.Sprintf("Squared %d to get %d", i.result, i.result*i.result)),
	}
}

func addOne(i numberWithLogs) numberWithLogs {
	return numberWithLogs{
		result: i.result + 1,
		logs:   append(i.logs, fmt.Sprintf("Added 1 to %d to get %d", i.result, i.result+1)),
	}
}
