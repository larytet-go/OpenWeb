package tree

import "github.com/SpotIM/BE-test/entities"

type Tree struct {
	children []*node
}

func NewTree() *Tree {
	return &Tree{
		children: make([]*node, 0),
	}
}

func (t *Tree) Add(msg *entities.Msg) {

}
