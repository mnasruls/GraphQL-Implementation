package entities

type User struct {
	ID       string    `json:"id" gorm:"column:Id;primary_key;auto_increment"`
	Name     string    `json:"name" gorm:"column:Name"`
	Dob      string    `json:"dob" gorm:"column:DOB"`
	Address  string    `json:"address" gorm:"column:Address"`
	Position *Position `json:"position"`
}
