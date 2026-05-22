// Package dto
package dto

type GetBranch struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CreateBranch struct {
	Name string `json:"name,omitempty"`
}
