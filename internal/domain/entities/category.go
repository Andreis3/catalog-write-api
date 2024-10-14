package entities

type Category struct {
	ID          int64
	CategoryID  int64
	Slug        string
	Description string
	Parent      string
}
