package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct{
	gorm.Model
	Name string `gorm:"type:varchar(20); not null"`
	Telephone string `gorm:"varchar(11); not null; unique"`
	Password string `gorm:"size:255; not null"`
}

func RandomString(n int) string{
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result{
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

func InitDB() *gorm.DB{
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "15089246141@163.com"
	charset := "utf8mb4"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
	charset)
	db, err := gorm.Open(driverName,args)
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}

	//自动创建数据表
	db.AutoMigrate(&User{})
	return db
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?",telephone).First(&user)

	if user.ID !=0 {
		return true
	}

	return false
}

func main(){
	db := InitDB()
	defer db.Close()

	r := gin.Default()
	r.POST("/api/auth/register",func(ctx *gin.Context) {
		// 获取参数
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")
		// 数据验证
		if len(telephone) != 11 {
			// 这里传输了一个http的一个状态常量
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422,
			"msg":"手机号必须为11位"})
			return //让其不再继续往后执行
		}

		if len(password) < 6 {
			// 这里传输了一个http的一个状态常量
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422,"msg":"密码不能少于6位"})
			return 
		}

		// 如果名称没有传，给一个10位的随机字符串
		if len(name) == 0 {
			name = RandomString(10)
			
		}

		log.Println(name, telephone, password)

		// 判断手机号是否存在
		if isTelephoneExist(db,telephone){
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422,"msg":"用户已经存在"})
			return 
		}
		//创建用户

		newUser := User{
			Name: name,
			Telephone: telephone,
			Password: password,
		}

		db.Create(&newUser)
		//返回结果

		ctx.JSON(200,gin.H{
			"msg":"注册成功",
		})
	})
	r.Run() //监听并在0.0.0.0:8080上启动服务
}


