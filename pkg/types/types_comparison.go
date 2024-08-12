package types

type comparison struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

func (c *comparison) MarkMax(name string, value int64) {
	if c.Value < value {
		c.Name = name
		c.Value = value
	}
}

func (c *comparison) MarkMin(name string, value int64) {
	if c.Value > value {
		c.Name = name
		c.Value = value
	}
}
