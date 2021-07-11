package index

// Index is the response object for the index route
type Index struct {
	Alive       bool     `json:"alive"`
	Author      string   `json:"author"`
	Maintainers []string `json:"maintainers,omitempty"`
	Email       string   `json:"email"`
	System      string   `json:"system"`
	Version     string   `json:"version"`
	Environment string   `json:"environment"`
}

