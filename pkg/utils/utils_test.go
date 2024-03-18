package utils_test

import (
	"context"
	"filmLibrary/pkg/utils"
	"testing"
)

func TestGetSearchQuery(t *testing.T) {
	tests := []struct {
		name     string
		elem     string
		elements []string
		expected bool
	}{
		{"in", "1", []string{"1", "2"}, true},
		{"multiple in", "1", []string{"1", "2", "1"}, true},
		{"single in", "1", []string{"1"}, true},
		{"not in", "1", []string{"2", "3"}, false},
		{"empty", "1", []string{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := utils.In(tt.elem, tt.elements)

			if actual != tt.expected {
				t.Errorf("In() = %v, want %v", actual, tt.expected)
			}
		})
	}
}
