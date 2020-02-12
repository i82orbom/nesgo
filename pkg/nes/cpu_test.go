package nes

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_CPU(t *testing.T) {
	cartridge, err := NewCartridge("../../testroms/nestest.nes")
	if err != nil {
		log.Fatalf(err.Error())
	}
	nes := NewConsole()
	nes.InsertCartridge(cartridge)
	nes.Reset()
	nes.cpu.pc = 0x00C0

	stop := false
	go func() {
		for {
			nes.Step()
			if stop {
				break
			}
		}
	}()

	time.Sleep(time.Second * 2)
	stop = true
	time.Sleep(time.Millisecond * 100)

	assert.Equal(t, uint8(0x0), nes.ram[0x02])
	assert.Equal(t, uint8(0x0), nes.ram[0x03])
}
