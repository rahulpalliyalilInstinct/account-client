package listing

type Page struct {
	Number int
	Size   int
}

func NewPage(Number, Size int) *Page {
	return &Page{
		Number: Number,
		Size:   Size,
	}
}
