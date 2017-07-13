package binarysearch

import (
	"strconv"
	"strings"
)

const testVersion = 1

var (
	messageTemplateSuccessBegin = "? found at beginning of slice"
	messageTemplateSuccessEnd   = "? found at end of slice"
	messageTemplateSuccess1     = "? found at index ?"
	messageTemplateFail         = "? > ? at index ?, < ? at index ?"
	messageTemplateNil          = "slice has no values"
	messageTemplateMin          = "? < all values"
	messageTemplateMax          = "? > all ? values"
)

func SearchInts(slices []int, key int) (x int) {
	i, j, k := 0, len(slices), 1
	mink := -1
	if j == 0 {
		return 0
	}
	if key > slices[j-1] {
		return j
	}
	if key < slices[0] {
		return 0
	}
	for k > 0 {
		k = (i + j) / 2
		if slices[k] == key {
			mink = k
			j = (i + j) / 2
		} else if key < slices[k] {
			j = (i + j) / 2
		} else {
			if i == k {
				break
			}
			i = k
		}
	}
	if mink == -1 {
		return k + 1
	}
	return mink
}
func Message(slices []int, key int) (res string) {
	if slices == nil {
		return messageTemplateNil
	}
	index := SearchInts(slices, key)
	if index == len(slices) {
		res = strings.Replace(messageTemplateMax, "?", strconv.Itoa(key), 1)
		res = strings.Replace(res, "?", strconv.Itoa(len(slices)), 1)
		return
	}
	if slices[index] == key {
		if index == 0 {
			res = strings.Replace(messageTemplateSuccessBegin, "?", strconv.Itoa(key), 1)
		} else if index == len(slices)-1 {
			res = strings.Replace(messageTemplateSuccessEnd, "?", strconv.Itoa(key), 1)
		} else {
			res = strings.Replace(messageTemplateSuccess1, "?", strconv.Itoa(key), 1)
			res = strings.Replace(res, "?", strconv.Itoa(index), 1)
		}
	} else {
		if index == 0 {
			res = strings.Replace(messageTemplateMin, "?", strconv.Itoa(key), 1)
		} else {
			res = strings.Replace(messageTemplateFail, "?", strconv.Itoa(key), 1)
			res = strings.Replace(res, "?", strconv.Itoa(slices[index-1]), 1)
			res = strings.Replace(res, "?", strconv.Itoa(index-1), 1)
			res = strings.Replace(res, "?", strconv.Itoa(slices[index]), 1)
			res = strings.Replace(res, "?", strconv.Itoa(index), 1)
		}
	}
	return
}
