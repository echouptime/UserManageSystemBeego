package froms

type UserAddFroms struct {
	Id         int    `form:"id" json:"id,omitempty"`
	Name       string `form:"name" json:"name,omitempty"`
	Department string `form:"department" json:"department,omitempty"`
	Addr       string `form:"addr" json:"addr,omitempty"`
	Sex        int    `form:"sex" json:"sex,omitempty"`
	Phone      string `form:"phone" json:"phone,omitempty"`
	Salary     int    `form:"salary" json:"salary,omitempty"`
}
