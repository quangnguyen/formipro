package api

import (
	"com.nguyenonline/formipro/internal/generator"
	"com.nguyenonline/formipro/pkg/model"
	"encoding/json"
	"net/http"
)

func DecodeModel(r *http.Request, m model.Model) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&m)
	return err
}

func WriteResponse(w http.ResponseWriter, m model.Model) {
	g, err := generator.New(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := g.Pdf(m.GetTemplateID(), m)
	if err != nil {
		http.Error(w, "Some thing wrong happened!", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(bytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
