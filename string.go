package jb

import (
	"sort"

	"github.com/spf13/cast"
)

// TruncateString truncates the given string limiting it by the given num
func TruncateString(str string, num int) string {
	s := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		s = str[0:num] + "..."
	}
	return s
}

// MakeSeqStr takes an int array and returns a string representation of all the values in a comma delimited format.
func MakeSeqStr(nums []int) string {
	seqMap := make(map[int][]int)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var mapCount int
	var done int
	var switchInt int
	seqMap[mapCount] = append(seqMap[mapCount], nums[done])
	done++
	switchInt = done
	for done < len(nums) {
		if nums[done] == ((seqMap[mapCount][(switchInt - 1)]) + 1) {
			seqMap[mapCount] = append(seqMap[mapCount], nums[done])
			switchInt++
		} else {
			mapCount++
			seqMap[mapCount] = append(seqMap[mapCount], nums[done])
			switchInt = 1
		}
		done++
	}
	var seqStr string
	for k, v := range seqMap {
		if k > 0 {
			seqStr += ","
		}
		if len(v) > 1 {
			seqStr += cast.ToString(v[0])
			seqStr += "-"
			seqStr += cast.ToString(v[len(v)-1])
		} else {
			seqStr += cast.ToString(v[0])
		}
	}
	return seqStr
}
