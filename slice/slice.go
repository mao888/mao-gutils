package slice

// DifferenceSet 返回两个slice切片的差集
func DifferenceSet(a []int, b []int) []int {
	var c []int
	temp := map[int]struct{}{}

	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}

	for _, val := range a {
		if _, ok := temp[val]; !ok {
			c = append(c, val)
		}
	}

	return c
}

// RemoveRepByMap 通过map主键唯一的特性过滤重复元素
// 结构体切片去重
func RemoveRepByMap(slc []*SeriesRes) []*SeriesRes {
	resultMap := make(map[string]*SeriesRes, len(slc))
	for _, v := range slc {
		resultMap[v.SeriesId] = v
	}
	var result []*SeriesRes
	for _, v := range resultMap {
		result = append(result, v)
	}
	return result
}

// RemoveStructRepByMap 通过map主键唯一的特性过滤重复元素
// 结构体切片去重
func RemoveStructRepByMap(slc []*SeriesRes) []*SeriesRes {
	resultMap := make(map[string]*SeriesRes, len(slc))
	for _, v := range slc {
		resultMap[v.SeriesId] = v
	}
	var result []*SeriesRes
	for _, v := range resultMap {
		result = append(result, v)
	}
	return result
}

type ListNode struct {
	Val  int
	next *ListNode
}

// SliceToLinkList 切片转链表
func SliceToLinkList(nums []int, head *ListNode) *ListNode {
	node := head
	for _, num := range nums {
		temp := ListNode{Val: num}
		head.next = &temp
		head = &temp
	}
	return node.next
}
