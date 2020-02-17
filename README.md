# NES-GO

Another NES Emulator implemented in Golang

# Dependencies

MacOS
```
brew install glfw
```

Linux/Unix Based Systems or Windows

Refer to https://www.glfw.org/


# Usage

```
go run cmd/nesgo/main.go <path-to-rom>
```

## Key bindings

Currently the keys are statically mapped, dynamic mapping through a configuration file will be supported in the future. Only the first controller is supported.

### Controller 1 (keyboard)

- A - N
- B - M
- Up/Down/Left/Right - WSAD
- Start - Enter
- Select - Backspace

### Special keys

- E - enable/disable emulation
- L - enable/disable disassembler in stdout
- Space - Step one frame
- O - Cycle through: PPU rendered output - Pattern table 1 - Pattern table 2
- P - Cycle through palette index 0-7 (when showing pattern tables)

## Progress:

- [X] - CPU implementation
- [X] - Basic cartridge / mapper implementation 
- [X] - GUI framework: glfw/gl
- [X] - Basic controller support
- [X] - PPU background rendering
- [ ] - PPU foreground rendering (sprites)
- [ ] - APU (audio processing unit)
- [ ] - Multiple controller mappings
- [ ] - Dynamic controller mapping
- [ ] - Save state
- [ ] - Battery support
- [ ] - More mappers... (contributions are appreciated) 

### Notes:

- ##### The code is set up as simple as possible, depicting all steps in emulation, specially when rendering the image in the PPU, optimisations are ommited intentionally

### References:

- http://wiki.nesdev.com/w/index.php/Nesdev_Wiki
- http://archive.6502.org/publications/pet_paper/pet_paper_v3_i2_i3.pdf
- https://www.cs.otago.ac.nz/cosc243/pdf/6502Poster.pdf

##### Special thanks to https://github.com/OneLoneCoder and his YouTube channel, this implementation it is based on his hard work creating such an amazing step-by-step guide to implement this emulator