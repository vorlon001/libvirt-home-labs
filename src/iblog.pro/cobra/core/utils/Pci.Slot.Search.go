package utils

import (
        "strconv"
        "strings"
)


func SearchSlotInt[Z int, T []Z](slots T) Z {
        var done Z
        done = 0
        for k, _ := range slots {
                if k < len(slots)-1 && done == 0 {
                        if slots[k+1]-slots[k] > 1 {
                                done = slots[k] + 1
                        }
                }
        }
        if done == 0 {
                done = slots[len(slots)-1] + 1
        }
        return done
}

func HexToInt(hex string) int {
        value, err := strconv.ParseInt(strings.Replace(hex, "0x", "", -1), 16, 64)
        if err != nil {
                return -1
        }
        return int(value)
}

func HexsToInts(hexs []string) []int {
        ints := make([]int, len(hexs))
        for k, v := range hexs {
                ints[k] = HexToInt(v)
        }
        return ints
}

func SearchSlot(hexs []string) int {
        return SearchSlotInt(HexsToInts(hexs))
}
