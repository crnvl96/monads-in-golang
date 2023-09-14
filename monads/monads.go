package monads

type NumberWithLogs struct {
	Logs   []string
	Result int
}

func RunWithLogs(i NumberWithLogs, transform func(int) NumberWithLogs) NumberWithLogs {
	factory := transform(i.Result)

	return NumberWithLogs{
		Result: factory.Result,
		Logs:   append(i.Logs, factory.Logs...),
	}
}
