package tree

import (
	"sync"
	"testing"

	"github.com/SpotIM/BE-test/entities"
)

func TestNewTreeConcurrent(t *testing.T) {
	for i := 0; i < 10000; i++ {
		tree := NewTree()
		wg := sync.WaitGroup{}
		wg.Add(len(TestData))
		for _, msg := range TestData {
			go func(msg *entities.Msg) {
				defer wg.Done()
				tree.Add(msg)
			}(msg)
		}
		wg.Wait()
		if !validTree(t, tree, TestData) {
			t.Fail()
		}
	}
}

func TestNewTreeLargeConcurrent(t *testing.T) {
	tests := readCsvTests()
	for i := 0; i < 100; i++ {
		tree := NewTree()
		wg := sync.WaitGroup{}
		wg.Add(len(tests))
		for _, msg := range tests {
			go func(msg *entities.Msg) {
				defer wg.Done()
				tree.Add(msg)
			}(msg)
		}
		wg.Wait()
		if !validTree(t, tree, tests) {
			t.Fail()
		}
	}
}
