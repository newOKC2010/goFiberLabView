package resultLabUtils

type LabResultRequest struct {
	CID       string `json:"cid"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type LabResultItem struct {
	PtName      string  `json:"pt_name"`
	OrderDate   string  `json:"order_date"`
	LabItemName string  `json:"lab_items_name"`
	LabResult   *string `json:"lab_order_result"`
}

type LabItem struct {
	LabItemName    string  `json:"lab_items_name"`
	LabResult      *string `json:"lab_order_result"`
	NormalValueRef *string `json:"lab_items_normal_value_ref"`
}

type LabCategory struct {
	GroupName string    `json:"group_name"`
	Items     []LabItem `json:"items"`
}

type LabGroupedResult struct {
	PtName    string        `json:"-"`
	OrderDate string        `json:"order_date"`
	Groups    []LabCategory `json:"groups"`
}

type LabResultResponse struct {
	Success bool               `json:"success"`
	Total   int                `json:"total"`
	PtName  string             `json:"pt_name"`
	Results []LabGroupedResult `json:"results"`
}
