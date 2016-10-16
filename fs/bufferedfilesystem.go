package fs

import (
	"bytes"
	"crypto/sha1"
	"io"
	"os"
	"path/filepath"
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
	dir := filepath.Dir(dst)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}
	// TODO copy and remove src if rename fails.
	return os.Rename(src, dst)
}

type fileOutput struct {
	Tmp  string
	Dst  string
	Perm os.FileMode
}

type bufferedFileSystem struct {
	files []fileOutput
	temp  *TempDir
}

func (o *bufferedFileSystem) InputFile(path string) DataInput {
	return &DataFile{Path: path}
}

func (o *bufferedFileSystem) OutputFile(path string, perm os.FileMode) DataOutput {
	temp_file, err := o.temp.TempFile("out_")
	if err != nil {
		return nil
	}
	name := temp_file.Name()
	o.files = append(o.files, fileOutput{Tmp: name, Dst: path, Perm: perm})
	return &DataFile{Path: name}
}

func (o *bufferedFileSystem) TempFile() Data {
	return &DataBuffer{}
}

func (o *bufferedFileSystem) Commit() error {
	for _, fout := range o.files {
		if !fileHaveSameContent(fout.Tmp, fout.Dst) {
			err := moveFile(fout.Tmp, fout.Dst)
			if err != nil {
				return err
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

func MakeBufferedFileSystem(temp *TempDir) BufferedFileSystem {
	return &bufferedFileSystem{temp: temp}
}
