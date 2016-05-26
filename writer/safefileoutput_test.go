package writer

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func writeStringToFileAndClose(t *testing.T, data string, f *os.File) {
	n, err := f.WriteString(data)
	assert.Equal(t, len(data), n)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, f.Close())
}

func assertFileContains(t *testing.T, path string, expected_data string) {
	actual_data, err := ioutil.ReadFile(path)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected_data, string(actual_data))
}

func TestUncommited(t *testing.T) {
	o, err := MakeSafeFileOutput()
	assert.Equal(t, nil, err)

	defer func() { assert.Equal(t, nil, o.Cleanup()) }()

	f, err := o.OutputFile("foo.txt", 0640)
	assert.Equal(t, nil, err)
	defer f.Close()
}

func TestCommitNew(t *testing.T) {
	temp_dir, err := ioutil.TempDir("", "test_output_")
	assert.Equal(t, nil, err)
	defer os.RemoveAll(temp_dir)

	o, err := MakeSafeFileOutput()
	assert.Equal(t, nil, err)
	defer func() { assert.Equal(t, nil, o.Cleanup()) }()

	out_file := filepath.Join(temp_dir, "foo.txt")
	payload := "bar"

	f, err := o.OutputFile(out_file, 0640)
	assert.Equal(t, nil, err)
	writeStringToFileAndClose(t, payload, f)

	// Not there yet.
	_, err = ioutil.ReadFile(out_file)
	assert.NotEqual(t, nil, err)

	assert.Equal(t, nil, o.Commit())

	// Now it's there.
	assertFileContains(t, out_file, payload)
	new_info, err := os.Stat(out_file)
	assert.Equal(t, nil, err)
	assert.Equal(t, os.FileMode(0640), new_info.Mode())
}

func TestCommitExistingDifferent(t *testing.T) {
	temp_dir, err := ioutil.TempDir("", "test_output_")
	assert.Equal(t, nil, err)
	defer os.RemoveAll(temp_dir)

	o, err := MakeSafeFileOutput()
	assert.Equal(t, nil, err)
	defer func() { assert.Equal(t, nil, o.Cleanup()) }()

	out_file := filepath.Join(temp_dir, "foo.txt")
	old_payload := "baz"
	ioutil.WriteFile(out_file, []byte(old_payload), 0600)
	old_info, err := os.Stat(out_file)
	assert.Equal(t, nil, err)

	payload := "bar"

	f, err := o.OutputFile(out_file, 0440)
	assert.Equal(t, nil, err)
	tmp_info, err := f.Stat()
	assert.Equal(t, nil, err)
	writeStringToFileAndClose(t, payload, f)

	// Old file still in place.
	assertFileContains(t, out_file, old_payload)

	assert.Equal(t, nil, o.Commit())

	// Now it's there.
	assertFileContains(t, out_file, payload)

	new_info, err := os.Stat(out_file)
	assert.Equal(t, nil, err)
	assert.Equal(t, os.FileMode(0440), new_info.Mode())
	assert.True(t, os.SameFile(tmp_info, new_info))
	assert.True(t, !os.SameFile(old_info, new_info))
	assert.Equal(t, old_info.ModTime(), new_info.ModTime())
}

func TestCommitExistingSame(t *testing.T) {
	temp_dir, err := ioutil.TempDir("", "test_output_")
	assert.Equal(t, nil, err)
	defer os.RemoveAll(temp_dir)

	o, err := MakeSafeFileOutput()
	assert.Equal(t, nil, err)
	defer func() { assert.Equal(t, nil, o.Cleanup()) }()

	out_file := filepath.Join(temp_dir, "foo.txt")
	payload := "bar"
	ioutil.WriteFile(out_file, []byte(payload), 0600)
	old_info, err := os.Stat(out_file)
	assert.Equal(t, nil, err)

	f, err := o.OutputFile(out_file, 0640)
	assert.Equal(t, nil, err)
	tmp_info, err := f.Stat()
	assert.Equal(t, nil, err)
	writeStringToFileAndClose(t, payload, f)

	assert.Equal(t, nil, o.Commit())

	// Now it's there.
	new_info, err := os.Stat(out_file)
	assert.Equal(t, nil, err)
	assert.Equal(t, os.FileMode(0640), new_info.Mode())
	assert.True(t, !os.SameFile(tmp_info, new_info))
	assert.True(t, os.SameFile(old_info, new_info))
	assert.Equal(t, old_info.ModTime(), new_info.ModTime())
}
