package types

import "encoding/binary"

const (
	// ModuleName defines the module name
	ModuleName = "billboard"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// RouterKey defines the module's message routing key
	BillboardKey = "Billboard/value/"

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_billboard"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func GetBillboardIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}
