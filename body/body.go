package body

type Body struct {
	Material    string
	Color       string
	TrunkVolume int
	Roof        bool
}

func NewBody(bodyMaterial, bodyColor string, bodyTrunkVolume int, bodyRoof bool) Body {
	return Body{
		Material:    bodyMaterial,
		Color:       bodyColor,
		TrunkVolume: bodyTrunkVolume,
		Roof:        bodyRoof,
	}
}

func (b *Body) ChangeBodyMaterial(material string) {
	b.Material = material
}

func (b *Body) ChangeBodyColor(color string) {
	b.Color = color
}

func (b *Body) CutRoof() {
	b.Roof = false
}
