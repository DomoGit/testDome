package boot

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func TestMysql() {
	var mysqlHost = os.Getenv("MYSQL_HOST")
	var mysqlPort = os.Getenv("MYSQL_PORT")
	var mysqlUser = os.Getenv("MYSQL_USER")
	var mysqlPassword = os.Getenv("MYSQL_PASSWORD")
	var mysqlDb = os.Getenv("MYSQL_DB")
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDb)
	DB, _ := sql.Open("mysql", dsn)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
}

func TestRedis() {
	// 创建客户端
	client := newClient()
	defer client.Close()

	// 设置key
	err := client.Set("name", "john", 0).Err()
	if err != nil {
		panic(err)
	}

	// 获取key
	val, err := client.Get("name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis connect success")
	fmt.Println("name", val)
}

// 创建redis客户端
func newClient() *redis.Client {
	var redisHost = os.Getenv("REDIS_HOST")
	var redisPort = os.Getenv("REDIS_PORT")
	var address = fmt.Sprintf("%s:%s", redisHost, redisPort)
	client := redis.NewClient(&redis.Options{
		Addr:     address, // redis地址
		Password: "",               // 密码
		DB:       0,                // 使用默认数据库
	})
	return client
}

