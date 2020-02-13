package BlocksModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "cos_blocks"

func Api_insert(height float64, time float64, chain_id float64, block_hash string, raw string) bool {
	db := tuuz.Db().Table(table)
	data := make(map[string]interface{})
	data["height"] = height
	data["time"] = time
	data["chain_id"] = chain_id
	data["block_hash"] = block_hash
	data["raw"] = raw
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_find(height float64) gorose.Data {
	db := tuuz.Db().Table(table)
	where := make(map[string]interface{})
	where["height"] = height
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		return nil
	} else {
		return ret
	}
}

func Api_find_last() gorose.Data {
	db := tuuz.Db().Table(table)
	db.OrderBy("height desc")
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select(limit int, page int) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Limit(limit)
	db.Page(page)
	ret, err := db.Get()
	if err != nil {
		return nil
	} else {
		return ret
	}
}
