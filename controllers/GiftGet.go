package controllers

import (
	"errors"
	beego "github.com/beego/beego/v2/server/web"
	"server/internel/conf"
	"server/internel/redis"
	"strconv"
	"strings"
)

type GiftGet struct {
	beego.Controller
}

func (this *GiftGet) Get() {
	userId := this.GetString("userId")
	code := this.GetString("code")
	changes, err := ProcessGiftGet(userId, code)

	ans := GiftGetRes{}
	if err != nil {
		ans.Message = err.Error()
	} else {
		ans.Message = "ok"
		ans.Changes = changes
	}
	this.Data["json"] = &ans
	this.ServeJSON()
}

func ProcessGiftGet(userId, code string) ( map[string]string, error) {
	var (
		err error
	)
	key := strings.Replace(conf.GITDETIL, "code", code, 1)
	detils:=redis.HGetAll(key)

	if err = validAndUpdateGift(userId,code);err!=nil{
		return nil, err
	}

	updateUserAsset(userId,detils)

	return detils,err
}

func updateUserAsset(userId string, detils map[string]string) {
	//todo 士兵，英雄，和道具区分种类,记录id
	for key,detil := range detils{
		rdKey:=key+":"+"{"+userId+"}"
		val,_:=strconv.Atoi(detil)
		redis.IncrBy(rdKey,val)
	}
}


func validAndUpdateGift(user ,code string) error {

	key := strings.Replace(conf.GIFTUSER, "code", code, 1)
	userId:=redis.Get(key)

	key = strings.Replace(conf.GIFTCOUNT, "code", code, 1)
	count:=redis.Get(key)

	if userId!=user{
		return errors.New("user not match")
	}

	// 保证count>=0 (次数有限）或 count==-1（次数无限）
	if count == "0" {
		return errors.New("the code can't use")
	}
	if count != "-1"{
		key = strings.Replace(conf.GIFTCOUNT, "code", code, 1)
		redis.IncrBy(key,-1)
	}
	return nil
}

type GiftGetRes struct {
	Changes map[string]string `json:"changes"`
	Message string         `json:"message"`
}

//func addGift(gift map[string]string, userId string) (map[string]map[string]int,error) {
//	var(
//		err error
//	)
//	giftType := gift["type"]
//	val := gift["member"]
//	member, _ := strconv.Atoi(val)
//	ans:=make(map[string]map[string]int)
//	switch giftType {
//	//var GiftType= map[string]int{"coin":0,"diamond":0,"Props":0,"hero":0,"army":0}
//	case "army":
//		army:=make(map[string]int)
//		ans["army"]=army
//		err=addArmy(userId, member,army)
//	case "hero":
//		hero:=make(map[string]int)
//		ans["hero"]=hero
//		err=addHero(userId, member,hero)
//	case "coin":
//		coin:=make(map[string]int)
//		ans["coin"]=coin
//		err=addCoin(userId, member,coin)
//	case "diamond":
//		diamond:=make(map[string]int)
//		ans["diamond"]=diamond
//		err=addDiamond(userId, member,diamond)
//	case "props":
//		props:=make(map[string]int)
//		ans["props"]=props
//		err=addProps(userId, member,props)
//	}
//	return ans,err
//}

//func addProps(id string, member int, changes map[string]int) error {
//	key := "prop:" + id
//	props := redis.Rdb.HGetAll(redis.Ctx, key).Val()
//
//	for i := 0; i < member; i++ {
//		propId := redis.Rdb.Incr(redis.Ctx, "prop").Val()
//		incr(props, propId)
//	}
//	err := redis.Rdb.HSet(redis.Ctx, key, props).Err()
//	return err
//}

//func incr(props map[string]string, id int64) {
//	var (
//		val string
//		ok  bool
//	)
//	key := fmt.Sprintf("%08d", id)
//	if val, ok = props[key]; !ok {
//		props[key] = "0"
//	} else {
//		i, _ := strconv.Atoi(val)
//		props[key] = fmt.Sprintf("%d", i+1)
//	}
//}

//func addDiamond(id string, member int, changes map[string]int) error {
//	key := "diamond:" + id
//	props := redis.Rdb.HGetAll(redis.Ctx, key).Val()
//
//	for i := 0; i < member; i++ {
//		propId := redis.Rdb.Incr(redis.Ctx, "prop").Val()
//		changes[fmt.Sprintf("%d",propId)]=1
//		incr(props, propId)
//	}
//	err := redis.Rdb.HSet(redis.Ctx, key, props).Err()
//	return err
//}
//
//func addCoin(id string, member int, changes map[string]int) error {
//
//	key := "coin:" + id
//	props := redis.Rdb.HGetAll(redis.Ctx, key).Val()
//
//	for i := 0; i < member; i++ {
//		propId := redis.Rdb.Incr(redis.Ctx, "prop").Val()
//		changes[fmt.Sprintf("%d",propId)]=1
//		incr(props, propId)
//	}
//	err := redis.Rdb.HSet(redis.Ctx, key, props).Err()
//	return err
//}
//
//func addHero(id string, member int, changes map[string]int) error {
//
//	key := "hero:" + id
//	props := redis.Rdb.HGetAll(redis.Ctx, key).Val()
//
//	for i := 0; i < member; i++ {
//		propId := redis.Rdb.Incr(redis.Ctx, "prop").Val()
//		changes[fmt.Sprintf("%d",propId)]=1
//		incr(props, propId)
//	}
//	err := redis.Rdb.HSet(redis.Ctx, key, props).Err()
//	return err
//}
//
//func addArmy(id string, member int, changes map[string]int) error {
//
//	key := "army:" + id
//	props := redis.Rdb.HGetAll(redis.Ctx, key).Val()
//
//	for i := 0; i < member; i++ {
//		propId := redis.Rdb.Incr(redis.Ctx, "prop").Val()
//		changes[fmt.Sprintf("%d",propId)]=1
//		incr(props, propId)
//	}
//	err := redis.Rdb.HSet(redis.Ctx, key, props).Err()
//	return err
//}
