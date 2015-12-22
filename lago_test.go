package lago

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

var (
	tStr0     = "Test string."
	tPre0     = "test_"
	tDirRoot0 = "./"
	tDir0     string
)

func SetupUnit() {
	var err error

	tDir0, err = ioutil.TempDir(tDirRoot0, tPre0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func TeardownUnit() {
	err := os.RemoveAll(tDir0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func TestUnitNewLJL(t *testing.T) {
	SetupUnit()
	defer TeardownUnit()

	l := newLumberjackLogger(nil)
	if l != nil {
		t.Errorf("want nil, got %T", l)
	}

	l = newLumberjackLogger(&Options{})
	if l != nil {
		t.Errorf("want nil, got %T", l)
	}

	f, err := ioutil.TempFile(tDir0, tPre0)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = f.Close()
	}()

	l = newLumberjackLogger(&Options{
		Filepath: f.Name(),
	})

	if l == nil {
		t.Errorf("don't want nil, got %T", l)
	}
}

func TestUnitJoinWriters(t *testing.T) {
	SetupUnit()
	defer TeardownUnit()

	var a, b io.Writer

	c := joinWriters(a, b)
	if c != nil {
		t.Errorf("want nil, got %T", c)
	}

	a = ioutil.Discard

	c = joinWriters(a, b)
	if c == nil {
		t.Errorf("don't want nil, got %T", c)
	}

	b = ioutil.Discard

	c = joinWriters(a, b)
	if c == nil {
		t.Errorf("don't want nil, got %T", c)
	}
}

func TestUnitDevNull(t *testing.T) {
	SetupUnit()
	defer TeardownUnit()

	l := NewDevNull()
	if l == nil {
		t.Errorf("don't want nil, got %T", l)
	}
}

func TestUnitNewFunc(t *testing.T) {
	SetupUnit()
	defer TeardownUnit()

	l := New(nil)
	if l == nil {
		t.Errorf("don't want nil, got %T", l)
	}

	l = New(&Options{})
	if l == nil {
		t.Errorf("don't want nil, got %T", l)
	}

	f, err := ioutil.TempFile(tDir0, tPre0)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = f.Close()
	}()

	l = New(&Options{
		Filepath: f.Name(),
	})

	if l == nil {
		t.Errorf("don't want nil, got %T", l)
	}

	l = New(&Options{
		StdStream: DevNull,
	})

	if l == nil {
		t.Errorf("don't want nil, got %T", l)
	}

	l = New(&Options{
		StdStream: Stdout,
	})

	if l == nil {
		t.Errorf("don't want nil, got %T", l)
	}

	l = New(&Options{
		StdStream: Stderr,
	})

	if l == nil {
		t.Errorf("don't want nil, got %T", l)
	}
}
