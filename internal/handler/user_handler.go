package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// DeleteUserResponse โครงสร้างสำหรับการตอบกลับเมื่อการลบสำเร็จ
type DeleteUserResponse struct {
	Message string `json:"message"`
	ID      int    `json:"id"`
}

var users = []User{}

// @Summary Get user by ID
// @Description Get details of a user by ID
// @Tags Users
// @Produce  json
// @Param   id   path      int     true  "User ID"
// @Success 200  {object}  User
// @Failure 404  {object}  ErrorResponse
// @Router  /users/{id} [get]
func GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid ID format"})
		return
	}

	// ค้นหาผู้ใช้ตาม ID
	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, ErrorResponse{Message: "User not found"})
}

// @Summary Delete user by ID
// @Description Delete a user by their ID
// @Tags Users
// @Produce  json
// @Param   id   path      int     true  "User ID"
// @Success 200 {object} DeleteUserResponse
// @Failure 404  {object}  ErrorResponse
// @Router  /users/{id} [delete]
func DeleteUserByID(c *gin.Context) {
	idParam := c.Param("id")
	//userFound := false
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid ID format"})
		return
	}

	for index, user := range users {
		if user.ID == id {
			// ลบผู้ใช้
			users = append(users[:index], users[index+1:]...)
			c.JSON(http.StatusOK, DeleteUserResponse{
				Message: "User deleted successfully",
				ID:      id,
			})
			return
		}
	}

	// หากไม่พบผู้ใช้
	c.JSON(http.StatusNotFound, ErrorResponse{Message: "Not found"})
}
