package main

import (
	"fmt"
)

func SearchIntToDiskName(id int) string {
        return fmt.Sprintf("sd%c", id+96)
}

func SearchDisk(disk []string) string {

        disksId := make([]int, len(disk))
        for k, v := range disk {
                if len(v) == 3 {
                        disksId[k] = int(v[2]) - 96
                }
        }
        return SearchIntToDiskName(SearchSlotInt(disksId))
        //SearchSlotInt(HexsToInts(hexs))
}
