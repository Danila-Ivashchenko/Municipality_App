package parser

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/http/keys"
	"reflect"
	"strconv"
)

func ParseMunicipalityID(c *gin.Context) (int64, error) {
	municipalityID := c.Param(keys.MunicipalityIdKey)

	municipalityIDInt, err := strconv.Atoi(municipalityID)
	if err != nil {
		return 0, fmt.Errorf("invalid value municipality_id")
	}
	if municipalityIDInt == 0 {
		return 0, fmt.Errorf("municipality_id is empty")
	}

	return int64(municipalityIDInt), nil
}

func ParsePassportID(c *gin.Context) (int64, error) {
	passportID := c.Param(keys.PassportID)

	passportIDInt, err := strconv.Atoi(passportID)
	if err != nil {
		return 0, fmt.Errorf("invalid value passport_id")
	}
	if passportIDInt == 0 {
		return 0, fmt.Errorf("passport_id is empty")
	}

	return int64(passportIDInt), nil
}

func ParseChapterID(c *gin.Context) (int64, error) {
	chapterID := c.Param(keys.ChapterID)

	chapterIDInt, err := strconv.Atoi(chapterID)
	if err != nil {
		return 0, fmt.Errorf("invalid value chapter_id")
	}
	if chapterIDInt == 0 {
		return 0, fmt.Errorf("chapter_id is empty")
	}

	return int64(chapterIDInt), nil
}

func ParsePartitionID(c *gin.Context) (int64, error) {
	partitionID := c.Param(keys.PartitionID)

	partitionIDInt, err := strconv.Atoi(partitionID)
	if err != nil {
		return 0, fmt.Errorf("invalid value partition_id")
	}
	if partitionIDInt == 0 {
		return 0, fmt.Errorf("partition_id is empty")
	}

	return int64(partitionIDInt), nil
}

func ParseObjectTemplateID(c *gin.Context) (int64, error) {
	partitionID := c.Param(keys.ObjectTemplateID)

	partitionIDInt, err := strconv.Atoi(partitionID)
	if err != nil {
		return 0, fmt.Errorf("invalid value object_template_id")
	}
	if partitionIDInt == 0 {
		return 0, fmt.Errorf("object_template_id is empty")
	}

	return int64(partitionIDInt), nil
}

func ParseEntityTemplateID(c *gin.Context) (int64, error) {
	partitionID := c.Param(keys.EntityTemplateID)

	partitionIDInt, err := strconv.Atoi(partitionID)
	if err != nil {
		return 0, fmt.Errorf("invalid value entity_template_id")
	}
	if partitionIDInt == 0 {
		return 0, fmt.Errorf("entity_template_id is empty")
	}

	return int64(partitionIDInt), nil
}

func Parse(c *gin.Context, v any) (context.Context, error) {
	ctx := c.Request.Context()

	err := c.ShouldBind(v)

	return ctx, err
}

func ParseGetParams(c *gin.Context, obj any) (context.Context, error) {
	ctx := c.Request.Context()

	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return ctx, fmt.Errorf("obj должен быть указателем на структуру")
	}

	v = v.Elem()
	typ := v.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		paramTag := field.Tag.Get("param")

		value, ok := c.GetQuery(paramTag)
		if !ok {
			continue
		}

		fmt.Println(value)
	}

	return ctx, nil
}

func Context(c *gin.Context) context.Context {
	return c.Request.Context()
}
