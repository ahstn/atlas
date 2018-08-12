package config

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var (
	validYAML = []byte(`---
  root: /tmp
  services:
    -
      docker:
        dockerfile: ./Dockerfile
        enabled: false
      name: atlas
      repo: https://github.com/ahstn/atlas.git
      tasks:
        - build
  `)

	invalidYAML = []byte(`
  root- :tmp&
  `)

	invalidRootYAML = []byte(`---
  root: /ThisDirectoryShouldNotExist
  `)
)

func TestSuccessfulRead(t *testing.T) {
	err := ioutil.WriteFile("./atlas.yaml", validYAML, 0644)
	if err != nil {
		_ = os.RemoveAll("./atlas.yaml")
		t.Fatal("Can't create test yaml, skipping TestReadConfig", err)
	}

	p, err := Read("./atlas.yaml")
	if err != nil {
		t.Fatal("Expected config read to be successful. Got:", err)
	}
	if p.Root != "/tmp" {
		_ = os.RemoveAll("./atlas.yaml")
		t.Fatal("Expected Project root to be '/tmp'. Got:", p.Root)
	}

	_ = os.RemoveAll("./atlas/.yaml")
}

func TestReadMissingFile(t *testing.T) {
	_, err := Read("./missing.yaml")
	if !strings.Contains(err.Error(), "missing.yaml: no such file or directory") {
		t.Fatal("Expected Read to return error. Got:", err)
	}
}

func TestReadInvalidFile(t *testing.T) {
	err := ioutil.WriteFile("./invalid.yaml", invalidYAML, 0644)
	if err != nil {
		_ = os.RemoveAll("./invalid.yaml")
		t.Fatal("Can't create test yaml, skipping TestReadInvalidFile", err)
	}

	_, err = Read("./invalid.yaml")
	if !strings.Contains(err.Error(), "cannot unmarshal") {
		t.Fatal("Expected Read to return error (invalid file). Got:", err)
	}

	_ = os.RemoveAll("./invalid.yaml")
}

func TestReadInvalidRootField(t *testing.T) {
	err := ioutil.WriteFile("./invalid-root.yaml", invalidRootYAML, 0644)
	if err != nil {
		_ = os.RemoveAll("./invalid-root.yaml")
		t.Fatal("Can't create test yaml, skipping TestReadInvalidRootField", err)
	}

	_, err = Read("./invalid-root.yaml")
	if !strings.Contains(err.Error(), "ThisDirectoryShouldNotExist: no such file or directory") {
		t.Fatal("Expected Read to return error (invalid root field). Got:", err)
	}

	_ = os.RemoveAll("./invalid-root.yaml")
}

func TestService_HasTask(t *testing.T) {
	type fields struct {
		Docker  DockerArtifact
		Package Package
		Name    string
		Repo    string
		Tasks   []string
		Test    bool
	}
	type args struct {
		x string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Does have task",
			args: args{x: "build"},
			fields: fields{
				Tasks: []string{"clean", "build"},
			},
			want: true,
		},
		{
			name: "Does not have task",
			args: args{x: "package"},
			fields: fields{
				Tasks: []string{"clean", "build"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				Docker:  tt.fields.Docker,
				Package: tt.fields.Package,
				Name:    tt.fields.Name,
				Repo:    tt.fields.Repo,
				Tasks:   tt.fields.Tasks,
				Test:    tt.fields.Test,
			}
			if got := s.HasTask(tt.args.x); got != tt.want {
				t.Errorf("Service.HasTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_HasPackageSubDir(t *testing.T) {
	type fields struct {
		Docker  DockerArtifact
		Package Package
		Name    string
		Repo    string
		Tasks   []string
		Test    bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Does have package sub directory",
			fields: fields{
				Package: Package{SubDir: "package/"},
			},
			want: true,
		},
		{
			name: "Does not have package sub directory",
			fields: fields{
				Package: Package{SubDir: ""},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				Docker:  tt.fields.Docker,
				Package: tt.fields.Package,
				Name:    tt.fields.Name,
				Repo:    tt.fields.Repo,
				Tasks:   tt.fields.Tasks,
				Test:    tt.fields.Test,
			}
			if got := s.HasPackageSubDir(); got != tt.want {
				t.Errorf("Service.HasPackageSubDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
