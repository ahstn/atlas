package git

import (
	"testing"
)

func TestIsShortLivedBranch(t *testing.T) {
	tests := []struct {
		name   string
		branch string
		want   bool
	}{
		{
			name:   "develop returns false",
			branch: "develop",
			want:   false,
		},
		{
			name:   "master returns false",
			branch: "master",
			want:   false,
		},
		{
			name:   "feature returns true",
			branch: "feature/TEAM-123",
			want:   true,
		},
		{
			name:   "bugfix returns true",
			branch: "bugfix/TEAM-123-fix-docker",
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsShortLivedBranch(tt.branch); got != tt.want {
				t.Errorf("IsShortLivedBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimBranchContext(t *testing.T) {
	tests := []struct {
		name   string
		branch string
		want   string
	}{
		{
			name:   "feature/TEAM-123 remains the same",
			branch: "feature/TEAM-123",
			want:   "feature/TEAM-123",
		},
		{
			name:   "feature/TEAM-123-implment-feature gets context removed",
			branch: "feature/TEAM-123-implment-feature",
			want:   "feature/TEAM-123",
		},
		{
			name:   "develop remains the same",
			branch: "develop",
			want:   "develop",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimBranchContext(tt.branch); got != tt.want {
				t.Errorf("TrimBranchContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
