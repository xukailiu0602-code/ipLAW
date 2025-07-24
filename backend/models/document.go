package models

type Document struct {
	ID        string   `bson:"_id" json:"id"`
	UserID    string   `bson:"user_id" json:"user_id"`
	Title     string   `bson:"title" json:"title"`
	Type      string   `bson:"type" json:"type"` // pdf, docx, txt
	Meta      any      `bson:"meta" json:"meta"`
	CreatedAt int64    `bson:"created_at" json:"created_at"`
	Slices    []Slice  `bson:"slices" json:"slices"`
}

type Slice struct {
	ID        string `bson:"id" json:"id"`
	Text      string `bson:"text" json:"text"`
	Embedding []float32 `bson:"embedding" json:"embedding"`
	Meta      any    `bson:"meta" json:"meta"`
}