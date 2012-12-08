package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Allocate memory.
	memorySize := 1024*64
	memory := make([]byte, memorySize)
	fmt.Printf("Memory has %d bytes.\n", len(memory))

	// Load ROM into memory.
	romFilename := "roms/level2.rom"
	rom, err := ioutil.ReadFile(romFilename)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ROM has %d bytes.\n", len(rom))

	// Copy ROM into memory.
	copy(memory, rom)

	// Make a CPU.
	cpu := &cpu{memory: memory, romSize: uint16(len(rom))}
	cpu.loadInstructions(instructionList)

	// Make it go.
	fmt.Println("Booting")
	cpu.run()
}