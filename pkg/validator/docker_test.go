package validator

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"
)

func TestValidateTag(t *testing.T) {
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
			args:    args{s: "ahstn:test"},
			wantErr: false,
		},
		{
			name:    "No Version",
			args:    args{s: "ahstn"},
			wantErr: true,
		},
		{
			name:    "Incorrectly Starts with ':'",
			args:    args{s: ":ahstn"},
			wantErr: true,
		},
		{
			name:    "Incorrectly Ends with '/'",
			args:    args{s: "ahstn/"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTag(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestValidateBuildArgs(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Happy Path",
			args:    args{s: []string{"VERSION=1.8.0", "LANG=GO"}},
			wantErr: false,
		},
		{
			name:    "Missing Value (without equal sign)",
			args:    args{s: []string{"VERSION"}},
			wantErr: true,
		},
		{
			name:    "Missing Value (with equal sign)",
			args:    args{s: []string{"VERSION="}},
			wantErr: true,
		},
		{
			name:    "Multiple Values",
			args:    args{s: []string{"VERSION=1.8.0=1.7.0"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateBuildArgs(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateBuildArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTryFindDockerfile(t *testing.T) {
	err := os.Mkdir("./testdir", 0777)
	if err != nil {
		t.Fatal("Can't create test dir, skipping TestTryFindDockerfile", err)
	}

	dockerfile := []byte("FROM golang\nRUN echo 'hello'\n")
	err = ioutil.WriteFile("./testdir/Dockerfile", dockerfile, 0644)
	if err != nil {
		_ = os.RemoveAll("./testdir")
		t.Fatal("Can't create Dockerfile, skipping TestTryFindDockerfile", err)
	}

	s, err := TryFindDockerfile(".")
	if s != "testdir/Dockerfile" {
		_ = os.RemoveAll("./testdir")
		t.Fatal("Expected FindDockerfile to return 'testdir/Dockerfile'. Got:", s)
	}
	if err != nil {
		_ = os.RemoveAll("./testdir")
		t.Fatal("Expected FindDockerfile to be successful. Got:", err)
	}

	err = os.RemoveAll("./testdir")
	if err != nil {
		t.Skip("Can't Remove Dockerfile, skpping TestTryFindDockerfile")
	}

	s, err = TryFindDockerfile(".")
	if err != errors.New(errDockerfile) {
		t.Fatal("Expected FindDockerfile to be unsuccessful. Got:", s)
	}
}
