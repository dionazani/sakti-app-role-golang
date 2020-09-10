package core

type appRoleModel struct {
	Id          int32  `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	IsActive    int8   `json:"isActive"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type myResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    *appRoleModel `json:"data"`
}
