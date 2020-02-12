package UserMemberModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "user_member"

func Api_insert(username string, password string, mid float64, access_token string, refresh_token string, cookie string, expire float64) bool {
	db := tuuz.Db().Table(table)
	data := make(map[string]interface{})
	data["username"] = username
	data["password"] = password
	data["mid"] = mid
	data["access_token"] = access_token
	data["refresh_token"] = refresh_token
	data["cookie"] = cookie
	data["expire"] = expire
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	}
	return true
}

func Api_find(username string) gorose.Data {
	db := tuuz.Db().Table(table)
	where := make(map[string]interface{})
	where["username"] = username
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	}
	return ret
}

func Api_select() []gorose.Data {
	db := tuuz.Db().Table(table)
	ret, _ := db.Get()
	return ret
}

func Api_update_by_username(username string, password string, mid float64, access_token string, refresh_token string, cookie string, expire float64) bool {
	db := tuuz.Db().Table(table)
	where := make(map[string]interface{})
	where["username"] = username
	db.Where(where)
	data := make(map[string]interface{})
	data["password"] = password
	data["mid"] = mid
	data["access_token"] = access_token
	data["refresh_token"] = refresh_token
	data["cookie"] = cookie
	data["expire"] = expire
	db.Data(data)
	_, err := db.Update()

	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	}
	return true
}
