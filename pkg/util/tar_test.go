package util

import (
	"archive/tar"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestCreateArchive(t *testing.T) {
	const testFile = "test.Dockerfile"
	tmp, err := ioutil.TempDir("", "TestCreateArchive")
	if err != nil {
		t.Fatal("TempDir failed: ", err)
	}
	err = os.Chdir(tmp)
	if err != nil {
		t.Fatal("Chdir failed: ", err)
	}

	f, err := os.OpenFile(testFile, os.O_CREATE, 0777)
	if err != nil {
		t.Fatal("OpenFile failed: ", err)
	}
	err = f.Close()
	if err != nil {
		t.Fatal("Close failed: ", err)
	}

	var buf bytes.Buffer
	err = CreateArchive(tmp, &buf)
	if err != nil {
		t.Fatal("Create Archive failed: ", err)
	}

	tr := tar.NewReader(&buf)
	for {
		header, err := tr.Next()
		switch {
		case err == io.EOF:
			return
		case err != nil:
			t.Fatal("Read Archive failed: ", err)
		case header == nil:
			continue
		}
		if header.Name != path.Join(tmp, testFile) {
			t.Fatal("test:", header.Name, err)
		}
	}
}
