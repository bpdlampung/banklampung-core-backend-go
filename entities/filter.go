package entities

import "strings"

type Filter struct {
	Sort []Sort
	Page int64
	Size int64
}

type Paging struct {
	Page        int64  `json:"page,default=1" form:"page,default=1"`
	ItemPerPage int64  `json:"item_per_page,default=10" form:"item_per_page,default=10"`
	SortBy      string `json:"sort_by" form:"sort_by"`
}

func (p Paging) BuildSortField() (sorts []Sort) {
	if len(p.SortBy) > 0 {
		splitedSort := strings.Split(p.SortBy, ",")

		for x := range splitedSort {
			splitedField := strings.Split(splitedSort[x], ":")

			if len(splitedField) != 2 {
				continue
			}

			if splitedField[1] == "asc" || splitedField[1] == "desc" {
				sorts = append(sorts, Sort{
					FieldName: splitedField[0],
					By:        splitedField[1],
				})
			}
		}
	}

	return sorts
}

func (p Paging) ToFilter() Filter {
	return Filter{
		Sort: p.BuildSortField(),
		Page: p.Page,
		Size: p.ItemPerPage,
	}
}
