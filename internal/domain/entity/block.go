package entity

const (
	Difficulty = 1
)

type (
	Block struct {
		Index      int64
		TimeStamp  string
		Data       int64
		Hash       string
		PrevHash   string
		Difficulty int64
		Nonce      string
	}
)
