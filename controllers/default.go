package controllers

import (
	"UserManagementSystem/froms"
	"UserManagementSystem/models"
	"UserManagementSystem/utils"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	InsertDataCmd = `insert into user_info (name,department,addr,sex,salary,phone) values (?,?,?,?,?,?)`
	SelectDataCmd = `select id,name,department,addr,sex,salary,phone from user_info`
	QueryDataCmd  = `select * from user_info where id = ?`
	DeleteDataCmd = `delete from user_info where id = ?`
	UpdateDataCmd = `update user_info set name=?,department=?,addr=?,sex=?,salary=?,phone=? where id=?`
)

var (
	from froms.UserAddFroms
	user models.User
)

type BaseController struct {
	beego.Controller
}

type UserController struct {
	beego.Controller
}

//访问根路径返回
func (c *BaseController) BaseInfo() {
	users := make([]models.User, 0, 10)
	rows, err := models.InitDB().Query(SelectDataCmd)
	if err == nil {
		for rows.Next() {
			var user models.User
			//将数据库中数据扫描到变量中
			err = rows.Scan(&user.Id, &user.Name, &user.Department, &user.Addr, &user.Sex, &user.Salary, &user.Phone)
			if err == nil {
				//将扫描到的数据添加到切片中
				users = append(users, user)
			} else {
				fmt.Println(err)
			}
		}
	}

	c.Data["users"] = users
	c.TplName = "index.html"
}

//添加用户
func (c *UserController) Add() {
	if c.Ctx.Input.IsPost() {
		from = froms.UserAddFroms{}
		if err := c.ParseForm(&from); err == nil {
			//解析用户数据
			name := strings.TrimSpace(from.Name)
			department := strings.TrimSpace(from.Department)
			addr := strings.TrimSpace(from.Addr)
			phone := strings.TrimSpace(from.Phone)
			sex := from.Sex
			salary := from.Salary

			//插入校验数据
			sal := strconv.Itoa(salary)
			utils.DataCheck("name", from.Name, 30)
			utils.DataCheck("department", from.Department, 50)
			utils.DataCheck("addr", from.Addr, 50)
			utils.DataCheck("phone", from.Phone, 11)
			salLength := utf8.RuneCountInString(sal)
			if salLength == 0 {
				utils.Errors["salary"] = "salary不能为空"
			} else if salLength > 11 {
				utils.Errors["salary"] = "salary字段不能超出11位"
			}
			//判断Errors长度为0则代表无错误,插入数据
			fmt.Println(utils.Errors)
			if len(utils.Errors) == 0 {
				if _, err := models.InitDB().Exec(InsertDataCmd, name, department, addr, sex, salary, phone); err == nil {
					fmt.Println("[+]数据插入成功")
					c.Ctx.Redirect(302, "/")
				}
			} else {
				c.Data["Errors"] = utils.Errors
				c.TplName = "add.html"
			}
			utils.Errors = map[string]string{}
		} else {
			fmt.Println(err)
		}
	}
	//Get请求调整到创建页面
	c.TplName = "add.html"
}

//删除用户
func (c *UserController) Delete() {
	from = froms.UserAddFroms{}
	c.ParseForm(&from)
	result, err := models.InitDB().Exec(DeleteDataCmd, from.Id)
	if err == nil {
		fmt.Println(result.RowsAffected())
	} else {
		fmt.Println(err)
	}
	c.Ctx.Redirect(302, "/")
}

//更新用户数据
func (c *UserController) Update() {
	if c.Ctx.Input.IsGet() {
		from = froms.UserAddFroms{}
		c.ParseForm(&from)
		rows := models.InitDB().QueryRow(QueryDataCmd, from.Id)
		err := rows.Scan(&user.Id, &user.Name, &user.Department, &user.Addr, &user.Sex, &user.Salary, &user.Phone)
		if err != nil {
			fmt.Println(err)
		}
		c.Data["User"] = user
		c.TplName = "update.html"

	} else {
		//解析更新数据
		from = froms.UserAddFroms{}
		c.ParseForm(&from)

		id := from.Id
		name := strings.TrimSpace(from.Name)
		department := strings.TrimSpace(from.Department)
		addr := strings.TrimSpace(from.Addr)
		phone := strings.TrimSpace(from.Phone)
		sex := from.Sex
		salary := from.Salary

		fmt.Println(id, name, department, addr, phone, sex, salary)
		//更新数据
		if result, err := models.InitDB().Exec(UpdateDataCmd, name, department, addr, sex, salary, phone, id); err == nil {
			fmt.Println("[+]数据更新成功")
			c.Ctx.Redirect(302, "/")
		} else {
			fmt.Println(result.RowsAffected())
		}
	}
}
