package api

import (
	"com.nguyenonline/formipro/internal/generator"
	"com.nguyenonline/formipro/pkg/model"
	"encoding/json"
	"net/http"
)

func DecodeModel(request *http.Request, model model.Model) error {
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&model)
	return err
}

func WriteResponse(w http.ResponseWriter, model model.Model) {
	pdfGenerator, err := generator.NewGenerator(model)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pwd, err := pdfGenerator.GeneratePdf(model.GetTemplateID(), model)
	if err != nil {
		http.Error(w, "Some thing wrong happened!", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(pwd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
