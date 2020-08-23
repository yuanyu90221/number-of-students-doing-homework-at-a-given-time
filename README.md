# number-of-students-doing-homework-at-a-given-time
## 題目解讀：

### 題目來源:
[number-of-students-doing-homework-at-a-given-time](https://leetcode.com/problems/number-of-students-doing-homework-at-a-given-time/)

### 原文:
Given two integer arrays startTime and endTime and given an integer queryTime.

The ith student started doing their homework at the time startTime[i] and finished it at time endTime[i].

Return the number of students doing their homework at time queryTime. More formally, return the number of students where queryTime lays in the interval [startTime[i], endTime[i]] inclusive.

### 解讀：

給定兩個 正整數陣列 startTime , endTime, 還有一個正整數queryTime
對於每個 startTime[i], endTime[i] 代表第i個學生做功課的起始時間

想要求出在給予的queryTime時間正在做功課的學生總數
也就是  queryTime 符合, startTime[i] <= queryTime <= endTime[i]

## 初步解法:
### 初步觀察:

對於每個 startTime[i], endTime[i]

逐步檢查 queryTime是否有在範圍中

即可得出解答


### 初步設計:
Given two integer arrays startTime, endTime, and an integer queryTime

step 0: let idx = 0, count = 0

step 1: if idx >= len(startTime) go to step 3

step 2: if startTime[idx] <= queryTime <= endTime[idx] set count = count + 1

step 3: return count
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
package busy_student

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	count := 0
	for idx := range startTime {
		if startTime[idx] <= queryTime && endTime[idx] >= queryTime {
			count++
		}
	}
	return count
}

```
## 測資的撰寫
```golang
package busy_student

import "testing"

func Test_busyStudent(t *testing.T) {
	type args struct {
		startTime []int
		endTime   []int
		queryTime int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				startTime: []int{1, 2, 3},
				endTime:   []int{3, 2, 7},
				queryTime: 4,
			},
			want: 1,
		},
		{
			name: "Example2",
			args: args{
				startTime: []int{4},
				endTime:   []int{4},
				queryTime: 4,
			},
			want: 1,
		},
		{
			name: "Example3",
			args: args{
				startTime: []int{4},
				endTime:   []int{4},
				queryTime: 5,
			},
			want: 0,
		},
		{
			name: "Example4",
			args: args{
				startTime: []int{1, 1, 1, 1},
				endTime:   []int{1, 3, 2, 4},
				queryTime: 7,
			},
			want: 0,
		},
		{
			name: "Example5",
			args: args{
				startTime: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
				endTime:   []int{10, 10, 10, 10, 10, 10, 10, 10, 10},
				queryTime: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := busyStudent(tt.args.startTime, tt.args.endTime, tt.args.queryTime); got != tt.want {
				t.Errorf("busyStudent() = %v, want %v", got, tt.want)
			}
		})
	}
}

```

## my self record
[golang leetcode 30 23thday](https://hackmd.io/NumvG0idTgy5N8fujgo2wg?view)
## 參考文章

[golang test](https://ithelp.ithome.com.tw/articles/10204692)