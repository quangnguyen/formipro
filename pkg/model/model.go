package model

type Model interface {
	Name() string
	GetTemplateId() string
	GetAttachments() map[string][]byte
}
