package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var DB = redis.NewClient(&redis.Options{
	Addr: "127.0.0.1:6379",
	//Password: "root",
	DB: 0,
})

func useLua(userid, prodid string) bool {
	//编写脚本 - 检查数值，是否够用，够用再减，否则返回减掉后的结果
	var luaScript = redis.NewScript(`
        local userid=KEYS[1];
        local prodid=KEYS[2];
        local qtKey="sk:"..prodid..":qt";
        local userKey="sk:"..prodid..":user";
        local userExists=redis.call("sismember",userKey,userid);
        if tonumber(userExists)==1 then
         return 2;
        end
        local num=redis.call("get",qtKey);
        if tonumber(num)<=0 then
         return 0;
        else
         redis.call("decr",qtKey);
         redis.call("SAdd",userKey,userid);
        end
        return 1;
    `)
	//执行脚本
	n, err := luaScript.Run(ctx, DB, []string{userid, prodid}).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	switch n {
	case int64(0):
		fmt.Println(userid, "抢购结束")
		return false
	case int64(1):
		fmt.Println(userid, "：抢购成功")
		return true
	case int64(2):
		fmt.Println(userid, "：已经抢购了")
		return false
	default:
		fmt.Println("发生未知错误！")
		return false
	}
	return true
}
func main() {
	// 并发的版本
	for i := 0; i < 80; i++ {
		go func(i int) {
			//uuids := uuid.New().String()
			uuids := strconv.Itoa(i)
			//fmt.Println(uuids)
			prodid := "1023"
			//time.Sleep(10 * time.Second)
			useLua(uuids, prodid)
		}(i)
	}
	time.Sleep(15 * time.Second)
}
