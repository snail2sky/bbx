package car

import (
	"log"
	"strconv"
)

const coefficient = 802018
const requireLength = 6

func GetArrizoPassword(serialNum string) string {
	if len(serialNum) < requireLength {
		log.Fatalf("serial number length is too short, must be at least %d", requireLength)
	}
	last6Str := serialNum[len(serialNum)-6:]
	last6Int, _ := strconv.Atoi(last6Str)
	result := last6Int * coefficient
	resultStr := strconv.Itoa(result)
	password := resultStr[len(resultStr)-6:]
	return password
}
