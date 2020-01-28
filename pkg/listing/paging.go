package listing

// Page struct consists of number and Size.
type Page struct {
	Number int
	Size   int
}

// NewPage returns a page with page number and size.
func NewPage(number, size int) *Page {
	return &Page{
		Number: number,
		Size:   size,
	}
}
