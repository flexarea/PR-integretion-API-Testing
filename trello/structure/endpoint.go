package structure

type Board struct {
	Endpoint string
}

func LoadEndpoint(endpt string) *Board {
	newBoard := &Board{
		Endpoint: endpt,
	}
	return newBoard
}
