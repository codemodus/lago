// Package lago provides a simple way to setup logging.
package lago

// Logger ...
type Logger interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
}

// NullLogger ...
type NullLogger struct{}

// NewNullLogger ...
func NewNullLogger() *NullLogger {
	return &NullLogger{}
}

// Error ...
func (l *NullLogger) Error(args ...interface{}) {}

// Errorf ...
func (l *NullLogger) Errorf(format string, args ...interface{}) {}

// Info ...
func (l *NullLogger) Info(args ...interface{}) {}

// Infof ...
func (l *NullLogger) Infof(format string, args ...interface{}) {}

// Warn ...
func (l *NullLogger) Warn(args ...interface{}) {}

// Warnf ...
func (l *NullLogger) Warnf(format string, args ...interface{}) {}
