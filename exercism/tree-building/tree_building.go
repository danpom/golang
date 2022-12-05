package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	n := make([]*Node, len(records))
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	// sort.Slice(records, func(i, j int) bool {
	// 	return records[i].Parent < records[j].Parent
	// })
	for i, r := range records {
		if r.ID != i || r.Parent > i || i > 0 && r.Parent == i {
			return nil, errors.New("invalild input")
		}

		n[i] = &Node{ID: i}
		if i > 0 {
			n[r.Parent].Children = append(n[r.Parent].Children, n[i])
		}
	}

	//return &Node{ID : records[0].ID, Children: n }, nil
	return n[0], nil
}
