package generator

import (
	"com.nguyenonline/formipro/internal"
	"com.nguyenonline/formipro/pkg/model"
)

type LetterGenerator struct {
}

func generate(templateID string, letter *model.Letter) ([]byte, error) {
	replaceSpecialCharacters(letter)
	tex, err := HTMLToLatex(letter.MainContent.HTML)
	if err != nil {
		return []byte{}, err
	}
	letter.MainContent.Tex = tex
	return PdfLatex(templateID, letter)
}

func replaceSpecialCharacters(l *model.Letter) {
	l.Reference.CustomerID = internal.ReplaceSpecialCharWithTexSymbol(l.Reference.CustomerID)
	l.Reference.ID = internal.ReplaceSpecialCharWithTexSymbol(l.Reference.ID)
	l.Reference.MailDate = internal.ReplaceSpecialCharWithTexSymbol(l.Reference.MailDate)
	l.Title = internal.ReplaceSpecialCharWithTexSymbol(l.Title)
	l.OpeningText = internal.ReplaceSpecialCharWithTexSymbol(l.OpeningText)
	l.ClosingText = internal.ReplaceSpecialCharWithTexSymbol(l.ClosingText)
}

func (l LetterGenerator) Pdf(templateID string, obj interface{}) ([]byte, error) {
	letter := obj.(model.Letter)
	return generate(templateID, &letter)
}
