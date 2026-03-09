package cli

import (
	"errors"
	"fmt"
)

type UC interface {
	Execute(outputDir string, remouteFiles []string) (string, error)
}

type CLI struct {
	uc UC
}

func New(uc UC) *CLI {
	return &CLI{
		uc: uc,
	}
}

var (
	ErrTooFewArguments = errors.New("Too few arguments.")
	ErrUc              = errors.New("UC")
)

func (c *CLI) Handle(args []string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("%w More than 2 expected. Now it is %d", ErrTooFewArguments, len(args))
	}

	res, err := c.uc.Execute(args[0], args[1:])
	if err != nil {
		return "", err
	}

	return res, nil
}
