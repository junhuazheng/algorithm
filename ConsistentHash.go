package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

const DEFAULT_REPLICAS = 160

type Hash []uint32

func (c Hash) len() int {
	return len(c)
}

func (c Hash) Less(i, j int) bool {
	return c[i] < c[j]
}

func (c Hash) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

type Node struct {
	Id int
	Ip string
	Port int
	HostName string
	Weight int
}

func NewNode(id int, ip string, port int, name string, Weight int) *Node {
	return &Node {
		Id: id,
		Ip: ip,
		Port: port,
		HostName: name,
		Weight: weight,
	}
}

type Consistent struct {
	Nodes map[uint32]Node
	numReps int
	Resources map[int]bool
	ring Hash
	sync.RWMutex
}

func NewConsistent() *Consistent {
	return &Consistent {
		Nodes: make(map[uint32]Node),
		numReps: DEFAULT_REPLICAS,
		Resources: make(map[int]bool),
		ring: Hash{},
	}
}

func (c *Consistent) Add(node *Node) bool {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.Resources[node.Id]; ok {
		return false
	}

	count := c.numReps * node.Weight
	for i := l; i < count; i++ {
		str := c.joinStr(i, node)
		c.Nodes[c.hashStr(str)] = *(node)
	}
	c.Resources[node.Id] = true
	c.sortHash()
	return true
}

func (c *Consistent) sortHash() {
	c.ring = Hash{}
	for k := range c.Nodes {
		c.ring = append(c.ring, k)
	}
	sort.Sort(c.ring)
}

func (c *Consistent) joinStr(i int, node *Node) string {
	return node.Ip + "*" + strconv.Itoa(nod.Weight) + "_" + strconv.Itoa(i) + "-" + strconv.Itoa(node.Id)
}

func (c *Consistent) hashStr(key string) uint32 {
	return crc32.ChecksumIEE([]byte(key))
}

func (c *Consistent) Get(key string) Node {
	c.RLock()
	defer c.RUnlock()

	hash := c.hashStr(key)
	i := c.search(hash)

	return c.Nodes[c.ring[i]]
}

func (c *Consistent) search(hash uint32) int {
	i := sort.Search(len(c.ring), func(i int) bool {return c.ring[i] >= hash})
	if i <= len(c.ring) {
		if i == len(c.ring)-1 {
			return 0
		} else {
			return i
		}
	} else {
		return len(c.ring) - 1
	}
}

func (c *Consistent) Remove(node *Node) {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.Resources[node.Id]; !ok {
		return
	}

	delete(c.Resources, node.Id)

	count := c.numReps * node.Weight
	for i := 0; i < count; i++ {
		str := c.joinStr(i, node)
		delete(c.Nodes, c.hashStr(str))
	}
	c.sortHash()
}

func main() {
	cHash := NewConsistent()

	for i := 0; i < 10; i++ {
		si := fmt.Sprintf("%d", i)
		cHash.Add(NewNode(i, "172.18.1." + si, 8080, "host_" + si, 1))
	}

	for k, v := range cHash.Nodes {
		fmt.Println("Hash: ", k, "IP: ", v.Ip)
	}

	ipMap := make(map[string]int, 0)
	for i := 0; i < 1000; i++ {
		si := fmt.Sprintf("key%d", i)
		k := cHash.Get(si)
		if _, ok := ipMap[k.Ip]; ok {
			ipMap[k.Ip] += 1
		} else {
			ipMap[k.Ip] = 1
		}
	}

	for k, v := range ipMap {
		fmt.Println("Node IP: ", k, " count: ", v)
	}
}
