package models

import (
	"errors"
	"time"

	"github.com/team/swe-project/middleware"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

type User struct {
	UserID       uint8     `json:"user_id" gorm:"primary_key"`
	FirstName    string    `json:"first_name,omitempty"`
	LastName     string    `json:"last_name,omitempty"`
	Email        string    `json:"email,omitempty"`
	Password     string    `json:"password,omitempty"`
	ProfilePic   string    `json:"profile_pic,omitempty"`
	Dark         bool      `json:"dark,omitempty"`
	Likes        uint8     `json:"likes,omitempty"`
	CreationDate time.Time `gorm:"-"`
}

func GetAllUsers() []User {
	var users []User

	middleware.DB.Find(&users)

	return users
}

func GetUserByID(id uint8) (User, error) {
	var user User

	err := middleware.DB.First(&user, id).Error

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func CreateUser(user User) (User, error) {
	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, errors.New("could not hash password")
	}

	user.Password = string(hashedPassword)
	result := middleware.DB.Where("email = ?", user.Email).FirstOrCreate(&user)

	// User already exists
	if result.RowsAffected == 0 {
		return User{}, errors.New("User Already Exists")
	}

	return user, nil
}

func DeleteUser(user_id uint8) (User, error) {

	user, err := GetUserByID(user_id)

	if err != nil {
		return user, err
	}

	for _, userClass := range GetAllUserClassRowsFromUser(user) {
		deletedUserClass := middleware.DB.Unscoped().Where("user_class_id = ?", userClass.UserClassID).Delete(&UserClasses{})

		if deletedUserClass.Error != nil {
			return user, deletedUserClass.Error
		}
	}

	for _, userMajor := range GetAllUserMajorsFromUser(user) {
		deletedUserMajor := middleware.DB.Unscoped().Where("user_major_id = ?", userMajor.UserMajorID).Delete(&UserMajors{})

		if deletedUserMajor.Error != nil {
			return user, deletedUserMajor.Error
		}
	}

	for _, userRole := range GetAllUserRoleRowsFromUser(user) {
		deletedUserRole := middleware.DB.Unscoped().Where("user_role_id = ?", userRole.UserRoleID).Delete(&UserRoles{})

		if deletedUserRole.Error != nil {
			return user, deletedUserRole.Error
		}
	}

	for _, thread := range user.GetThreads() {
		err := middleware.DB.Model(&thread).Clauses(clause.Returning{}).Where("thread_id = ?", thread.ThreadID).Updates(Thread{UserID: 0})

		if err.Error != nil {
			return user, err.Error
		}
	}

	for _, post := range user.GetPosts() {
		err := middleware.DB.Model(&post).Clauses(clause.Returning{}).Where("post_id = ?", post.PostID).Updates(Post{UserID: 0})

		if err.Error != nil {
			return user, err.Error
		}
	}

	result := middleware.DB.Where("user_id = ?", user.UserID).Delete(&User{})

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func Update() {
	// figure out criteria for that later, gonna be multiple update functions
}

func CheckSignIn(email string, password string) (User, error) {
	var user User

	err := middleware.DB.First(&user, "email = ?", email).Error

	// Database Error
	if err != nil {
		return User{}, err
	}

	// Check Password
	passwordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	// Invalid login
	if passwordErr != nil {
		return User{}, passwordErr
	}

	return user, nil
}

func (u *User) GetThreads() []Thread {
	var threads []Thread

	for _, thread := range GetAllThreads() {
		if thread.UserID == u.UserID {
			threads = append(threads, thread)
		}
	}

	return threads
}

func (u *User) GetPosts() []Post {
	var posts []Post

	for _, post := range GetAllPosts() {
		if post.UserID == u.UserID {
			posts = append(posts, post)
		}
	}

	return posts
}

func (u *User) GetClasses() []Class {
	var classes []Class

	for _, userClass := range GetAllUserClassRows() {
		if userClass.UserID == u.UserID {
			class, err := GetClassByID(userClass.ClassID)
			if err == nil {
				classes = append(classes, class)
			}
		}
	}

	return classes
}

func (u *User) GetRoles() []Role {
	var roles []Role

	for _, userRole := range GetAllUserRoleRows() {
		if userRole.UserID == u.UserID {
			role, err := GetRoleByID(userRole.RoleID)
			if err == nil {
				roles = append(roles, role)
			}
		}
	}

	return roles
}

func (u *User) GetMajors() []Major {
	var majors []Major

	for _, userMajor := range GetAllUserMajorRows() {
		if userMajor.UserID == u.UserID {
			major, err := GetMajorByID(userMajor.MajorID)
			if err == nil {
				majors = append(majors, major)
			}
		}
	}

	return majors
}
