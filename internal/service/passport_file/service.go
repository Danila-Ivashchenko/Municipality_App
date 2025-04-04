package passport_file

import (
	"context"
	"fmt"
	"municipality_app/internal/domain/entity"
)

func (svc *passportFileService) Create(ctx context.Context, passport *entity.PassportEx) error {
	var (
		err error
	)

	fileBuilder := NewFileBuilder()

	err = fileBuilder.UploadFont("./font/timesnewromanpsmt.ttf", "base")
	if err != nil {
		return err
	}

	err = fileBuilder.WriteH1(passport.Name)
	if err != nil {
		return err
	}

	for i, chapter := range passport.Chapters {
		err = fileBuilder.WriteH2(fmt.Sprintf("%d ", i+1) + chapter.Name)
		if err != nil {
			return err
		}

		err = fileBuilder.WriteCommonText(chapter.Text)
		if err != nil {
			return err
		}

		for j, partition := range chapter.Partitions {
			var (
				columns          []string
				rows             [][]string
				attributeToOrder = make(map[int64]int)
			)

			err = fileBuilder.WriteH3(fmt.Sprintf("%d.%d ", i+1, j+1) + partition.Name)
			if err != nil {
				return err
			}
			err = fileBuilder.WriteCommonText(partition.Text)
			if err != nil {
				return err
			}

			columns = append(columns, "Название")

			for _, templateEx := range partition.Objects {
				for _, attribute := range templateEx.Attributes {
					if attribute.ToShow {
						columns = append(columns, attribute.Name)
						attributeToOrder[attribute.ID] = len(columns)
					}
				}

				for _, object := range templateEx.Objects {
					row := make([]string, len(columns))
					row[0] = object.Name

					for _, attribute := range object.AttributeValues {
						order, exists := attributeToOrder[attribute.Value.ObjectAttributeID]
						if !exists {
							continue
						}

						row[order-1] = attribute.Value.Value
					}

					rows = append(rows, row)
				}

			}

			err = fileBuilder.CreateTable(columns, rows)
		}
	}

	for c := 0; c < 50; c++ {
		bigText := "Очень длинный текст, который должен автоматически переноситься на следующую строку. Очень длинный текст, который должен автоматически переноситься на следующую строку. Очень длинный текст, который должен автоматически переноситься на следующую строку. Очень длинный текст, который должен автоматически переноситься на следующую строку."
		err = fileBuilder.WriteCommonText(bigText)
		if err != nil {
			return err
		}
	}

	err = fileBuilder.Save("file.pdf")
	if err != nil {
		return err
	}

	return nil
}
