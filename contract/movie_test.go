package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"fmt"
	"github.com/bitly/go-simplejson"
)

var bluemixUrl string = "https://f78040cda74f4ab5bfe4e3c5c53a9a4b-vp1.us.blockchain.ibm.com:5001/chaincode"
var contractHash string = "327771fb6f34cc8fb0f55aa10695aa49eefcc0e73414bb0606b2a9fd6bdb495c14fb2940de5075db751723cf758dd93a3d218c58b88faac481838060cdbac76f"
var secureContext string = "user_type1_1"
var id uint = 0

type req struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  params `json:"params"`
	Id uint `json:"id"`
}

type ctorMsg struct {
	Function string `json:"function"`
	Args     []string `json:"args"`
}

type chaincodeID struct {
	Name string `json:"name"`
}

type params struct {
	Type          int `json:"type"`
	ChaincodeID   chaincodeID `json:"chaincodeID"`
	CtorMsg       ctorMsg `json:"ctorMsg"`
	SecureContext string `json:"secureContext"`
}

func TestRegisterCinema(t *testing.T) {
	cinema := queryCinema([]string{"翠苑电影大世界"})
	if cinema != nil {
		fmt.Printf("%v\n", cinema)
	}
	registerCinema([]string{"翠苑电影大世界", "万达"})
}

func registerCinema(args []string)  {
	resp, _ := request("invoke", "registerCinema", args)
	txid, _ := parseInvoke(resp)
	fmt.Printf("registerCinema txid = %s\n", txid)
}

func queryCinema(args []string) *Cinema {
	resp, _ := request("query", "queryCinema", args)
	var cinema Cinema
	if parseQuery(resp, &cinema) {
		return &cinema
	}
	return nil
}

func request(method string, function string, args []string) (*simplejson.Json, error) {
	buf, _ := json.Marshal(req{
		Jsonrpc: "2.0",
		Method: method,
		Params: params{
			Type: 1,
			ChaincodeID: chaincodeID{
				Name: contractHash,
			},
			CtorMsg: ctorMsg{
				Function: function,
				Args: args,
			},
			SecureContext: secureContext,
		},
		Id: id,
	})
	id += 1
	resp, _ := http.Post(bluemixUrl, "application/json", bytes.NewBuffer(buf))
	return simplejson.NewFromReader(resp.Body)
}

func parseQuery(resp *simplejson.Json, v interface{}) bool {
	if resp != nil {
		result := resp.Get("result")
		if result != nil {
			message := result.Get("message")
			if message != nil {
				bs, err := message.Bytes()
				if err == nil {
					json.Unmarshal(bs, v)
					return true
				}
			}
		}
	}
	return false
}

func parseInvoke(resp *simplejson.Json) (string, error) {
	if resp != nil {
		result := resp.Get("result")
		if result != nil {
			message := result.Get("message")
			if message != nil {
				return message.String()
			}
		}
	}
	return "", nil
}
