package lago

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestNewLJL(t *testing.T) {
	l := newLumberjackLogger(nil)
	if l != nil {
		t.Errorf("want nil, got %T", l)
	}

	l = newLumberjackLogger(&Options{})
	if l != nil {
		t.Errorf("want nil, got %T", l)
	}

	f, err := ioutil.TempFile("./", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = f.Close()

		err := os.Remove(f.Name())
		if err != nil {
			fmt.Printf("WARN: tmp file not removed: %s\n", err)
		}
	}()

	l = newLumberjackLogger(&Options{
		Filepath: f.Name(),
	})

	if l == nil {
		t.Errorf("don't want nil, got %T", l)
	}
}

func TestUnitJoinWriters(t *testing.T) {
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
	l := NewDevNull()
	if l == nil {
		t.Errorf("don't want nil, got %T", l)
	}
}

func TestUnitNewFunc(t *testing.T) {
	l := New(nil)
	if l == nil {
		t.Errorf("don't want nil, got %T", l)
	}

	l = New(&Options{})
	if l == nil {
		t.Errorf("don't want nil, got %T", l)
	}

	f, err := ioutil.TempFile("./", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = f.Close()

		err := os.Remove(f.Name())
		if err != nil {
			fmt.Printf("WARN: tmp file not removed: %s\n", err)
		}
	}()

	l = New(&Options{
		Filepath: f.Name(),
	})

	if l == nil {
		t.Errorf("don't want nil, got %T", l)
	}
}
