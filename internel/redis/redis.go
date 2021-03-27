package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)
var Ctx = context.Background()

var Rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func Get(key string)string  {
	return Rdb.Get(Ctx,key).Val()
}

func Set(key string ,val interface{}) error{
	err:=Rdb.Set(Ctx,key,val,0).Err()
	return err
}

func Hset(key string,vals map[string]interface{}) error{
	err:=Rdb.HSet(Ctx,key,vals).Err()
	return err
}

func HGetAll(key string) map[string]string{
	return Rdb.HGetAll(Ctx,key).Val()
}


func IncrBy(key string,val int) int64{
	return Rdb.IncrBy(Ctx,key, int64(val)).Val()
}