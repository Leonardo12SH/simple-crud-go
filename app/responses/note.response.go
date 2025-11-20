package responses

type NoteResponse struct {
	ID    *int    `json:"id"`
	Judul *string `json:"judul"`
	Notes *string `json:"notes"`
}
