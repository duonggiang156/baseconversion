// ConvertToLaTeX chuyển đổi văn bản thành định dạng LaTeX cho Excel
func ConvertToLaTeX(text string) string {
	// Bảo vệ từ "hexadecimal" trước khi thực hiện các thay thế
	text = preserveHexadecimalWord(text)

	// Kiểm tra nếu đã có từ LaTeX
	if strings.Contains(text, "\\") {
		return restoreHexadecimalWord(text)
	}

	// Map các mẫu cần tìm kiếm tới các xử lý tương ứng
	basePatterns := map[string]struct{
		base    int
		oldBase int
		format  string
	}{
		"binary to decimal":        {base: 10, oldBase: 2, format: "%d_{10}"},
		"octal to decimal":         {base: 10, oldBase: 8, format: "%d_{10}"},
		"he--adecimal to decimal":  {base: 10, oldBase: 16, format: "%d_{10}"},
		"decimal to binary":        {base: 2, oldBase: 10, format: "%s_{2}"},
		"decimal to octal":         {base: 8, oldBase: 10, format: "%s_{8}"},
		"decimal to he--adecimal":  {base: 16, oldBase: 10, format: "%s_{16}"},
		"binary to octal":          {base: 8, oldBase: 2, format: "%s_{8}"},
		"binary to he--adecimal":   {base: 16, oldBase: 2, format: "%s_{16}"},
		"octal to binary":          {base: 2, oldBase: 8, format: "%s_{2}"},
		"octal to he--adecimal":    {base: 16, oldBase: 8, format: "%s_{16}"},
		"he--adecimal to binary":   {base: 2, oldBase: 16, format: "%s_{2}"},
		"he--adecimal to octal":    {base: 8, oldBase: 16, format: "%s_{8}"},
	}
	
	// Kiểm tra nếu văn bản chứa các mẫu chuyển đổi cơ số
	// Đầu tiên phục hồi từ "hexadecimal" trước khi kiểm tra các mẫu chuỗi đơn giản
	tempText := restoreHexadecimalWord(text)
	
	for pattern, config := range basePatterns {
		// Phục hồi từ "hexadecimal" trong mẫu để so sánh với tempText
		patternRestored := restoreHexadecimalWord(pattern)
		shortPattern := strings.Replace(patternRestored, " to ", " ", 1)
		
		if strings.Contains(tempText, patternRestored) || strings.Contains(tempText, shortPattern) {
			// Tiến hành xử lý chuyển đổi cơ số
			result := formatNumbersWithBase(text, config.base, config.oldBase, config.format)
			return restoreHexadecimalWord(result)
		}
	}

	// Xử lý đặc biệt cho \\_{16}
	text = strings.Replace(text, "\\_{16}", "_{16}", -1)
	
	// Nếu không phải chuyển đổi cơ số, áp dụng chuyển đổi LaTeX thông thường
	text = ConvertLineToLaTeX(text)
	
	// Phục hồi từ "hexadecimal" trước khi trả về kết quả
	return restoreHexadecimalWord(text)
}

// ConvertLineToLaTeX chuyển đổi một dòng văn bản sang dạng LaTeX
func ConvertLineToLaTeX(text string) string {
	// Thay thế từng ký tự đặc biệt
	text = strings.Replace(text, "<=", "\\leq", -1)
	text = strings.Replace(text, ">=", "\\geq", -1)
	text = strings.Replace(text, "<", "\\lt", -1)
	text = strings.Replace(text, ">", "\\gt", -1)
	
	// Thay thế phép nhân (x) chỉ khi nằm giữa hai số
	re := regexp.MustCompile(`(\d+)\s*[xX]\s*(\d+)`)
	text = re.ReplaceAllString(text, "$1 \\times $2")
	
	// Bảo tồn khoảng trắng cho LaTeX
	text = strings.Replace(text, " ", "~", -1)
	
	return text
}

// Các hàm bổ trợ để xử lý từ "hexadecimal"
func preserveHexadecimalWord(input string) string {
	return strings.ReplaceAll(input, "hexadecimal", "he--adecimal")
}

func restoreHexadecimalWord(input string) string {
	return strings.ReplaceAll(input, "he--adecimal", "hexadecimal")
}

// formatNumbersWithBase định dạng LaTeX cho các số có cơ số
func formatNumbersWithBase(text string, base int, oldBase int, format string) string {
	// Thực hiện chuyển đổi cơ số
	if base != oldBase && !strings.Contains(text, "$") {
		// Tìm và định dạng các số trong văn bản
		re := regexp.MustCompile(`([0-9a-fA-F]+)`)
		text = re.ReplaceAllStringFunc(text, func(numStr string) string {
			// Chuyển đổi từ cơ số cũ sang cơ số mới
			var num int64
			var err error
			
			// Phân tích số theo cơ số cũ
			num, err = strconv.ParseInt(numStr, oldBase, 64)
			if err != nil {
				return numStr // Nếu không phải số hợp lệ, giữ nguyên
			}
			
			// Định dạng theo cơ số mới
			var result string
			if base == 10 {
				result = fmt.Sprintf(format, num)
			} else {
				// Chuyển đổi sang cơ số 2, 8, hoặc 16
				switch base {
				case 2:
					result = fmt.Sprintf(format, strconv.FormatInt(num, 2))
				case 8:
					result = fmt.Sprintf(format, strconv.FormatInt(num, 8))
				case 16:
					result = fmt.Sprintf(format, strconv.FormatInt(num, 16))
				default:
					result = numStr
				}
			}
			return "$" + result + "$"
		})
		return text
	}
	
	// Định dạng cho số thập phân có cơ số dạng "123_10"
	baseRegex := regexp.MustCompile(`([0-9a-fA-F]+)_([0-9]+)`)
	text = baseRegex.ReplaceAllString(text, "$$$1_{$2}$$$")
	
	// Định dạng cho số hệ 16 (hexa) dạng "ABCh"
	hexRegex := regexp.MustCompile(`([0-9a-fA-F]+)h`)
	text = hexRegex.ReplaceAllString(text, "$$$1_{16}$$$")
	
	// Định dạng cho số hệ 2 (binary) dạng "1010b"
	binRegex := regexp.MustCompile(`([01]+)b`)
	text = binRegex.ReplaceAllString(text, "$$$1_{2}$$$")
	
	// Định dạng cho số hệ 8 (octal) dạng "123o"
	octRegex := regexp.MustCompile(`([0-7]+)o`)
	text = octRegex.ReplaceAllString(text, "$$$1_{8}$$$")
	
	// Thay thế các ký tự đặc biệt
	text = strings.Replace(text, "<=", "\\leq", -1)
	text = strings.Replace(text, ">=", "\\geq", -1)
	text = strings.Replace(text, "<", "\\lt", -1)
	text = strings.Replace(text, ">", "\\gt", -1)
	text = strings.Replace(text, "...", "\\dots", -1)
	
	// Thay thế "x" với dấu nhân chỉ khi nằm giữa hai số
	re := regexp.MustCompile(`(\d+)\s*[xX]\s*(\d+)`)
	text = re.ReplaceAllString(text, "$1 \\times $2")
	
	// Định dạng mũ
	if strings.Contains(text, "^") {
		expRe := regexp.MustCompile(`\^(\d+)`)
		text = expRe.ReplaceAllString(text, "^{$1}")
	}
	
	// Bảo tồn khoảng trắng
	text = strings.Replace(text, " ", "~", -1)
	
	return text
} 