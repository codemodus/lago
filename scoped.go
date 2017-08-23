package lago

import (
	"fmt"
	"strconv"
)

// ScopedLogger ...
type ScopedLogger struct {
	log Logger
	scp string
	fmt string
}

// NewScopedLogger ...
func NewScopedLogger(log Logger, scope string, padding int) *ScopedLogger {
	format := "-%" + intToStringWithSign(padding) + "s: %s"

	l := &ScopedLogger{
		log: log,
		scp: scope,
		fmt: format,
	}

	return l
}

// Errorf ...
func (l *ScopedLogger) Errorf(format string, args ...interface{}) {
	l.log.Errorf(l.fmt, l.scp, fmt.Sprintf(format, args...))
}

// Infof ...
func (l *ScopedLogger) Infof(format string, args ...interface{}) {
	l.log.Infof(l.fmt, l.scp, fmt.Sprintf(format, args...))
}

// Warnf ...
func (l *ScopedLogger) Warnf(format string, args ...interface{}) {
	l.log.Warnf(l.fmt, l.scp, fmt.Sprintf(format, args...))
}

// Fatalf ...
func (l *ScopedLogger) Fatalf(format string, args ...interface{}) {
	l.log.Fatalf(l.fmt, l.scp, fmt.Sprintf(format, args...))
}

func intToStringWithSign(i int) string {
	a := strconv.Itoa(i)
	s := "+"
	if i < 0 {
		s = "-"
	}
	a = s + a

	return a
}
