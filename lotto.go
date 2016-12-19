package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

func thisweekLotto() {
	arrs := rand.New(rand.NewSource(time.Now().UnixNano()))
	r := arrs.Perm(45)

	for i := 0; i < 6; i++ {
		fmt.Print(r[rand.Intn(45)]+1, " ")
	}
	fmt.Println()

}

func main() {
	fmt.Println("금주번호 : ")
	for i := 0; i < 5; i++ {
		thisweekLotto()
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()

	bs, err := ioutil.ReadFile("d:/Temp/lotto/number.txt")
	if err != nil {
		panic(err)
	}

	lottoStr := string(bs)
	lottoArray := make([]int, 46)
	oldLottoNum := strings.Split(lottoStr, "\r\n")

	for _, oldNums := range oldLottoNum {
		oldNumItem := strings.Split(oldNums, ",")

		for _, item := range oldNumItem {
			index, err := strconv.Atoi(item)
			if err != nil {
				fmt.Println(err)
			}
			lottoArray[index]++

		}

	}

	copyLottoArray := make([]int, 46)
	for k := 0; k < 46; k++ {
		copyLottoArray[k] = lottoArray[k]
	}
	sort.Ints(copyLottoArray)
	var maxLottoCount int

	fmt.Print("확률번호 : ")
	for i := 0; i < 6; i++ {
		maxLottoCount = copyLottoArray[len(copyLottoArray)-1-i]
		for count, arrayNum := range lottoArray {
			if arrayNum == maxLottoCount {
				fmt.Print(count, " ")
				break
			}
		}
	}
	fmt.Println()

}
