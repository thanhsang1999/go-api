package common

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcutil/base58"
	"strconv"
	"strings"
)

type UID struct {
	localID    uint32
	objectType int
	shardID    uint32
}

func NewUID(localID uint32, objectType int, shardID uint32) UID {
	return UID{
		localID:    localID,
		objectType: objectType,
		shardID:    shardID,
	}
}

func (uid UID) String() string {
	val := uint64(uid.localID)<<28 | uint64(uid.objectType)<<18 | uint64(uid.shardID)<<0
	return base58.Encode([]byte(fmt.Sprintf("%v", val)))
}
func (uid UID) GetLocalID() uint32 {
	return uid.localID
}
func (uid UID) GetObjectType() int {
	return uid.objectType
}
func (uid UID) GetShardId() uint32 {
	return uid.shardID
}
func DecomposeUID(s string) (UID, error) {
	uid, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		return UID{}, err
	}
	if (1 << 18) > uid {
		return UID{}, errors.New("wrong uid")
	}
	u := UID{
		localID:    uint32(uid >> 28),
		objectType: int(uid >> 18 & 0x3FF),
		shardID:    uint32(uid >> 0 & 0x3FFFF),
	}
	return u, nil
}
func FromBase58(s string) (UID, error) {
	return DecomposeUID(string(base58.Decode(s)))
}
func (uid UID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", uid.String())), nil
}
func (uid *UID) UnmarshalJSON(data []byte) error {
	decodeUID, err := FromBase58(strings.Replace(string(data), "\"", "", -1))

	if err != nil {
		return err
	}
	uid.localID = decodeUID.localID
	uid.objectType = decodeUID.objectType
	uid.shardID = decodeUID.shardID

	return nil
}

func (uid *UID) Value() (driver.Value, error) {
	if uid == nil {
		return nil, nil
	}
	return int64(uid.localID), nil
}
func (uid *UID) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	var i uint32

	switch t := value.(type) {
	case int:
	case int8:
	case int16:
	case int32:
	case int64:
	case uint8:
	case uint16:
	case uint64:
		i = uint32(t)
	case uint32:
		i = t
	case []byte:
		a, err := strconv.Atoi(string(t))
		if err != nil {
			return err
		}
		i = uint32(a)
	default:
		return errors.New("invalid Scan Source")
	}
	*uid = NewUID(i, 0, 1)
	return nil
}
