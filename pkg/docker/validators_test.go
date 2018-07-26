package docker

import (
	"reflect"
	"testing"
)

func TestValidateTag(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Happy Path",
			args:    args{s: "ahstn:test"},
			want:    "ahstn:test",
			wantErr: false,
		},
		{
			name:    "No Version",
			args:    args{s: "ahstn"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Incorrectly Starts with ':'",
			args:    args{s: ":ahstn"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Incorrectly Ends with '/'",
			args:    args{s: "ahstn/"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateTag(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateTag() = %v, want %v", got, tt.want)
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
		want    []string
		wantErr bool
	}{
		{
			name:    "Happy Path",
			args:    args{s: []string{"VERSION=1.8.0", "LANG=GO"}},
			want:    []string{"VERSION=1.8.0", "LANG=GO"},
			wantErr: false,
		},
		{
			name:    "Missing Value (without equal sign)",
			args:    args{s: []string{"VERSION"}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Missing Value (with equal sign)",
			args:    args{s: []string{"VERSION="}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Multiple Values",
			args:    args{s: []string{"VERSION=1.8.0=1.7.0"}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateBuildArgs(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateBuildArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateBuildArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
