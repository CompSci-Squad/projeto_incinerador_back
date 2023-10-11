package services

type OrganizationServicesImpl interface {
	Store(body *dtos.OrganizationStoreRequest) (*dtos.OrganizationStoreResponse, error)
	Index(payload *dtos.OrganizationIndexRequest) (*dtos.OrganizationIndexResponse, error)
	Show(id string) (*dtos.OrganizationShowResponse, error)
	Update(payload *dtos.OrganizationUpdateRequest) (*dtos.OrganizationUpdateResponse, error)
	Remove(payload *dtos.PlanRequest) (*dtos.OrganizationUpdateResponse, error)
}
