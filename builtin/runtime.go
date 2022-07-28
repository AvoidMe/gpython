package builtin

import (
	"hash/maphash"
)

var globalSeed *maphash.Seed

func GetPyHashSeed() *maphash.Seed {
	if globalSeed == nil {
		newSeed := maphash.MakeSeed()
		globalSeed = &newSeed
	}
	return globalSeed
}
