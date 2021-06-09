package database

type BaseModel interface {
	GetFieldNames() string
	GetValueList() string
	GetUpdateList(fields []string) string
	Deserialize(db_output string)
}
