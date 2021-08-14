package client

import (
	"fmt"
)

func NewEncoder(c *Client) *Encoder {
	return &Encoder{c}
}

type Encoder struct {
	*Client
}

func (e *Encoder) Encode() string {
	return fmt.Sprintf("%d_%d", e.userId, e.osType)
}
