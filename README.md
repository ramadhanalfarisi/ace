# go-pagination
The go-pagination package used to generate pagination response for a much of data result

### Install
`go get github.com/ramadhanalfarisi/go-pagination`

### Simple Example
```go
package main

import (
	"encoding/json"
	"fmt"
	go_pagination "github.com/ramadhanalfarisi/go-pagination"
)


func main(){
	get_pagination := go_pagination.GetPagination(1,10,40)
	create_pagiantion, err := go_pagination.CreatePagination(get_pagination)
	if err != nil {
		panic(err)
	}
	response,_ := json.Marshal(create_pagiantion)
	fmt.Println(response)
}
```
