package codec

import "github.com/ethereum/go-ethereum/rlp"

type RLPEncoder struct {}

func (r *RLPEncoder) Marshal(v interface{}) ([]byte, error) {
	a, err := rlp.EncodeToBytes(v)
	return a, err
}

func (r *RLPEncoder) Unmarshal(data []byte, v interface{}) error {
	a := rlp.DecodeBytes(data, v)
	return a
}

func (r *RLPEncoder) Name() string {
	return "rlp"
}
