package config

import (
	"database/sql"

	"github.com/jdcd/account_balance/internal/application"
	"github.com/jdcd/account_balance/internal/domain/service/file"
	notificationService "github.com/jdcd/account_balance/internal/domain/service/notification"
	notificationAdapter "github.com/jdcd/account_balance/internal/infrastructure/adapter/notification"
	"github.com/jdcd/account_balance/internal/infrastructure/adapter/report/postgresSQL"
	"github.com/jdcd/account_balance/internal/infrastructure/adapter/template"
	"github.com/jdcd/account_balance/internal/infrastructure/http/server"
	"github.com/jdcd/account_balance/internal/infrastructure/http/server/controller"
)

// GetRouterDependencies configures dependencies injection
func GetRouterDependencies(config *AppConfiguration) *server.RouterDependencies {
	return &server.RouterDependencies{
		CheckController:  &controller.PingController{},
		ReportController: getReportDependencies(config),
	}
}

func getReportDependencies(config *AppConfiguration) *controller.ReportController {
	finderImpl := getFinderService()
	return &controller.ReportController{
		ProcessUseCase: &application.ProcessTransactionsUseCase{
			Sender:    getNotificationService(config),
			Finder:    finderImpl,
			Reader:    &file.ReaderService{},
			Processor: &file.ProcessService{},
			Repository: &postgresSQL.Repository{
				Writer: getPostgresSQLConnection(config.DbUrl)},
		},
		PickerUseCase: &application.FilePickerUseCase{Finder: finderImpl},
	}

}

func getPostgresSQLConnection(dbUrl string) *sql.DB {
	return postgresSQL.GetConnection(dbUrl)
}

func getFinderService() file.IFinder {
	return &file.FinderService{
		PendingDirectory: "/csv_files/pending",
		SuccessDirectory: "/csv_files/processed",
		ErrorDirectory:   "/csv_files/error",
	}
}

func getNotificationService(config *AppConfiguration) notificationService.ISender {
	return &notificationService.SenderService{
		NotificationDeliveryService: &notificationAdapter.SmtpEmailSender{
			EmailSender:     config.EmailSender,
			EmailPwd:        config.EmailPwd,
			EmailSenderName: config.EmailSenderName,
			SmtpServer:      config.SmtpServer,
			SmtpPort:        config.SmtpPort,
		},
		SummaryTemplateGenerator: &template.HtmlSummaryGenerator{},
	}
}
