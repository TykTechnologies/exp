package example

type Logger interface {
	Info(msg string)
}

type S struct {
}

func (S) Logger() Logger {
	return nil
}

func (s S) WorkBad() bool {
	// ruleid: find-repeated-logger-calls
	s.Logger().Info("One")
	result := s.Work()
	s.Logger().Info("Two")

	return result
}

func (s S) WorkGood() bool {
	logger := s.Logger()

	logger.Info("One")
	result := s.Work()
	logger.Info("Two")

	return result
}

func (s S) Work() bool {
	return true
}
