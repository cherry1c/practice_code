package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var redBall []int
var blueBall []int

const (
	redAmount  = 6  // 红球数量
	blueAmount = 1  // 篮球数量
	redMax     = 33 // 红球最大值
	blueMax    = 16 // 篮球最大值
)

func init() {
	redBall = make([]int, redMax)
	blueBall = make([]int, blueMax)
}

// getRandNum 获取随机值x (x < max)
func getRandNum(max int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return int(n.Int64())
}

// getRedBallNum 随机获取一个未被选中的红球
func getRedBallNum() int {
	for {
		n := getRandNum(redMax)
		if redBall[n] == 0 {
			redBall[n] = 1
			return n + 1
		}
	}
}

// getBlueBallNum 随机获取一个未被选中的蓝球
func getBlueBallNum() int {
	for {
		n := getRandNum(blueMax)
		if blueBall[n] == 0 {
			blueBall[n] = 1
			return n + 1
		}
	}
}

// getRedBallList 获取选中红球列表
func getRedBallList() []int {
	ret := make([]int, redAmount)
	for i := 0; i < redAmount; i++ {
		ret[i] = getRedBallNum()
	}
	// 清空红球选中记录
	redBall = make([]int, redMax)
	return ret
}

// getBlueBallList 获取选中蓝球列表
func getBlueBallList() []int {
	ret := make([]int, blueAmount)
	for i := 0; i < blueAmount; i++ {
		ret[i] = getBlueBallNum()
	}
	// 清空蓝球选中记录
	blueBall = make([]int, blueMax)
	return ret
}

// getDoubleBallList 随机获取一组双色球数据
func getDoubleBallList() []int {
	redList := getRedBallList()
	blueList := getBlueBallList()
	return append(redList, blueList...)
}

// printDoubleBallResult 打印双色球结果
func printDoubleBallResult(balls []int) {
	if len(balls) != redAmount+blueAmount {
		return
	}
	fmt.Printf("红：")
	for i := 0; i < redAmount; i++ {
		fmt.Printf(" %d", balls[i])
	}
	fmt.Printf("  蓝：")
	for i := redAmount; i < redAmount+blueAmount; i++ {
		fmt.Printf(" %d", balls[i])
	}
	fmt.Printf("\n")
}

func getGroupsDoubleBall(groupsNum int) {
	for i := 0; i < groupsNum; i++ {
		printDoubleBallResult(getDoubleBallList())
	}
}

func main() {
	getGroupsDoubleBall(5)
}
