package AddressModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "cos_address"

var Db = tuuz.Db()

func Api_insert(name string, typ string, address string, pubkey string, mnemonic string, raw string) bool {
	db := Db.Table(table)
	data := make(map[string]interface{})
	data["name"] = name
	data["type"] = typ
	data["address"] = address
	data["pubkey"] = pubkey
	data["mnemonic"] = mnemonic
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

func Api_find(address string) gorose.Data {
	db := Db.Table(table)
	where := make(map[string]interface{})
	where["address"] = address
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
