package analysis

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectLanguage(t *testing.T) {
	tmp, err := ioutil.TempDir("", "TestInitialiseCommand")
	assert.NoError(t, err)

	wd, err := os.Getwd()
	assert.NoError(t, err)

	err = os.Chdir(tmp)
	assert.NoError(t, err)

	defer os.Chdir(wd)

	contents := []byte("package main\nfunc main() {}\n")
	err = ioutil.WriteFile("./main.go", contents, 0644)
	assert.NoError(t, err)

	repo, err := DetectLanguage(wd)
	assert.NoError(t, err)

	assert.Equal(t, Go, repo.Language)
	assert.Equal(t, float64(100), repo.Percentages[Go])
}

func TestDetectLanguageError(t *testing.T) {
	_, err := DetectLanguage("/invalid_dir")

	assert.Error(t, err)
}

func TestGetExtension(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name     string
		filename string
		want     Language
	}{
		{
			name:     "detects Go",
			filename: "/tmp/main.go",
			want:     Go,
		},
		{
			name:     "detects Java",
			filename: "/tmp/Application.java",
			want:     Java,
		},
		{
			name:     "detects JavaScript",
			filename: "/tmp/index.js",
			want:     JavaScript,
		},
		{
			name:     "defaults to 'Unknown'",
			filename: "/tmp/file.txt",
			want:     Unknown,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getExtension(tt.filename); got != tt.want {
				t.Errorf("getExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculatePercentages(t *testing.T) {
	tests := []struct {
		name      string
		languages map[Language]int
		total     int
		want      RepoResult
		wantErr   bool
	}{
		{
			name: "only Golang",
			languages: map[Language]int{
				Go: 100,
			},
			total: 100,
			want: RepoResult{
				Language: Go,
				Percentages: Percentages{
					Go: 100.0,
				},
			},
			wantErr: false,
		},
		{
			name: "invalid Langugage input data",
			languages: map[Language]int{
				Go:   450,
				Java: 250,
			},
			total:   75,
			want:    RepoResult{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculatePercentages(tt.languages, tt.total)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculatePercentages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculatePercentages() = %v, want %v", got, tt.want)
			}
		})
	}
}
