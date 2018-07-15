package cmd

import "testing"

func Test_processRepoURL(t *testing.T) {
	type args struct {
		r string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "GitHub SSH URL",
			args:    args{r: "git@github.com:ahstn/atlas.git"},
			want:    "https://github.com/ahstn/atlas.git",
			wantErr: false,
		},
		{
			name:    "GitHub HTTPS URL",
			args:    args{r: "https://github.com/ahstn/atlas.git"},
			want:    "https://github.com/ahstn/atlas.git",
			wantErr: false,
		},
		{
			name:    "GitLab SSH URL",
			args:    args{r: "git@gitlab.com:fdroid/fdroidclient.git"},
			want:    "https://gitlab.com/fdroid/fdroidclient.git",
			wantErr: false,
		},
		{
			name:    "GitLab HTTPS URL",
			args:    args{r: "https://gitlab.com/fdroid/fdroidclient.git"},
			want:    "https://gitlab.com/fdroid/fdroidclient.git",
			wantErr: false,
		},
		{
			name:    "Nonsense URL",
			args:    args{r: "ftp://invalid.git"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := processRepoURL(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("processRepoURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("processRepoURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
