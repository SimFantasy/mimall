package models

type Manager struct {
	Id        int
	Username  string
	Password  string
	Mobile    int
	Email     string
	Status    int
	RolaId    int
	CreatedAt int
	IsSuper   int
}

func (Manager) TableName() string {
	return "manager"
}
