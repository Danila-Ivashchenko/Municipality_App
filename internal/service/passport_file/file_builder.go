package passport_file

import (
	"github.com/signintech/gopdf"
	"log/slog"
	"strings"
)

const (
	LeftField  = 60
	RightField = 30
	Top        = 40
	BotFiled   = 40

	W = 842
	H = 595

	CommonFontSize = 14
	H1             = 24
	H2             = 20
	H3             = 18

	H1Padding         float64 = 20
	H2Padding         float64 = 15
	H3Padding         float64 = 12
	CommonPadding     float64 = 10
	TablePadding      float64 = 5
	TableBottomMargin float64 = 20

	RedLine float64 = 25

	MinCellWidth = 59.5
)

type PageParams struct {
	LeftField      float64
	RightField     float64
	TopField       float64
	BotField       float64
	W              float64
	H              float64
	baseFontFamily string
}

type FileBuilder struct {
	pdf        *gopdf.GoPdf
	iterator   *PositionIterator
	pageParams PageParams
}

func NewFileBuilder() *FileBuilder {
	f := new(FileBuilder)

	f.pdf = &gopdf.GoPdf{}

	f.pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4Landscape})

	f.pageParams.LeftField = LeftField
	f.pageParams.RightField = RightField
	f.pageParams.TopField = Top
	f.pageParams.BotField = BotFiled

	f.pageParams.W = W
	f.pageParams.H = H

	f.iterator = NewPositionIterator(LeftField, Top, f.pageParams)

	f.NewPage()

	f.pdf.SetXY(f.iterator.Position())

	return f
}

func (f *FileBuilder) UploadFont(filepath, fontFamily string) error {
	err := f.pdf.AddTTFFont(fontFamily, filepath)
	if err != nil {
		return err
	}

	if f.pageParams.baseFontFamily == "" {
		f.pageParams.baseFontFamily = fontFamily
	}

	return nil
}

func (f *FileBuilder) SetFontSize(size float64) error {
	err := f.pdf.SetFont(f.pageParams.baseFontFamily, "", size)
	if err != nil {
		return err
	}

	return nil
}

func (f *FileBuilder) NewPage() {
	f.pdf.AddPage()
}

func (f *FileBuilder) WriteCommonText(text string) error {
	alignment := Justify
	fontSize := CommonFontSize
	padding := CommonPadding
	useRedLine := true

	params := &TextParams{
		Alignment:  &alignment,
		FontSize:   &fontSize,
		LineSize:   &padding,
		UseRedLine: &useRedLine,
	}

	return f.WriteText(text, params)
}

func (f *FileBuilder) WriteH1(text string) error {
	alignment := Center
	fontSize := H1
	padding := H1Padding

	params := &TextParams{
		Alignment: &alignment,
		FontSize:  &fontSize,
		LineSize:  &padding,
	}

	return f.WriteText(text, params)
}

func (f *FileBuilder) WriteH2(text string) error {
	alignment := Right
	fontSize := H2
	padding := H2Padding

	params := &TextParams{
		Alignment: &alignment,
		FontSize:  &fontSize,
		LineSize:  &padding,
	}

	return f.WriteText(text, params)
}

func (f *FileBuilder) WriteH3(text string) error {
	alignment := Right
	fontSize := H3
	padding := H3Padding

	params := &TextParams{
		Alignment: &alignment,
		FontSize:  &fontSize,
		LineSize:  &padding,
	}

	return f.WriteText(text, params)
}

func (f *FileBuilder) WriteText(text string, params *TextParams) error {
	var (
		alignment       = Right
		fontSize        = CommonFontSize
		currentLineSize = CommonPadding
		useRedLine      = false
	)

	if params.Alignment != nil {
		alignment = *params.Alignment
	}

	if params.FontSize != nil {
		fontSize = *params.FontSize
	}

	if params.UseRedLine != nil {
		useRedLine = *params.UseRedLine
	}

	if params.LineSize != nil {
		currentLineSize = *params.LineSize
	}

	err := f.SetFontSize(float64(fontSize))
	if err != nil {
		return err
	}

	wrappedLines, err := f.pdf.SplitTextWithWordWrap(text, f.iterator.LineWidth())
	if err != nil {
		return err
	}

	for i, line := range wrappedLines {
		if i == len(wrappedLines)-1 {
			alignment = Right
		}

		if i > 0 && useRedLine {
			useRedLine = false
		}

		f.IncrLine(currentLineSize)
		f.writeTextWithAlignment(line, alignment, useRedLine)
		f.IncrLine(currentLineSize)
	}

	return nil
}

