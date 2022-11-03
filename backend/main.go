package main

import (
   "github.com/gin-gonic/gin"
   "net/http"
)

type project struct {
	ID		string `json:"id"`
	Name    string `json:"name"`
	Authors string `json:"authors"`
	Rating  int    `json:"rating"`
	//time added probs implemented via external lib
}

func getProject(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, projects)
}

func postProject (c *gin.Context) {
	var newProject project

	if err := c.BindJSON(&newProject); err != nil {
		return
	}
	projects = append(projects, newProject)
	c.IndentedJSON(http.StatusCreated, newProject)
}

var projects = []project{
}

func main() {
	router := gin.Default()
	router.GET("/projects", getProject)
	router.POST("/projects", postProject) //first argument can be any address
	router.Run("localhost:8080")
}
