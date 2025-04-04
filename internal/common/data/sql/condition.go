package sql

import (
	"fmt"
	"strings"
)

func NewCondition() *Condition {
	return &Condition{
		addSeparator: false,
		argsCount:    0,
		condition:    "",
	}
}

type Condition struct {
	condition    string
	addSeparator bool
	argsCount    int
	args         []interface{}
}

func (c *Condition) ArgsCount() int { return c.argsCount }

func (c *Condition) Condition() string { return c.condition }

func (c *Condition) Args() []interface{} { return c.args }

func (c *Condition) AddCondition(subCondition, separator string, arg any) {
	subCondition = strings.TrimSpace(subCondition)
	separator = strings.TrimSpace(separator)

	if c.addSeparator {
		c.condition += " " + separator
	}

	c.condition += " " + fmt.Sprintf(subCondition, c.argsCount+1)

	if arg != nil {
		c.args = append(c.args, arg)
		c.argsCount++
	}

	c.addSeparator = len(c.condition) > 0
}
