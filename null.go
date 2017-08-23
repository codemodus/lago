package lago

import "os"

// NullLogger ...
type NullLogger struct{}

// NewNullLogger ...
func NewNullLogger() *NullLogger {
	return &NullLogger{}
}

// Errorf ...
func (l *NullLogger) Errorf(format string, args ...interface{}) {}

// Infof ...
func (l *NullLogger) Infof(format string, args ...interface{}) {}

// Warnf ...
func (l *NullLogger) Warnf(format string, args ...interface{}) {}

// Fatalf ...
func (l *NullLogger) Fatalf(format string, args ...interface{}) {
	os.Exit(1)
}
