package tree

import "github.com/SpotIM/BE-test/entities"

type node struct {
	ID       string
	ParentID string
	data     *entities.Msg
	children []*node
}

func newNode(msg *entities.Msg) *node {
	return &node{
		ID:       msg.ID,
		ParentID: msg.ParentID,
		data:     msg,
		children: make([]*node, 0),
	}
}
