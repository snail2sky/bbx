package car

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

const coefficient = 802018
const requireLength = 6

func GetArrizoPassword(serialNum string) (string, error) {
	if len(serialNum) < requireLength {
		errMsg := fmt.Sprintf("Serial number length is too short, must be at least %d", requireLength)
		log.Println(errMsg)
		return "", errors.New(errMsg)
	}
	last6Str := serialNum[len(serialNum)-6:]
	last6Int, err := strconv.Atoi(last6Str)
	if err != nil {
		errMsg := fmt.Sprintln("Serial number is invalid, must contain 6 digits")
		log.Println(errMsg)
		return "", errors.New(errMsg)
	}
	result := last6Int * coefficient
	resultStr := strconv.Itoa(result)
	password := resultStr[len(resultStr)-6:]
	return password, nil
}
