package domain

import "errors"

type OutputDir struct {
	value string
}

var (
	ErrOutputDirCannotBeEmpty = errors.New("cannot be an empty string")
)

func NewOutputDir(str string) (OutputDir, error) {
	if str == "" {
		return OutputDir{}, ErrOutputDirCannotBeEmpty
	}

	return OutputDir{
		value: str,
	}, nil
}

func (vo OutputDir) Value() string {
	return vo.value
}
