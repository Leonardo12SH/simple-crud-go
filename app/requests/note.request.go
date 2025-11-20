package requests

type NoteRequest struct {
	Judul string `json:"judul" form:"judul" binding:"required"`
	Notes string `json:"notes" form:"notes" binding:"required"`
}
