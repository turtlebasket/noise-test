package types

import (
	"encoding/json"
	"fmt"
)

type PStatus int

const (
	Unpaid PStatus = iota
	Pending
	Verified
	Error
)

type PMessageResponse struct {
	Status PStatus
	Amount float32
	To     string
}

func (p *PMessageResponse) Print() {
	fmt.Printf("Status: %s", fmt.Sprint(p.Status))
}

func (m PMessageResponse) Marshal() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func UnmarshalPMessageResponse(buf []byte) (PMessageResponse, error) {
	res := PMessageResponse{}
	if err := json.Unmarshal(buf, &res); err != nil {
		return res, err
	}
	return res, nil
}
