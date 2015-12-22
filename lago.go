// Package lago provides a simple way to setup logging.
package lago

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

type StdStreamer interface {
	Writer() io.Writer
}

type devNil struct{}

func (s *devNil) Writer() io.Writer {
	return nil
}

type devNull struct{}

func (s *devNull) Writer() io.Writer {
	return ioutil.Discard
}

type stdout struct{}

func (s *stdout) Writer() io.Writer {
	return os.Stdout
}

type stderr struct{}

func (s *stderr) Writer() io.Writer {
	return os.Stderr
}

var (
	DevNull StdStreamer = &devNull{}
	Stdout  StdStreamer = &stdout{}
	Stderr  StdStreamer = &stderr{}
)

type Options struct {
	Filepath   string
	StdStream  StdStreamer
	LogWriter  io.Writer
	WithTime   bool
	MaxSize    int
	MaxAge     int
	MaxBackups int
	LocalTime  bool
}

// Logger wraps a std lib logger.
type Logger struct {
	*log.Logger
}

func New(opts *Options) (l *Logger) {
	if opts == nil {
		return NewDevNull()
	}

	if opts.StdStream == nil {
		opts.StdStream = &devNil{}
	}

	f := newLumberjackLogger(opts)

	w := joinWriters(f, opts.StdStream.Writer())

	r := joinWriters(w, opts.LogWriter)

	if r == nil {
		return NewDevNull()
	}

	fs := log.LstdFlags
	if !opts.WithTime {
		fs = 0
	}

	return &Logger{
		log.New(r, "", fs),
	}
}

func NewDevNull() *Logger {
	return &Logger{
		log.New(ioutil.Discard, "", 0),
	}
}

func newLumberjackLogger(opts *Options) io.Writer {
	if opts == nil || opts.Filepath == "" {
		return nil
	}

	return &lumberjack.Logger{
		Filename:   opts.Filepath,
		MaxSize:    opts.MaxSize,
		MaxAge:     opts.MaxAge,
		MaxBackups: opts.MaxBackups,
		LocalTime:  opts.LocalTime,
	}
}

func joinWriters(a, b io.Writer) (c io.Writer) {
	if a != nil && b != nil {
		c = io.MultiWriter(a, b)
	}

	if c == nil {
		c = a
	}

	if c == nil {
		c = b
	}

	return c
}
