package controller

import (
	"fmt"
	redis2 "gin-demo/core/redis"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"net/http"
	"strings"
)

const (
	db0 uint8 = iota
	db1
	db2
	db3
	db4
	db5
	db6
	db7
	db8
	db9
	db10
	db11
	db12
	db13
	db14
	db15
)

var databases = map[uint8]string{
	db0:  "DB0",
	db1:  "DB1",
	db2:  "DB2",
	db3:  "DB3",
	db4:  "DB4",
	db5:  "DB5",
	db6:  "DB6",
	db7:  "DB7",
	db8:  "DB8",
	db9:  "DB9",
	db10: "DB10",
	db11: "DB11",
	db12: "DB12",
	db13: "DB13",
	db14: "DB14",
	db15: "DB15",
}

func RdsIndex(c *gin.Context) {
	conn := redis2.Instance().GetRds()
	info, err := redis.String(conn.Do("INFO"))
	if err != nil {
		c.Error(err)
		return
	}
	infoSlice := strings.Split(info, "\r\n\r\n")

	c.HTML(http.StatusOK, "redis/index", gin.H{
		"name":      "body",
		"info":      infoSlice,
		"databases": databases,
	})
	return
}

func SelectDB(c *gin.Context) {
	conn := redis2.Instance().GetRds()
	db := c.PostForm("id")
	re, err := redis.String(conn.Do("select", db))
	if err != nil {
		c.Error(err)
		return
	}
	data := make(map[string]interface{})
	data["isOk"] = re
	data["db"] = db
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    data,
		"message": "操作成功",
	})
}

// Search key
func Search(c *gin.Context) {
	command := c.PostForm("command")
	conn := redis2.Instance().GetRds()
	//KEYS w3c*
	re, err := redis.Values(conn.Do("KEYS", "*"+command+"*"))
	if err != nil {
		c.Error(err)
		return
	}
	keySlice := make([]string, len(re))
	for k, v := range re {
		keySlice[k] = string(v.([]uint8))
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": keySlice,
	})
}

type albums struct {
	RoleId       string `redis:"role_id"`
	IsLocked     string `reids:"is_locked"`
	GemId        string `redis:"gem_id"`
	GemSuitId    string `redis:"gem_suit_id"`
	ArtifactId   string `redis:"artifact_id"`
	Level        string `redis:"level"`
	AbilityStep  string `redis:"ability_step"`
	HpMax        string `redis:"hp_max"`
	Mp           string `redis:"mp"`
	Attack       string `redis:"attack"`
	Armor        string `redis:"armor"`
	Speed        string `redis:"speed"`
	SkillDamageR string `redis:"skill_damage_r"`
	Aid          string `redis:"aid"`
	Name         string `redis:"name"`
}

func Excuse(c *gin.Context) {

	//command := c.PostForm("command")
	conn := redis2.Instance().GetRds()

	// set redis
	b, err := msgpack.Marshal(56428)
	if err != nil {
		fmt.Println(err)
		return
	}

	re, e := conn.Do("hset", "111111111", "aid", string(b))
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println("set result", re)

	data, err := redis.Values(conn.Do("HGETAll", "appz:u:base:1804019036556625920"))
	if err != nil {
		fmt.Println("~~~", err)
		return
	}

	object := albums{}
	err = redis.ScanStruct(data, &object)

	//core.Logger.Error([]byte(object.RoleId))

	//byteData1:=biu.BinaryStringToBytes(data)
	//byteData := fmt.Sprintf("%X", data)
	//byteData1, err := hex.DecodeString(data)
	//if err != nil {
	//	fmt.Println("-----", err)
	//	return
	//}
	var a int
	err = msgpack.Unmarshal([]byte(object.Level), &a)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("结果", a)
	return

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": "value",
	})
}
func RdsDemo(c *gin.Context) {

	//fmt.Printf("T%",rec1)
	//rec1 := make(map[string]interface{})
	//err = msgpack.Unmarshal(list, &rec1)

	/*	rec1, err := redis.String(conn.Do("GET", "appz:uz:hook:res:seconds:1:1540547975346917376"))
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(rec1)*/

	/*	re := make(map[string]interface{})
		if rec1Slice ,ok :=rec1.([]byte);ok{

			err = msgpack.Unmarshal(rec1Slice, &re)
			if err != nil {
				log.Fatal(err)
			}

		}



		data, _ := json.Marshal(re)*/

	/*	reply, err := redis.Values(conn.Do("GET","appz:db:amd:basearenarobotitems:all:100047"))
		if err != nil {
			panic("查询失败")
		}

		var s string
		if _, err := redis.Scan(reply, &s); err != nil {
			panic("解析失败")
		}*/
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": "",
	})
}
