package fs

import (
	"os"
	"path/filepath"
)

type FileSystem interface {
	InputFile(path string) DataInput
	OutputFile(path string, mode os.FileMode) DataOutput
	TempFile() Data
}

type BufferedFileSystem interface {
	FileSystem
	Commit() error
}

type relativeFileSystem struct {
	Parent           FileSystem
	WorkingDirectory string
}

func (fs *relativeFileSystem) InputFile(path string) DataInput {
	return fs.Parent.InputFile(filepath.Join(fs.WorkingDirectory, path))
}

func (fs *relativeFileSystem) OutputFile(path string, mode os.FileMode) DataOutput {
	return fs.Parent.OutputFile(filepath.Join(fs.WorkingDirectory, path), mode)
}

func (fs *relativeFileSystem) TempFile() Data {
	return fs.Parent.TempFile()
}

func MakeRelative(parent FileSystem, path string) FileSystem {
	return &relativeFileSystem{Parent: parent, WorkingDirectory: path}
}