func (f *FileBuilder) writeTextWithAlignment(text string, alignment TextAlignment, redLine bool) {
	switch alignment {
	case Justify:
		if redLine {
			f.JustifyTextWithRedLine(text)
		} else {
			f.JustifyText(text)
		}
	case Center:
		f.CenterText(text)
	case Right:
		if redLine {
			f.RightTextWithRedLine(text)
		} else {
			f.RightText(text)
		}
	default:
		if redLine {
			f.RightTextWithRedLine(text)
		} else {
			f.RightText(text)
		}
	}
	return
}

func (f *FileBuilder) Save(filename string) error {
	return f.pdf.WritePdf(filename)
}

func (f *FileBuilder) IncrLine(lineSize float64) {
	if f.iterator.IsNewPagePosition() {
		return
	}

	if toMakeNewPage := f.iterator.IncrLine(lineSize); toMakeNewPage {
		f.NewPage()
		f.iterator.SetNewPagePosition()
	}
}

func (f *FileBuilder) IncrLineWithPositionY(lineSize float64) {
	if f.iterator.IsNewPagePosition() {
		return
	}

	if toMakeNewPage := f.iterator.IncrLineWithPositionY(lineSize); toMakeNewPage {
		f.NewPage()
		f.iterator.SetNewPagePosition()
	}
}

func (f *FileBuilder) RightText(text string) {
	var (
		err error
	)

	f.pdf.SetXY(f.iterator.Position())
	err = f.pdf.Text(text)
	if err != nil {
		slog.Warn("error writing text to pdf")
	}
}

func (f *FileBuilder) RightTextWithRedLine(text string) {
	var (
		err error
	)

	f.iterator.IncrXStep(RedLine)
	f.pdf.SetXY(f.iterator.Position())

	err = f.pdf.Text(text)
	if err != nil {
		slog.Warn("error writing text to pdf")
	}
}

func (f *FileBuilder) CenterText(text string) {
	var (
		totalWordWidth float64
		step           float64
		err            error
	)

	words := strings.Fields(text)

	for _, word := range words {
		lineWidth, _ := f.pdf.MeasureTextWidth(word)
		totalWordWidth += lineWidth
	}

	spaceDelta := f.iterator.LineWidth() - totalWordWidth
	step = spaceDelta / 2

	f.pdf.SetXY(f.iterator.WithXStep(step).Position())
	err = f.pdf.Text(text)
	if err != nil {
		slog.Warn("error writing text to pdf")
	}
}

func (f *FileBuilder) JustifyText(text string) {
	var (
		totalWordWidth float64
		step           float64
		err            error
	)

	f.pdf.SetXY(f.iterator.Position())

	words := strings.Fields(text)
	if len(words) < 2 {
		err = f.pdf.Text(text)
		if err != nil {
			slog.Warn("error writing text to pdf")
		}
	}

	for _, word := range words {
		lineWidth, _ := f.pdf.MeasureTextWidth(word)
		totalWordWidth += lineWidth
	}

	spaceDelta := f.iterator.LineWidth() - totalWordWidth
	step = spaceDelta / float64(len(words)-1)

	for _, word := range words {
		wordWidth, _ := f.pdf.MeasureTextWidth(word)
		err = f.pdf.Text(word)
		if err != nil {
			slog.Warn("error writing text to pdf")
		}

		f.pdf.SetXY(f.iterator.WithXStep(step + wordWidth).Position())
	}
}

func (f *FileBuilder) JustifyTextWithRedLine(text string) {
	var (
		totalWordWidth float64
		step           float64
		err            error
	)

	f.iterator.IncrXStep(RedLine)
	f.pdf.SetXY(f.iterator.Position())

	words := strings.Fields(text)
	if len(words) < 2 {
		err = f.pdf.Text(text)
		if err != nil {
			slog.Warn("error writing text to pdf")
		}
	}

	for _, word := range words {
		lineWidth, _ := f.pdf.MeasureTextWidth(word)
		totalWordWidth += lineWidth
	}

	spaceDelta := f.iterator.LineWidth() - totalWordWidth
	step = spaceDelta / float64(len(words)-1)

	for _, word := range words {
		wordWidth, _ := f.pdf.MeasureTextWidth(word)
		err = f.pdf.Text(word)
		if err != nil {
			slog.Warn("error writing text to pdf")
		}

		f.pdf.SetXY(f.iterator.WithXStep(step + wordWidth).Position())
	}
}
