package application

import "github.com/jdcd/account_balance/internal/domain/service/file"

type IFilePickerUseCase interface {
	SelectFileHandler() (filePath string, err error)
}

type FilePickerUseCase struct {
	Finder file.IFinder
}

func (c *FilePickerUseCase) SelectFileHandler() (string, error) {
	return c.Finder.GiveNextFileName()
}
