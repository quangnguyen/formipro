package generator

import (
	"com.nguyenonline/formipro/pkg/model"
	"errors"
)

type Generator interface {
	Pdf(templateID string, obj interface{}) ([]byte, error)
}

func New(obj interface{}) (Generator, error) {
	_, ok := obj.(model.Letter)
	if ok {
		return LetterGenerator{}, nil
	}

	return nil, errors.New("could not find a suitable generator")
}
