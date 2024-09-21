package utils

import (
	"fmt"
)

func VMOctetAddIncrement(octet string, increment int) string {
	zerovm := int(octet[len(octet)-1]) - 48 + increment
	newOctet := fmt.Sprintf("%s%d", octet[0:len(octet)-1], zerovm)
	return newOctet
}
