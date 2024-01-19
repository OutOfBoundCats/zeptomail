package zmodels

type EmailAddress struct {
	Address string `json:"address" validate:"required"`
	Name    string `json:"name" validate:"required"`
}

type Attachments []struct {
	Content      string `json:"content"`
	MimeType     string `json:"mime_type,omitempty"`
	Name         string `json:"name"`
	FileCacheKey string `json:"file_cache_key,omitempty"`
}

type SendEmailTo struct {
	EmailAddress EmailAddress `json:"email_address" validate:"required"`
}
