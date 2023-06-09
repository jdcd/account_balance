package notification

import (
	"errors"
	"github.com/jdcd/account_balance/internal/mock/port"
	"testing"

	"github.com/jdcd/account_balance/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestTemplateGeneratorIsWrongThenSendNotificationReturnError(t *testing.T) {
	generatorHtmlMock := &port.GeneratorMock{}
	ms := SenderService{SummaryTemplateGenerator: generatorHtmlMock}
	emailList := []string{"example@gmail.com"}
	expectedError := errors.New("an error")
	summary := domain.Summary{}
	generatorHtmlMock.On("FormatSummary", summary).Return("", expectedError).Once()

	err := ms.SendSummaryNotification(summary, emailList)

	assert.EqualValues(t, expectedError, err)
	generatorHtmlMock.AssertExpectations(t)
}

func TestWhenDeliveryServiceANdTemplateGeneratorIsOkThenSendNotificationReturnOk(t *testing.T) {
	deliveryMock := &port.DeliveryMock{}
	generatorHtmlMock := &port.GeneratorMock{}
	ms := SenderService{
		NotificationDeliveryService: deliveryMock,
		SummaryTemplateGenerator:    generatorHtmlMock,
	}
	emailList := []string{"example@gmail.com"}
	content := "fancy html content"
	summary := domain.Summary{}
	notification := domain.Notification{
		Content:    content,
		Subject:    notificationSubject,
		Recipients: emailList,
		Branding:   notificationBranding,
	}
	generatorHtmlMock.On("FormatSummary", summary).Return(content, nil).Once()
	deliveryMock.On("SendNotification", notification).Return(nil).Once()

	err := ms.SendSummaryNotification(summary, emailList)

	assert.Nil(t, err)
	deliveryMock.AssertExpectations(t)
	generatorHtmlMock.AssertExpectations(t)
}
