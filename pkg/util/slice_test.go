package util

import "testing"

func TestStringSliceContains(t *testing.T) {
	type args struct {
		slice  []string
		search string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "happy path",
			args: args{
				slice:  []string{"git", "maven", "docker"},
				search: "docker",
			},
			want: true,
		},
		{
			name: "sad path",
			args: args{
				slice:  []string{"git", "maven"},
				search: "docker",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringSliceContains(tt.args.slice, tt.args.search); got != tt.want {
				t.Errorf("StringSliceContains() = %v, want %v", got, tt.want)
			}
		})
	}
}
