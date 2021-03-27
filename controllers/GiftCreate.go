package controllers

import (
	"encoding/json"
	"errors"
	beego "github.com/beego/beego/v2/server/web"
	"log"
	"server/internel/conf"
	"server/internel/model"
	"server/internel/redis"
	"server/internel/utils"
	"strings"
	"time"
)

type GiftCodeCreate struct {
	beego.Controller
}

func (this *GiftCodeCreate) Post() {
	req := &GiftCodeReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, req); err != nil {
		log.Fatal(err)
	}
	code, err := processGiftCodeCreatre(req)
	ans := GiftCreateRes{}
	if err != nil {
		ans.Message = err.Error()
	} else {
		ans.Message = "ok"
		ans.Code = code
	}
	this.Data["json"] = &ans
	this.ServeJSON()
}

func processGiftCodeCreatre(req *GiftCodeReq) (string, error) {
	var (
		err error
	)

	detils := make(map[string]interface{})

	for _, val := range req.Gifts {
		err = validGiftType(val)
		if err != nil {
			return "", err
		}
		detils[val.Type] = val.Member
	}

	code := utils.GenerateCode(time.Now().Unix())
	//todo code重复处理

	key := strings.Replace(conf.GIFTCOUNT, "code", code, 1)
	if err = redis.Set(key, req.Count); err != nil {
		return code, err
	}

	key = strings.Replace(conf.GIFTUSER, "code", code, 1)
	if err = redis.Set(key, req.Uid); err != nil {
		return code, err
	}

	key = strings.Replace(conf.GITDETIL, "code", code, 1)
	if err = redis.Hset(key, detils); err != nil {
		return code, err
	}

	return code, nil
}

func validGiftType(gift GiftStruct) error {
	if _, ok := model.GiftType[gift.Type]; !ok {
		return errors.New("type error")
	}
	if gift.Member <= 0 {
		return errors.New("member should bigger than 0")
	}
	return nil
}

type GiftCodeReq struct {
	Uid   string       `json:"userId"`
	Count int          `json:"count"`
	Gifts []GiftStruct `json:"gifts"`
}

type GiftStruct struct {
	Type   string `json:"type"`
	Member int    `json:"member"`
}

type GiftCreateRes struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
