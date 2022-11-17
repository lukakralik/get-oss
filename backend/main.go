package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/golang-module/carbon/v2"
	_"encoding/json"
)

type project struct {
	ID		int			   `json:"id"`
	Name    string         `json:"name"`
	Author	string	       `json:"author"`
	Rating  float32		   `json:"rating"`
	Time	carbon.Carbon  `json:"time"`
}

func getProject(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, projects)
}

func postProject (c *gin.Context) {
	var newProject project //initialize a new entry to our array
	newProject.Time = carbon.Now() //specify the time parameter

	if err := c.BindJSON(&newProject); err != nil {
		return
	}
	projects = append(projects, newProject)
	c.IndentedJSON(http.StatusCreated, newProject) //return http status and json of entry
}

var projects = []project{
}

func main() {
	router := gin.Default()
	router.GET("/projects", getProject)
	router.POST("/projects", postProject) //first argument can be any address
	router.Run("localhost:8080")
}
