package mycount

type Sansi struct {
	TempCurrentHeight uint64
	TempCurrentTotalBlock uint64
	TempCurrentTotalCount uint64

	TempF03176TotalCount uint64
	TempF03176TotalBlock uint64
}

var TempMinerArray  = make(map[string]string)