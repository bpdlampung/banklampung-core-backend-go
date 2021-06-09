package entities

const (
	SortAscending  = `asc`
	SortDescending = `desc`
)

type Sort struct {
	FieldName string
	By        string
}

func (s Sort) BuildSortBy() int {
	if s.By == SortDescending {
		return -1
	}

	return 1
}
