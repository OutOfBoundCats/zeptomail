package zmodels

type EmailData struct {
	BounceAddress   string                 `json:"bounce_address,omitempty"`
	From            EmailAddress           `json:"from" validate:"required"`
	To              []SendEmailTo          `json:"to" validate:"required"`
	ReplyTo         []EmailAddress         `json:"reply_to"`
	Subject         string                 `json:"subject" validate:"required"`
	Htmlbody        string                 `json:"htmlbody" validate:"required_without=Textbody"`
	Textbody        string                 `json:"textbody" validate:"required_without=Htmlbody"`
	Cc              []EmailAddress         `json:"cc,omitempty"`
	Bcc             []EmailAddress         `json:"bcc,omitempty"`
	TrackClicks     bool                   `json:"track_clicks,omitempty"`
	TrackOpens      bool                   `json:"track_opens,omitempty"`
	ClientReference string                 `json:"client_reference,omitempty"`
	MimeHeaders     map[string]interface{} `json:"mime_headers,omitempty"`
	Attachments     []Attachments          `json:"attachments,omitempty"`
	Name            string                 `json:"name"`
	InlineImage     []string               `json:"inline_images,omitempty"`
	Cid             string                 `json:"cid,omitempty"`
}
