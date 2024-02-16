package requests

type TenantCreateRequest struct {
	Name       string
	TenantCode string
}

type TenantUpdateRequest struct {
	Name   string
	Active bool
}
