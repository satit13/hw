package util

func SimpleChecksum(data []byte) uint8 {
	var checksum uint8 = 0
	for _, item := range data {
		checksum += uint8(item)
	}
	return -checksum
}

func IsValidChecksum(dataWithChecksum []byte) bool {
	validChecksum := byte(SimpleChecksum(dataWithChecksum[:len(dataWithChecksum)-1]))
	dataChecksum := dataWithChecksum[len(dataWithChecksum)-1]
	if validChecksum == dataChecksum {
		return true
	} else {
		return false
	}
}

func CheckSum8Modulo256(data []byte) byte {
	var checksum uint16 = uint16(data[0])
	for _, item := range data[1:] {
		checksum += uint16(item)
	}
	return byte(checksum % 256)
}
