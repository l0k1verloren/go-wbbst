// Package wbbst - A Top Heavy Weight Balanced Binary Search Tree for sorting small, largely random blobs of data
package wbbst

// Wbbst - Top Heavy Weight Balanced Binary Search Tree
// Tree is encoded by an array with uint32 indices. Data type is unspecified and defined by the implementation.
type Wbbst interface {
	// Comparator functions
	IsLeft(interface{}) bool
	IsRight(interface{}) bool
	// Insert a new node in the tree and balance if necessary
	Insert(interface{}) uint32
	// Removes item at index and rebalances if necessary
	DeleteByIndex(uint32)
	// Searches for a node with specified data and rebalances if necessary
	DeleteByData(interface{}) error
	// Walking functions. These take an index and return the correct index of the data on this path
	WalkUp(uint32) uint32
	WalkRight(uint32) uint32
	WalkLeft(uint32) uint32
}

// Tree - a generic structure for storing a short data type for the search tree
type Tree struct {
	Weight [2]uint32
	Depth  uint8
	// The store is tree structured only by the inherent structure of the tree (it is just rows of progressively doubling strings)
	Store []interface{}
}