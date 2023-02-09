package trace

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

const dsn = "root:liu@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func Test_GormTracing(t *testing.T) {
	closer := InitJaeger("gormTracing")
	defer closer.Close()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	_ = db.Use(&OpentracingPlugin{})

	// 迁移 schema
	_ = db.AutoMigrate(&Product{})

	// 生成新的Span - 注意将span结束掉，不然无法发送对应的结果
	span := opentracing.StartSpan("gormTracing unit test")
	defer span.Finish()

	// 把生成的Root Span写入到Context上下文，获取一个子Context
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	session := db.WithContext(ctx)

	// Create
	session.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	session.First(&product, 1)                 // 根据整形主键查找
	session.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	session.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	session.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	session.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	session.Delete(&product, 1)
}

func Test_GormTracing2(t *testing.T) {
	closer := InitJaeger("gormTracing")
	defer closer.Close()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	_ = db.Use(&OpentracingPlugin{})

	rand.Seed(time.Now().UnixNano())

	num, wg := 1<<10, &sync.WaitGroup{}

	wg.Add(num)

	for i := 0; i < num; i++ {
		go func(t int) {
			span := opentracing.StartSpan(fmt.Sprintf("gormTracing unit test %d", t))
			defer span.Finish()

			ctx := opentracing.ContextWithSpan(context.Background(), span)
			session := db.WithContext(ctx)

			p := &Product{Code: strconv.Itoa(t), Price: uint(rand.Intn(1 << 10))}

			session.Create(p)

			session.First(p, p.ID)

			session.Delete(p, p.ID)

			wg.Done()
		}(i)
	}

	wg.Wait()
}
