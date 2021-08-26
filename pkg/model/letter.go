package model

func (l Letter) Name() string {
	return "letter"
}

func (l Letter) GetTemplateId() string {
	return l.TemplateId
}

func (l Letter) GetAttachments() map[string][]byte {
	return l.Attachments
}

type Sender struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Street    string `json:"street"`
	Zipcode   string `json:"zipcode"`
	City      string `json:"city"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Signature []byte
}

type Receiver struct {
	Name    string `json:"name"`
	Street  string `json:"street"`
	Zipcode string `json:"zipcode"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type Reference struct {
	Id         string `json:"id"`
	CustomerId string `json:"customerId"`
	MailDate   string `json:"mailDate"`
}

type Letter struct {
	TemplateId  string      `json:"templateId"`
	Sender      Sender      `json:"sender"`
	Receiver    Receiver    `json:"receiver"`
	Reference   Reference   `json:"reference"`
	Title       string      `json:"title"`
	OpeningText string      `json:"openingText"`
	ClosingText string      `json:"closingText"`
	MainContent MainContent `json:"mainContent"`
	Language    string      `json:"language"`
	Attachments map[string][]byte
}
