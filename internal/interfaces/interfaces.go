package interfaces

type PasteRequest struct {
	Text   string `json:"text"`
	Expiry string `json:"expiry,omitempty"`
}

type PasteResponse struct {
	Url string `json:"url"`
}
