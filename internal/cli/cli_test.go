package cli

import (
	"errors"
	"testing"
)

type mockUC struct {
	result string
	err    error
}

func (m *mockUC) Execute(outputDir string, remouteFiles []string) (string, error) {
	return m.result, m.err
}

func TestCLI_Handle(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		ucResult   string
		ucErr      error
		wantResult string
		wantErr    error
	}{
		{
			name:       "valid args returns result",
			args:       []string{"/tmp", "https://example.com/file.txt"},
			ucResult:   "success",
			ucErr:      nil,
			wantResult: "success",
			wantErr:    nil,
		},
		{
			name:    "too few args should error",
			args:    []string{"/tmp"},
			ucResult: "",
			ucErr:   nil,
			wantResult: "",
			wantErr:   ErrTooFewArguments,
		},
		{
			name:    "empty args should error",
			args:    []string{},
			ucResult: "",
			ucErr:   nil,
			wantResult: "",
			wantErr:   ErrTooFewArguments,
		},
		{
			name:       "uc error propagates",
			args:       []string{"/tmp", "https://example.com/file.txt"},
			ucResult:   "",
			ucErr:      errors.New("uc error"),
			wantResult: "",
			wantErr:    errors.New("uc error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &mockUC{result: tt.ucResult, err: tt.ucErr}
			cli := New(uc)

			got, err := cli.Handle(tt.args)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("Handle() expected error but got none")
					return
				}
				if tt.wantErr == ErrTooFewArguments && !errors.Is(err, ErrTooFewArguments) {
					t.Errorf("Handle() error = %v, want %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("Handle() unexpected error = %v", err)
				return
			}

			if got != tt.wantResult {
				t.Errorf("Handle() result = %v, want %v", got, tt.wantResult)
			}
		})
	}
}

func TestNew(t *testing.T) {
	uc := &mockUC{}
	cli := New(uc)

	if cli == nil {
		t.Errorf("New() returned nil")
	}
	if cli.uc != uc {
		t.Errorf("New() uc not set correctly")
	}
}
