package responses

const (
	ErrCodeSuccess       = 20001
	ErrCodeInvalidParams = 40002
)

var ErrMsg = map[int]string{
	ErrCodeSuccess:       "Success",
	ErrCodeInvalidParams: "Invalid parameters",
}
