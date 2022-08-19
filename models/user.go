package models

//type userContextType string
type Operations string

const (
	SetOperations    Operations = "set"
	GetOperations    Operations = "get"
	DeleteOperations Operations = "delete"
	ExitOperations   Operations = "exit"
)

//type User struct {
//	//	UserId   int    `db:"userid" json:"userId"`
//	Name     string `db:"name" json:"Name"`
//	Email    string `db:"email" json:"email"`
//	Password string `db:"password" json:"password"`
//	//CreatedAt time.Time `db:"created_at" json:"createdAt"`
//}
//
//type LoginUser struct {
//	Email    string `db:"email" json:"email"`
//	Password string `db:"password" json:"password"`
//}
//
//type UserContext struct {
//	ID     int         `db:"id"`
//	AuthID string      `db:"auth_id"`
//	Email  null.String `db:"email"`
//	Name   string      `db:"name"`
//	Phone  null.String `db:"phone"`
//}
//
//type Todo struct {
//	Userid    int    `json:"userId"`
//	ID        int    `json:"id"`
//	Title     string `json:"title"`
//	Completed bool   `json:"completed"`
//}
