package lago_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/codemodus/lago"
)

var (
	ts = "test"
)

func TestIntegNewFuncNilOpts(t *testing.T) {
	l := lago.New(nil)
	if l == nil {
		t.Errorf("don't want nil, got %T", l)
	}

	l.Println("test")
}

func TestIntegNewFuncEmptyOpts(t *testing.T) {
	l := lago.New(&lago.Options{})
	if l == nil {
		t.Errorf("don't want nil, got %T", l)
	}
}

func TestIntegNewFuncCompleteOpts(t *testing.T) {
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

	l := lago.New(&lago.Options{
		Filepath: f.Name(),
	})

	if l == nil {
		t.Fatalf("don't want nil, got %T", l)
	}

	l.Printf(ts)

	tf, err := os.Open(f.Name())
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(tf)
	if err != nil {
		t.Fatal(err)
	}

	if string(b)[:len(b)-1] != ts {
		t.Errorf("want %s, got %s", ts, string(b)[:len(b)-1])
	}
}
