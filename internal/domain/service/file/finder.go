package file

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jdcd/account_balance/pkg"
)

// IFinder provides contracts related to find report files on filesystem
type IFinder interface {
	// GiveNextFileName returns the name of the next csv file to process
	GiveNextFileName() (fileName string, err error)

	// Relocate changes csv file location
	Relocate(filePath string, date time.Time, newFolder Directory)
}

type Directory string

const (
	fileTimeFormat           = "2006-01-02-15-04-05"
	PendingFolder  Directory = "pending"
	ErrorFolder    Directory = "error"
	SuccessFolder  Directory = "success"
)

type FinderService struct {
	PendingDirectory string
	SuccessDirectory string
	ErrorDirectory   string
}

func (s *FinderService) GiveNextFileName() (string, error) {
	var fileName string
	files, err := filepath.Glob(filepath.Join(s.PendingDirectory, "*.csv"))
	if err != nil {
		pkg.ErrorLogger().Printf("error searching for csv files %s", err)
		return fileName, err
	}

	if len(files) > 0 {
		fileName = files[0]
	}

	return fileName, err
}

func (s *FinderService) Relocate(filePath string, date time.Time, newFolder Directory) {
	fileName := filepath.Base(filePath)
	formattedDate := date.Format(fileTimeFormat)
	newPath := fmt.Sprintf("%s/%s-%s", s.getDirectory(newFolder), fileName, formattedDate)

	err := os.Rename(filePath, newPath)
	if err != nil {
		pkg.ErrorLogger().Printf("error moving %s file to %s folder: %s", filePath, newPath, err)
		return
	}
}

func (s *FinderService) getDirectory(folder Directory) string {
	switch folder {
	case ErrorFolder:
		return s.ErrorDirectory
	case SuccessFolder:
		return s.SuccessDirectory
	case PendingFolder:
		return s.PendingDirectory
	default:
		return ""
	}
}
