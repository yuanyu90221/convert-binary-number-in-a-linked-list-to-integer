# convert-binary-number-in-a-linked-list-to-integer
## 題目解讀：

### 題目來源:
[convert-binary-number-in-a-linked-list-to-integer](https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/)

### 原文:

Given head which is a reference node to a singly-linked list. The value of each node in the linked list is either 0 or 1. The linked list holds the binary representation of a number.

Return the decimal value of the number in the linked list.
### 解讀：

給定一個 鏈節串列的起始節點 head

每個串列節點的值 都是 0或是1 代表二元數的位元

舉例來說就是 1->0->1 => b101 = 1* 2^2+ 0* 2^1 + 1* 2^0 = 4 + 0 + 1 =5

求解 每個鏈節串列所代表的10進制的值
## 初步解法:
### 初步觀察:

首先 這個鏈節串列是從最高的digits往後排

因此 無法直接遞加

遞加必須等到走完所有節點才知道 digit所代表的次方數

因此可以考慮把每個位元轉換成string

最後再把string 轉回 integer

### 初步設計:

Given an LinkedList root node head

Step 0: let a string result = ""

Step 1: if ptr == nil go to step 5

Step 2: let ptr = head

Step 3: set string += strconv.Itoa(ptr.Val)

Step 4: ptr = ptr.Next

Step 5: return strconv.Atoi(result, 2, 10)
## 遇到的困難
### 題目上理解的問題
因為英文不是筆者母語

所以在題意解讀上 容易被英文用詞解讀給搞模糊

### pseudo code撰寫

一開始不習慣把pseudo code寫下來

因此 不太容易把自己的code做解析

### golang table driven test不熟
對於table driven test還不太熟析

所以對於寫test還是耗費不少時間
## 我的github source code
```golang
package decimal_linkedlist

import (
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	result := 0
	binaryStr := ""
	for head != nil {
		binaryStr += strconv.Itoa(head.Val)
		head = head.Next
	}
	result64, _ := strconv.ParseInt(binaryStr, 2, 0)
	result = int(result64)
	return result
}

```
## 測資的撰寫
```golang
package decimal_linkedlist

import (
	"testing"
)

func Test_getDecimalValue(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				head: BuildList([]int{1, 0, 1}),
			},
			want: 5,
		},
		{
			name: "Example2",
			args: args{
				head: BuildList([]int{0}),
			},
			want: 0,
		},
		{
			name: "Example3",
			args: args{
				head: BuildList([]int{1}),
			},
			want: 1,
		},
		{
			name: "Example4",
			args: args{
				head: BuildList([]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}),
			},
			want: 18880,
		},
		{
			name: "Example5",
			args: args{
				head: BuildList([]int{0, 0}),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDecimalValue(tt.args.head); got != tt.want {
				t.Errorf("getDecimalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BuildList(nums []int) *ListNode {
	var result *ListNode
	var currentNode *ListNode
	for idx, val := range nums {
		tempNode := ListNode{Val: val, Next: nil}
		if idx == 0 {
			result = &tempNode
			currentNode = &tempNode
		} else {
			currentNode.Next = &tempNode
			currentNode = currentNode.Next
		}
	}
	return result
}
```
## 參考文章
## my self record
[golang leetcode 30day 18th day](https://hackmd.io/v9oK2KEQSdOVmMi94NixlA?view)
[golang test](https://ithelp.ithome.com.tw/articles/10204692)