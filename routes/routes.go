package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nathancordeiro0/api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/students", controllers.GetStudents)
	r.GET("/students/:id", controllers.GetStudentById)
	r.POST("/students", controllers.CreateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.EditStudent)
	r.GET("/students/cpf/:cpf", controllers.SearchStudentByCPF)

	r.Run()
}
