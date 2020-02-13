package BlocksModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "cos_blocks"

func Api_insert(height int, time string, chain_id string, block_hash string, from_address string, to_address string, memo string, amount string, fee string, raw string) bool {
	db := tuuz.Db().Table(table)
	data := make(map[string]interface{})
	data["height"] = height
	data["time"] = time
	data["chain_id"] = chain_id
	data["block_hash"] = block_hash
	data["from_address"] = from_address
	data["to_address"] = to_address
	data["memo"] = memo
	data["amount"] = amount
	data["fee"] = fee
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
	db.OrderBy("id desc")
	ret, err := db.Get()
	if err != nil {
		return nil
	} else {
		return ret
	}
}

func Api_select_byToAddress(to_address string, limit int, page int) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := make(map[string]interface{})
	where["to_address"] = to_address
	db.Where(where)
	db.Limit(limit)
	db.Page(page)
	db.OrderBy("id desc")
	ret, err := db.Get()
	if err != nil {
		return nil
	} else {
		return ret
	}
}
