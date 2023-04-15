package controller

// RecipeList receives the information from the body for the request to create a report
type RecipeList struct {
	EmailList []string `json:"email_list"`
}

// ApiResponse is the format to response from report Api
type ApiResponse struct {
	Message      string       `json:"message"`
	Notification Notification `json:"notification,omitempty"`
	Error        string       `json:"error,omitempty"`
}

// Notification shows which emails was valid an invalid to send notification
type Notification struct {
	ValidList   []string `json:"to,omitempty"`
	InvalidList []string `json:"discarded,omitempty"`
}
