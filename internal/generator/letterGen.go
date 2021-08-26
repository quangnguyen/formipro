package generator

import (
	"com.nguyenonline/formipro/internal"
	"com.nguyenonline/formipro/pkg/model"
)

type LetterGenerator struct {
}

func letterPdf(templateId string, letter *model.Letter) ([]byte, error) {
	replaceSpecialCharacters(letter)
	tex, err := HtmlToLatex(letter.MainContent.HTML)
	if err != nil {
		return []byte{}, err
	}
	letter.MainContent.Tex = tex
	return GeneratePdf(templateId, letter)
}

func replaceSpecialCharacters(letter *model.Letter) {
	letter.Reference.CustomerId = internal.ReplaceToTex(letter.Reference.CustomerId)
	letter.Reference.Id = internal.ReplaceToTex(letter.Reference.Id)
	letter.Reference.MailDate = internal.ReplaceToTex(letter.Reference.MailDate)
	letter.Title = internal.ReplaceToTex(letter.Title)
	letter.OpeningText = internal.ReplaceToTex(letter.OpeningText)
	letter.ClosingText = internal.ReplaceToTex(letter.ClosingText)
}

func (l LetterGenerator) GeneratePdf(templateId string, obj interface{}) ([]byte, error) {
	letter := obj.(model.Letter)
	return letterPdf(templateId, &letter)
}
