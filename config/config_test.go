package config

import (
	"testing"
)

func TestService_HasTask(t *testing.T) {
	type fields struct {
		Docker  Docker
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
		Docker  Docker
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
