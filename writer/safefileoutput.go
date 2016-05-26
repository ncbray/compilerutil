package writer

import (
	"bytes"
	"crypto/sha1"
	"io"
	"io/ioutil"
	"os"
)

func hashFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	hasher := sha1.New()
	_, err = io.Copy(hasher, f)
	if err != nil {
		return nil, err
	}
	return hasher.Sum(nil), nil
}

func fileHaveSameContent(path1 string, path2 string) bool {
	// Check size, first.
	i1, err := os.Stat(path1)
	if err != nil {
		return false
	}
	i2, err := os.Stat(path1)
	if err != nil {
		return false
	}
	if i1.Size() != i2.Size() {
		return false
	}

	// Check contents.
	h1, err := hashFile(path1)
	if err != nil {
		return false
	}
	h2, err := hashFile(path2)
	if err != nil {
		return false
	}
	return bytes.Equal(h1, h2)
}

func moveFile(src string, dst string) error {
	// TODO copy and remove src if rename fails.
	return os.Rename(src, dst)
}

type fileOutput struct {
	Tmp  string
	Dst  string
	Perm os.FileMode
}

type SafeFileOutput struct {
	files   []fileOutput
	tempDir string
}

func (o *SafeFileOutput) OutputFile(path string, perm os.FileMode) (*os.File, error) {
	if o.tempDir == "" {
		panic("uninitialized")
	}
	temp_file, err := ioutil.TempFile(o.tempDir, "out_")
	if err != nil {
		return nil, err
	}
	o.files = append(o.files, fileOutput{Tmp: temp_file.Name(), Dst: path, Perm: perm})
	return temp_file, err
}

func (o *SafeFileOutput) Commit() error {
	if o.tempDir == "" {
		panic("uninitialized")
	}
	for _, fout := range o.files {
		if !fileHaveSameContent(fout.Tmp, fout.Dst) {
			err := moveFile(fout.Tmp, fout.Dst)
			if err != nil {
				return nil
			}
		}
		err := os.Chmod(fout.Dst, fout.Perm)
		if err != nil {
			return err
		}
	}
	o.files = nil
	return nil
}

func (o *SafeFileOutput) Cleanup() (err error) {
	if o.tempDir != "" {
		err = os.RemoveAll(o.tempDir)
	}
	o.tempDir = ""
	return
}

func MakeSafeFileOutput() (*SafeFileOutput, error) {
	temp_dir, err := ioutil.TempDir("", "compiler_outputs_")
	if err != nil {
		return nil, err
	}
	return &SafeFileOutput{tempDir: temp_dir}, nil
}
