package domain

type Users struct {
	Id          int64  `json:"id"`
	UserName    string `json:"user_name"`
	Age         int    `json:"age"`
	AvatarUrl   string `json:"avatar_url"`
	Role        int    `json:"role"`
	IsActive    int    `json:"is_active"`
	ExpiredTime int    `json:"expired_time"`
	CreatedAt   int    `json:"created_at"`
	UpdatedAt   int    `json:"updated_at"`
}

type UsersReqByForm struct {
	Id          int64  `form:"id"`
	UserName    string `form:"user_name"`
	Age         int    `form:"age"`
	AvatarUrl   string `form:"avatar_url"`
	Role        int    `form:"role"`
	IsActive    int    `form:"is_active"`
	ExpiredTime int    `form:"expired_time"`
	CreatedAt   int    `form:"created_at"`
	UpdatedAt   int    `form:"updated_at"`
}
type UserUpdate struct {
	UserName    string `json:"user_name"`
	Age         int    `json:"age"`
	AvatarUrl   string `json:"avatar_url"`
	Role        int    `json:"role"`
	IsActive    int    `json:"is_active"`
	ExpiredTime int    `json:"expired_time"`
	UpdatedAt   int    `json:"updated_at"`
}

type RepositoryUser interface {
	AddUser(user *Users) error
	GetAllUserStaffs(user *UsersReqByForm) ([]*Users, error)
	DeleteUserByUsernameStaff(userName string) error
	UpdateUserByUsernameStaff(user *UserUpdate) error
}
