package passport_file

import "log/slog"

func (f *FileBuilder) CalculateCellHeight(text string, fontSize, width, padding float64) float64 {
	err := f.pdf.SetFontSize(fontSize)
	if err != nil {
		slog.Warn(err.Error())
	}

	textWidth, err := f.pdf.MeasureTextWidth(text)
	if err != nil {
		slog.Warn(err.Error())
	}

	if textWidth <= width-padding*2 {
		return fontSize + padding*2
	}

	rows, err := f.pdf.SplitTextWithWordWrap(text, width-padding*2)
	if err != nil {
		slog.Warn(err.Error())
	}

	return fontSize*float64(len(rows)) + padding*2
}

func (f *FileBuilder) WriteTextToCell(text string, x, y, width, fontSize float64) {
	xLast, yLast := f.iterator.Position()

	f.iterator.SetNewPosition(x, y)

	f.pdf.SetXY(f.iterator.Position())
	err := f.pdf.SetFontSize(fontSize)
	if err != nil {
		slog.Warn(err.Error())
	}

	rows, err := f.pdf.SplitTextWithWordWrap(text, width)
	if err != nil {
		slog.Warn(err.Error())
	}

	for _, row := range rows {

		f.iterator.IncrYStep(fontSize)
		if err != nil {
			slog.Warn(err.Error())
		}
		err = f.pdf.Text(row)
		if err != nil {
			slog.Warn("error writing text to pdf")
		}

		f.pdf.SetXY(f.iterator.Position())
	}

	f.iterator.SetNewPosition(xLast, yLast)
}

type TableHeadCell struct {
	Title string
	width float64
}

type TableCell struct {
	X       float64
	Y       float64
	Width   float64
	Height  float64
	Text    string
	Padding float64
}

func (f *FileBuilder) CreateTable(tableName string, head []string, rows [][]string) error {
	var (
		headCells []TableCell
	)

	if len(head) == 0 || len(rows) == 0 {
		return nil
	}

	columnSize := f.iterator.LineWidth() / float64(len(head))

	for _, h := range head {
		headCells = append(headCells, TableCell{
			Text:    h,
			Width:   columnSize,
			Padding: TablePadding,
		})
	}

	f.WriteTableNameText(tableName)

	err := f.CreateRow(headCells)
	if err != nil {
		return err
	}

	for _, row := range rows {
		rowCells := make([]TableCell, 0, len(row))
		for _, c := range row {
			rowCells = append(rowCells, TableCell{
				Text:    c,
				Width:   columnSize,
				Padding: TablePadding,
			})
		}

		err = f.CreateRow(rowCells)
		if err != nil {
			return err
		}
	}

	f.IncrLine(TableBottomMargin)

	return nil
}

func (f *FileBuilder) CreateRow(row []TableCell) error {
	var (
		maxHeight float64
	)

	for _, cell := range row {
		cellHeight := f.CalculateCellHeight(cell.Text, 12, cell.Width, cell.Padding)
		if cellHeight > maxHeight {
			maxHeight = cellHeight
		}
	}

	f.IncrLineWithPositionY(maxHeight)

	for i := range row {
		row[i].Height = maxHeight

		x, y := f.iterator.Position()

		row[i].X = x
		row[i].Y = y

		f.iterator.IncrXStep(row[i].Width)
	}

	for _, cell := range row {
		f.pdf.RectFromLowerLeft(cell.X, cell.Y, cell.Width, cell.Height)
		f.WriteTextToCell(cell.Text, cell.X+cell.Padding, cell.Y-cell.Height+cell.Padding+12, cell.Width-cell.Padding*2, 12)
	}

	return nil
}
