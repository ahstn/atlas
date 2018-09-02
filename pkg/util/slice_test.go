package util

import (
	"reflect"
	"testing"
)

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

func TestStringContainsAny(t *testing.T) {
	type args struct {
		s       string
		targets []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "happy path",
			args: args{
				s:       "Hello World",
				targets: []string{"World", "Planet"},
			},
			want: true,
		},
		{
			name: "sad path",
			args: args{
				s:       "Hello World",
				targets: []string{"Planet", "Ocean"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringContainsAny(tt.args.s, tt.args.targets...); got != tt.want {
				t.Errorf("StringContainsAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSliceEquals(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "happy path - returns true",
			args: args{
				a: []string{"Hello", "World"},
				b: []string{"Hello", "World"},
			},
			want: true,
		},
		{
			name: "happy path - returns false",
			args: args{
				a: []string{"Hello", "World"},
				b: []string{"Hello"},
			},
			want: false,
		},
		{
			name: "happy path - returns false",
			args: args{
				a: []string{"Hello", "World"},
				b: []string{"Hello", "Adam"},
			},
			want: false,
		},
		{
			name: "happy path - one nil slice",
			args: args{
				a: []string{"Hello", "World"},
				b: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringSliceEquals(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("StringSliceEquals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSliceRemove(t *testing.T) {
	type args struct {
		s []string
		r string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "happy path",
			args: args{
				s: []string{"Hello", "World"},
				r: "World",
			},
			want: []string{"Hello"},
		},
		{
			name: "slice doesn't contain element",
			args: args{
				s: []string{"Hello", "World"},
				r: "Planet",
			},
			want: []string{"Hello", "World"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringSliceRemove(tt.args.s, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringSliceRemove() = %v, want %v", got, tt.want)
			}
		})
	}
}
