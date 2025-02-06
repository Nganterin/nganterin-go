package helpers

func StringPointer(b []byte) *string {
	str := string(b)
	return &str
}