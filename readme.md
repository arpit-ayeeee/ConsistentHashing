# Consistent Hashing in Go

This package provides a simple implementation of consistent hashing in Go. Consistent hashing is a technique for distributing items (like files or requests) across nodes in a way that minimizes reassignments when nodes are added or removed, making it useful in distributed systems.

## Overview

This package allows items (e.g., files) to be distributed across a ring of storage nodes based on a consistent hash function. It includes methods to add or remove nodes from the ring while ensuring minimal redistribution of items across nodes.

The code includes:
- **StorageNode**: A struct representing a storage node in the system.
- **ConsistentHash**: The main struct implementing the consistent hashing ring.
- **Main functions**: `AddNode`, `RemoveNode`, and `Assign`, which handle node management and item assignment.

## Functions

### `NewConsistentHash(totalSlots int) *ConsistentHash`

Initializes a `ConsistentHash` object with a specified number of slots.

- **Parameters**: `totalSlots` - The total number of slots in the hash ring.
- **Returns**: A pointer to a new `ConsistentHash` instance.

### `AddNode(node *StorageNode) int`

Adds a new `StorageNode` to the hash ring at a hashed position.

- **Parameters**: `node` - A pointer to the `StorageNode` to add.
- **Returns**: The key (position in the ring) where the node was placed.

### `RemoveNode(node *StorageNode)`

Removes a `StorageNode` from the hash ring.

- **Parameters**: `node` - A pointer to the `StorageNode` to remove.

### `Assign(item string) *StorageNode`

Finds the appropriate node for an item based on the item’s hash value.

- **Parameters**: `item` - A string representing the item to assign (e.g., a file name).
- **Returns**: A pointer to the `StorageNode` where the item should reside.

## Usage

1. Initialize a consistent hash ring with a set number of slots.
2. Add storage nodes to the hash ring.
3. Assign items to nodes based on consistent hashing.


### Output Example

```
Added node A at key 5
Added node B at key 12
Added node C at key 28
Added node D at key 37
Added node E at key 45
File f1.txt is assigned to node C
File f2.txt is assigned to node B
File f3.txt is assigned to node E
File f4.txt is assigned to node A
File f5.txt is assigned to node D
```

## Hashing Details

This implementation uses SHA-256 to generate a hash from the node's host and the item’s name. The hash value is then mapped to one of the slots in the ring by taking a modulo of the `totalSlots`.

## Notes

- This implementation assumes `totalSlots` is a manageable number for the system, where `hashFn` maps nodes and items uniformly across the ring.
- For production-grade systems, consider adding virtual nodes or replicas for better load balancing across nodes.

## License

This code is open-source and available under the MIT License.