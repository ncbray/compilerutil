package fs

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
)

type DataInput interface {
	GetReader() (io.ReadCloser, error)
	GetBytes() ([]byte, error)
}

type DataOutput interface {
	GetWriter() (io.WriteCloser, error)
	SetBytes(data []byte) error
}

type Data interface {
	DataInput
	DataOutput
}

type DataFile struct {
	Path string
}

func (f *DataFile) GetReader() (io.ReadCloser, error) {
	return os.OpenFile(f.Path, os.O_RDONLY, 0640)
}

func (f *DataFile) GetBytes() ([]byte, error) {
	return ioutil.ReadFile(f.Path)
}

func (f *DataFile) GetWriter() (io.WriteCloser, error) {
	return os.OpenFile(f.Path, os.O_WRONLY, 0640)
}

func (f *DataFile) SetBytes(data []byte) error {
	return ioutil.WriteFile(f.Path, data, 0640)
}

type DataBuffer struct {
	Buffer bytes.Buffer
}

func (d *DataBuffer) Read(p []byte) (n int, err error) {
	return d.Buffer.Read(p)
}

func (d *DataBuffer) Write(p []byte) (n int, err error) {
	return d.Buffer.Write(p)
}

func (d *DataBuffer) Close() error {
	return nil
}

func (d *DataBuffer) GetReader() (io.ReadCloser, error) {
	return d, nil
}

func (d *DataBuffer) GetBytes() ([]byte, error) {
	return d.Buffer.Bytes(), nil
}

func (d *DataBuffer) GetWriter() (io.WriteCloser, error) {
	return d, nil
}

func (d *DataBuffer) SetBytes(data []byte) error {
	d.Buffer.Reset()
	_, err := d.Buffer.Write(data)
	return err
}
