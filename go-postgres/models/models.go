package models

// User schema of the user table
type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Age      int64  `json:"age"`
}

// Skill schema of the skills table
type Skill struct {
	ID          int64  `json:"id"`
	Major       string `json:"major"`
	Minor       string `json:"minor"`
	Sub         string `json:"sub"`
	Description string `json:"description"`
}
