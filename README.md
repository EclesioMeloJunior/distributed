# distributed

A implementation of kademlia dht (distributed hash table) for p2p file sharing. The structure can be defined as nodes with a node id of 20 bits and files with file id of 20 bits. The node has a routing table that can track the neighboors by the closest id using the **XOR** (bitwise or) between 2 nodes id to define the distance.