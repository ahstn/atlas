package util

import (
	"testing"
)

func TestProcessRepoURL(t *testing.T) {
	tests := []struct {
		name    string
		repo    string
		want    string
		wantErr bool
	}{
		{
			name:    "GitHub SSH URL",
			repo:    "git@github.com:ahstn/atlas",
			want:    "https://github.com/ahstn/atlas",
			wantErr: false,
		},
		{
			name:    "GitHub HTTPS URL",
			repo:    "https://github.com/ahstn/atlas",
			want:    "https://github.com/ahstn/atlas",
			wantErr: false,
		},
		{
			name:    "GitLab SSH URL",
			repo:    "git@gitlab.com:fdroid/fdroidclient.git",
			want:    "https://gitlab.com/fdroid/fdroidclient.git",
			wantErr: false,
		},
		{
			name:    "GitLab HTTPS URL",
			repo:    "https://gitlab.com/fdroid/fdroidclient.git",
			want:    "https://gitlab.com/fdroid/fdroidclient.git",
			wantErr: false,
		},
		{
			name:    "invalid URL",
			repo:    "tcp://ahstn/atlas",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ProcessRepoURL(tt.repo)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessRepoURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ProcessRepoURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
