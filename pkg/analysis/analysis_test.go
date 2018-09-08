package analysis

import (
	"reflect"
	"testing"
)

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
