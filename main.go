package main

import "fmt"

type numberWithLogs struct {
	logs   []string
	result int
}

func runWithLogs(i numberWithLogs, transform func(int) numberWithLogs) numberWithLogs {
	factory := transform(i.result)

	return numberWithLogs{
		result: factory.result,
		logs:   append(i.logs, factory.logs...),
	}
}

func square(i int) numberWithLogs {
	return numberWithLogs{
		result: i * i,
		logs:   []string{fmt.Sprintf("Squared %d to get %d", i, i*i)},
	}
}

func addOne(i int) numberWithLogs {
	return numberWithLogs{
		result: i + 1,
		logs:   []string{fmt.Sprintf("Added one to %d to get %d", i, i+1)},
	}
}

func main() {
	fmt.Println(runWithLogs(addOne(2), addOne))
}
