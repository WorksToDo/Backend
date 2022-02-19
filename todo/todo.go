package todo

import "fmt"

type Todo struct {
	ID int `json:"id" gorm:"primaryKey;column:id"`
	Task string `json:"task" gorm:"column:task"`
}

func (Todo) TableName() string {
	return fmt.Sprintf("%s.%s",SchemeName,TableName)
}
const(
	SchemeName = "todos"
	TableName = "todos"
)