package entities

type Position struct {
	ID          string `json:"id" gorm:"column:Id;primary_key;auto_increment"`
	UserID      string `json:"user_id" gorm:"column:UserId"`
	PostionName string `json:"postionName" gorm:"column:PositionName"`
	Salary      string `json:"salary" gorm:"column:Salary"`
}
