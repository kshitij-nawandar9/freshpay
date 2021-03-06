package base

type Model struct {
	ID        string `gorm:"type:varchar(20)"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type Error struct {
	Code        string
	Description string
	Source      string
	Reason      string
	Step        string
	Metadata    string
}
type Failure struct {
	Error Error
}
