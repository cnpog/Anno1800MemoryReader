package anno

type Inhabitant struct {
	Name     string
	Addr     []uintptr
	GuidAddr []uintptr
	Guid     int
}

var Inhabitants []Inhabitant = []Inhabitant{
	{
		Name:     "Bauern",
		Addr:     []uintptr{0x05832918, 0x50, 0x8, 0x4},
		GuidAddr: []uintptr{0x05832918, 0x50, 0x8, 0x8},
		Guid:     15000000,
	},
	{
		Name:     "Arbeiter",
		Addr:     []uintptr{0x05832918, 0x50, 0x8, 0xc},
		GuidAddr: []uintptr{0x05832918, 0x50, 0x8, 0x10},
		Guid:     15000001,
	},
	{
		Name:     "Handwerker",
		Addr:     []uintptr{0x05832918, 0x50, 0x8, 0x14},
		GuidAddr: []uintptr{0x05832918, 0x50, 0x8, 0x18},
		Guid:     15000002,
	},
	{
		Name:     "Ingenieure",
		Addr:     []uintptr{0x05832918, 0x50, 0x8, 0x1c},
		GuidAddr: []uintptr{0x05832918, 0x50, 0x8, 0x20},
		Guid:     15000003,
	},
	{
		Name:     "Investoren",
		Addr:     []uintptr{0x05832918, 0x50, 0x8, 0x24},
		GuidAddr: []uintptr{0x05832918, 0x50, 0x8, 0x28},
		Guid:     15000004,
	},
	{
		Name:     "Jornaleros",
		Addr:     []uintptr{0x05832918, 0x50, 0x8, 0x2c},
		GuidAddr: []uintptr{0x05832918, 0x50, 0x8, 0x30},
		Guid:     15000005,
	},
	{
		Name:     "Oberos",
		Addr:     []uintptr{0x05832918, 0x50, 0x8, 0x34},
		GuidAddr: []uintptr{0x05832918, 0x50, 0x8, 0x38},
		Guid:     15000006,
	},
}
