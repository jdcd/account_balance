package template

import (
	"bytes"
	"html/template"

	"github.com/jdcd/account_balance/internal/domain"
	"github.com/jdcd/account_balance/pkg"
)

type HtmlSummaryGenerator struct{}

const (
	templatePath          = "../../../../resources/html/template.html"
	templateNotFoundError = "template %s could not be opened"
	templateExecError     = "error parsing the html template: %s"
)

func (g *HtmlSummaryGenerator) FormatSummary(summary domain.Summary) (string, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		pkg.ErrorLogger().Printf(templateNotFoundError, err)
		return "", err
	}

	var bodyContent bytes.Buffer
	err = tmpl.Execute(&bodyContent, summary)
	if err != nil {
		pkg.ErrorLogger().Printf(templateExecError, err)
		return "", err
	}

	return bodyContent.String(), nil
}
