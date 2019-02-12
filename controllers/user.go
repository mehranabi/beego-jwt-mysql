package controllers

import (
	"beego_jwt_sample/models"
	"beego_jwt_sample/services"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type UserController struct {
	beego.Controller
}

// Define a struct to return when user is authorized
type AuthorizedResponse struct {
	Message string			`json:"message"`
	User    *models.User	`json:"user"`
	Token   string			`json:"token"`
}

// Define a struct to return when there is an error
type ErrorResponse struct {
	Message string `json:"message"`
}

// @Title Register User
// @Description Register a new User in system
// @Param	user	body	{InputUser}	true	"User initial data"
// @router /register [post]
func (cont *UserController) RegisterUser() {
	// Parse input data
	var iu models.InputUser
	_ = json.Unmarshal(cont.Ctx.Input.RequestBody, &iu)

	fmt.Println("Input User: ", iu)

	// Create User
	id, err := models.CreateNew(iu.Email, iu.Password, iu.Name)
	if err != nil {
		// Return response
		errResponse := ErrorResponse{
			Message: err.Error(),
		}
		cont.Data["json"] = errResponse
		cont.ServeJSON()
		cont.StopRun()
	}

	// Get User from database
	user, err := models.FindById(id)
	if err != nil {
		errResponse := ErrorResponse{
			Message: err.Error(),
		}
		cont.Data["json"] = errResponse
		cont.ServeJSON()
		cont.StopRun()
	}

	// Get Token
	token, err := services.MakeToken(id)
	if err != nil {
		errResponse := ErrorResponse{
			Message: err.Error(),
		}
		cont.Data["json"] = errResponse
		cont.ServeJSON()
		cont.StopRun()
	}

	// Return result
	successRes := AuthorizedResponse{
		Message: "User created successfully",
		User:    user,
		Token:   token,
	}
	cont.Data["json"] = successRes
	cont.ServeJSON()
}

// @Title Login User
// @Description Log in an existing User with credentials
// @Param	credentials	body	{BasicCredentials}	true	"User credentials"
// @router /login [post]
func (cont *UserController) LoginUser() {
	// Parse input data
	var credentials models.BasicCredentials
	_ = json.Unmarshal(cont.Ctx.Input.RequestBody, &credentials)

	// Try to login
	user, err := models.Login(credentials.Email, credentials.Password)
	if err != nil {
		errResponse := ErrorResponse{
			Message: err.Error(),
		}
		cont.Data["json"] = errResponse
		cont.ServeJSON()
		cont.StopRun()
	}

	// Get Token
	token, err := services.MakeToken(user.Id)
	if err != nil {
		errResponse := ErrorResponse{
			Message: err.Error(),
		}
		cont.Data["json"] = errResponse
		cont.ServeJSON()
		cont.StopRun()
	}

	// Return result
	successRes := AuthorizedResponse{
		Message: "User logged in successfully",
		User:    user,
		Token:   token,
	}
	cont.Data["json"] = successRes
	cont.ServeJSON()
}

// @Title Index Users
// @Description Index all of users when request is authorized
// @Param	authorization	header	string	true	"Authorization Token"
// @router /all [get]
func (cont *UserController) IndexAll(authorization string) {
	// Get token
	token := authorization[strings.IndexByte(authorization, ' ') + 1:]

	// Validate token (now, just user with ID=1 will be validated)
	valid, err := services.ValidateToken(token, 1)
	if err != nil {
		errResponse := ErrorResponse{
			Message: err.Error(),
		}
		cont.Data["json"] = errResponse
		cont.ServeJSON()
		cont.StopRun()
	}

	if !valid {
		errResponse := ErrorResponse{
			Message: "authentication token is not valid",
		}
		cont.Data["json"] = errResponse
		cont.ServeJSON()
		cont.StopRun()
	}

	// Get all of users
	users, err := models.IndexAll()
	if err != nil {
		errResponse := ErrorResponse{
			Message: err.Error(),
		}
		cont.Data["json"] = errResponse
		cont.ServeJSON()
		cont.StopRun()
	}

	// No error - Show response (users)
	/////// If you see, the ID and Created_on and Updated_on doesnt have real values
	////// because we just get NAME and EMAIL field from database
	cont.Data["json"] = users
	cont.ServeJSON()
}
