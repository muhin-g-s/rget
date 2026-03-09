package app

import (
	"github.com/muhin-g-s/rget/internal/cli"
	"github.com/muhin-g-s/rget/internal/usecase"
)

func Run(args []string) string {
	uc := usecase.New()

	cli := cli.New(uc)

	msg, err := cli.Handle(args)
	if err != nil {
		return err.Error()
	}

	return msg
}
