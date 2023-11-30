package action

type User struct {
	ID         uint   `json:"id"`
	UserName   string `json:"username"`
	Password   string `json:"password"`
	NickName   string `json:"nickname"`
	Face       string `json:"face"`
	Sex        uint   `json:"sex"`
	Phone      string `json:"phone"`
	CreateTime int32  `json:"create_time"`
	UpdateTime int32  `json:"update_time"`
}
