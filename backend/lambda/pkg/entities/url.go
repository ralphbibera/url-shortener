package entities

type Url struct {
	UrlId       *string `json:"urlId,omitempty"`
	RedirectUrl *string `json:"redirectUrl,omitempty"`
}
