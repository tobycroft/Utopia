package LogRaffleModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "log_raffle"

func Api_insert(raffleid float64, roomid float64, title string, typ string, asset_animation_pic string, asset_tips_pic string, json string) bool {
	db := tuuz.Db().Table(table)
	data := make(map[string]interface{})
	data["raffleid"] = raffleid
	data["roomid"] = roomid
	data["title"] = title
	data["type"] = typ
	data["asset_animation_pic"] = asset_animation_pic
	data["asset_tips_pic"] = asset_tips_pic
	data["json"] = json
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_find(raffleid float64) gorose.Data {
	db := tuuz.Db().Table(table)
	where := make(map[string]interface{})
	where["raffleid"] = raffleid
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select(raffleid float64) []gorose.Data {
	db := tuuz.Db().Table(table)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
