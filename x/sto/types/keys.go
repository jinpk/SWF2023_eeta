package types

import (
	"encoding/binary"
	"fmt"
)

const (
	// ModuleName defines the module name
	ModuleName = "sto"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	StoKey = "Sto/value/"

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_sto"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func GetStoKeyPrefix(billboardId uint64) []byte {
	bz := []byte(StoKey)
	bz = append(bz, GetIDBytes(billboardId)...)
	bz = append(bz, []byte("/")...)
	return bz
}

func GetIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func GetStoTokenDenom(id uint64) string {
	return fmt.Sprintf("sto/coin/%d", id)
}
