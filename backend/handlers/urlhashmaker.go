package handlers

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"strconv"
)

func encodeURLHash(shard int, urlID int64) (string, error) {
	if shard < 0 || shard > 9 {
		return "", errors.New("shard only one char allowed")
	}

	hash := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d%d", shard, urlID)))
	return hash, nil
}

func decodeURLHash(hash string) (shard int, urlID int64, err error) {
	if hash == "" {
		err = errors.New("hash not provided")
		return
	}
	sDec, err := base64.StdEncoding.DecodeString(hash)
	log.Println("s", string(sDec), err, hash)
	if err != nil {
		return
	}

	if len(sDec) < 2 {
		err = errors.New("wrong hash")
		return
	}

	shardString := string(sDec[:1])
	urlIDString := string(sDec[1:])

	shard_, err := strconv.ParseInt(shardString, 10, 0)
	if err != nil {
		return
	}
	shard = int(shard_)

	urlID, err = strconv.ParseInt(urlIDString, 10, 64)
	return
}
