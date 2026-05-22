package dto

type GetProfile struct {
	ID       int64  `json:"id,omitempty"`
	BranchID int64  `json:"branch_id,omitempty"`
	FullName string `json:"full_name,omitempty"`
	Login    string `json:"login,omitempty"`
	Role     string `json:"role,omitempty"`
}

type CreateProfile struct {
	BranchID int64  `json:"branch_id,omitempty"`
	FullName string `json:"full_name,omitempty"`
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role,omitempty"`
}
