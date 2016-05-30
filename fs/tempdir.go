package fs

import (
	"io/ioutil"
	"os"
)

type TempDir struct {
	path string
}

func (t *TempDir) Path() string {
	return t.path
}

func (t *TempDir) TempFile(prefix string) (*os.File, error) {
	return ioutil.TempFile(t.path, prefix)
}

func (t *TempDir) Cleanup() (err error) {
	if t.path != "" {
		err = os.RemoveAll(t.path)
	}
	t.path = ""
	return
}

func MakeTempDir(prefix string) (*TempDir, error) {
	path, err := ioutil.TempDir("", prefix)
	if err != nil {
		return nil, err
	}
	return &TempDir{path: path}, nil
}
