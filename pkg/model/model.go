package model

type Model interface {
	Name() string
	GetTemplateID() string
	GetAttachments() map[string][]byte
}
