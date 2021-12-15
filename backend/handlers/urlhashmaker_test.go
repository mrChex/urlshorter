package handlers

import (
	"testing"
)

func TestEncDec(t *testing.T) {
	hash, err := encodeURLHash(7, 10123123123)
	if err != nil {
		t.Error(err)
	}

	shardID, urlID, err := decodeURLHash(hash)
	if err != nil {
		t.Error(err)
	}
	if shardID != 7 {
		t.Error("shard != 7")
	}
	if urlID != 10123123123 {
		t.Error("urlID != 10123123123")
	}
}
