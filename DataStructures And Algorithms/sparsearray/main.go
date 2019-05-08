// Package main provides ...
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//稀疏数组结构体
type ValNode struct {
	row int
	col int
	val int
}

func main() {
	//创建一个原始数组
	var chessMap [11][11]int
	//黑棋
	chessMap[1][2] = 1
	//白棋
	chessMap[2][3] = 2
	//输出原始数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
	var sparseSlice []ValNode
	//添加初始行列值
	vnode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseSlice = append(sparseSlice, vnode)
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				vnode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseSlice = append(sparseSlice, vnode)
			}
		}
	}
	filename := "chess.data"
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("写入文件失败错误消息:", err)
	}
	writer := bufio.NewWriter(file)
	for _, v := range sparseSlice {
		str := fmt.Sprintf("%d,%d,%d\n", v.row, v.col, v.val)
		writer.WriteString(str)
	}
	writer.Flush()
	file.Close()
	var chessMap2 [11][11]int
	var sparseSlice2 []ValNode
	file1, _ := os.Open(filename)
	reader := bufio.NewReader(file1)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		str = strings.Replace(str, "\n", "", -1)
		strArr := strings.Split(str, ",")
		var vnode ValNode
		if strArr[0] != "11" {
			vnode.row, _ = strconv.Atoi(strArr[0])
			vnode.col, _ = strconv.Atoi(strArr[1])
			vnode.val, _ = strconv.Atoi(strArr[2])
			sparseSlice2 = append(sparseSlice2, vnode)
		}
	}
	file1.Close()
	fmt.Println(sparseSlice2)
	for _, v := range sparseSlice2 {
		chessMap2[v.row][v.col] = v.val
	}
	fmt.Println("恢复后的数据")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d \t", v2)
		}
		fmt.Println()
	}
}
