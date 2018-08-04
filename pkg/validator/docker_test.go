package validator

import (
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
