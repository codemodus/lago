// Package lago provides a simple way to setup logging.
package lago

// Logger ...
type Logger interface {
	Errorf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}
