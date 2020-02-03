package DocStract

func StripSeperators(s string) string {
	iBytes := []byte(s)
	oBytes := []byte{}

	for _, b := range iBytes {
		if b != '0' {
			oBytes = append(oBytes, b)
		}
	}

	return string(oBytes)
}
