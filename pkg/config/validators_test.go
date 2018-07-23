package config

import (
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

// TODO : Assert Error message as well
func Test_validateDockerTag(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Happy Path",
			args:    args{t: "atlas:0.1.0"},
			wantErr: false,
		},
		{
			name:    "Missing Version",
			args:    args{t: "atlas"},
			wantErr: true,
		},
		{
			name:    "Invalid Format",
			args:    args{t: "atlas:"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateDockerTag(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("validateDockerTag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
