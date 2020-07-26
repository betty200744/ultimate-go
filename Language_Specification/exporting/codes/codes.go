package codes

var (
	ArchiveTypeForbidAdd = 10001 // 该类型不支持投稿
	ArchiveExist         = 10002 // 已经存在该稿件了
	ArchiveNotExist      = 10003 // 不存在该稿件
	ArchiveAlreadyDel    = 10004 // 稿件已经被删除
)

type alertCode int

// ------------------------------------------
// Access a value of an unexported identifier
// ------------------------------------------
func New(value int) alertCode {
	return alertCode(value)
}

// -----------------------------------------
// Unexported fields from an exported struct
// -----------------------------------------
type Status struct {
	Name    string
	message string
}
