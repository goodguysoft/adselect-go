package adselect

type ServiceResult struct {
	Success bool        `json:"Success"`
	Data    interface{} `json:"Data,omitempty"`
	Errors  []string    `json:"Errors,omitempty"`
}
