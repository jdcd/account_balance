package controller

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/jdcd/account_balance/internal/application"
	"github.com/jdcd/account_balance/pkg"
)

type ReportController struct {
	PickerUseCase  application.IFilePickerUseCase
	ProcessUseCase application.IProcessTransactionsUseCase
}

const (
	errorParsingNotificationList = "error reading the request"
	errorReadingFile             = "error reading transactions file"
	wrongNotificationList        = "wrong mailing list, check data and try again"
	emptyFileFolder              = "there is no file to process"
	startProcessMessage          = "file %s will be processed, process notification will be sent to non-discarded destinations"
)

func (r *ReportController) PostReport(c *gin.Context) {
	var rList RecipeList
	if err := c.Bind(&rList); err != nil {
		response := ApiResponse{Message: errorParsingNotificationList, Error: err.Error()}
		c.IndentedJSON(http.StatusBadRequest, response)
		return
	}

	validList, invalidList := pkg.IsValidateEmailList(rList.EmailList)

	if len(validList) == 0 {
		response := ApiResponse{Message: wrongNotificationList}
		c.IndentedJSON(http.StatusBadRequest, response)
		return
	}

	fileToProcess, err := r.PickerUseCase.SelectFileHandler()
	if err != nil {
		response := ApiResponse{Message: errorReadingFile}
		c.IndentedJSON(http.StatusBadRequest, response)
		return
	}

	if fileToProcess == "" {
		response := ApiResponse{Message: emptyFileFolder}
		c.IndentedJSON(http.StatusOK, response)
		return
	}

	response := ApiResponse{
		Message: fmt.Sprintf(startProcessMessage, filepath.Base(fileToProcess)),
		Notification: Notification{
			ValidList:   validList,
			InvalidList: invalidList,
		},
	}
	c.IndentedJSON(http.StatusAccepted, response)
	go r.ProcessUseCase.TransactionHandler(fileToProcess, validList)
}
