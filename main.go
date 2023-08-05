package main

//https://play.golang.org/p/VE10ZkgGKMM
import (
	"fmt"

	f "play.ground/Core"
)

func main() {
	core := f.GetCore()

	CSICeph := core.GetCSICeph()
	CSICeph.Print("GetCSICeph main()")

	CSILocal := core.GetCSILocal()
	CSILocal.Print("GetCSILocal main()")

	CSIZfs := core.GetCSIZfs()
        CSIZfs.Print("GetCSIZfs main()")


        CNIOvn := core.GetCNIOvn()
        CNIOvn.Print("GetCNIOvn main()")

        CNIOvs := core.GetCNIOvs()
        CNIOvs.Print("GetCNIOvs main()")

        CNIBridge := core.GetCNIBridge()
        CNIBridge.Print("GetCNIBridge main()")

	random := core.GetRand()
	fmt.Printf("%d %d\n", random.Intn(100), random.Intn(100))
}
