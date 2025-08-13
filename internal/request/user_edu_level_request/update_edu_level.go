package useredulevelrequest

type UpdateMultiEduLevelRequest struct {
	Levels []UpdateEduLevelRequest `json:"levels"`
}

type UpdateEduLevelRequest struct {
	ID           uint   `json:"id"`
	UserID       uint   `json:"user_id"`
	EduLevel     string `json:"edu_level"`
	StudentCount int    `json:"student_count"`
	EduYear      int    `json:"edu_year"`
}
