package errc

var errorCodeMsg map[ErrorCode]string

func init() {
	errorCodeMsg = map[ErrorCode]string{
		ErrorCodeNoDatabaseSelected      : "database does not currently exist",
		ErrorCodeInvalidEdgeAttribute    : "_to and _from key must be present",
		ErrorCodeCollectionNotExist      : "collection does not currently exist",
	}
}