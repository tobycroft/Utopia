package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gohouse/gorose"

	"os/exec"
)

func IndexController(route *gin.RouterGroup) {
	route.Any("/", index)
	route.Any("/login", loginss)
	route.Any("/transfer", transfer)
}

func index(c *gin.Context) {
	c.String(0, "index")
}

func loginss(c *gin.Context) {
	password := c.Query("password")
	username := c.Query("username")
	json := map[string]string{}
	json["username"] = username
	json["password"] = password
	gorose.Open()
	c.JSON(0, json)
}

func transfer(c *gin.Context) {
	cmd := exec.Command("ltcli.exe", "tx", "send", "cosmos13v60v23sheck50em6jlvdqmmgkmp2n0qqrchsv",
		"cosmos19yfkv45mlly4n2u8609w6vda678kxgphk60q6t", "1stake", "--chain-id=lt",
		"--memo", "dhu2387e8n2x 2yr 2o38r ow e f", "-y", "-o", "json")
	buf, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(buf))
	//ret,err:= Jsong.JObject(string(buf))
	c.String(200, string(buf))
}
