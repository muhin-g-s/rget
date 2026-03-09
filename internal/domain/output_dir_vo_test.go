package domain

import (
	"testing"
)

func TestNewOutputDir(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantValue string
		wantErr   error
	}{
		{
			name:      "valid directory path",
			input:     "/tmp/downloads",
			wantValue: "/tmp/downloads",
			wantErr:   nil,
		},
		{
			name:      "valid relative path",
			input:     "./downloads",
			wantValue: "./downloads",
			wantErr:   nil,
		},
		{
			name:    "empty string should error",
			input:   "",
			wantErr: ErrOutputDirCannotBeEmpty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewOutputDir(tt.input)

			if tt.wantErr != nil {
				if err != tt.wantErr {
					t.Errorf("NewOutputDir() error = %v, want %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("NewOutputDir() unexpected error = %v", err)
				return
			}

			if got.Value() != tt.wantValue {
				t.Errorf("NewOutputDir().Value() = %v, want %v", got.Value(), tt.wantValue)
			}
		})
	}
}
