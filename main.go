package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

var people []Person

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":  "People List",
			"people": people,
		})
	})

	router.GET("/add", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add.html", gin.H{
			"title": "Add Person",
		})
	})

	router.POST("/create", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.PostForm("id"))
		firstName := c.PostForm("first_name")
		lastName := c.PostForm("last_name")
		email := c.PostForm("email")
		person := Person{id, firstName, lastName, email}
		people = append(people, person)
		c.Redirect(http.StatusFound, "/")
	})

	router.GET("/edit", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("id"))
		var person Person
		for _, p := range people {
			if p.ID == id {
				person = p
				break
			}
		}
		c.HTML(http.StatusOK, "edit.html", gin.H{
			"title":  "Edit Person",
			"person": person,
		})
	})

	router.POST("/update", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.PostForm("id"))
		firstName := c.PostForm("first_name")
		lastName := c.PostForm("last_name")
		email := c.PostForm("email")
		for i, p := range people {
			if p.ID == id {
				people[i] = Person{id, firstName, lastName, email}
				break
			}
		}
		c.Redirect(http.StatusFound, "/")
	})

	router.GET("/delete", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("id"))
		var index int
		for i, p := range people {
			if p.ID == id {
				index = i
				break
			}
		}
		people = append(people[:index], people[index+1:]...)
		c.Redirect(http.StatusFound, "/")
	})

	router.Run()
}
