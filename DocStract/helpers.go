package DocStract

func StripSeperators(s string) string {
	iBytes := []byte(s)
	oBytes := []byte{}

	if len(iBytes) >= 3 {
		offset := 0
		if iBytes[0] == iBytes[2] && iBytes[0] == byte(0) {
			offset = 1
		}

		for i := offset; i < len(iBytes); i += 2 {
			oBytes = append(oBytes, iBytes[i])
		}
	}

	return string(oBytes)
}
