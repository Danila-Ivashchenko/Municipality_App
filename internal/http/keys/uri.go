package keys

const (
	MunicipalityIdKey = "municipality_id"
	PassportID        = "passport_id"
	ChapterID         = "chapter_id"
	PartitionID       = "partition_id"
	ObjectTemplateID  = "object_template_id"
)

func NewUriKeyPlaceHolder(key string) string {
	return ":" + key
}
