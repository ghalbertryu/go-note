package _range

import "strconv"

type Dto struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func buildCollection() []Dto {
	dtos := make([]Dto, 0)
	for i := 0; i < 10; i++ {
		dtos = append(dtos, Dto{
			Id:   i,
			Name: "aaa" + strconv.Itoa(i),
		})
	}
	return dtos
}
