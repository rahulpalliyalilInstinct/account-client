package listing

// Page struct consists of number and Size.
type Page struct {
	Number int
	Size   int
}

// NewPage returns a page with page number and size.
func NewPage(Number, Size int) *Page {
	return &Page{
		Number: Number,
		Size:   Size,
	}
}
