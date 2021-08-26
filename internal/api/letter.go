package api

import (
	"bytes"
	"com.nguyenonline/formipro/pkg/model"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const FormLetterContentJson = "content"
const FormAttachment = "attachments"

func Letter(w http.ResponseWriter, r *http.Request) {
	var letter model.Letter
	err := DecodeModel(r, &letter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteResponse(w, letter)
}

func FormLetter(w http.ResponseWriter, r *http.Request) {
	var letter model.Letter
	content := r.FormValue(FormLetterContentJson)
	if content == "" {
		http.Error(w, "Expected letter content from multipart form not found.", http.StatusBadRequest)
		return
	}
	err := json.Unmarshal([]byte(content), &letter)
	if err != nil {
		http.Error(w, "Content of letter was not conform", http.StatusBadRequest)
		return
	}

	err = attachment(r, &letter)
	if err != nil {
		http.Error(w, "Could not process attachments", http.StatusInternalServerError)
		return
	}

	WriteResponse(w, letter)
}

func attachment(r *http.Request, letter *model.Letter) error {
	files := r.MultipartForm.File[FormAttachment]
	letter.Attachments = make(map[string][]byte)
	for _, att := range files {
		f, err := att.Open()
		if err != nil {
			return err
		}

		bytes := bytes.NewBuffer(nil)
		if _, err := io.Copy(bytes, f); err != nil {
			return err
		}

		mimeType := http.DetectContentType(bytes.Bytes())

		if mimeType != "application/pdf" {
			return errors.New("attachment must be pdf")
		}

		letter.Attachments[att.Filename] = bytes.Bytes()
	}
	return nil
}
