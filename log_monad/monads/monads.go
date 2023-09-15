package monads

// Wrapper Type => The contract type between the monad ecosystem and the external ecosystem
type NumberWithLogs struct {
	Logs   []string
	Result int
}

// Wrapper Function =>The entry to the monad ecosystem
func WrapNumberWithLogs(i int) NumberWithLogs {
	return NumberWithLogs{
		Result: i,
		Logs:   []string{},
	}
}

// Run Function (Transformer) => Runs transformations into monadic values
func RunNumberWithLogs(i NumberWithLogs, transform func(int) NumberWithLogs) NumberWithLogs {
	factory := transform(i.Result)

	return NumberWithLogs{
		Result: factory.Result,
		Logs:   append(i.Logs, factory.Logs...),
	}
}
