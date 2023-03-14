package main

import (
	"testing"
)

func TestCreatePagination(testing *testing.T) {
	var req_pagination RequestPagination
	req_pagination.Offset = 0
	req_pagination.Size = 10
	req_pagination.Page = 1
	req_pagination.Total = 20
	pagination, err := CreatePagination(req_pagination)
	testing.Log(pagination)
	if err != nil {
		testing.Errorf("error pagination test")
	}
}
