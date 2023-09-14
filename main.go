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
	numberWithLogsFactory := numberWithLogs{
		result: i.result * i.result,
		logs:   []string{fmt.Sprintf("Squared %d to get %d", i.result, i.result*i.result)},
	}

	return numberWithLogs{
		result: numberWithLogsFactory.result,
		logs:   append(i.logs, numberWithLogsFactory.logs...),
	}
}

func addOne(i numberWithLogs) numberWithLogs {
	numberWithLogsFactory := numberWithLogs{
		result: i.result + 1,
		logs:   []string{fmt.Sprintf("Added 1 to %d to get %d", i.result, i.result+1)},
	}

	return numberWithLogs{
		result: numberWithLogsFactory.result,
		logs:   append(i.logs, numberWithLogsFactory.logs...),
	}
}
