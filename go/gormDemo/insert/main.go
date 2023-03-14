package main

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/big"
)

var db *gorm.DB

func Init() {
	dsn := "root:123456@tcp(localhost:3306)/learn_test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("open db failed")
		panic("")
		return
	}
}

var nameList = []string{"张三", "李四", "王五", "wesly", "simmous", "rey", "gene", "bill", "evin", "clare"}
var addressList = []string{"美食", "娱乐", "房", "车", "聚会"}

// getRandNum returns a uniform random value in [0, max)
func getRandNum(max int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return int(n.Int64())
}

func getRandName() string {
	return nameList[getRandNum(len(nameList))]
}

func getRandAddress() string {
	return addressList[getRandNum(len(addressList))]
}

func AddShop1Data(num int) error {
	var shop1 []*Shops1
	for i := 0; i < num; i++ {
		shop1 = append(shop1, &Shops1{
			Name:     getRandName(),
			Type:     uint64(getRandNum(10)),
			Address:  getRandAddress(),
			IsDelete: sql.NullInt64{Int64: int64(getRandNum(2)), Valid: true},
			Status:   sql.NullInt64{Int64: int64(getRandNum(4)), Valid: true},
		})
	}
	result := db.Create(shop1)
	return result.Error
}

func AddShop2Data(num int) error {
	var shop1 []*Shops2
	for i := 0; i < num; i++ {
		shop1 = append(shop1, &Shops2{
			Name:     getRandName(),
			Type:     uint64(getRandNum(10)),
			Address:  getRandAddress(),
			IsDelete: sql.NullInt64{Int64: int64(getRandNum(2)), Valid: true},
			Status:   sql.NullInt64{Int64: int64(getRandNum(4)), Valid: true},
		})
	}
	result := db.Create(shop1)
	return result.Error
}

func AddData(num int) {
	for i := 0; i <= num/1000; i++ {
		if err := AddShop1Data(1000); err != nil {
			fmt.Println(err.Error())
			return
		}
		if err := AddShop2Data(1000); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func main() {
	Init()
	AddData(100000)
}
