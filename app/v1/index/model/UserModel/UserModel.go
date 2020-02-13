package UserModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "cos_user"

func Api_insert(username string, password string, paypass string) bool {
	db := tuuz.Db().Table(table)
	data := make(map[string]interface{})
	data["username"] = username
	data["password"] = password
	data["paypass"] = paypass
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_find_byUsername(username string) gorose.Data {
	db := tuuz.Db().Table(table)
	where := make(map[string]interface{})
	where["username"] = username
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find(username string, password string) gorose.Data {
	db := tuuz.Db().Table(table)
	where := make(map[string]interface{})
	where["username"] = username
	where["password"] = password
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find_byPaypass(username string, paypass string) gorose.Data {
	db := tuuz.Db().Table(table)
	where := make(map[string]interface{})
	where["username"] = username
	where["paypass"] = paypass
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
