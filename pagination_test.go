package go_pagination

import (
	"encoding/json"
	"testing"
)

func TestCreatePagination(testing *testing.T) {
	req_pagination := GetPagination(1, 10, 20)
	pagination, err := req_pagination.CreatePagination()
	json,_ := json.Marshal(pagination)
	testing.Log(string(json))
	if err != nil {
		testing.Errorf("error pagination test")
	}
}
