package zmodels

// temaplte email
type ZeptoTemplateEmail[MergeDataType any] struct {
	TemplateKey     string                 `json:"template_key" validate:"required_without=TemplateAlias"`
	TemplateAlias   string                 `json:"template_alias,omitempty" validate:"required_without=TemplateKey"`
	BounceAddress   string                 `json:"bounce_address,omitempty"`
	From            EmailAddress           `json:"from" validate:"required"`
	To              []SendEmailTo          `json:"to" validate:"required"`
	MergeInfo       MergeDataType          `json:"merge_info" validate:"required"`
	ReplyTo         []EmailAddress         `json:"reply_to"`
	TrackClicks     bool                   `json:"track_clicks,omitempty"`
	TrackOpens      bool                   `json:"track_opens,omitempty"`
	ClientReference string                 `json:"client_reference,omitempty"`
	MimeHeaders     map[string]interface{} `json:"mime_headers,omitempty"`
	Attachments     []Attachments          `json:"attachments,omitempty"`
	Cid             string                 `json:"cid,omitempty"`
}
