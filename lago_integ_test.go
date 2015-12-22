package lago_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/codemodus/lago"
)

var (
	tStr0     = "Test string."
	tPre0     = "test_"
	tDirRoot0 = "./"
	tDir0     string
)

func SetupInteg() {
	var err error

	tDir0, err = ioutil.TempDir(tDirRoot0, tPre0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func TeardownInteg() {
	err := os.RemoveAll(tDir0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func TestIntegDiscard(t *testing.T) {
	SetupInteg()
	defer TeardownInteg()

	f, err := ioutil.TempFile(tDir0, tPre0)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = f.Close()
	}()

	stdout := os.Stdout
	os.Stdout = f
	defer func() {
		os.Stdout = stdout
	}()

	l := lago.New(nil)

	l.Println(tStr0)

	tfNil, err := os.Open(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = tfNil.Close()
	}()

	b, err := ioutil.ReadAll(tfNil)
	if err != nil {
		t.Fatal(err)
	}

	if len(b) > 0 {
		t.Errorf("want empty string, got %s", string(b)[:len(b)-1])
	}

	l = lago.New(&lago.Options{})

	l.Println(tStr0)

	tfEmpty, err := os.Open(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = tfEmpty.Close()
	}()

	b, err = ioutil.ReadAll(tfEmpty)
	if err != nil {
		t.Fatal(err)
	}

	if len(b) > 0 {
		t.Errorf("want empty string, got %s", string(b)[:len(b)-1])
	}
}

func TestIntegFilepathStdStreamWriter(t *testing.T) {
	SetupInteg()
	defer TeardownInteg()

	ff, err := ioutil.TempFile(tDir0, tPre0)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = ff.Close()
	}()

	sf, err := ioutil.TempFile(tDir0, tPre0)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = sf.Close()
	}()

	stdout := os.Stdout
	os.Stdout = sf
	defer func() {
		os.Stdout = stdout
	}()

	wf, err := ioutil.TempFile(tDir0, tPre0)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = wf.Close()
	}()

	l := lago.New(&lago.Options{
		Filepath:  ff.Name(),
		StdStream: lago.Stdout,
		LogWriter: wf,
	})

	l.Println(tStr0)

	tfF, err := os.Open(ff.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = tfF.Close()
	}()

	b, err := ioutil.ReadAll(tfF)
	if err != nil {
		t.Fatal(err)
	}

	if len(b) == 0 {
		t.Errorf("want %s, got empty string", tStr0)
	}

	if len(b) > 0 && string(b)[:len(b)-1] != tStr0 {
		t.Errorf("want %s, got %s", tStr0, string(b)[:len(b)-1])
	}

	tfS, err := os.Open(sf.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = tfS.Close()
	}()

	b, err = ioutil.ReadAll(tfS)
	if err != nil {
		t.Fatal(err)
	}

	if len(b) == 0 {
		t.Errorf("want %s, got empty string", tStr0)
	}

	if len(b) > 0 && string(b)[:len(b)-1] != tStr0 {
		t.Errorf("want %s, got %s", tStr0, string(b)[:len(b)-1])
	}

	tfW, err := os.Open(wf.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = tfW.Close()
	}()

	b, err = ioutil.ReadAll(tfW)
	if err != nil {
		t.Fatal(err)
	}

	if len(b) == 0 {
		t.Errorf("want %s, got empty string", tStr0)
	}

	if len(b) > 0 && string(b)[:len(b)-1] != tStr0 {
		t.Errorf("want %s, got %s", tStr0, string(b)[:len(b)-1])
	}
}
