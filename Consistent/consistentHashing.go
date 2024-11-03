package main

import (
	"crypto/sha256"
	"fmt"
	"sort"
)

// StorageNode represents a storage node in the hash ring.
type StorageNode struct {
	Name string
	Host string
}

// ConsistentHash implements consistent hashing.
type ConsistentHash struct {
	Keys       []int                // Sorted keys in the hash ring.
	Nodes      map[int]*StorageNode // Node mapping with hashed keys.
	TotalSlots int                  // Total slots in the ring (hash space).
}

// NewConsistentHash initializes a ConsistentHash struct.
func NewConsistentHash(totalSlots int) *ConsistentHash {
	return &ConsistentHash{
		Nodes:      make(map[int]*StorageNode),
		TotalSlots: totalSlots,
	}
}

// hashFn calculates a hash for a given key and maps it to the total slots.
func hashFn(key string, totalSlots int) int {
	h := sha256.New()
	h.Write([]byte(key))
	hashValue := int(h.Sum(nil)[0]) // Using the first byte for simplicity
	return hashValue % totalSlots
}

// AddNode adds a node to the hash ring.
func (ch *ConsistentHash) AddNode(node *StorageNode) int {
	key := hashFn(node.Host, ch.TotalSlots)
	// Avoid collision by finding a new slot if the key is already taken.
	for {
		if _, exists := ch.Nodes[key]; !exists {
			break
		}
		key = (key + 1) % ch.TotalSlots
	}
	ch.Nodes[key] = node
	ch.Keys = append(ch.Keys, key)
	sort.Ints(ch.Keys) // Keep the ring sorted
	return key
}

// RemoveNode removes a node from the hash ring.
func (ch *ConsistentHash) RemoveNode(node *StorageNode) {
	key := hashFn(node.Host, ch.TotalSlots)
	index := sort.SearchInts(ch.Keys, key)
	if index < len(ch.Keys) && ch.Keys[index] == key {
		ch.Keys = append(ch.Keys[:index], ch.Keys[index+1:]...)
		delete(ch.Nodes, key)
	}
}

// Assign assigns an item to a node based on consistent hashing.
func (ch *ConsistentHash) Assign(item string) *StorageNode {
	key := hashFn(item, ch.TotalSlots)
	// Find the first node to the right of the hash
	index := sort.SearchInts(ch.Keys, key)
	if index == len(ch.Keys) { // Wrap around if necessary
		index = 0
	}
	return ch.Nodes[ch.Keys[index]]
}

func main() {
	storageNodes := []*StorageNode{
		{Name: "A", Host: "239.67.52.72"},
		{Name: "B", Host: "137.70.131.229"},
		{Name: "C", Host: "98.5.87.182"},
		{Name: "D", Host: "11.225.158.95"},
		{Name: "E", Host: "203.187.116.210"},
	}

	ch := NewConsistentHash(50)
	for _, node := range storageNodes {
		key := ch.AddNode(node)
		fmt.Printf("Added node %s at key %d\n", node.Name, key)
	}

	files := []string{"f1ejbfihwefnbeiuweboucweoie.txt", "f2.txt", "f3.txt", "f4.txt", "f5.txt"}
	for _, file := range files {
		node := ch.Assign(file)
		fmt.Printf("File %s is assigned to node %s\n", file, node.Name)
	}
}
