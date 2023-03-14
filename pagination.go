package go_pagination

import (
	"fmt"
	"math"
)

type requestPagination struct {
	Page   int64 `json:"page"`   // current page
	Size   int64 `json:"size"`   // limit
	Offset int64 `json:"offset"` // offset
	Total  int64 `json:"total"`  // total row
}

type pagination struct {
	Page       int64 `json:"page"`       // current page
	Size       int64 `json:"size"`       // limit
	Offset     int64 `json:"offset"`     // offset
	TotalPages int64 `json:"totalPages"` // total page
	Total      int64 `json:"total"`      // total row
	Visible    int64 `json:"visible"`    // total row in current page
	Last       bool  `json:"last"`       // is last page
	First      bool  `json:"first"`      // is first page
}

// GetPagination is used to define pagination object and return it
func GetPagination(page int64, size int64, offset int64, total int64) requestPagination {
	return requestPagination{Page: page, Size: size, Offset: offset, Total: total}
}

// CreatePagination is used to create pagination response format using pagination object as parameter
func CreatePagination(req_pagination requestPagination) (pagination, error) {
	if req_pagination.Size <= req_pagination.Total {
		if req_pagination.Offset <= req_pagination.Total {
			var pagination pagination
			if req_pagination.Page == 0 {
				pagination.Page = 1
			} else {
				pagination.Page = req_pagination.Page
			}
			pagination.Size = req_pagination.Size
			pagination.Offset = req_pagination.Offset
			pagination.Total = req_pagination.Total
			if pagination.Total <= pagination.Size {
				pagination.Visible = pagination.Total
			} else if pagination.Total > pagination.Size {
				current_total := pagination.Page * pagination.Size
				if pagination.Total > current_total {
					pagination.Visible = pagination.Size
				} else {
					mod_total := pagination.Total % pagination.Size
					pagination.Visible = mod_total
				}
			}
			total_pages := math.Ceil(float64(pagination.Total / pagination.Size))
			pagination.TotalPages = int64(total_pages)
			if pagination.Page == 1 {
				pagination.First = true
				pagination.Last = false
			} else if pagination.Page == pagination.TotalPages {
				pagination.First = false
				pagination.Last = true
			} else {
				pagination.First = false
				pagination.Last = false
			}
			return pagination, nil
		} else {
			return pagination{}, fmt.Errorf("offset have to less than or equals to total")
		}
	} else {
		return pagination{}, fmt.Errorf("size have to less than or equals to total")
	}

}
