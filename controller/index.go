package controller

import (
	"fmt"
	"gin-demo/core"
	"github.com/gin-gonic/gin"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"net/http"
	"time"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
		"data":    "no data",
	})
}

func IndexTest(c *gin.Context) {
	db := core.DbInstance().GetDB()
	chan1 := make(chan int64)
	chan2 := make(chan int64)
	go func() {
		tx, _ := db.Begin()
		re, err := tx.Update("u_friend", dbx.Params{"f_love": 1, "receipt_time": time.Now().Unix()}, dbx.HashExp{"uid": 1757575766767310848, "fuid": 1702645567076702208}).Execute()

		if err != nil {
			fmt.Println("错误信息", err)
			tx.Rollback()
		} else {
			time.Sleep(time.Second * 6)
			tx.Commit()
		}
		row, err1 := re.RowsAffected()

		if err1 != nil {
			fmt.Println(err1)
			row = 0
		}
		chan1 <- row
	}()

	go func() {
		tx, _ := db.Begin()

		re, err := tx.Update("u_friend", dbx.Params{"f_love": 1, "receipt_time": time.Now().Unix()}, dbx.HashExp{"uid": 1702645567076702208, "fuid": 1757575766767310848}).Execute()
		if err != nil {
			fmt.Println("错误信息", err)
			tx.Rollback()
		} else {
			time.Sleep(time.Second * 6)
			tx.Commit()
		}
		row, err1 := re.RowsAffected()

		if err1 != nil {
			fmt.Println(err1)
			row = 0
		}

		chan2 <- row
	}()

	r1 := <-chan1
	r2 := <-chan2
	c.JSON(http.StatusOK, gin.H{
		"chan1": r1,
		"chan2": r2,
	})
	fmt.Println("受影响的行数", r1, r2)
}
