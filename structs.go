package goflip

const (
	// BoxDefaultOutlineColor default color for text border color
	BoxDefaultOutlineColor = "#000000"

	// BoxDefaultColor default color for text fill color
	BoxDefaultColor = "#ffffff"
)

// Textbox defines a textbox to put into an ImgFlip meme template
type Textbox struct {
	Text         string `json:"text"`
	X            uint   `json:"x"`
	Y            uint   `json:"y"`
	Width        uint   `json:"width"`
	Height       uint   `json:"height"`
	Color        string `json:"color"`
	OutlineColor string `json:"outline_color"`
}

// APIRequest a request object to create a new meme on the ImgFlip API
type APIRequest struct {
	TemplateID string    `json:"template_id"`
	Text0      string    `json:"text0"`
	Text1      string    `json:"text1"`
	Boxes      []Textbox `json:"boxes"`
}

// APIResponse a response object returned by the ImgFlip API
type APIResponse struct {
	Success bool               `json:"success"`
	Data    APIResponsePayload `json:"data"`
	Error   string             `json:"error_message"`
}

// APIResponsePayload the data returned from a successful ImgFlip API response
type APIResponsePayload struct {
	URL     string `json:"url"`
	PageURL string `json:"page_url"`
}
