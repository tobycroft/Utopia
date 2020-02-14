package CosCore

import (
	"fmt"
	"main.go/config/app_conf"
	Calc2 "main.go/tuuz/Calc"
	"main.go/tuuz/Net"
	Calc "main.go/tuuz/Str"
	"os/exec"
	"time"
)

func Blocks(blocks int) (string, error) {
	_, ret, err := Net.Get(app_conf.CosMos_addr+"/blocks/"+Calc.Any2String(blocks), nil, nil, nil)
	if err != nil {
		fmt.Println("err:", err)
		return "", err
	} else {
		//fmt.Println("suc", ret)
		return Calc.Any2String(ret), err
	}
}

func Txs_Decode(txs string) (string, error) {
	postData := make(map[string]interface{})
	postData["tx"] = txs
	_, ret, err := Net.PostJson(app_conf.CosMos_addr+"/txs/decode", nil, postData, nil, nil)
	if err != nil {
		fmt.Println("err:", err)
		return "", err
	} else {
		return ret.(string), err
	}
}

func NewAccount(username string) (string, error) {
	num := Calc2.Rand(100000, 999999)
	name := Calc.Any2String(num)
	ts := time.Now().Unix()
	cmd := exec.Command("ltcli", "keys", "add", "test_"+username+"_"+name+"_"+Calc.Any2String(ts), "-o", "json")
	buf, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(buf), err
}
