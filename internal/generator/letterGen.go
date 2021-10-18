package generator

import (
	"com.nguyenonline/formipro/internal"
	"com.nguyenonline/formipro/pkg/model"
)

type LetterGenerator struct {
}

func letterPdf(templateID string, letter *model.Letter) ([]byte, error) {
	replaceSpecialCharacters(letter)
	tex, err := HTMLToLatex(letter.MainContent.HTML)
	if err != nil {
		return []byte{}, err
	}
	letter.MainContent.Tex = tex
	return GeneratePdf(templateID, letter)
}

func replaceSpecialCharacters(letter *model.Letter) {
	letter.Reference.CustomerID = internal.ReplaceToTex(letter.Reference.CustomerID)
	letter.Reference.ID = internal.ReplaceToTex(letter.Reference.ID)
	letter.Reference.MailDate = internal.ReplaceToTex(letter.Reference.MailDate)
	letter.Title = internal.ReplaceToTex(letter.Title)
	letter.OpeningText = internal.ReplaceToTex(letter.OpeningText)
	letter.ClosingText = internal.ReplaceToTex(letter.ClosingText)
}

func (l LetterGenerator) GeneratePdf(templateID string, obj interface{}) ([]byte, error) {
	letter := obj.(model.Letter)
	return letterPdf(templateID, &letter)
}
