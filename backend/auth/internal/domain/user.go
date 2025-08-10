package domain

import "time"

type User struct {
	ID           int       `json:"id"`
	Login        string    `json:"login"`
	Name         string    `json:"name"`
	Surname      string    `json:"surname"`
	Patronymic   string    `json:"patronymic"`
	Password     string    `json:"password"`
	Role         string    `json:"role"`
	Section      string    `json:"section"`
	Section_id   int       `json:"section_id"`
	StudentGroup string    `json:"student_group"`
	Visits       int       `json:"visits"`
	Paid         bool      `json:"paid"`
	Last_scanned time.Time `json:"last_scanned"`
	QrCode       string    `json:"qrcode"`
}

type LoginInfo struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RoleInfo struct {
	Role string `json:"role"`
}