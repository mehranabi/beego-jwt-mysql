package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// Id will automatically considered as Auto Increment Key
// Unique string (Email field) cannot have the default size (that's too long)
// Created field will be updated at every save
// Updated field will be updated at the first save

// User model struct (database & response)
type User struct {
	Id			int64		`json:"id"`
	Email		string		`json:"email" orm:"unique;index;size(191)"`
	Password	string		`json:"-"`
	Name		string		`json:"name"`
	Created		time.Time	`json:"created_on" orm:"auto_now_add;type(datetime)"`
	Updated		time.Time	`json:"updated_on" orm:"auto_now;type(datetime)"`
}

// User model struct (input object)
type InputUser struct {
	Email		string	`json:"email"`
	Password	string	`json:"password"`
	Name		string	`json:"name"`
}

// Define basic credentials struct
type BasicCredentials struct {
	Email		string	`json:"email"`
	Password	string	`json:"password"`
}

func init() {
	// Register this model for database
	orm.RegisterModel(new(User))
}

// Custom table name
func (u *User) TableName() string {
	return "users"
}

func IndexAll() (users []User, err error) {
	// New ORM object
	o := orm.NewOrm()

	// Define empty users
	var us []User

	// Query table (Just get NAME and EMAIL) - Limit to 50 rows
	count, e := o.QueryTable(new(User)).Limit(50).All(&us, "Name", "Email")
	if e != nil {
		return nil, e
	}

	if count <= 0 {
		return nil, errors.New("nothing found")
	}

	return us, nil
}

// Create a new user
func CreateNew(email, password, name string) (id int64, err error) {
	// New ORM object
	o := orm.NewOrm()

	// Calculate password hash to save in database
	passHash, hashErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if hashErr != nil {
		return -1, errors.New("cannot generate hash from password")
	}

	// Init User and set data
	user := User{}
	user.Email = email
	user.Password = string(passHash)
	user.Name = name

	// Insert object to database
	uId, insertErr := o.Insert(&user)
	if insertErr != nil {
		return -1, errors.New("failed to insert user to database")
	}

	// Return result
	return uId, nil
}

// Find an user by ID
func FindById(id int64) (user *User, err error)  {
	// New ORM object
	o := orm.NewOrm()

	// Init user with Id
	u := User{Id: id}

	// Read from database
	e := o.Read(&u)

	// Check for errors
	if e == orm.ErrNoRows {
		return nil, errors.New("user not found")
	} else if e == nil {
		return &u, nil
	} else {
		return nil, errors.New("unknown error occurred")
	}
}

// Find an user by email
func FindByEmail(email string) (user *User, err error)  {
	// New ORM object
	o := orm.NewOrm()

	// Init user with Email
	u := User{Email: email}

	// Read from database
	e := o.Read(&u, "Email")

	// Check for errors
	if e == orm.ErrNoRows {
		return nil, errors.New("user not found")
	} else if e == nil {
		return &u, nil
	} else {
		return nil, errors.New("unknown error occurred")
	}
}

// Login method for user
func Login(email, password string) (user *User, err error) {
	// Get user
	u, e := FindByEmail(email)

	// Check for errors
	if e == nil {
		// No error -> Check credentials
		if pErr := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); pErr != nil {
			return nil, errors.New("email and password doesn't match")
		}
		return u, nil
	} else {
		// Return error
		return nil, e
	}
}
