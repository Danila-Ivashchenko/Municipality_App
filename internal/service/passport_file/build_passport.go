package passport_file

import (
	"fmt"
	"municipality_app/internal/domain/entity"
)

func (svc *passportFileService) BuildPassportFile(path string, passport *entity.PassportEx) error {
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
			err = fileBuilder.WriteH3(fmt.Sprintf("%d.%d ", i+1, j+1) + partition.Name)
			if err != nil {
				return err
			}
			err = fileBuilder.WriteCommonText(partition.Text)
			if err != nil {
				return err
			}

			for _, templateEx := range partition.Objects {
				var (
					columns          []string
					rows             [][]string
					attributeToOrder = make(map[int64]int)
					tableName        = templateEx.Template.Name
				)
				columns = append(columns, "Название")
				columns = append(columns, "Адрес")

				for _, attribute := range templateEx.Attributes {
					if attribute.ToShow {
						columns = append(columns, attribute.Name)
						attributeToOrder[attribute.ID] = len(columns)
					}
				}

				for _, object := range templateEx.Objects {
					row := make([]string, len(columns))
					row[0] = object.Name

					if object.LocationID != nil {
						if len(object.LocationID.Address) > 0 {
							row[1] = object.LocationID.Address
						}
					}

					for _, attribute := range object.AttributeValues {
						order, exists := attributeToOrder[attribute.Value.ObjectAttributeID]
						if !exists {
							continue
						}

						row[order-1] = attribute.Value.Value
					}

					rows = append(rows, row)
				}

				err = fileBuilder.CreateTable(tableName, columns, rows)
			}

			for _, templateEx := range partition.Entities {
				var (
					columns          []string
					rows             [][]string
					attributeToOrder = make(map[int64]int)
					tableName        = templateEx.Template.Name
				)
				columns = append(columns, "Название")

				for _, attribute := range templateEx.Attributes {
					if attribute.ToShow {
						columns = append(columns, attribute.Name)
						attributeToOrder[attribute.ID] = len(columns)
					}
				}

				for _, e := range templateEx.Entities {
					row := make([]string, len(columns))
					row[0] = e.Name

					for _, attribute := range e.AttributeValues {
						order, exists := attributeToOrder[attribute.Value.EntityAttributeID]
						if !exists {
							continue
						}

						row[order-1] = attribute.Value.Value
					}

					rows = append(rows, row)
				}

				err = fileBuilder.CreateTable(tableName, columns, rows)
			}

		}
	}

	err = fileBuilder.Save(path)
	if err != nil {
		return err
	}

	return nil
}
