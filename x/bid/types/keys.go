package types

import "encoding/binary"

const (
	// ModuleName defines the module name
	ModuleName = "bid"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	AuctionKey = "Auction/value/"

	BideKey = "Bid/value/"

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_bid"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func GetAuctionKeyPrefix(billboardId uint64) []byte {
	bz := []byte(AuctionKey)
	bz = append(bz, GetIDBytes(billboardId)...)
	bz = append(bz, []byte("/")...)
	return bz
}

func GetBidKeyPrefix(billboardId uint64, auctionId uint64) []byte {
	bz := []byte(BideKey)
	bz = append(bz, GetIDBytes(billboardId)...)
	bz = append(bz, []byte("/")...)
	bz = append(bz, GetIDBytes(auctionId)...)
	bz = append(bz, []byte("/")...)
	return bz
}

func GetIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}
