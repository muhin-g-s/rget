package usecase

import (
	"fmt"
	"strings"

	"github.com/muhin-g-s/rget/internal/domain"
)

type FetchAndSaveUc struct{}

func New() *FetchAndSaveUc {
	return &FetchAndSaveUc{}
}

func (uc *FetchAndSaveUc) Execute(outputDirStr string, remouteFilesStr []string) (string, error) {
	outputDir, remouteFiles, err := parseInput(outputDirStr, remouteFilesStr)
	if err != nil {
		return "", err
	}

	var msg strings.Builder

	fmt.Fprintf(&msg, "OutputDir: %s\n", outputDir.Value())

	for index, url := range remouteFiles {
		fmt.Fprintf(&msg, "%d url %s \n", index, url.Value())
	}

	return msg.String(), nil
}

func parseInput(outputDirStr string, remouteFilesStr []string) (*domain.OutputDir, []domain.RemoteFileAddr, error) {
	outputDir, err := domain.NewOutputDir(outputDirStr)
	if err != nil {
		return nil, nil, err
	}

	remouteFiles, err := domain.ParseRemoteFileAddrs(remouteFilesStr)
	if err != nil {
		return nil, nil, err
	}

	return &outputDir, remouteFiles, nil
}
