package validator

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

func TestValidateConfig(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Happy Path",
			args:    args{s: "atlas.yaml"},
			wantErr: false,
		},
		{
			name:    "Invalid Format",
			args:    args{s: "invalid"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateConfig(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("ValidateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExists(t *testing.T) {
	oldPath := os.Getenv("PWD")
	defer os.Setenv("PWD", oldPath)

	const testFile = "atlas.yaml"
	tmp, err := ioutil.TempDir("", "TestValidateArguments")
	if err != nil {
		t.Fatal("TempDir failed: ", err)
	}
	os.Chdir(tmp)
	os.Setenv("PWD", tmp)

	f, err := os.OpenFile(testFile, os.O_CREATE, 0777)
	if err != nil {
		t.Fatal("OpenFile failed: ", err)
	}
	err = f.Close()
	if err != nil {
		t.Fatal("Close failed: ", err)
	}

	p, err := ValidateExists(testFile)
	if p != path.Join(tmp, testFile) {
		t.Fatal("Expected return to be dir. Got", p, err)
	}
}

func TestExistsReturnsErr(t *testing.T) {
	_, err := ValidateExists("should_not_exist.yaml")
	if err.Error() != errCfgMissing {
		t.Fatal("Expected return to be err. Got", err)
	}
}

func TestValidateConfigBaseDir(t *testing.T) {
	home := os.Getenv("HOME")
	tmp, err := ioutil.TempDir(home, "TestConfigDir")
	p, err := ValidateConfigBaseDir(strings.Replace(tmp, home, "~", 1))

	if p != tmp {
		t.Fatal("Expected return to be temp dir path. Got: ", p, err)
	}

	p, err = ValidateConfigBaseDir("should_not_exist")
	if err == nil {
		t.Fatal("Expected return to be err. Got: ", p, err)
	}
}
