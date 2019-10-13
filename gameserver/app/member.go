package gameserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	// pw, exist := c.Get("username")
	name := c.PostForm("username")
	pw := c.PostForm("password")
	fmt.Printf("name[%v], pw[%v]\n", name, pw)

	memberID, err := memberService.Login(name)

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
