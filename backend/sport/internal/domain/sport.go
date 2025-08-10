package domain

type Section struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Info     string `json:"info"`
	Schedule string `json:"schedule"`
}

// type Sections struct {
// 	SectionArray []string `json:"sections"`
// }
