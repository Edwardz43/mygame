package gameserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type member struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}

func login(c *gin.Context) {
	var m *member
	c.BindJSON(&m)
	fmt.Printf("name[%v], pw[%v]\n", m.Name, m.Password)
	memberID, err := memberService.Login(m.Name)
	if err != nil {
		fmt.Println(err)
		c.JSON(401, err.Error())
	}
	c.JSON(200, memberID)
}

func register(c *gin.Context) {
	name := c.PostForm("username")
	mail := c.PostForm("email")
	pw := c.PostForm("password")

	ok, err := memberService.Register(name, mail, pw)
	if !ok && err != nil {
		fmt.Println(err)
		c.JSON(401, err.Error())
	}

	c.JSON(200, ok)
}
