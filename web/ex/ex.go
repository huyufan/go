package ex

type Show struct {
	ID int
}

func (id Show) GetShow() int {
	return id.ID
}
