package main

import "log"

type RouteTrie[T any] struct {
	root RouteTrieNode[T]
}

type RouteTrieNode[T any] struct {
	children map[rune]RouteTrieNode[T]
	hasValue bool
	value    T
}

func (t *RouteTrie[T]) put(route string) {
	log.Fatal("function not implemented")

}

func (t *RouteTrie[T]) get(route string) (T, int, []string, bool) {
    //todo when run into ? or numbers, separate out arguments appropriately
    //todo stop when # is hit
    found := false
    var val T
    var id int
    args := make([]string, 0)

    currentNode := t.root
    for _, r := range route {
        if currentNode.hasValue {
            found = true 
            val = currentNode.value
        }
        child, ok := currentNode.children[r]
        if ok {
            currentNode = child
        }
    }

	return val, id, args, found
}
