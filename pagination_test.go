package go_pagination

import (
	"testing"
)

func TestCreatePagination(testing *testing.T) {
	req_pagination := GetPagination(1, 10, 20)
	pagination, err := CreatePagination(req_pagination)
	testing.Log(pagination)
	if err != nil {
		testing.Errorf("error pagination test")
	}
}
