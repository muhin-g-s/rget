package domain

import (
	"testing"
)

func TestNewRemoteFileAddr(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantValue string
		wantErr   error
	}{
		{
			name:      "valid https url",
			input:     "https://example.com/file.txt",
			wantValue: "https://example.com/file.txt",
			wantErr:   nil,
		},
		{
			name:      "valid http url",
			input:     "http://example.com/file.txt",
			wantValue: "http://example.com/file.txt",
			wantErr:   nil,
		},
		{
			name:      "valid url with path",
			input:     "https://example.com/path/to/file.txt",
			wantValue: "https://example.com/path/to/file.txt",
			wantErr:   nil,
		},
		{
			name:    "empty string should error",
			input:   "",
			wantErr: ErrRemoteFileEmptyStr,
		},
		{
			name:    "invalid protocol ftp should error",
			input:   "ftp://example.com/file.txt",
			wantErr: ErrRemoteFileAddrInvalidProtocol,
		},
		{
			name:    "invalid url format should error",
			input:   "not-a-valid-url",
			wantErr: ErrRemoteFileAddrNotValidUrl,
		},
		{
			name:    "url without host should error",
			input:   "http:///file.txt",
			wantErr: ErrRemoteFileAddrNotValidUrl,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRemoteFileAddr(tt.input)

			if tt.wantErr != nil {
				if err != tt.wantErr {
					t.Errorf("NewRemoteFileAddr() error = %v, want %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("NewRemoteFileAddr() unexpected error = %v", err)
				return
			}

			if got.Value() != tt.wantValue {
				t.Errorf("NewRemoteFileAddr().Value() = %v, want %v", got.Value(), tt.wantValue)
			}
		})
	}
}

func TestParseRemoteFileAddrs(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		wantLen  int
		wantErr  bool
	}{
		{
			name:    "multiple valid urls",
			input:   []string{"https://example.com/file1.txt", "http://example.com/file2.txt"},
			wantLen: 2,
			wantErr: false,
		},
		{
			name:    "empty list",
			input:   []string{},
			wantLen: 0,
			wantErr: false,
		},
		{
			name:    "one invalid url in list should error",
			input:   []string{"https://example.com/file1.txt", "invalid-url"},
			wantLen: 0,
			wantErr: true,
		},
		{
			name:    "invalid protocol in list should error",
			input:   []string{"ftp://example.com/file.txt"},
			wantLen: 0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseRemoteFileAddrs(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Errorf("ParseRemoteFileAddrs() expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("ParseRemoteFileAddrs() unexpected error = %v", err)
				return
			}

			if len(got) != tt.wantLen {
				t.Errorf("ParseRemoteFileAddrs() len = %v, want %v", len(got), tt.wantLen)
			}
		})
	}
}
