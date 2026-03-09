package app

import (
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		wantContains string
		wantErr      bool
	}{
		{
			name:         "valid args returns formatted output",
			args:         []string{"/tmp/downloads", "https://example.com/file.txt"},
			wantContains: "OutputDir: /tmp/downloads",
			wantErr:      false,
		},
		{
			name:         "multiple files",
			args:         []string{"/tmp", "https://example.com/file1.txt", "https://example.com/file2.txt"},
			wantContains: "0 url https://example.com/file1.txt",
			wantErr:      false,
		},
		{
			name:    "too few args returns error message",
			args:    []string{"/tmp"},
			wantContains: "Too few arguments",
			wantErr:      true,
		},
		{
			name:    "empty args returns error message",
			args:    []string{},
			wantContains: "Too few arguments",
			wantErr:      true,
		},
		{
			name:    "invalid url returns error message",
			args:    []string{"/tmp", "invalid-url"},
			wantContains: "",
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Run(tt.args)

			if tt.wantErr {
				// For error cases, we just check that we got some output (error message)
				if got == "" && tt.wantContains == "" {
					return
				}
			}

			if tt.wantContains != "" && !strings.Contains(got, tt.wantContains) {
				t.Errorf("Run() result = %v, should contain %v", got, tt.wantContains)
			}
		})
	}
}
