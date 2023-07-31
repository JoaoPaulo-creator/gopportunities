package handler

import (
	"fmt"
)

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	IsRemote *bool  `json:"isRemote"`
	RoleLink string `json:"roleLink"`
	Salary   int64  `json:"salary"`
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.IsRemote == nil && r.Salary <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.IsRemote == nil {
		return errParamIsRequired("isRemote", "boolean")
	}
	if r.RoleLink == "" {
		return errParamIsRequired("roleLink", "string")
	}

	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}

	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	IsRemote *bool  `json:"isRemote"`
	RoleLink string `json:"roleLink"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Role != "" || r.Company != "" || r.Location != "" || r.IsRemote != nil || r.RoleLink != "" || r.Salary > 0 {
		return nil
	}

	return fmt.Errorf("al least one valid field must be provided")
}
