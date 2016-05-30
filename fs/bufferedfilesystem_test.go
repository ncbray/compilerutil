package fs

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func writeStringToFileAndClose(t *testing.T, data string, f DataOutput) {
	err := f.SetBytes([]byte(data))
	assert.Equal(t, nil, err)
}

func assertFileContains(t *testing.T, path string, expected_data string) {
	actual_data, err := ioutil.ReadFile(path)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected_data, string(actual_data))
}

func TestUncommited(t *testing.T) {
	temp, err := MakeTempDir("test_")
	assert.Equal(t, nil, err)
	defer func() { assert.Equal(t, nil, temp.Cleanup()) }()
	fsys := MakeBufferedFileSystem(temp)

	fsys.OutputFile("fofsys.txt", 0640)
}

func TestCommitNew(t *testing.T) {
	temp, err := MakeTempDir("test_")
	assert.Equal(t, nil, err)
	defer func() { assert.Equal(t, nil, temp.Cleanup()) }()
	fsys := MakeBufferedFileSystem(temp)

	out_file := filepath.Join(temp.Path(), "fofsys.txt")
	payload := "bar"

	f := fsys.OutputFile(out_file, 0640)
	writeStringToFileAndClose(t, payload, f)

	// Not there yet.
	_, err = ioutil.ReadFile(out_file)
	assert.NotEqual(t, nil, err)

	assert.Equal(t, nil, fsys.Commit())

	// Now it's there.
	assertFileContains(t, out_file, payload)
	new_info, err := os.Stat(out_file)
	assert.Equal(t, nil, err)
	assert.Equal(t, os.FileMode(0640), new_info.Mode())
}

func TestCommitExistingDifferent(t *testing.T) {
	temp, err := MakeTempDir("test_")
	assert.Equal(t, nil, err)
	defer func() { assert.Equal(t, nil, temp.Cleanup()) }()
	fsys := MakeBufferedFileSystem(temp)

	out_file := filepath.Join(temp.Path(), "fofsys.txt")
	old_payload := "baz"
	ioutil.WriteFile(out_file, []byte(old_payload), 0600)
	old_info, err := os.Stat(out_file)
	assert.Equal(t, nil, err)

	payload := "bar"

	f := fsys.OutputFile(out_file, 0440)
	writeStringToFileAndClose(t, payload, f)

	// Old file still in place.
	assertFileContains(t, out_file, old_payload)

	assert.Equal(t, nil, fsys.Commit())

	// Now it's there.
	assertFileContains(t, out_file, payload)

	new_info, err := os.Stat(out_file)
	assert.Equal(t, nil, err)
	assert.Equal(t, os.FileMode(0440), new_info.Mode())
	assert.True(t, !os.SameFile(old_info, new_info))
	assert.Equal(t, old_info.ModTime(), new_info.ModTime())
}

func TestCommitExistingSame(t *testing.T) {
	temp, err := MakeTempDir("test_")
	assert.Equal(t, nil, err)
	defer func() { assert.Equal(t, nil, temp.Cleanup()) }()
	fsys := MakeBufferedFileSystem(temp)

	out_file := filepath.Join(temp.Path(), "fofsys.txt")
	payload := "bar"
	ioutil.WriteFile(out_file, []byte(payload), 0600)
	old_info, err := os.Stat(out_file)
	assert.Equal(t, nil, err)

	f := fsys.OutputFile(out_file, 0640)
	writeStringToFileAndClose(t, payload, f)

	assert.Equal(t, nil, fsys.Commit())

	// Now it's there.
	new_info, err := os.Stat(out_file)
	assert.Equal(t, nil, err)
	assert.Equal(t, os.FileMode(0640), new_info.Mode())
	assert.True(t, os.SameFile(old_info, new_info))
	assert.Equal(t, old_info.ModTime(), new_info.ModTime())
}
