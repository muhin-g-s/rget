package usecase

import (
	"strings"
	"testing"
)

func TestFetchAndSaveUc_Execute(t *testing.T) {
	tests := []struct {
		name           string
		outputDir      string
		remoteFiles    []string
		wantContains   string
		wantErr        bool
	}{
		{
			name:         "valid input returns formatted output",
			outputDir:    "/tmp/downloads",
			remoteFiles:  []string{"https://example.com/file.txt"},
			wantContains: "OutputDir: /tmp/downloads",
			wantErr:      false,
		},
		{
			name:         "multiple files",
			outputDir:    "/tmp",
			remoteFiles:  []string{"https://example.com/file1.txt", "https://example.com/file2.txt"},
			wantContains: "0 url https://example.com/file1.txt",
			wantErr:      false,
		},
		{
			name:        "empty output dir should error",
			outputDir:   "",
			remoteFiles: []string{"https://example.com/file.txt"},
			wantErr:     true,
		},
		{
			name:        "invalid url should error",
			outputDir:   "/tmp",
			remoteFiles: []string{"invalid-url"},
			wantErr:     true,
		},
		{
			name:        "invalid protocol should error",
			outputDir:   "/tmp",
			remoteFiles: []string{"ftp://example.com/file.txt"},
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := New()

			got, err := uc.Execute(tt.outputDir, tt.remoteFiles)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Execute() expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Execute() unexpected error = %v", err)
				return
			}

			if !strings.Contains(got, tt.wantContains) {
				t.Errorf("Execute() result = %v, should contain %v", got, tt.wantContains)
			}
		})
	}
}

func TestNew(t *testing.T) {
	uc := New()
	if uc == nil {
		t.Errorf("New() returned nil")
	}
}
