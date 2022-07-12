package main

import (
	"encoding/binary"
	"fmt"
	"log"

	"anno1800mem/pkg/anno"
	"anno1800mem/pkg/mem"

	"golang.org/x/sys/windows"
)

func main() {
	procID, err := mem.GetProcessIDByProcessName("Anno1800.exe")
	if err != nil {
		log.Fatal(err)
	}
	modBaseAddr, err := mem.GetModuleEntryByProcID(procID)
	if err != nil {
		log.Fatal(err)
	}
	handle, err := windows.OpenProcess(windows.PROCESS_VM_READ, false, procID)
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range anno.Inhabitants {
		guidAddr, err := mem.FindDMAAddy(handle, modBaseAddr+i.GuidAddr[0], i.GuidAddr[1:]...)
		if err != nil {
			log.Fatal(err)
		}
		var buf [4]byte
		if err := windows.ReadProcessMemory(handle, guidAddr, &buf[0], 4, nil); err != nil {
			log.Fatal(err)
		}
		if int(binary.LittleEndian.Uint32(buf[:])) == i.Guid {
			inhabAddr, err := mem.FindDMAAddy(handle, modBaseAddr+i.Addr[0], i.Addr[1:]...)
			if err != nil {
				log.Fatal(err)
			}
			if err := windows.ReadProcessMemory(handle, inhabAddr, &buf[0], 4, nil); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s: %d\n", i.Name, binary.LittleEndian.Uint32(buf[:]))
		}
	}
}
