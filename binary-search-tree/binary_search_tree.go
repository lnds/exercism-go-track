package binarysearchtree

type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

func Bst(data int) *SearchTreeData {
	return &SearchTreeData{
		data: data,
	}
}

func (bst *SearchTreeData) Insert(data int) {
	if data <= bst.data {
		if bst.left == nil {
			bst.left = Bst(data)
		} else {
			bst.left.Insert(data)
		}
	} else if data > bst.data {
		if bst.right == nil {
			bst.right = Bst(data)
		} else {
			bst.right.Insert(data)
		}
	}
}

func (bst *SearchTreeData) MapString(fn func(int) string) []string {
	result := []string{}
	if bst.left != nil {
		result = bst.left.MapString(fn)
	}
	result = append(result, fn(bst.data))
	if bst.right != nil {
		result = append(result, bst.right.MapString(fn)...)
	}
	return result
}

func (bst *SearchTreeData) MapInt(fn func(int) int) []int {
	result := []int{}
	if bst.left != nil {
		result = bst.left.MapInt(fn)
	}
	result = append(result, fn(bst.data))
	if bst.right != nil {
		result = append(result, bst.right.MapInt(fn)...)
	}
	return result
}
