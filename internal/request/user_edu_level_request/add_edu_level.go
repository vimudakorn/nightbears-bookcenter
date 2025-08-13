package useredulevelrequest

type AddMultiEduLevelRequest struct {
	Levels []AddEduLevelRequest `json:"levels"`
}

type AddEduLevelRequest struct {
	UserID       uint   `json:"user_id"`
	EduLevel     string `json:"edu_level"`
	StudentCount int    `json:"student_count"`
	EduYear      int    `json:"edu_year"`
}
