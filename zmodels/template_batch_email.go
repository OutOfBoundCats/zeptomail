package zmodels

// bactch temaplte email
type BatchTemplateTo[MergeDataType any] struct {
	ToEmailAddress EmailAddress  `json:"email_address" validate:"required"`
	MergeInfo      MergeDataType `json:"merge_info" validate:"required"`
}

type TemplateEmailBatch[MergeDataType any] struct {
	TemplateKey     string                           `json:"template_key" validate:"required_without=TemplateAlias"`
	TemplateAlias   string                           `json:"template_alias" validate:"required_without=TemplateKey"`
	BounceAddress   string                           `json:"bounce_address,omitempty"`
	From            EmailAddress                     `json:"from" validate:"required"`
	To              []BatchTemplateTo[MergeDataType] `json:"to" validate:"required"`
	ReplyTo         []EmailAddress                   `json:"reply_to"`
	TrackClicks     bool                             `json:"track_clicks,omitempty"`
	TrackOpens      bool                             `json:"track_opens,omitempty"`
	ClientReference string                           `json:"client_reference,omitempty"`
	MimeHeaders     map[string]interface{}           `json:"mime_headers,omitempty"`
	Attachments     []Attachments                    `json:"attachments,omitempty"`
	Cid             string                           `json:"cid,omitempty"`
}
