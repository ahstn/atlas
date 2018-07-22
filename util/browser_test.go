package util

import "testing"

func TestOpenBrowser(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := OpenBrowser(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("OpenBrowser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
