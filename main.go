package main

import (
	mypackage "github.com/Enderstyles/mypackage"
)
func main(){
	var a int = 1;
	var b int = 2;
	
	mypackage.MyPrintInt(mypackage.MaxInt(a,b))

	pc := new(mypackage.Pc)
	pc = mypackage.NewPc();
	mypackage.AddRam(pc, 8)

	mypackage.MyPrintStr("\nCpu_name:")
	mypackage.MyPrintStr(pc.Cpu_name)
	mypackage.MyPrintStr("\nGpu_name:")
	mypackage.MyPrintStr(pc.Gpu_name)
	mypackage.MyPrintStr("\nRam_amount:")
	mypackage.MyPrintInt(pc.Ram_amount);


	ar := new(mypackage.Ar)
	ar = mypackage.NewAr();
	mypackage.AddMagazineCapacity(ar,30)
	mypackage.ChangeScope(ar,4)
}