package BlockSync

import (
	"fmt"
	"main.go/app/v1/cosmos/model/BlocksModel"
	"main.go/extend/CosMos/CosCore"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Log"
	"time"
)

func Syncdata() {
	defer Syncdata()
	BlocksModel.Db = tuuz.Db()
	lastblock := BlocksModel.Api_find_last()
	height := 1
	if len(lastblock) > 0 {
		height = Calc.Any2Int(lastblock["height"]) + 1
	}
	ret, err := CosCore.Blocks(height)
	if err != nil {
		panic(err)
	} else {
		rtt, err := Jsong.JObject(ret)
		timer := ""
		chain_id := ""
		block_hash := ""
		from_address := ""
		to_address := ""
		memo := ""
		amount := ""
		fee := ""
		if err != nil {
			Log.Errs(err, tuuz.FUNCTION_ALL())
			time.Sleep(time.Second)
			return
		} else {
			block_meta, err := Jsong.ParseObject(rtt["block_meta"])
			if err != nil {
				Log.Errs(err, tuuz.FUNCTION_ALL())
			}
			block_id, err := Jsong.ParseObject(block_meta["block_id"])
			if err != nil {
				Log.Errs(err, tuuz.FUNCTION_ALL())
			} else {
				block_hash = Calc.Any2String(block_id["hash"])
			}

			block, err := Jsong.ParseObject(rtt["block"])
			if err != nil {
				Log.Errs(err, tuuz.FUNCTION_ALL())
			} else {
				header, err := Jsong.ParseObject(block["header"])
				if err != nil {
					Log.Errs(err, tuuz.FUNCTION_ALL())
				} else {
					timer = Calc.Any2String(header["time"])
					height = Calc.Any2Int(header["height"])
					if height < Calc.Any2Int(lastblock["height"]) {
						time.Sleep(time.Second)
						return
					}
					chain_id = Calc.Any2String(header["chain_id"])
				}

				data, err := Jsong.ParseObject(block["data"])
				if err != nil {
					Log.Errs(err, tuuz.FUNCTION_ALL())
				} else {
					txss, err := Jsong.ParseSlice(data["txs"])

					if err != nil || len(txss) < 1 {
						//fmt.Println(err, tuuz.FUNCTION_ALL())
					} else {
						txs, err := txs_format(Calc.Any2String(txss[0]))
						if err != nil {
							fmt.Println(err, tuuz.FUNCTION_ALL())
							time.Sleep(time.Second)
							return
						} else {
							from_address = Calc.Any2String(txs["from_address"])
							to_address = Calc.Any2String(txs["to_address"])
							memo = Calc.Any2String(txs["memo"])
							amount = Calc.Any2String(txs["amount"])
							fee = Calc.Any2String(txs["fee"])
						}
					}
				}
			}
		}
		if height != 0 {
			BlocksModel.Db = tuuz.Db()
			BlocksModel.Api_insert(height, timer, chain_id, block_hash, from_address, to_address, memo, amount, fee, ret)
		} else {
			time.Sleep(time.Second)
		}

	}
}

func txs_format(txs string) (map[string]interface{}, error) {
	arr := make(map[string]interface{})
	str, err := CosCore.Txs_Decode(txs)
	if err != nil {
		fmt.Println("txs.err")
		return nil, err
	} else {
		json, err := Jsong.JObject(str)
		if err == nil {
			result, err := Jsong.ParseObject(json["result"])
			if err != nil {
				fmt.Println("result,为空")
				Log.Errs(err, tuuz.FUNCTION_ALL())
				return nil, err
			}
			msgs, err := Jsong.ParseSlice(result["msg"])
			if err != nil {
				fmt.Println("msgs,为空")
				Log.Errs(err, tuuz.FUNCTION_ALL())
				return nil, err
			}
			msg, err := Jsong.ParseObject(msgs[0])
			if err != nil {
				fmt.Println("msg,为空")
				Log.Errs(err, tuuz.FUNCTION_ALL())
				return nil, err
			}
			value, err := Jsong.ParseObject(msg["value"])
			if err != nil {
				fmt.Println("value,,为空")
				Log.Errs(err, tuuz.FUNCTION_ALL())
				return nil, err
			}
			arr["from_address"] = value["from_address"]
			arr["to_address"] = value["to_address"]
			arr["amount"], err = Jsong.Encode(value["amount"])
			arr["fee"], err = Jsong.Encode(result["fee"])
			arr["memo"] = result["memo"]
		}
	}
	return arr, err
}
