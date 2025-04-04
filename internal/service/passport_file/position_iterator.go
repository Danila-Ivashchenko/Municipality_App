package passport_file

type PositionIterator struct {
	X       float64
	Y       float64
	newPage bool

	PageParams PageParams
}

func NewPositionIterator(x, y float64, pageParams PageParams) *PositionIterator {
	return &PositionIterator{
		X:          x,
		Y:          y,
		PageParams: pageParams,
		newPage:    true,
	}
}

func (p *PositionIterator) IncrLine(lineSize float64) bool {
	p.X = p.PageParams.LeftField

	if p.Y+lineSize >= p.PageParams.H-p.PageParams.BotField {
		p.Y = p.PageParams.TopField
		return true
	} else {
		p.Y += lineSize
		return false
	}
}

func (p *PositionIterator) IncrLineWithPositionY(lineSize float64) bool {
	p.X = p.PageParams.LeftField

	if p.Y+lineSize >= p.PageParams.H-p.PageParams.BotField {
		p.Y = p.PageParams.TopField + lineSize
		return true
	} else {
		p.Y += lineSize
		return false
	}
}

func (p *PositionIterator) Position() (float64, float64) {
	p.newPage = false
	return p.X, p.Y
}

func (p *PositionIterator) SetNewPosition(x, y float64) {
	p.X = x
	p.Y = y
}

func (p *PositionIterator) IsNewPagePosition() bool {
	return p.newPage
}

func (p *PositionIterator) SetNewPagePosition() {
	p.newPage = true
}

func (p *PositionIterator) WithYStep(step float64) *PositionIterator {
	p.Y += step
	return p
}

func (p *PositionIterator) WithXStep(step float64) *PositionIterator {
	p.X += step
	return p
}

func (p *PositionIterator) IncrXStep(step float64) {
	p.X += step
}

func (p *PositionIterator) IncrYStep(step float64) {
	p.Y += step
}

func (p *PositionIterator) LineWidth() float64 {
	return W - LeftField - RightField
}
