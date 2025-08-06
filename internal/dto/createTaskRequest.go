package dto

type CreateTaskRequest struct {
	Title            string `json:"title"`
	Description      string `json:"description"`
	ProjectID        string `json:"projectId"`
	UserID           string `json:"userId"`
	StartDate        string `json:"startDate"`
	EstimatedEndDate string `json:"estimatedEndDate"`
}
