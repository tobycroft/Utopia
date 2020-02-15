package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/index/model/AddressModel"
	"main.go/app/v1/index/model/UserModel"
	"main.go/extend/CosMos/CosCore"
	"main.go/tuuz"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/RET"
)

func IndexController(route *gin.RouterGroup) {
	route.Any("/", index)
	route.Any("/register", register)
}

func index(c *gin.Context) {
	c.String(0, "index")
}

func register(c *gin.Context) {
	username, is := c.GetPostForm("username")
	if is == false {
		c.JSON(200, RET.Ret_fail(400, "username"))
		c.Abort()
		return
	}
	password, is := c.GetPostForm("password")
	if is == false {
		c.JSON(200, RET.Ret_fail(400, "password"))
		c.Abort()
		return
	}
	UserModel.Db = tuuz.Db()
	if len(UserModel.Api_find_byUsername(username)) > 0 {
		c.JSON(200, RET.Ret_fail(400, "用户名已经被注册"))
		c.Abort()
		return
	}
	ret, err := CosCore.NewAccount(username)
	if err != nil {
		c.JSON(200, RET.Ret_fail(500, "出现错误："+err.Error()))
		c.Abort()
		return
	} else {
		data, err := Jsong.JObject(ret)
		if err != nil {
			c.JSON(200, RET.Ret_fail(500, "数据解析错误："+err.Error()))
			c.Abort()
			return
		} else {
			name := data["name"].(string)
			typ := data["type"].(string)
			address := data["address"].(string)
			pubkey := data["pubkey"].(string)
			mnemonic := data["mnemonic"].(string)
			db := tuuz.Db()
			UserModel.Db, AddressModel.Db = db, db
			db.Begin()
			if UserModel.Api_insert(username, password, "", address) != true {
				db.Rollback()
			} else {
				if AddressModel.Api_insert(name, typ, address, pubkey, mnemonic, ret) != true {
					db.Rollback()
				}
			}
			db.Commit()
		}
		c.JSON(200, RET.Ret_succ(0, data))
	}

}
