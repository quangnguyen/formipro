package generator

import (
	"com.nguyenonline/formipro/pkg/model"
	"errors"
)

type Generator interface {
	GeneratePdf(templateId string, obj interface{}) ([]byte, error)
}

func NewGenerator(obj interface{}) (Generator, error) {
	_, ok := obj.(model.Letter)
	if ok {
		return LetterGenerator{}, nil
	}

	return nil, errors.New("could not find a suitable generator")
}
