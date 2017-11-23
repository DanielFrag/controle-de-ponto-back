package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"math"
	"math/rand"
	"time"
)

//FormatJSON format the map[string]interface{} to json
func FormatJSON(m map[string]interface{}) []byte {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(m)
	return b.Bytes()
}

//EncryptWithXor apply xor between the payload and key content using fillers's values to complete the difference between length of payload and key
func EncryptWithXor(payload, key, fillers []byte) []byte {
	rand.Seed(time.Now().UnixNano())
	if len(payload) <= len(key) {
		result := make([]byte, len(key))
		for i := 0; i < len(key); i++ {
			if i < len(payload) {
				result[i] = payload[i] ^ key[i]
			} else {
				fillerIndex := rand.Intn(len(fillers))
				result[i] = fillers [fillerIndex] ^ key[i]
			}
		}
		return result
	} else {
		steps := int(math.Ceil(float64(len(payload)) / float64(len(key))))
		var result []byte
		for i := 0; i < steps; i++ {
			chunkStartIndex := i * len(key)
			chunkFinishIndex := (i + 1) * len(key)
			if chunkFinishIndex > len(payload) {
				chunkFinishIndex = len(payload)
			}
			result = append(result, EncryptWithXor(payload[chunkStartIndex:chunkFinishIndex], key, fillers)...)
		}
		return result
	}
}

//ConvertInt64ToByte parse int64 to byte slice
func ConvertInt64ToByte(num int64) []byte {
	intBuffer := make([]byte, binary.MaxVarintLen64)
	length := binary.PutVarint(intBuffer, num)
	return intBuffer[:length]
} 

//GetInt64FromByte parse slice to int64
func GetInt64FromByte(b []byte) (int64, error) {
	num, readedBytes := binary.Varint(b)
	if readedBytes != len(b) {
		return num, errors.New("The byte slice length does not match the int64 length")
	}
	return num, nil
}

//ConvertUint64ToByte parse uint64 to byte slice
func ConvertUint64ToByte(num uint64) []byte {
	intBuffer := make([]byte, binary.MaxVarintLen64)
	length := binary.PutUvarint(intBuffer, num)
	return intBuffer[:length]
} 

//GetUint64FromByte parse slice to uint64
func GetUint64FromByte(b []byte) (uint64, error) {
	num, readedBytes := binary.Uvarint(b)
	if readedBytes != len(b) {
		return num, errors.New("The byte slice length does not match the uint64 length")
	}
	return num, nil
}
