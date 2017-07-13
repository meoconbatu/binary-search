package binarysearch

import (
	"fmt"
)

const testVersion = 1

var (
	messageTemplateSuccessBegin = "%d found at beginning of slice"
	messageTemplateSuccessEnd   = "%d found at end of slice"
	messageTemplateSuccess1     = "%d found at index %d"
	messageTemplateFail         = "%d > %d at index %d, < %d at index %d"
	messageTemplateNil          = "slice has no values"
	messageTemplateMin          = "%d < all values"
	messageTemplateMax          = "%d > all %d values"
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
func Message(slices []int, key int) string {
	if slices == nil {
		return messageTemplateNil
	}
	index := SearchInts(slices, key)
	if index == len(slices) {
		return fmt.Sprintf(messageTemplateMax, key, len(slices))
	}
	if slices[index] == key {
		if index == 0 {
			return fmt.Sprintf(messageTemplateSuccessBegin, key)
		} else if index == len(slices)-1 {
			return fmt.Sprintf(messageTemplateSuccessEnd, key)
		}
		return fmt.Sprintf(messageTemplateSuccess1, key, index)
	} else if index == 0 {
		return fmt.Sprintf(messageTemplateMin, key)
	}
	return fmt.Sprintf(messageTemplateFail, key, slices[index-1], index-1, slices[index], index)
}
