package types

import (
	"encoding/json"
	"fmt"
)

type PMessage struct {
	To     string
	Amount float32
}

func (p *PMessage) Print() {
	fmt.Printf("To: %s\nAmount: %f", p.To, p.Amount)
}

func (m PMessage) Marshal() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func UnmarshalPMessage(buf []byte) (PMessage, error) {
	res := PMessage{}
	if err := json.Unmarshal(buf, &res); err != nil {
		return res, err
	}

	return res, nil
}
