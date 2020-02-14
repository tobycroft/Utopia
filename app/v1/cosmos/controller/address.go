package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/cosmos/model/BlocksModel"
	"main.go/tuuz/Calc"
	"main.go/tuuz/RET"
)

func AddressController(route *gin.RouterGroup) {
	route.Any("/", index)
	route.Any("/to", to)
	route.Any("/from", from)
	route.Any("/to_clean", to_clean)
	route.Any("/from_clean", from_clean)
	route.Any("/to_memo", to_memo)
}

func to(c *gin.Context) {
	address, had := c.GetPostForm("address")
	if had == false {
		c.JSON(200, RET.Fail(404, "address"))
		c.Abort()
		return
	}
	limit, had := c.GetPostForm("limit")
	if had == false {
		limit = "30"
	}
	page, had := c.GetPostForm("page")
	if had == false {
		page = "1"
	}
	to := BlocksModel.Api_select_byToAddress(address, Calc.Any2Int(limit), Calc.Any2Int(page))
	c.JSON(200, to)
}

func to_clean(c *gin.Context) {
	address, had := c.GetPostForm("address")
	if had == false {
		c.JSON(200, RET.Fail(404, "address"))
		c.Abort()
		return
	}
	limit, had := c.GetPostForm("limit")
	if had == false {
		limit = "30"
	}
	page, had := c.GetPostForm("page")
	if had == false {
		page = "1"
	}
	to := BlocksModel.Api_select_byToAddress(address, Calc.Any2Int(limit), Calc.Any2Int(page))
	if len(to) > 0 {
		for i, data := range to {
			delete(data, "raw")
			to[i] = data
		}
	}
	c.JSON(200, to)
}

func from(c *gin.Context) {
	address, had := c.GetPostForm("address")
	if had == false {
		c.JSON(200, RET.Fail(404, "address"))
		c.Abort()
		return
	}
	limit, had := c.GetPostForm("limit")
	if had == false {
		limit = "30"
	}
	page, had := c.GetPostForm("page")
	if had == false {
		page = "1"
	}
	from := BlocksModel.Api_select_byFromAddress(address, Calc.Any2Int(limit), Calc.Any2Int(page))
	c.JSON(200, from)
}

func from_clean(c *gin.Context) {
	address, had := c.GetPostForm("address")
	if had == false {
		c.JSON(200, RET.Fail(404, "address"))
		c.Abort()
		return
	}
	limit, had := c.GetPostForm("limit")
	if had == false {
		limit = "30"
	}
	page, had := c.GetPostForm("page")
	if had == false {
		page = "1"
	}
	from := BlocksModel.Api_select_byFromAddress(address, Calc.Any2Int(limit), Calc.Any2Int(page))
	if len(from) > 0 {
		for i, data := range from {
			delete(data, "raw")
			from[i] = data
		}
	}
	c.JSON(200, from)
}

func to_memo(c *gin.Context) {
	address, had := c.GetPostForm("address")
	if had == false {
		c.JSON(200, RET.Fail(404, "address"))
		c.Abort()
		return
	}
	limit, had := c.GetPostForm("limit")
	if had == false {
		limit = "30"
	}
	page, had := c.GetPostForm("page")
	if had == false {
		page = "1"
	}
	to := BlocksModel.Api_select_byToAddress(address, Calc.Any2Int(limit), Calc.Any2Int(page))
	arr := []map[string]interface{}{}
	if len(to) > 0 {
		for i, data := range to {
			delete(data, "raw")
			arr = append(arr, map[string]interface{}{"height": data["height"], "memo": data["memo"], "time": data["time"]})
			to[i] = data
		}
	}
	c.JSON(200, arr)
}
