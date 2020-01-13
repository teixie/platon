package types

type CollectionInt64 struct {
	elements []int64
	unique   []int64
	exists   map[int64]bool
}

func (c *CollectionInt64) Add(args ...int64) {
	if c.exists == nil {
		c.exists = make(map[int64]bool)
	}
	for _, v := range args {
		if _, ok := c.exists[v]; !ok {
			c.unique = append(c.unique, v)
			c.exists[v] = true
		}
		c.elements = append(c.elements, v)
	}
}

func (c *CollectionInt64) All() []int64 {
	return c.elements
}

func (c *CollectionInt64) Unique() []int64 {
	return c.unique
}
