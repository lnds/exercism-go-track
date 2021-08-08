package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children [](*Node)
}

func Build(records []Record) (*Node, error) {
	size := len(records)
	if size == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	if records[0].ID > 0 || records[size-1].ID >= size {
		return nil, fmt.Errorf("INVALID RECORDS")
	}

	nodes := make([]*Node, size)
	for i, record := range records {
		if i > 0 && records[i-1] == records[i] {
			return nil, fmt.Errorf("DUPLICATE RECORDS")
		}

		id := record.ID
		parent := record.Parent
		if id < parent {
			return nil, fmt.Errorf("INVALID PARENT")
		}
		if id > 0 && id == parent {
			return nil, fmt.Errorf("CYCLE")
		}
		nodes[i] = &Node{ID: id}
		if id != parent {
			parent := nodes[parent]
			parent.Children = append(parent.Children, nodes[id])
		}
	}
	return nodes[0], nil
}
