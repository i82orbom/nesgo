package nes

import (
	"os"

	"github.com/i82orbom/nesgo/pkg/gui"
)

// Console represents the NES
type Console struct {
	// Console RAM
	ram [2048]uint8

	// Status
	clockCounter int

	// Devices connected to the Console
	cpu         *CPU
	ppu         *PPU
	apu         *APU
	cart        *Cartridge
	controller1 *controller
	controller2 *controller

	// DMA control variables
	dmaDataBuffer uint8
	dmaTransfer   bool
	dmaSync       bool
	dmaPage       uint8
	dmaAddress    uint8
}

// NewConsole creates an instance of a NES
func NewConsole() *Console {
	cpu := NewCPU()
	ppu := NewPPU()
	apu := NewAPU()

	bus := &Console{
		cpu:         cpu,
		ppu:         ppu,
		apu:         apu,
		controller1: newController(),
		controller2: newController(),

		dmaDataBuffer: 0x00,
		dmaPage:       0x00,
		dmaSync:       true,
		dmaTransfer:   false,
		dmaAddress:    0x00,
	}
	// Connect devices
	cpu.ConnectBus(bus)
	return bus
}

// Reset resets the console
func (b *Console) Reset() {
	b.cpu.reset()
	b.clockCounter = 0
	b.dmaDataBuffer = 0x0
	b.dmaPage = 0x0
	b.dmaTransfer = false
	b.dmaSync = true
	b.ppu.reset()
}

// InsertCartridge connects the cartridge to the console
func (b *Console) InsertCartridge(c *Cartridge) {
	b.cart = c
	b.ppu.InsertCartridge(c)
}

// Step steps the console a single cycle
func (b *Console) Step() {
	b.ppu.Step()
	if (b.clockCounter % 3) == 0 {
		if b.dmaTransfer {
			if b.dmaSync { // Synchronize for the correct clock cycle to init the DMA (odd clock cycle)
				if b.clockCounter%2 == 1 {
					b.dmaSync = false
					// Clean current oam data
					b.dmaAddress = 0
				}
			} else {
				if b.clockCounter%2 == 0 { // DMA Read
					addressToRead := uint16(b.dmaPage)<<8 | uint16(b.dmaAddress)
					b.dmaDataBuffer = b.cpuRead(addressToRead, false)
				} else { // DMA Write
					b.ppu.oamMemory[b.dmaAddress] = b.dmaDataBuffer
					b.dmaAddress++
					if b.dmaAddress == 0x00 { // Address overflow, EO-Transfer
						b.dmaTransfer = false
						b.dmaSync = true
					}
				}
			}
		} else {
			b.cpu.Step()
		}
	}

	// Trigger a NMI if the PPU requests it
	if b.ppu.nmiSignaled {
		b.ppu.nmiSignaled = false
		b.cpu.nmi()
	}
	b.clockCounter++
}

// StepFrame steps the console enough to generate one frame
func (b *Console) StepFrame() {
	for {
		b.Step()
		if b.ppu.frameComplete {
			b.ppu.frameComplete = false
			break
		}
	}
}

// Disassemble shows the currently executed code
func (b *Console) Disassemble() {
	b.cpu.DissasembleCurrentPC(os.Stdout)
}

func (b *Console) cpuRead(address uint16, readOnly bool) uint8 {
	data := uint8(0x00)
	// The cartridge can veto reads to other devices
	if b.cart.CPURead(address, &data) {
		return data
	}

	if address >= 0x0000 && address <= 0x1FFF {
		data = b.ram[address&0x07FF] // Mask for mirroring
	} else if address >= 0x2000 && address <= 0x3FFF { // PPU
		data = b.ppu.CPURead(address&0x0007, readOnly)
	} else if address == 0x4016 {
		data = b.controller1.Data()
	} else if address == 0x4017 {
		data = b.controller2.Data()
	}

	// Read from other devices
	return data
}

func (b *Console) cpuWrite(address uint16, data uint8) {
	// The cartridge can veto writes to other devices
	if b.cart.CPUWrite(address, data) {
		return
	}

	if address >= 0x0000 && address <= 0x1FFF {
		b.ram[address&0x07FF] = data // Mask for mirroring
	} else if address >= 0x2000 && address <= 0x3FFF { // PPU
		b.ppu.CPUWrite(address&0x0007, data)
	} else if (address >= 0x4000 && address <= 0x4013) || address == 0x4015 || address == 0x4017 {
		b.apu.CPUWrite(address, data)
	} else if address == 0x4014 { // DMA transfer
		b.dmaPage = data
		b.dmaTransfer = true
	} else if address == 0x4016 {
		b.controller1.Write()
	} else if address == 0x4017 {
		b.controller2.Write()
	}
}

// TextureProvider returns the texture provider, in this case the PPU
func (b *Console) TextureProvider() gui.TextureProvider {
	return b.ppu
}

// Controller1 returns a handle to the first NES controller
func (b *Console) Controller1() InputController {
	return b.controller1
}

// Controller2 returns a handle to the second NES controller
func (b *Console) Controller2() InputController {
	return b.controller2
}

// AudioSource provides the audio source to the outside
func (b *Console) AudioSource() *APU {
	return b.apu
}
