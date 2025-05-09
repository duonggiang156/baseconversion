package stepbystep

import (
	"fmt"
	"strconv"
	"regexp"
	"strings"
	"bufio"
	"os"
	"sort"
	"math"
	
	"github.com/xuri/excelize/v2"
	"github.com/clarketm/ncalc/utils"
)

// StepByStepResult là cấu trúc để lưu kết quả chuyển đổi từng bước
type StepByStepResult struct {
	Input      string   // Giá trị đầu vào
	InputBase  string   // Cơ số của đầu vào
	Output     string   // Kết quả đầu ra
	OutputBase string   // Cơ số của đầu ra
	Steps      []string // Các bước chuyển đổi
}

// Binary2DecimalSteps chuyển đổi số nhị phân sang thập phân với các bước chi tiết
func Binary2DecimalSteps(s string) *StepByStepResult {
	result := &StepByStepResult{
		Input:      s,
		InputBase:  utils.BINARY,
		Output:     "",
		OutputBase: utils.DECIMAL,
		Steps:      []string{},
	}
	
	// Thêm tiêu đề cho giải pháp
	result.Steps = append(result.Steps, fmt.Sprintf("Converting binary number %s to decimal:", s))
	
	// Thêm công thức tổng quát
	result.Steps = append(result.Steps, "Formula: \\text{Decimal} = d_1 \\times 2^{n-1} + d_2 \\times 2^{n-2} + \\dots + d_n \\times 2^{0}\\\\")
	result.Steps = append(result.Steps, fmt.Sprintf("For %s:", s)) 
	
	// Tính toán từng bước
	var sum int64
	for i, digit := range s {
		position := len(s) - i - 1
		digitValue, _ := strconv.ParseInt(string(digit), 2, 64)
		
		powerValue := int64(1)
		for j := 0; j < position; j++ {
			powerValue *= 2
		}
		
		value := digitValue * powerValue
		sum += value
		
		result.Steps = append(result.Steps, fmt.Sprintf("  %s x 2^%d = %d x %d = %d", 
			string(digit), position, digitValue, powerValue, value))
	}
	
	// Thêm tổng kết
	result.Steps = append(result.Steps, fmt.Sprintf("Sum: %d", sum))
	result.Output = strconv.FormatInt(sum, 10)
	
	return result
}

// Octal2DecimalSteps chuyển đổi số bát phân sang thập phân với các bước chi tiết
func Octal2DecimalSteps(s string) *StepByStepResult {
	result := &StepByStepResult{
		Input:      s,
		InputBase:  utils.OCTAL,
		Output:     "",
		OutputBase: utils.DECIMAL,
		Steps:      []string{},
	}
	
	// Thêm tiêu đề cho giải pháp
	result.Steps = append(result.Steps, fmt.Sprintf("Converting octal number %s to decimal:", s))
	
	// Thêm công thức tổng quát
	result.Steps = append(result.Steps, "Formula: \\text{Decimal} = d_1 \\times 8^{n-1} + d_2 \\times 8^{n-2} + \\dots + d_n \\times 8^{0}\\\\")
	result.Steps = append(result.Steps, fmt.Sprintf("For %s:", s))
	
	// Tính toán từng bước
	var sum int64
	for i, digit := range s {
		position := len(s) - i - 1
		digitValue, _ := strconv.ParseInt(string(digit), 8, 64)
		
		powerValue := int64(1)
		for j := 0; j < position; j++ {
			powerValue *= 8
		}
		
		value := digitValue * powerValue
		sum += value
		
		result.Steps = append(result.Steps, fmt.Sprintf("  %s x 8^%d = %d x %d = %d", 
			string(digit), position, digitValue, powerValue, value))
	}
	
	// Thêm tổng kết
	result.Steps = append(result.Steps, fmt.Sprintf("Sum: %d", sum))
	result.Output = strconv.FormatInt(sum, 10)
	
	return result
}

// Hexadecimal2DecimalSteps chuyển đổi số thập lục phân sang thập phân với các bước chi tiết
func Hexadecimal2DecimalSteps(s string) *StepByStepResult {
	result := &StepByStepResult{
		Input:      s,
		InputBase:  utils.HEXADECIMAL,
		Output:     "",
		OutputBase: utils.DECIMAL,
		Steps:      []string{},
	}
	
	// Thêm tiêu đề cho giải pháp
	result.Steps = append(result.Steps, fmt.Sprintf("Converting hexadecimal number %s to decimal:", s))
	
	// Thêm công thức tổng quát
	result.Steps = append(result.Steps, "Formula: \\text{Decimal} = d_1 \\times 16^{n-1} + d_2 \\times 16^{n-2} + \\dots + d_n \\times 16^{0}\\\\")
	result.Steps = append(result.Steps, "Note: A=10, B=11, C=12, D=13, E=14, F=15")
	result.Steps = append(result.Steps, fmt.Sprintf("For %s:", s))
	
	// Tính toán từng bước
	var sum int64
	for i, digit := range s {
		position := len(s) - i - 1
		
		var digitValue int64
		var digitStr string
		
		// Xử lý chữ cái A-F
		if digit >= 'A' && digit <= 'F' {
			digitValue = int64(digit - 'A' + 10)
			digitStr = string(digit) + " (=" + strconv.FormatInt(digitValue, 10) + ")"
		} else if digit >= 'a' && digit <= 'f' {
			digitValue = int64(digit - 'a' + 10)
			digitStr = string(digit) + " (=" + strconv.FormatInt(digitValue, 10) + ")"
		} else {
			digitValue, _ = strconv.ParseInt(string(digit), 16, 64)
			digitStr = string(digit)
		}
		
		powerValue := int64(1)
		for j := 0; j < position; j++ {
			powerValue *= 16
		}
		
		value := digitValue * powerValue
		sum += value
		
		result.Steps = append(result.Steps, fmt.Sprintf("  %s x 16^%d = %d x %d = %d", 
			digitStr, position, digitValue, powerValue, value))
	}
	
	// Thêm tổng kết
	result.Steps = append(result.Steps, fmt.Sprintf("Sum: %d", sum))
	result.Output = strconv.FormatInt(sum, 10)
	
	return result
}

// Decimal2BinarySteps chuyển đổi số thập phân sang nhị phân với các bước chi tiết
func Decimal2BinarySteps(s string) *StepByStepResult {
	i, _ := strconv.ParseInt(s, 10, 64)
	
	result := &StepByStepResult{
		Input:      s,
		InputBase:  utils.DECIMAL,
		Output:     "",
		OutputBase: utils.BINARY,
		Steps:      []string{},
	}
	
	// Bỏ tiêu đề "Converting decimal number..." theo yêu cầu mới
	// Chỉ giữ lại phần Method
	result.Steps = append(result.Steps, "Method: Divide continuously by 2, note the remainders, read the result from bottom to top.")
	
	// Tính toán từng bước
	temp := i
	var remainders []int64
	var steps []string
	
	for temp > 0 {
		remainder := temp % 2
		quotient := temp / 2
		
		steps = append(steps, fmt.Sprintf("%d ÷ 2 = %d remainder %d", temp, quotient, remainder))
		remainders = append(remainders, remainder)
		
		temp = quotient
	}
	
	// Không đảo ngược kết quả, giữ nguyên thứ tự từ lớn đến nhỏ theo yêu cầu mới
	for i := 0; i < len(steps); i++ {
		result.Steps = append(result.Steps, steps[i])
	}
	
	// Ghép các số dư từ dưới lên để được kết quả
	var binary string
	for i := len(remainders) - 1; i >= 0; i-- {
		binary += strconv.FormatInt(remainders[i], 10)
	}
	
	result.Steps = append(result.Steps, fmt.Sprintf("Result: %s", binary))
	result.Output = binary
	
	return result
}

// Decimal2OctalSteps chuyển đổi số thập phân sang bát phân với các bước chi tiết
func Decimal2OctalSteps(s string) *StepByStepResult {
	i, _ := strconv.ParseInt(s, 10, 64)
	
	result := &StepByStepResult{
		Input:      s,
		InputBase:  utils.DECIMAL,
		Output:     "",
		OutputBase: utils.OCTAL,
		Steps:      []string{},
	}
	
	// Bỏ tiêu đề "Converting decimal number..." theo yêu cầu mới
	// Chỉ giữ lại phần Method
	result.Steps = append(result.Steps, "Method: Divide continuously by 8, note the remainders, read the result from bottom to top.")
	
	// Tính toán từng bước
	temp := i
	var remainders []int64
	var steps []string
	
	for temp > 0 {
		remainder := temp % 8
		quotient := temp / 8
		
		steps = append(steps, fmt.Sprintf("%d ÷ 8 = %d remainder %d", temp, quotient, remainder))
		remainders = append(remainders, remainder)
		
		temp = quotient
	}
	
	// Không đảo ngược kết quả, giữ nguyên thứ tự từ lớn đến nhỏ
	for i := 0; i < len(steps); i++ {
		result.Steps = append(result.Steps, steps[i])
	}
	
	// Ghép các số dư từ dưới lên để được kết quả
	var octal string
	for i := len(remainders) - 1; i >= 0; i-- {
		octal += strconv.FormatInt(remainders[i], 10)
	}
	
	result.Steps = append(result.Steps, fmt.Sprintf("Result: %s", octal))
	result.Output = octal
	
	return result
}

// Decimal2HexadecimalSteps chuyển đổi số thập phân sang thập lục phân với các bước chi tiết
func Decimal2HexadecimalSteps(s string) *StepByStepResult {
	i, _ := strconv.ParseInt(s, 10, 64)
	
	result := &StepByStepResult{
		Input:      s,
		InputBase:  utils.DECIMAL,
		Output:     "",
		OutputBase: utils.HEXADECIMAL,
		Steps:      []string{},
	}
	
	// Bỏ tiêu đề "Converting decimal number..." theo yêu cầu mới
	// Chỉ giữ lại phần Method và Note
	result.Steps = append(result.Steps, "Method: Divide continuously by 16, note the remainders, read the result from bottom to top.")
	result.Steps = append(result.Steps, "Note: 10=A, 11=B, 12=C, 13=D, 14=E, 15=F")
	
	// Tính toán từng bước
	temp := i
	var remainders []int64
	var steps []string
	
	for temp > 0 {
		remainder := temp % 16
		quotient := temp / 16
		
		var remainderStr string
		if remainder < 10 {
			remainderStr = strconv.FormatInt(remainder, 10)
		} else {
			// Chuyển đổi 10-15 thành A-F
			remainderStr = string(rune('A' + remainder) - 10)
		}
		
		steps = append(steps, fmt.Sprintf("%d ÷ 16 = %d remainder %s (%s)", 
			temp, quotient, remainderStr, remainderStr))
		remainders = append(remainders, remainder)
		
		temp = quotient
	}
	
	// Không đảo ngược kết quả, giữ nguyên thứ tự từ lớn đến nhỏ
	for i := 0; i < len(steps); i++ {
		result.Steps = append(result.Steps, steps[i])
	}
	
	// Ghép các số dư từ dưới lên để được kết quả
	var hex string
	for i := len(remainders) - 1; i >= 0; i-- {
		if remainders[i] < 10 {
			hex += strconv.FormatInt(remainders[i], 10)
		} else {
			// Chuyển đổi 10-15 thành A-F
			hex += string(rune('A' + int(remainders[i]) - 10))
		}
	}
	
	result.Steps = append(result.Steps, fmt.Sprintf("Result: %s", hex))
	result.Output = hex
	
	return result
}

// Binary2OctalSteps chuyển đổi số nhị phân sang bát phân thông qua thập phân
func Binary2OctalSteps(s string) *StepByStepResult {
	result := &StepByStepResult{
		Input:      s,
		InputBase:  utils.BINARY,
		Output:     "",
		OutputBase: utils.OCTAL,
		Steps:      []string{},
	}
	
	// Bước 1: Chuyển nhị phân sang thập phân
	decResult := Binary2DecimalSteps(s)
	result.Steps = append(result.Steps, "Step 1: Convert binary to decimal:")
	result.Steps = append(result.Steps, decResult.Steps...)
	
	// Bước 2: Chuyển thập phân sang bát phân
	octResult := Decimal2OctalSteps(decResult.Output)
	result.Steps = append(result.Steps, "Step 2: Convert decimal to octal:")
	result.Steps = append(result.Steps, octResult.Steps...)
	
	result.Output = octResult.Output
	return result
}

// Binary2HexadecimalSteps chuyển đổi số nhị phân sang thập lục phân thông qua thập phân
func Binary2HexadecimalSteps(s string) *StepByStepResult {
	result := &StepByStepResult{
		Input:      s,
		InputBase:  utils.BINARY,
		Output:     "",
		OutputBase: utils.HEXADECIMAL,
		Steps:      []string{},
	}
	
	// Bước 1: Chuyển nhị phân sang thập phân
	decResult := Binary2DecimalSteps(s)
	result.Steps = append(result.Steps, "Step 1: Convert binary to decimal:")
	result.Steps = append(result.Steps, decResult.Steps...)
	
	// Bước 2: Chuyển thập phân sang thập lục phân
	hexResult := Decimal2HexadecimalSteps(decResult.Output)
	result.Steps = append(result.Steps, "Step 2: Convert decimal to hexadecimal:")
	result.Steps = append(result.Steps, hexResult.Steps...)
	
	result.Output = hexResult.Output
	return result
}

// Octal2BinarySteps chuyển đổi số bát phân sang nhị phân thông qua thập phân
func Octal2BinarySteps(s string) *StepByStepResult {
	result := &StepByStepResult{
		Input:      s,
		InputBase:  utils.OCTAL,
		Output:     "",
		OutputBase: utils.BINARY,
		Steps:      []string{},
	}
	
	// Bước 1: Chuyển bát phân sang thập phân
	decResult := Octal2DecimalSteps(s)
	result.Steps = append(result.Steps, fmt.Sprintf("Converting octal number %s to binary:", s))
	result.Steps = append(result.Steps, "Step 1: Convert octal to decimal:")
	result.Steps = append(result.Steps, decResult.Steps...)
	
	// Bước 2: Chuyển thập phân sang nhị phân
	binResult := Decimal2BinarySteps(decResult.Output)
	result.Steps = append(result.Steps, "Step 2: Convert decimal to binary:")
	result.Steps = append(result.Steps, binResult.Steps...)
	
	result.Output = binResult.Output
	return result
}

// Octal2HexadecimalSteps chuyển đổi số bát phân sang thập lục phân thông qua thập phân
func Octal2HexadecimalSteps(s string) *StepByStepResult {
	result := &StepByStepResult{
		Input:      s,
		InputBase:  utils.OCTAL,
		Output:     "",
		OutputBase: utils.HEXADECIMAL,
		Steps:      []string{},
	}
	
	// Bước 1: Chuyển bát phân sang thập phân
	decResult := Octal2DecimalSteps(s)
	result.Steps = append(result.Steps, fmt.Sprintf("Converting octal number %s to hexadecimal:", s))
	result.Steps = append(result.Steps, "Step 1: Convert octal to decimal:")
	result.Steps = append(result.Steps, decResult.Steps...)
	
	// Bước 2: Chuyển thập phân sang thập lục phân
	hexResult := Decimal2HexadecimalSteps(decResult.Output)
	result.Steps = append(result.Steps, "Step 2: Convert decimal to hexadecimal:")
	result.Steps = append(result.Steps, hexResult.Steps...)
	
	result.Output = hexResult.Output
	return result
}

// Hexadecimal2BinarySteps chuyển đổi số thập lục phân sang nhị phân thông qua thập phân
func Hexadecimal2BinarySteps(s string) *StepByStepResult {
	result := &StepByStepResult{
		Input:      s,
		InputBase:  utils.HEXADECIMAL,
		Output:     "",
		OutputBase: utils.BINARY,
		Steps:      []string{},
	}
	
	// Bước 1: Chuyển thập lục phân sang thập phân
	decResult := Hexadecimal2DecimalSteps(s)
	result.Steps = append(result.Steps, fmt.Sprintf("Converting hexadecimal number %s to binary:", s))
	result.Steps = append(result.Steps, "Step 1: Convert hexadecimal to decimal:")
	result.Steps = append(result.Steps, decResult.Steps...)
	
	// Bước 2: Chuyển thập phân sang nhị phân
	binResult := Decimal2BinarySteps(decResult.Output)
	result.Steps = append(result.Steps, "Step 2: Convert decimal to binary:")
	result.Steps = append(result.Steps, binResult.Steps...)
	
	result.Output = binResult.Output
	return result
}

// Hexadecimal2OctalSteps chuyển đổi số thập lục phân sang bát phân thông qua thập phân
func Hexadecimal2OctalSteps(s string) *StepByStepResult {
	result := &StepByStepResult{
		Input:      s,
		InputBase:  utils.HEXADECIMAL,
		Output:     "",
		OutputBase: utils.OCTAL,
		Steps:      []string{},
	}
	
	// Bước 1: Chuyển thập lục phân sang thập phân
	decResult := Hexadecimal2DecimalSteps(s)
	result.Steps = append(result.Steps, fmt.Sprintf("Converting hexadecimal number %s to octal:", s))
	result.Steps = append(result.Steps, "Step 1: Convert hexadecimal to decimal:")
	result.Steps = append(result.Steps, decResult.Steps...)
	
	// Bước 2: Chuyển thập phân sang bát phân
	octResult := Decimal2OctalSteps(decResult.Output)
	result.Steps = append(result.Steps, "Step 2: Convert decimal to octal:")
	result.Steps = append(result.Steps, octResult.Steps...)
	
	result.Output = octResult.Output
	return result
}

// FormatBaseName trả về tên dễ đọc của cơ số
func FormatBaseName(base string) string {
	switch base {
	case utils.BINARY:
		return "2"
	case utils.OCTAL:
		return "8"
	case utils.DECIMAL:
		return "10"
	case utils.HEXADECIMAL:
		return "16"
	case utils.ASCII:
		return "ASCII"
	default:
		return base
	}
}

// ConvertToLaTeX chuyển các ký tự thông thường sang biểu thức LaTeX
func ConvertToLaTeX(input string) string {
	// Bảo vệ từ "hexadecimal" trước khi thực hiện các thay thế
	input = preserveHexadecimalWord(input)

	// Kiểm tra nếu đã có từ LaTeX
	if strings.Contains(input, "\\") {
		return restoreHexadecimalWord(input)
	}

	// Phục hồi từ "hexadecimal" tạm thời để kiểm tra
	tempInput := restoreHexadecimalWord(input)
	
	// Map các mẫu cần tìm kiếm tới hàm xử lý tương ứng
	conversionPatterns := map[string]func(string) string{
		"binary to decimal":     convertBinary2DecimalToLaTeX,
		"binary decimal":        convertBinary2DecimalToLaTeX,
		"octal to decimal":      convertOctal2DecimalToLaTeX,
		"octal decimal":         convertOctal2DecimalToLaTeX,
		"hexadecimal to decimal": convertHexadecimal2DecimalToLaTeX,
		"hexadecimal decimal":    convertHexadecimal2DecimalToLaTeX,
		"decimal to binary":     convertDecimal2BinaryToLaTeX,
		"decimal binary":        convertDecimal2BinaryToLaTeX,
		"decimal to octal":      convertDecimal2OctalToLaTeX,
		"decimal octal":         convertDecimal2OctalToLaTeX,
		"decimal to hexadecimal": convertDecimal2HexadecimalToLaTeX,
		"decimal hexadecimal":    convertDecimal2HexadecimalToLaTeX,
		"binary to octal":       convertBinary2OctalToLaTeX,
		"binary octal":          convertBinary2OctalToLaTeX,
		"binary to hexadecimal": convertBinary2HexadecimalToLaTeX,
		"binary hexadecimal":    convertBinary2HexadecimalToLaTeX,
		"octal to binary":       convertOctal2BinaryToLaTeX,
		"octal binary":          convertOctal2BinaryToLaTeX,
		"octal to hexadecimal":  convertOctal2HexadecimalToLaTeX,
		"octal hexadecimal":     convertOctal2HexadecimalToLaTeX,
		"hexadecimal to binary": convertHexadecimal2BinaryToLaTeX,
		"hexadecimal binary":    convertHexadecimal2BinaryToLaTeX,
		"hexadecimal to octal":  convertHexadecimal2OctalToLaTeX,
		"hexadecimal octal":     convertHexadecimal2OctalToLaTeX,
	}
	
	// Tìm mẫu phù hợp và áp dụng hàm chuyển đổi
	for pattern, conversionFunc := range conversionPatterns {
		if strings.Contains(tempInput, pattern) {
			// Nếu mẫu chứa "hexadecimal", đảm bảo khôi phục từ gốc
			if strings.Contains(pattern, "hexadecimal") {
				input = restoreHexadecimalWord(input)
			}
			return conversionFunc(input)
		}
	}

	// Nếu không khớp với bất kỳ mẫu nào, sử dụng hàm chuyển đổi mặc định
	// Đảm bảo khôi phục từ "hexadecimal" trước khi trả về
	return convertLineToLaTeX(restoreHexadecimalWord(input))
}

// preserveHexadecimalWord tạm thay thế từ "hexadecimal" để bảo vệ nó khỏi việc thay thế
func preserveHexadecimalWord(input string) string {
	return strings.ReplaceAll(input, "hexadecimal", "he--adecimal")
}

// restoreHexadecimalWord khôi phục từ "hexadecimal"
func restoreHexadecimalWord(input string) string {
	return strings.ReplaceAll(input, "he--adecimal", "hexadecimal")
}

// convertBinary2DecimalToLaTeX chuyển đổi giải thích từ nhị phân sang thập phân sang định dạng LaTeX
func convertBinary2DecimalToLaTeX(input string) string {
	// Phân tích các bước từ input
	lines := strings.Split(input, "\n")
	var result strings.Builder
	
	// Thêm công thức tổng quát
	result.WriteString("\\begin{enumerate}\n")
	result.WriteString("\\item Method: To convert a binary number to its decimal equivalent, assign each binary digit a positional weight based on its position, where the rightmost digit has a weight of \\(2^0\\), the next digit to the left has a weight of \\(2^1\\), and so on, up to the leftmost digit with a weight of \\(2^{n-1}\\) for a number with \\(n\\) digits. Multiply each digit by its corresponding power of 2, then sum all the products to obtain the decimal value, as expressed by the formula: \\begin{center} \\(\\text{Decimal} = d_1 \\times 2^{n-1} + d_2 \\times 2^{n-2} + \\  \\dots + d_n \\times 2^{0}\\) \\end{center} where \\(d_i\\) is the \\(i\\)-th digit from the left. \n")
	
	// Tìm số nhị phân trong input
	var binaryNumber string
	for _, line := range lines {
		if strings.Contains(line, "For") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				binaryNumber = strings.TrimSpace(parts[1])
				break
			}
		}
	}
	
	_ = binaryNumber // Tránh lỗi biến chưa sử dụng
	
	// Thu thập các bước tính toán chi tiết
	var calculationSteps []string
	var sumResult string
	
	for _, line := range lines {
		if strings.Contains(line, "x") && strings.Contains(line, "=") && !strings.Contains(line, "Formula") {
			// Đây là dòng tính toán chi tiết
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "  ") {
				line = strings.TrimPrefix(line, "  ")
			}
			
			// Chuẩn bị dòng LaTeX - chỉ thay thế "x" khi nó là toán tử nhân
			line = regexp.MustCompile(`\b(\d+) x (\d+)\b`).ReplaceAllString(line, "$1 \\times $2")
			
			// Xử lý định dạng số mũ - loại bỏ dấu {} và thay đổi thành dạng đơn giản
			re := regexp.MustCompile(`(\d+) \\\times 2\^{(\d+).*?}`)
			if matches := re.FindStringSubmatch(line); len(matches) > 0 {
				digit := matches[1]
				power := matches[2]
				result := ""
				
				// Tìm phần kết quả sau dấu =
				reResult := regexp.MustCompile(`= (.*)`)
				if resultMatches := reResult.FindStringSubmatch(line); len(resultMatches) > 0 {
					result = resultMatches[1]
				}
				
				// Tạo lại dòng với định dạng mới
				line = fmt.Sprintf("%s \\times 2^%s = %s", digit, power, result)
			}
			
			calculationSteps = append(calculationSteps, "\\item $"+line+"$")
		} else if strings.Contains(line, "Sum:") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				sumResult = strings.TrimSpace(parts[1])
			}
		}
	}
	
	// Thêm các bước tính toán vào kết quả dạng danh sách
	if len(calculationSteps) > 0 {
		result.WriteString("\\item Calculate each part: \n")
		result.WriteString("\\begin{itemize}\n")
		for _, step := range calculationSteps {
			result.WriteString(step + "\n")
		}
		result.WriteString("\\end{itemize}\n")
	}
	
	// Thêm kết quả cuối cùng
	if sumResult != "" {
		result.WriteString(fmt.Sprintf("\\text{Final Answer: } $%s_{10}$ \\\\\n", sumResult))
		result.WriteString("\\end{enumerate}\n")
	}
	
	return result.String()
}

// convertDecimal2BinaryToLaTeX chuyển đổi giải thích từ thập phân sang nhị phân sang định dạng LaTeX
func convertDecimal2BinaryToLaTeX(input string) string {
	// Tạo kết quả dựa trên định dạng LaTeX cho thập phân sang nhị phân
	// Phân tích các bước từ input
	lines := strings.Split(input, "\n")
	var result strings.Builder
	
	// Thêm tiêu đề
	result.WriteString("\\begin{enumerate}\n")
	result.WriteString("\\item Method: Divide continuously by 2, note the remainders, read the result from bottom to top. \n")
	
	// Thu thập các bước phép chia và số kết quả
	var divisionSteps []string
	var binaryResult string
	
	for _, line := range lines {
		if strings.Contains(line, "÷ 2 =") {
			// Phân tích dòng phép chia
			parts := strings.Split(line, "remainder")
			if len(parts) != 2 {
				continue
			}
			
			divParts := strings.Split(parts[0], "=")
			if len(divParts) != 2 {
				continue
			}
			
			leftSide := strings.TrimSpace(divParts[0])
			rightSide := strings.TrimSpace(divParts[1])
			remainder := strings.TrimSpace(parts[1])
			
			// Phân tách số chia và 2
			divNumbers := strings.Split(leftSide, "÷")
			if len(divNumbers) != 2 {
				continue
			}
			
			dividend := strings.TrimSpace(divNumbers[0])
			
			// Tạo chuỗi LaTeX cho dòng này dạng bullet point
			latexLine := fmt.Sprintf("\\item Divide %s by 2 = %s remainder %s", 
				dividend, rightSide, remainder)
			divisionSteps = append(divisionSteps, latexLine)
		} else if strings.Contains(line, "Result:") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				binaryResult = strings.TrimSpace(parts[1])
			}
		}
	}
	
	// Thêm các bước phép chia vào kết quả dạng danh sách
	result.WriteString("\\begin{itemize}\n")
	for _, step := range divisionSteps {
		result.WriteString(step + "\n")
	}
	result.WriteString("\\end{itemize}\n")
	
	// Thêm dòng kết quả 
	if binaryResult != "" {
		result.WriteString(fmt.Sprintf("\\item Read the result from bottom to top. Result: %s \\\\\n", binaryResult))
		result.WriteString("\\end{enumerate}\n")
	}
	
	return result.String()
}

// convertDecimal2OctalToLaTeX chuyển đổi giải thích từ thập phân sang bát phân sang định dạng LaTeX
func convertDecimal2OctalToLaTeX(input string) string {
	// Tạo kết quả dựa trên định dạng LaTeX cho thập phân sang bát phân
	// Phân tích các bước từ input
	lines := strings.Split(input, "\n")
	var result strings.Builder
	
	// Thêm tiêu đề
	result.WriteString("\\begin{enumerate}\n")
	result.WriteString("\\item Understand the Method: Divide the decimal number by 8 repeatedly, record the remainders, and read the remainders from bottom to top to form the octal number.  \\\\")
	result.WriteString("\\textbf{Note}: Octal digits range from 0 to 7.")
	result.WriteString("\\item Perform the Division")
	// Thu thập các bước phép chia và số kết quả
	var divisionSteps []string
	var octalResult string
	
	for _, line := range lines {
		if strings.Contains(line, "÷ 8 =") {
			// Phân tích dòng phép chia
			parts := strings.Split(line, "remainder")
			if len(parts) != 2 {
				continue
			}
			
			divParts := strings.Split(parts[0], "=")
			if len(divParts) != 2 {
				continue
			}
			
			leftSide := strings.TrimSpace(divParts[0])
			rightSide := strings.TrimSpace(divParts[1])
			remainder := strings.TrimSpace(parts[1])
			
			// Phân tách số chia và 8
			divNumbers := strings.Split(leftSide, "÷")
			if len(divNumbers) != 2 {
				continue
			}
			
			dividend := strings.TrimSpace(divNumbers[0])
			
			// Tạo chuỗi LaTeX cho dòng này dạng bullet point
			latexLine := fmt.Sprintf("\\item Divide %s by 8 = %s remainder %s", 
				dividend, rightSide, remainder)
			divisionSteps = append(divisionSteps, latexLine)
		} else if strings.Contains(line, "Result:") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				octalResult = strings.TrimSpace(parts[1])
			}
		}
	}
	
	// Thêm các bước phép chia vào kết quả dạng danh sách
	result.WriteString("\\begin{itemize}\n")
	for _, step := range divisionSteps {
		result.WriteString(step + "\n")
	}
	result.WriteString("\\end{itemize}\n")
	
	// Thêm dòng kết quả 
	if octalResult != "" {
		// Thu thập các chữ số của kết quả
		var octalDigits []string
		for _, char := range octalResult {
			octalDigits = append(octalDigits, string(char))
		}
		result.WriteString("\\item Read the Result \\\\")
		// Hiển thị từng chữ số trong ngoặc và phân tách bằng dấu phẩy
		result.WriteString("Reading the remainders from bottom to top: ")
		for i, digit := range octalDigits {
			if i > 0 {
				result.WriteString(", ")
			}
			result.WriteString(fmt.Sprintf("\\(%s\\)", digit))
		}
		result.WriteString(". \\\\\n")
		
		// Thêm dòng kết luận
		result.WriteString(fmt.Sprintf("Thus, the octal number is \\(%s_8\\).\n\n", octalResult))
		
		// Tìm số thập phân gốc từ input
		var decimalValue string
		for _, line := range lines {
			if strings.Contains(line, "Converting decimal number") {
				re := regexp.MustCompile(`Converting decimal number\s+(\S+)\s+to`)
				if matches := re.FindStringSubmatch(line); len(matches) > 0 {
					decimalValue = matches[1]
					break
				}
			}
		}
		
		// Nếu không tìm thấy trong tiêu đề, tìm trong các dòng khác
		if decimalValue == "" {
			// Lấy từ phép chia đầu tiên
			if len(divisionSteps) > 0 {
				re := regexp.MustCompile(`Divide (\d+) by 8`)
				firstStep := divisionSteps[0]
				if matches := re.FindStringSubmatch(firstStep); len(matches) > 0 {
					decimalValue = matches[1]
				}
			}
		}
		
		// Thêm kết quả cuối cùng trong thẻ center
		if decimalValue != "" {
			result.WriteString("\\begin{center}\n")
			result.WriteString(fmt.Sprintf("\\textbf{Final Answer:} \\(%s_{10} = %s_8\\)\n", decimalValue, octalResult))
			result.WriteString("\\end{center}\n")
			result.WriteString("\\end{enumerate}\n")
		}
	}
	
	return result.String()
}

// convertDecimal2HexadecimalToLaTeX chuyển đổi giải thích từ thập phân sang thập lục phân sang định dạng LaTeX
func convertDecimal2HexadecimalToLaTeX(input string) string {
	// Tạo kết quả dựa trên định dạng LaTeX cho thập phân sang thập lục phân
	// Phân tích các bước từ input
	lines := strings.Split(input, "\n")
	var result strings.Builder
	
	// Thêm tiêu đề và ghi chú
	result.WriteString("\\begin{enumerate}\n")
	result.WriteString("\\item Understand the Method: Divide the decimal number by 16 repeatedly, record the remainders, and read the remainders from bottom to top to form the hexadecimal number. \\\\")
	result.WriteString("\\textbf{Note}: Hexadecimal digits range from 0 to 15, where:")
	result.WriteString("\\begin{itemize}\n")
	result.WriteString("\\item \\(10 = A\\), \\(11 = B\\), \\(12 = C\\), \\(13 = D\\), \\(14 = E\\), \\(15 = F\\)")
	result.WriteString("\\end{itemize}\n")
	result.WriteString("\\item Perform the Division")
	
	// Thu thập các bước phép chia và số kết quả
	var divisionSteps []string
	var hexResult string
	
	for _, line := range lines {
		if strings.Contains(line, "÷ 16 =") {
			// Phân tích dòng phép chia
			parts := strings.Split(line, "remainder")
			if len(parts) != 2 {
				continue
			}
			
			divParts := strings.Split(parts[0], "=")
			if len(divParts) != 2 {
				continue
			}
			
			leftSide := strings.TrimSpace(divParts[0])
			rightSide := strings.TrimSpace(divParts[1])
			
			// Xử lý remainder có thể có dạng "11 (B)" cho số lớn hơn 9
			remainderPart := strings.TrimSpace(parts[1])
			var remainder string
			var hexChar string
			
			if strings.Contains(remainderPart, "(") {
				// Trường hợp có chữ cái trong ngoặc
				reParts := strings.Split(remainderPart, "(")
				if len(reParts) == 2 {
					remainder = strings.TrimSpace(reParts[0])
					hexChar = strings.TrimSpace(strings.TrimSuffix(reParts[1], ")"))
					remainderPart = fmt.Sprintf("%s (%s)", remainder, hexChar)
				}
			} else {
				remainder = remainderPart
			}
			
			// Phân tách số chia và 16
			divNumbers := strings.Split(leftSide, "÷")
			if len(divNumbers) != 2 {
				continue
			}
			
			dividend := strings.TrimSpace(divNumbers[0])
			
			// Tạo chuỗi LaTeX cho dòng này dạng bullet point
			var latexLine string
			if hexChar != "" {
				latexLine = fmt.Sprintf("\\item Divide %s by 16 = %s remainder %s (%s)", 
					dividend, rightSide, remainder, hexChar)
			} else {
				latexLine = fmt.Sprintf("\\item Divide %s by 16 = %s remainder %s", 
					dividend, rightSide, remainder)
			}
			
			divisionSteps = append(divisionSteps, latexLine)
		} else if strings.Contains(line, "Result:") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				hexResult = strings.TrimSpace(parts[1])
			}
		}
	}
	
	// Thêm các bước phép chia vào kết quả dạng danh sách
	result.WriteString("\\begin{itemize}\n")
	for _, step := range divisionSteps {
		result.WriteString(step + "\n")
	}
	result.WriteString("\\end{itemize}\n")
	
	// Thêm dòng kết quả 
	if hexResult != "" {
		// Thu thập các chữ số của kết quả
		var hexDigits []string
		for _, char := range hexResult {
			hexDigits = append(hexDigits, string(char))
		}
		result.WriteString("\\item Read the Result} \\\\")
		// Hiển thị từng chữ số trong ngoặc và phân tách bằng dấu phẩy
		result.WriteString("Reading the remainders from bottom to top: ")
		for i, digit := range hexDigits {
			if i > 0 {
				result.WriteString(", ")
			}
			result.WriteString(fmt.Sprintf("\\(%s\\)", digit))
		}
		result.WriteString(". \\\\\n")
		
		// Thêm dòng kết luận
		result.WriteString(fmt.Sprintf("Thus, the hexadecimal number is \\(%s_{16}\\).\n\n", hexResult))
		
		// Tìm số thập phân gốc từ input
		var decimalValue string
		for _, line := range lines {
			if strings.Contains(line, "Converting decimal number") {
				re := regexp.MustCompile(`Converting decimal number\s+(\S+)\s+to`)
				if matches := re.FindStringSubmatch(line); len(matches) > 0 {
					decimalValue = matches[1]
					break
				}
			}
		}
		
		// Nếu không tìm thấy trong tiêu đề, tìm trong các dòng khác
		if decimalValue == "" {
			// Lấy từ phép chia đầu tiên
			if len(divisionSteps) > 0 {
				re := regexp.MustCompile(`Divide (\d+) by 16`)
				firstStep := divisionSteps[0]
				if matches := re.FindStringSubmatch(firstStep); len(matches) > 0 {
					decimalValue = matches[1]
				}
			}
		}
		
		// Thêm kết quả cuối cùng trong thẻ center
		if decimalValue != "" {
			result.WriteString("\\begin{center}\n")
			result.WriteString(fmt.Sprintf("\\textbf{Final Answer:} \\(%s_{10} = %s_{16}\\)\n", decimalValue, hexResult))
			result.WriteString("\\end{center}\n")
			result.WriteString("\\end{enumerate}\n")
		}
	}
	
	return result.String()
}

// convertOctal2DecimalToLaTeX chuyển đổi giải thích từ bát phân sang thập phân sang định dạng LaTeX
func convertOctal2DecimalToLaTeX(input string) string {
	// Phân tích các bước từ input
	lines := strings.Split(input, "\n")
	var result strings.Builder
	

	
	// Thêm công thức tổng quát
	result.WriteString("\\begin{enumerate}\n")
	result.WriteString("\\item Apply the formula: $\\text{Decimal} = d_1 \\times 8^{n-1} + d_2 \\times 8^{n-2} + \\dots + d_n \\times 8^{0}$, where $d_i$ represents each digit of the octal number, and $n$ is the number of digits. \n")
	
	// Tìm số bát phân trong input
	var octalNumber string
	for _, line := range lines {
		if strings.Contains(line, "For") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				octalNumber = strings.TrimSpace(parts[1])
				break
			}
		}
	}
	
	// Tìm thêm trong tiêu đề nếu không có trong phần "For"
	if octalNumber == "" {
		for _, line := range lines {
			if strings.Contains(line, "Converting") && strings.Contains(line, "to decimal") {
				re := regexp.MustCompile(`Converting octal number\s+(\S+)\s+to`)
				if matches := re.FindStringSubmatch(line); len(matches) > 0 {
					octalNumber = matches[1]
					break
				}
			}
		}
	}
	
	// Nếu vẫn không tìm thấy, cố gắng trích xuất từ các phép tính
	if octalNumber == "" {
		// Tìm các chữ số từ phép tính và vị trí của chúng
		type digitPosition struct {
			digit    string
			position int
		}
		var digits []digitPosition
		
		for _, line := range lines {
			if strings.Contains(line, "x 8^") && strings.Contains(line, "=") {
				re := regexp.MustCompile(`(\d+)\s+x\s+8\^(\d+)`)
				if matches := re.FindStringSubmatch(line); len(matches) > 0 {
					digit := matches[1]
					position, _ := strconv.Atoi(matches[2])
					digits = append(digits, digitPosition{digit: digit, position: position})
				}
			}
		}
		
		// Sắp xếp chữ số theo vị trí từ cao xuống thấp
		sort.Slice(digits, func(i, j int) bool {
			return digits[i].position > digits[j].position
		})
		
		// Ghép các chữ số lại với nhau
		var digitBuilder strings.Builder
		for _, d := range digits {
			digitBuilder.WriteString(d.digit)
		}
		octalNumber = digitBuilder.String()
	}
	
	// Kiểm tra và đặt giá trị mặc định nếu cần thiết
	if octalNumber == "" {
		// Không tìm thấy số, dùng giá trị mặc định generic
		octalNumber = "dndn-1...d2d1"
	}
	
	// Đếm số chữ số trong số bát phân - chỉ đếm nếu là số thực sự
	numDigits := len(octalNumber)
	re := regexp.MustCompile(`^[0-7]+$`)
	isNumeric := re.MatchString(octalNumber)
	
	// // Thêm bước 2 - áp dụng công thức cho số cụ thể
	// if isNumeric {
	// 	// Đầu tiên, xác định n (số chữ số)
	// 	result.WriteString(fmt.Sprintf("2. For the octal number $%s_8$: first, determine the number of digits. In this case, $%s_8$ has %d digits, so $n=%d$. \\\\\n", 
	// 		octalNumber, octalNumber, numDigits, numDigits))
		
	// 	// Tiếp theo, áp dụng công thức
	// 	result.WriteString("Now apply the formula: \\\\\n")
	// } else {
	// 	result.WriteString("2. For the octal number, apply the formula: \\\\\n")
	// }
	
	// // Tạo công thức với các biến
	// if isNumeric {
	// 	// Mở rộng công thức với chỉ số cụ thể
	// 	var formulaTerms []string
	// 	for i := 1; i <= numDigits; i++ {
	// 		formulaTerms = append(formulaTerms, fmt.Sprintf("d_%d \\times 8^{%d-%d}", i, numDigits, i))
	// 	}
	// 	result.WriteString("$\\text{Decimal} = " + strings.Join(formulaTerms, " + ") + "$. \\\\\n")
	// }
	
	// Thêm bước 3 - tính toán chi tiết
	result.WriteString("\\item Calculate each term:\n")
	result.WriteString("\\begin{itemize}\n")
	
	// Thêm dòng mở rộng công thức nếu có số cụ thể
	if isNumeric {
		// Tạo công thức mở rộng với các biến và số mũ cụ thể
		var formulaExpanded []string
		for i := 1; i <= numDigits; i++ {
			formulaExpanded = append(formulaExpanded, fmt.Sprintf("d_%d \\times 8^{%d-%d}", i, numDigits, i))
		}
		
		// Tạo công thức sau khi tính toán số mũ
		var expandedTerms []string
		for i := 1; i <= numDigits; i++ {
			expandedTerms = append(expandedTerms, fmt.Sprintf("d_%d \\times 8^%d", i, numDigits-i))
		}
		
		// Mở rộng công thức với giá trị cụ thể
		var valueTerms []string
		for i := 0; i < numDigits; i++ {
			digit := string(octalNumber[i])
			power := numDigits - i - 1
			valueTerms = append(valueTerms, fmt.Sprintf("%s \\times 8^%d", digit, power))
		}
		
		// Hiển thị công thức với cả dạng chưa tính và đã tính số mũ
		result.WriteString("    \\item The formula becomes: $\\text{Decimal} = " + 
			strings.Join(formulaExpanded, " + ") + " = " + 
			strings.Join(expandedTerms, " + ") + "$.\n")
	} else {
		result.WriteString("    \\item The formula becomes: $\\text{Decimal} = d_1 \\times 8^{n-1} + d_2 \\times 8^{n-2} + ... + d_n \\times 8^0$.\n")
	}
	
	// Thu thập các bước tính toán chi tiết
	var calculationSteps []string
	var values []int64
	var sumResult string
	
	// Phân tích các bước tính toán từ input
	for _, line := range lines {
		if strings.Contains(line, "x 8^") && strings.Contains(line, "=") && !strings.Contains(line, "Formula") {
			// Đây là dòng tính toán chi tiết
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "  ") {
				line = strings.TrimPrefix(line, "  ")
			}
			
			// Phân tích dòng để lấy thông tin chi tiết
			parts := strings.Split(line, "=")
			
			// Tìm chữ số và vị trí
			digitPart := strings.TrimSpace(parts[0])
			digitMatch := regexp.MustCompile(`^(\d+)\s+x\s+8\^(\d+)`).FindStringSubmatch(digitPart)
			
			if len(digitMatch) >= 3 {
				digit := digitMatch[1]
				position, _ := strconv.Atoi(digitMatch[2])
				
				// Tính vị trí từ trái sang, dựa vào vị trí từ phải qua
				var positionFromLeft int
				if isNumeric {
					positionFromLeft = numDigits - position
				} else {
					// Nếu không xác định được numDigits, ước lượng từ vị trí cao nhất
					positionFromLeft = position + 1
				}
				
				// Tìm kết quả tính toán
				resultParts := strings.Split(parts[len(parts)-1], "=")
				finalResult := strings.TrimSpace(resultParts[len(resultParts)-1])
				numValue, _ := strconv.ParseInt(finalResult, 10, 64)
				values = append(values, numValue)
				
				// Xây dựng dòng theo định dạng mới
				calculationItem := fmt.Sprintf("    \\item %s digit: $d_%d = %s$, position %d (from left), so compute $%s \\times 8^%d$:\n    \\[\n    %s\n    \\]", 
					getOrdinalText(positionFromLeft), positionFromLeft, digit, positionFromLeft, digit, position, line)
				
				// Sửa định dạng cho phép tính (chỉ thay thế x khi nó là phép nhân)
				calculationItem = strings.Replace(calculationItem, " x ", " \\times ", -1)
				
				calculationSteps = append(calculationSteps, calculationItem)
			}
		} else if strings.Contains(line, "Sum:") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				sumResult = strings.TrimSpace(parts[1])
			}
		}
	}
	
	// Sắp xếp các bước tính toán theo vị trí từ trái qua phải
	sort.Slice(calculationSteps, func(i, j int) bool {
		reI := regexp.MustCompile(`position (\d+) \(from left\)`)
		reJ := regexp.MustCompile(`position (\d+) \(from left\)`)
		matchesI := reI.FindStringSubmatch(calculationSteps[i])
		matchesJ := reJ.FindStringSubmatch(calculationSteps[j])
		if len(matchesI) > 1 && len(matchesJ) > 1 {
			posI, _ := strconv.Atoi(matchesI[1])
			posJ, _ := strconv.Atoi(matchesJ[1])
			return posI < posJ
		}
		return false
	})
	
	// Thêm các bước tính toán vào kết quả
	for _, step := range calculationSteps {
		result.WriteString(step + "\n")
	}
	
	// Thêm bước tính tổng
	if len(values) > 0 && sumResult != "" {
		valueStrings := make([]string, len(values))
		for i, v := range values {
			valueStrings[i] = fmt.Sprintf("%d", v)
		}
		result.WriteString(fmt.Sprintf("    \\item Sum the results: $%s = %s$.\n", 
			strings.Join(valueStrings, " + "), sumResult))
	}
	
	result.WriteString("\\end{itemize}\n")
	
	// Thêm kết quả cuối cùng với định dạng mới
	if sumResult != "" {
		result.WriteString("\\begin{center}\n")
		result.WriteString(fmt.Sprintf("\\textbf{Final Answer:} $%s_{10}$\n", sumResult))
		result.WriteString("\\end{center}\n")
		result.WriteString("\\end{enumerate}\n")
	}
	
	return result.String()
}

// Hàm hỗ trợ để chuyển đổi số sang dạng thứ tự bằng văn bản
func getOrdinalText(n int) string {
	switch n {
	case 1:
		return "First"
	case 2:
		return "Second"
	case 3:
		return "Third"
	case 4:
		return "Fourth"
	case 5:
		return "Fifth"
	case 6:
		return "Sixth"
	case 7:
		return "Seventh"
	case 8:
		return "Eighth"
	case 9:
		return "Ninth"
	case 10:
		return "Tenth"
	default:
		return fmt.Sprintf("%dth", n)
	}
}

// convertHexadecimal2DecimalToLaTeX chuyển đổi giải thích từ thập lục phân sang thập phân sang định dạng LaTeX
func convertHexadecimal2DecimalToLaTeX(input string) string {
	// Phân tích các bước từ input
	lines := strings.Split(input, "\n")
	var result strings.Builder
	
	// Thêm công thức tổng quát
	result.WriteString("\\begin{enumerate}\n")
	result.WriteString("\\item Apply the formula: $\\text{Decimal} = d_1 \\times 16^{n-1} + d_2 \\times 16^{n-2} + \\dots + d_n \\times 16^{0}$, where $d_i$ represents each digit of the hexadecimal number, and $n$ is the number of digits. \n")
	
	// Thêm giải thích về các chữ số thập lục phân
	result.WriteString("\\item In hexadecimal, digits range from 0 to 15. Digits 0--9 are represented as is, while letters A--F represent values 10--15, respectively:\n")
	result.WriteString("\\begin{center}\n")
	result.WriteString("$A = 10$, $B = 11$, $C = 12$, $D = 13$, $E = 14$, $F = 15$")
	result.WriteString("\\end{center}\n")
	
	// Tìm số thập lục phân trong input và xác định số chữ số n
	var hexNumber string
	for _, line := range lines {
		if strings.Contains(line, "For") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				hexNumber = strings.TrimSpace(parts[1])
				break
			}
		}
	}
	
	// Nếu không tìm thấy số thập lục phân trong input, tìm từ dòng tính toán
	if hexNumber == "" {
		for _, line := range lines {
			if strings.Contains(line, "x 16^") && !strings.Contains(line, "Formula") {
				// Tìm chữ số đầu tiên
				re := regexp.MustCompile(`^(\w+)\s+.*?x 16\^`)
				if matches := re.FindStringSubmatch(line); len(matches) > 0 {
					// Tìm thêm các dòng tính toán khác để xác định tất cả các chữ số
					var digits []string
					digits = append(digits, matches[1])
					
					for _, l := range lines {
						if strings.Contains(l, "x 16^") && !strings.Contains(l, "Formula") && !strings.Contains(l, matches[1]) {
							reOther := regexp.MustCompile(`^(\w+)\s+.*?x 16\^`)
							if otherMatches := reOther.FindStringSubmatch(l); len(otherMatches) > 0 {
								digits = append(digits, otherMatches[1])
							}
						}
					}
					
					// Nối lại các chữ số để tạo số thập lục phân
					for _, d := range digits {
						hexNumber += d
					}
					break
				}
			}
		}
	}
	
	// Nếu vẫn không tìm thấy, thử tìm trong kết quả (có thể là FF)
	if hexNumber == "" {
		for _, line := range lines {
			re := regexp.MustCompile(`Converting\s+(\w+)\s+from`)
			if matches := re.FindStringSubmatch(line); len(matches) > 0 {
				hexNumber = matches[1]
				break
			}
		}
	}
	
	if hexNumber == "" {
		// Mặc định nếu không tìm thấy
		hexNumber = "FF"
	}
	
	// Đảm bảo không có ký tự trắng trong số thập lục phân
	hexNumber = strings.ReplaceAll(hexNumber, " ", "")
	
	_ = hexNumber // Tránh lỗi biến chưa sử dụng
	
	// === Xác định numDigits đúng ===
	// Tính số chữ số trong số thập lục phân
	hexNumber = strings.TrimSpace(hexNumber)
	// Loại bỏ tiền tố 0x nếu có
	if strings.HasPrefix(hexNumber, "0x") || strings.HasPrefix(hexNumber, "0X") {
		hexNumber = hexNumber[2:]
	}
	
	// Đếm số chữ số thực tế (không phải length của chuỗi)
	numDigits := len(hexNumber)
	
	// Kiểm tra xem numDigits có hợp lý không
	if numDigits > 10 {
		// Có thể là chuỗi công thức hoặc giá trị khác, không phải số thập lục phân
		// Đặt về giá trị mặc định
		numDigits = 2
	}
	
	if numDigits == 0 {
		numDigits = 2 // Giả sử là 2 nếu không xác định được
	}
	
	// Thu thập thông tin về các chữ số và giá trị của chúng
	var digits []string
	var values []string
	var positions []int
	var powerTerms []string
	var calculations []string
	var digitResults []string
	
	for _, line := range lines {
		if strings.Contains(line, "x 16^") && !strings.Contains(line, "Formula") {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "  ") {
				line = strings.TrimPrefix(line, "  ")
			}
			
			// Tìm chữ số và giá trị
			reDigit := regexp.MustCompile(`^(\w+)\s*(?:\(=(\d+)\))?`)
			if digitMatches := reDigit.FindStringSubmatch(line); len(digitMatches) > 0 {
				digit := digitMatches[1]
				var value string
				if len(digitMatches) > 2 && digitMatches[2] != "" {
					value = digitMatches[2]
				} else if digit >= "A" && digit <= "F" {
					// Chuyển đổi A-F thành 10-15
					value = strconv.Itoa(int(digit[0]) - int('A') + 10)
				} else if digit >= "a" && digit <= "f" {
					value = strconv.Itoa(int(digit[0]) - int('a') + 10)
				} else {
					value = digit
				}
				
				digits = append(digits, digit)
				values = append(values, value)
				positions = append(positions, len(digits))
			}
			
			// Tìm số mũ
			rePower := regexp.MustCompile(`16\^(\d+)`)
			if powerMatches := rePower.FindStringSubmatch(line); len(powerMatches) > 0 {
				power := powerMatches[1]
				powerTerms = append(powerTerms, fmt.Sprintf("16^%s", power))
			}
			
			// Tìm kết quả tính toán
			reResult := regexp.MustCompile(`= (.*?)$`)
			if resultMatches := reResult.FindStringSubmatch(line); len(resultMatches) > 0 {
				calculation := resultMatches[1]
				calculations = append(calculations, calculation)
				
				// Trích xuất kết quả cuối cùng của phép tính
				parts := strings.Split(calculation, "=")
				if len(parts) > 0 {
					finalPart := strings.TrimSpace(parts[len(parts)-1])
					digitResults = append(digitResults, finalPart)
				}
			}
		} else if strings.Contains(line, "Sum:") {
			// Đã xử lý ở phần cuối
		}
	}
	
	// Tìm tổng kết quả
	var sumResult string
	for _, line := range lines {
		if strings.Contains(line, "Sum:") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				sumResult = strings.TrimSpace(parts[1])
				break
			}
		}
	}
	
	// Tạo dạng viết gọn của công thức áp dụng cho số cụ thể
	var formulaSpecific string
	if numDigits > 0 {
		// Công thức tổng quát với n
		formulaSpecific = fmt.Sprintf("\\text{Decimal} = d_1 \\times 16^{%d-1} + d_2 \\times 16^{%d-2}", numDigits, numDigits)
		if numDigits > 2 {
			formulaSpecific += " + \\dots"
		}
		formulaSpecific += fmt.Sprintf(" + d_%d \\times 16^{%d-%d}", numDigits, numDigits, numDigits)
		
		// Công thức đã tính toán giá trị mũ cụ thể
		formulaSpecific += " = "
		
		// Tạo chuỗi giá trị mũ đã tính
		var simplifiedTerms []string
		simplifiedTerms = append(simplifiedTerms, fmt.Sprintf("d_1 \\times 16^%d", numDigits-1))
		simplifiedTerms = append(simplifiedTerms, fmt.Sprintf("d_2 \\times 16^%d", numDigits-2))
		
		// Thêm các phần tử giữa nếu cần
		if numDigits > 2 {
			for i := 3; i < numDigits; i++ {
				simplifiedTerms = append(simplifiedTerms, fmt.Sprintf("d_%d \\times 16^%d", i, numDigits-i))
			}
			// Thêm phần tử cuối
			simplifiedTerms = append(simplifiedTerms, fmt.Sprintf("d_%d \\times 16^0", numDigits))
			
			// Ghép với dấu +
			formulaSpecific += strings.Join(simplifiedTerms[:2], " + ")
			formulaSpecific += " + \\dots + " + simplifiedTerms[len(simplifiedTerms)-1]
		} else {
			// Nếu chỉ có 2 phần tử, ghép trực tiếp
			formulaSpecific += strings.Join(simplifiedTerms, " + ")
		}
	}
	
	// Tạo giải thích chi tiết cho từng chữ số
	result.WriteString(fmt.Sprintf("\\item Calculate each term for the hexadecimal number $%s$ (%d digits, so $n=%d$):\n", hexNumber, numDigits, numDigits))
	result.WriteString("\\begin{itemize}\n")
	
	// Thêm công thức áp dụng cho số cụ thể
	if formulaSpecific != "" {
		result.WriteString("\\item The formula becomes: \n\\begin{center}\n$" + formulaSpecific + "$\n\\end{center}\n")
	}
	
	// Tạo giải thích chi tiết cho từng chữ số
	for i := 0; i < len(digits) && i < len(values) && i < len(positions) && i < len(powerTerms) && i < len(calculations); i++ {
		position := positions[i]
		ordinal := getOrdinalSuffix(position)
		result.WriteString(fmt.Sprintf("\\item %s digit: $d_%d = %s = %s$, position %d (from left), so compute $%s \\times %s$:\n", 
			ordinal, position, digits[i], values[i], position, values[i], powerTerms[i]))
		
		result.WriteString("\\[\n")
		if i < len(calculations) {
			// Sửa lại định dạng, dùng \times thay vì x
			calculations[i] = strings.Replace(calculations[i], "x", "\\times", -1)
			result.WriteString(calculations[i] + "\n")
		}
		result.WriteString("\\]\n")
	}
	
	// Thêm phần tính tổng
	if len(digitResults) > 0 && sumResult != "" {
		resultStr := "\\item Sum the results: "
		for i, dr := range digitResults {
			if i > 0 {
				resultStr += " + "
			}
			resultStr += dr
		}
		resultStr += " = " + sumResult + ".\n"
		
		// Đảm bảo không có ký tự 'a' thừa
		resultStr = strings.Replace(resultStr, "a+", "+", -1)
		result.WriteString(resultStr)
	}
	
	result.WriteString("\\end{itemize}\n")

	
	// Thêm kết quả cuối cùng
	if sumResult != "" {
		result.WriteString("\\begin{center}\n")
		result.WriteString(fmt.Sprintf("\\textbf{Final Answer:} $%s_{10}$\n", sumResult))
		result.WriteString("\\end{center}\n")
		result.WriteString("\\end{enumerate}\n")
	}
	
	return result.String()
}

// getOrdinalSuffix trả về hậu tố số thứ tự dạng tiếng Anh (1st, 2nd, 3rd, 4th, v.v.)
func getOrdinalSuffix(n int) string {
	if n <= 0 {
		return fmt.Sprintf("%dth", n)
	}
	
	switch n {
	case 1:
		return "First"
	case 2:
		return "Second"
	case 3:
		return "Third"
	case 4:
		return "Fourth"
	case 5:
		return "Fifth"
	case 6:
		return "Sixth"
	case 7:
		return "Seventh"
	case 8:
		return "Eighth"
	case 9:
		return "Ninth"
	case 10:
		return "Tenth"
	default:
		suffix := "th"
		if n%10 == 1 && n%100 != 11 {
			suffix = "st"
		} else if n%10 == 2 && n%100 != 12 {
			suffix = "nd"
		} else if n%10 == 3 && n%100 != 13 {
			suffix = "rd"
		}
		return fmt.Sprintf("%d%s", n, suffix)
	}
}

// convertGenericToLaTeX xử lý các trường hợp chung
func convertGenericToLaTeX(input string) string {
	// Xử lý chung cho các loại chuyển đổi
	// Sử dụng RegEx để đảm bảo chỉ thay thế "x" đứng một mình giữa các ký tự khác (toán tử nhân)
	re := regexp.MustCompile(`\b([0-9a-zA-Z]+)\s+x\s+([0-9a-zA-Z]+)\b`)
	input = re.ReplaceAllString(input, "$1 \\times $2")
	
	// Chuyển đổi dấu mũ
	input = strings.Replace(input, "^", "^{", -1)
	input = strings.Replace(input, "_", "_{", -1)
	re = regexp.MustCompile(`\^{(\d+)}`)
	input = re.ReplaceAllString(input, "^{$1}")
	re = regexp.MustCompile(`_{(\d+)}`)
	input = re.ReplaceAllString(input, "_{$1}")
	
	// Thêm $ ở đầu và cuối các dòng có công thức toán học
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		if strings.Contains(line, "\\times") || strings.Contains(line, "^{") || strings.Contains(line, "_{") {
			// Kiểm tra nếu dòng chưa có dấu $
			if !strings.Contains(line, "$") {
				lines[i] = "$" + line + "$"
			}
		}
	}
	input = strings.Join(lines, "\n")
	
	return input
}

// convertBinary2OctalToLaTeX chuyển đổi số nhị phân sang số bát phân và định dạng kết quả ở dạng LaTeX
func convertBinary2OctalToLaTeX(input string) string {
	// Trích xuất số nhị phân từ input
	var binaryStr string
	parts := strings.Fields(input)
	for _, part := range parts {
		if utils.IsValidBinary(part) {
			binaryStr = part
			break
		}
	}

	if binaryStr == "" {
		return convertGenericToLaTeX(input)
	}

	// Chuyển đổi số nhị phân sang thập phân
	decimal, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		return convertGenericToLaTeX(input)
	}

	// Chuyển số thập phân sang bát phân
	octalStr := strconv.FormatInt(decimal, 8)

	// Xây dựng kết quả LaTeX
	var result strings.Builder

	// Bước 1: Chuyển từ nhị phân sang thập phân
	result.WriteString("\\begin{enumerate}\n")
	result.WriteString("\\item Convert Binary to Decimal \\\\\n")
	result.WriteString("Use the formula:\n")
	result.WriteString("\\[\n")
	result.WriteString("\\text{Decimal} = d_1 \\times 2^{n-1} + d_2 \\times 2^{n-2} + \\dots + d_n \\times 2^0\n")
	result.WriteString("\\]\n")
	result.WriteString("where \\(d_i\\) is the \\(i\\)-th digit of the binary number (0 or 1), and \\(n\\) is the number of digits. \\\\\n")
	
	// Ứng dụng công thức cho số nhị phân cụ thể
	result.WriteString(fmt.Sprintf("For \\(%s_2\\) (\\(n = %d\\) digits):\n", binaryStr, len(binaryStr)))
	result.WriteString("\\begin{itemize}\n")
	
	// Tính toán giá trị của từng bit
	var sum int64
	digits := []rune(binaryStr)
	for i, digit := range digits {
		//position := len(digits) - i
		digitValue, _ := strconv.ParseInt(string(digit), 2, 64)
		
		power := len(digits) - i - 1
		powerValue := int64(math.Pow(2, float64(power)))
		value := digitValue * powerValue
		sum += value
		
		ordinal := getOrdinalText(i+1)
		result.WriteString(fmt.Sprintf("    \\item %s digit (\\(d_%d = %s\\)): \\(%s \\times 2^{%d-%d} = %s \\times 2^%d = %s \\times %d = %d\\)\n", 
			ordinal, i+1, string(digit), string(digit), len(digits), i+1, string(digit), power, string(digit), powerValue, value))
	}
	
	// Tổng kết bước 1
	result.WriteString(fmt.Sprintf("    \\item Sum: \\(%s\\)\n", formatSum(digits)))
	result.WriteString("\\end{itemize}\n")
	result.WriteString(fmt.Sprintf("So, \\(%s_2 = %d_{10}\\).\n\n", binaryStr, decimal))
	
	// Bước 2: Chuyển từ thập phân sang bát phân
	result.WriteString("\\item Convert Decimal to Octal \\\\\n")
	result.WriteString("Divide the decimal number by 8 repeatedly, record the remainders, and read the remainders from bottom to top. \\\\\n")
	result.WriteString(fmt.Sprintf("For \\(%d_{10}\\):\n", decimal))
	result.WriteString("\\begin{itemize}\n")
	
	// Tạo phép chia liên tiếp
	var divisions []string
	var remainders []int
	tempDecimal := decimal
	
	for tempDecimal > 0 {
		remainder := tempDecimal % 8
		quotient := tempDecimal / 8
		divisions = append(divisions, fmt.Sprintf("    \\item \\(%d \\div 8 = %d\\), remainder \\(%d\\)", tempDecimal, quotient, remainder))
		remainders = append([]int{int(remainder)}, remainders...)
		tempDecimal = quotient
	}
	
	// Hiển thị các phép chia
	for _, division := range divisions {
		result.WriteString(division + "\n")
	}
	
	result.WriteString("\\end{itemize}\n")
	
	// Đọc các số dư từ dưới lên trên
	var remainderStr strings.Builder
	for _, r := range remainders {
		remainderStr.WriteString(strconv.Itoa(r))
	}
	
	result.WriteString(fmt.Sprintf("Reading the remainders from bottom to top: \\(%s_8\\).\n\n", remainderStr.String()))
	
	// Kết quả cuối cùng
	result.WriteString("\\begin{center}\n")
	result.WriteString(fmt.Sprintf("\\textbf{Final Answer:} \\(%s_2 = %s_8\\)\n", binaryStr, octalStr))
	result.WriteString("\\end{center}\n")
	result.WriteString("\\end{enumerate}\n")
	
	return result.String()
}

// formatSum tạo chuỗi hiển thị phép cộng của các giá trị bit
func formatSum(digits []rune) string {
	var terms []string
	var sum int64
	
	for i, digit := range digits {
		digitValue, _ := strconv.ParseInt(string(digit), 2, 64)
		power := len(digits) - i - 1
		powerValue := int64(math.Pow(2, float64(power)))
		value := digitValue * powerValue
		
		if value > 0 {
			terms = append(terms, fmt.Sprintf("%d", value))
			sum += value
		}
	}
	
	return strings.Join(terms, " + ") + " = " + strconv.FormatInt(sum, 10)
}

// ExportToExcelWithLaTeX xuất kết quả sang Excel với định dạng LaTeX
func ExportToExcelWithLaTeX(results []*StepByStepResult, filename string) error {
	f := excelize.NewFile()
	
	// Tạo sheet mới
	sheetName := "Chuyển đổi cơ số"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return err
	}
	
	// Đặt tiêu đề cột
	f.SetCellValue(sheetName, "A1", "Input")
	f.SetCellValue(sheetName, "B1", "Solution")
	f.SetCellValue(sheetName, "C1", "Output")
	
	// Đổ dữ liệu
	for i, result := range results {
		row := i + 2
		
		// Input - định dạng theo yêu cầu mới
		inputCell := fmt.Sprintf("A%d", row)
		inputValue := formatInputQuestion(result)
		f.SetCellValue(sheetName, inputCell, inputValue)
		
		// Solution (các bước với định dạng LaTeX)
		solutionCell := fmt.Sprintf("B%d", row)
		
		// Xử lý tất cả các trường hợp chuyển đổi
		var solutionValue string
		
		// Biến đổi từng bước thành định dạng LaTeX
		stepsStr := strings.Join(result.Steps, "\n")
		
		// Áp dụng chuyển đổi LaTeX dựa trên loại chuyển đổi
		if result.InputBase == utils.DECIMAL && result.OutputBase == utils.BINARY {
			// Sử dụng hàm đã được cải tiến để chuyển đổi từ thập phân sang nhị phân
			solutionValue = convertDecimal2BinaryToLaTeX(stepsStr)
		} else if result.InputBase == utils.DECIMAL && result.OutputBase == utils.OCTAL {
			// Sử dụng hàm chuyển đổi từ thập phân sang bát phân
			solutionValue = convertDecimal2OctalToLaTeX(stepsStr)
		} else if result.InputBase == utils.DECIMAL && result.OutputBase == utils.HEXADECIMAL {
			// Sử dụng hàm chuyển đổi từ thập phân sang thập lục phân
			solutionValue = convertDecimal2HexadecimalToLaTeX(stepsStr)
		} else if result.InputBase == utils.BINARY && result.OutputBase == utils.DECIMAL {
			// Sử dụng hàm chuyển đổi từ nhị phân sang thập phân
			solutionValue = convertBinary2DecimalToLaTeX(stepsStr)
		} else if result.InputBase == utils.BINARY && result.OutputBase == utils.OCTAL {
			// Sử dụng hàm chuyển đổi từ nhị phân sang bát phân
			solutionValue = convertBinary2OctalToLaTeX(stepsStr)
		} else if result.InputBase == utils.BINARY && result.OutputBase == utils.HEXADECIMAL {
			// Sử dụng hàm chuyển đổi từ nhị phân sang thập lục phân
			solutionValue = convertBinary2HexadecimalToLaTeX(stepsStr)
		} else if result.InputBase == utils.OCTAL && result.OutputBase == utils.DECIMAL {
			// Sử dụng hàm chuyển đổi từ bát phân sang thập phân
			solutionValue = convertOctal2DecimalToLaTeX(stepsStr)
		} else if result.InputBase == utils.HEXADECIMAL && result.OutputBase == utils.DECIMAL {
			// Sử dụng hàm chuyển đổi từ thập lục phân sang thập phân
			solutionValue = convertHexadecimal2DecimalToLaTeX(stepsStr)
		} else if result.InputBase == utils.HEXADECIMAL && result.OutputBase == utils.BINARY {
			// Sử dụng hàm chuyển đổi từ thập lục phân sang nhị phân
			solutionValue = convertHexadecimal2BinaryToLaTeX(stepsStr)
		} else if result.InputBase == utils.OCTAL && result.OutputBase == utils.BINARY {
			// Sử dụng hàm chuyển đổi từ bát phân sang nhị phân
			solutionValue = convertOctal2BinaryToLaTeX(stepsStr)
		} else if result.InputBase == utils.OCTAL && result.OutputBase == utils.HEXADECIMAL {
			// Sử dụng hàm chuyển đổi từ bát phân sang thập lục phân
			solutionValue = convertOctal2HexadecimalToLaTeX(stepsStr)
		} else if result.InputBase == utils.HEXADECIMAL && result.OutputBase == utils.OCTAL {
			// Sử dụng hàm chuyển đổi từ thập lục phân sang bát phân
			solutionValue = convertHexadecimal2OctalToLaTeX(stepsStr)
		} else {
			// Xử lý các trường hợp còn lại bằng cách xử lý từng dòng
			for _, step := range result.Steps {
				latexStep := ConvertToLaTeX(step)
				if latexStep != "" {
					solutionValue += latexStep + "\n"
				}
			}
		}
		
		f.SetCellValue(sheetName, solutionCell, solutionValue)
		
		// Output - định dạng theo yêu cầu mới
		outputCell := fmt.Sprintf("C%d", row)
		outputValue := formatOutputAnswer(result)
		f.SetCellValue(sheetName, outputCell, outputValue)
	}
	
	// Điều chỉnh độ rộng cột
	f.SetColWidth(sheetName, "A", "A", 45)
	f.SetColWidth(sheetName, "B", "B", 80)
	f.SetColWidth(sheetName, "C", "C", 35)
	
	// Đặt sheet này làm mặc định
	f.SetActiveSheet(index)
	
	// Lưu file
	if err := f.SaveAs(filename); err != nil {
		return err
	}
	
	return nil
}

// formatInputQuestion định dạng câu hỏi chuyển đổi theo định dạng yêu cầu
func formatInputQuestion(result *StepByStepResult) string {
	switch result.InputBase {
	case utils.BINARY:
		if result.OutputBase == "all" {
			return fmt.Sprintf("Convert the binary number $%s_{2}$ to decimal, octal, and hexadecimal.", 
				result.Input)
		}
		return fmt.Sprintf("Convert the binary number $%s_{2}$ to %s.", 
			result.Input, getReadableBaseName(result.OutputBase))
	case utils.DECIMAL:
		if result.OutputBase == "all" {
			return fmt.Sprintf("Convert the decimal number $%s_{10}$ to binary, octal, and hexadecimal.", 
				result.Input)
		}
		return fmt.Sprintf("Convert the decimal number $%s_{10}$ to %s.", 
			result.Input, getReadableBaseName(result.OutputBase))
	case utils.OCTAL:
		if result.OutputBase == "all" {
			return fmt.Sprintf("Convert the octal number $%s_{8}$ to binary, decimal, and hexadecimal.", 
				result.Input)
		}
		return fmt.Sprintf("Convert the octal number $%s_{8}$ to %s.", 
			result.Input, getReadableBaseName(result.OutputBase))
	case utils.HEXADECIMAL:
		if result.OutputBase == "all" {
			return fmt.Sprintf("Convert the hexadecimal number $%s_{16}$ to binary, octal, and decimal.", 
				result.Input)
		}
		return fmt.Sprintf("Convert the hexadecimal number $%s_{16}$ to %s.", 
			result.Input, getReadableBaseName(result.OutputBase))
	default:
		if result.OutputBase == "all" {
			return fmt.Sprintf("Convert %s (base %s) to all other number bases.",
				result.Input, FormatBaseName(result.InputBase))
		}
		return fmt.Sprintf("Convert %s (base %s) to %s.",
			result.Input, FormatBaseName(result.InputBase), 
			getReadableBaseName(result.OutputBase))
	}
}

// formatOutputAnswer định dạng kết quả chuyển đổi theo định dạng yêu cầu
func formatOutputAnswer(result *StepByStepResult) string {
	// Xử lý trường hợp đặc biệt nếu đầu ra của result là "all"
	if result.OutputBase == "all" {
		// Trong thực tế, chúng ta không tạo result với OutputBase là "all"
		// nhưng vẫn xử lý phòng trường hợp trong tương lai
		var outputs []string
		if result.InputBase != utils.BINARY {
			outputs = append(outputs, fmt.Sprintf("$%s_2$", 
				getFunctionResult(result.Input, result.InputBase, utils.BINARY)))
		}
		if result.InputBase != utils.OCTAL {
			outputs = append(outputs, fmt.Sprintf("$%s_8$", 
				getFunctionResult(result.Input, result.InputBase, utils.OCTAL)))
		}
		if result.InputBase != utils.DECIMAL {
			outputs = append(outputs, fmt.Sprintf("$%s_{10}$", 
				getFunctionResult(result.Input, result.InputBase, utils.DECIMAL)))
		}
		if result.InputBase != utils.HEXADECIMAL {
			outputs = append(outputs, fmt.Sprintf("$%s_{16}$", 
				getFunctionResult(result.Input, result.InputBase, utils.HEXADECIMAL)))
		}
		return strings.Join(outputs, ", ")
	}

	// Xử lý các trường hợp thông thường
	switch result.OutputBase {
	case utils.BINARY:
		return fmt.Sprintf("$%s_{2}$", result.Output)
	case utils.DECIMAL:
		return fmt.Sprintf("$%s_{10}$", result.Output)
	case utils.OCTAL:
		return fmt.Sprintf("$%s_{8}$", result.Output)
	case utils.HEXADECIMAL:
		return fmt.Sprintf("$%s_{16}$", result.Output)
	default:
		return fmt.Sprintf("$%s$ (%s)",
			result.Output, getReadableBaseName(result.OutputBase))
	}
}

// getFunctionResult gọi hàm chuyển đổi tương ứng để lấy kết quả
func getFunctionResult(input, inputBase, outputBase string) string {
	var result string
	
	// Map các hàm chuyển đổi
	conversionMap := map[string]func(string) string{
		"binary|decimal":      func(s string) string { r := Binary2DecimalSteps(s); return r.Output },
		"binary|octal":        func(s string) string { r := Binary2OctalSteps(s); return r.Output },
		"binary|hexadecimal":  func(s string) string { r := Binary2HexadecimalSteps(s); return r.Output },
		"decimal|binary":      func(s string) string { r := Decimal2BinarySteps(s); return r.Output },
		"decimal|octal":       func(s string) string { r := Decimal2OctalSteps(s); return r.Output },
		"decimal|hexadecimal": func(s string) string { r := Decimal2HexadecimalSteps(s); return r.Output },
		"octal|binary":        func(s string) string { r := Octal2BinarySteps(s); return r.Output },
		"octal|decimal":       func(s string) string { r := Octal2DecimalSteps(s); return r.Output },
		"octal|hexadecimal":   func(s string) string { r := Octal2HexadecimalSteps(s); return r.Output },
		"hexadecimal|binary":  func(s string) string { r := Hexadecimal2BinarySteps(s); return r.Output },
		"hexadecimal|decimal": func(s string) string { r := Hexadecimal2DecimalSteps(s); return r.Output },
		"hexadecimal|octal":   func(s string) string { r := Hexadecimal2OctalSteps(s); return r.Output },
	}
	
	key := inputBase + "|" + outputBase
	if convFunc, exists := conversionMap[key]; exists {
		result = convFunc(input)
	} else {
		result = "N/A"
	}
	
	return result
}

// getReadableBaseName trả về tên dễ đọc của cơ số dành cho câu hỏi
func getReadableBaseName(base string) string {
	switch base {
	case utils.BINARY:
		return "binary"
	case utils.OCTAL:
		return "octal"
	case utils.DECIMAL:
		return "decimal"
	case utils.HEXADECIMAL:
		return "hexadecimal"
	case "all":
		return "binary, octal, decimal, and hexadecimal"
	default:
		return base
	}
}

// ReadInputFromTxt đọc danh sách các số cần chuyển đổi từ file txt
func ReadInputFromTxt(filename string) ([]struct{Input string; FromBase string; ToBase string}, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	var inputs []struct{Input string; FromBase string; ToBase string}
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		
		// Format mỗi dòng: <số> <cơ số đầu> <cơ số đích>
		if len(fields) >= 3 {
			item := struct{Input string; FromBase string; ToBase string}{
				Input:    fields[0],
				FromBase: fields[1],
				ToBase:   fields[2],
			}
			inputs = append(inputs, item)
		}
	}
	
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	
	return inputs, nil
}

// ExportToExcel xuất kết quả ra file Excel
func ExportToExcel(results []*StepByStepResult, filename string) error {
	f := excelize.NewFile()
	
	// Tạo sheet mới
	sheetName := "Chuyển đổi cơ số"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return err
	}
	
	// Đặt tiêu đề cột
	f.SetCellValue(sheetName, "A1", "Input")
	f.SetCellValue(sheetName, "B1", "Solution")
	f.SetCellValue(sheetName, "C1", "Output")
	
	// Đổ dữ liệu
	for i, result := range results {
		row := i + 2
		
		// Input
		inputCell := fmt.Sprintf("A%d", row)
		f.SetCellValue(sheetName, inputCell, fmt.Sprintf("%s (cơ số %s)", 
			result.Input, FormatBaseName(result.InputBase)))
		
		// Solution (các bước)
		solutionCell := fmt.Sprintf("B%d", row)
		solutionValue := ""
		for _, step := range result.Steps {
			solutionValue += step + "\n"
		}
		f.SetCellValue(sheetName, solutionCell, solutionValue)
		
		// Output
		outputCell := fmt.Sprintf("C%d", row)
		f.SetCellValue(sheetName, outputCell, fmt.Sprintf("%s (cơ số %s)", 
			result.Output, FormatBaseName(result.OutputBase)))
	}
	
	// Điều chỉnh độ rộng cột
	f.SetColWidth(sheetName, "A", "A", 25)
	f.SetColWidth(sheetName, "B", "B", 60)
	f.SetColWidth(sheetName, "C", "C", 25)
	
	// Đặt sheet này làm mặc định
	f.SetActiveSheet(index)
	
	// Lưu file
	if err := f.SaveAs(filename); err != nil {
		return err
	}
	
	return nil
} 

// convertBinary2HexadecimalToLaTeX chuyển đổi giải thích từ nhị phân sang thập lục phân sang định dạng LaTeX
func convertBinary2HexadecimalToLaTeX(input string) string {
	// Trích xuất số nhị phân từ input
	var binaryStr string
	parts := strings.Fields(input)
	for _, part := range parts {
		if utils.IsValidBinary(part) {
			binaryStr = part
			break
		}
	}

	if binaryStr == "" {
		return convertGenericToLaTeX(input)
	}

	// Chuyển đổi số nhị phân sang thập phân
	decimal, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		return convertGenericToLaTeX(input)
	}

	// Chuyển số thập phân sang thập lục phân
	hexStr := strings.ToUpper(strconv.FormatInt(decimal, 16))

	// Bảng chuyển đổi nhị phân sang thập lục phân
	binaryToHex := map[string]string{
		"0000": "0", "0001": "1", "0010": "2", "0011": "3",
		"0100": "4", "0101": "5", "0110": "6", "0111": "7",
		"1000": "8", "1001": "9", "1010": "A", "1011": "B",
		"1100": "C", "1101": "D", "1110": "E", "1111": "F",
	}

	// Xây dựng kết quả LaTeX
	var result strings.Builder

	// Tiêu đề và phương pháp
	result.WriteString("\\textbf{Method:} Starting from the rightmost bit, divide the binary number into groups of 4 bits, adding leading zeros to the leftmost group if necessary to complete a group of 4, then convert each 4-bit group into its corresponding hexadecimal digit using a conversion table, and concatenate the digits from left to right to form the final hexadecimal number. \\\\\n")
	
	// Chuẩn bị số nhị phân để chia nhóm (thêm 0 ở đầu nếu cần để tạo nhóm 4 bit đầy đủ)
	paddedBinary := binaryStr
	if len(binaryStr) % 4 != 0 {
		padding := 4 - (len(binaryStr) % 4)
		paddedBinary = strings.Repeat("0", padding) + binaryStr
	}
	
	// Chia nhóm và chuyển đổi
	var groups []string
	var hexDigits []string
	
	for i := 0; i < len(paddedBinary); i += 4 {
		end := i + 4
		if end > len(paddedBinary) {
			end = len(paddedBinary)
		}
		group := paddedBinary[i:end]
		groups = append(groups, group)
		hexDigits = append(hexDigits, binaryToHex[group])
	}
	
	// Hiển thị bảng chuyển đổi
	result.WriteString("\\begin{center}\n")
	result.WriteString("\\begin{tabular}{|c|c|} \\hline\n")
	result.WriteString("\\text{Binary Group} & \\text{Hexadecimal} \\\\ \\hline\n")
	
	for i := 0; i < len(groups); i++ {
		result.WriteString(fmt.Sprintf("%s & %s \\\\ \\hline\n", groups[i], hexDigits[i]))
	}
	
	result.WriteString("\\end{tabular}\n")
	result.WriteString("\\end{center}\n")
	
	// Kết quả cuối cùng
	result.WriteString("\\begin{center}\n")
	result.WriteString(fmt.Sprintf("\\textbf{Final Answer:} \\(%s_2 = %s_{16}\\)\n", binaryStr, hexStr))
	result.WriteString("\\end{center}\n")
	
	return result.String()
}

// convertOctal2BinaryToLaTeX chuyển đổi giải thích từ bát phân sang nhị phân sang định dạng LaTeX
func convertOctal2BinaryToLaTeX(input string) string {
	// Trích xuất số bát phân từ input
	var octalStr string
	parts := strings.Fields(input)
	for _, part := range parts {
		if matched, _ := regexp.MatchString(`^[0-7]+$`, part); matched {
			octalStr = part
			break
		}
	}

	if octalStr == "" {
		return convertGenericToLaTeX(input)
	}

	// Kiểm tra cú pháp đúng của số bát phân
	if _, err := strconv.ParseInt(octalStr, 8, 64); err != nil {
		return convertGenericToLaTeX(input)
	}

	// Bảng chuyển đổi bát phân sang nhị phân
	octalToBinary := map[rune]string{
		'0': "000", '1': "001", '2': "010", '3': "011",
		'4': "100", '5': "101", '6': "110", '7': "111",
	}

	// Xây dựng kết quả LaTeX
	var result strings.Builder

	// Tiêu đề và phương pháp
	result.WriteString("\\begin{enumerate}\n")
	result.WriteString("\\item Method: Transform each octal digit into its corresponding 3-bit binary representation using a conversion table, then concatenate the binary sequences from left to right to form the final binary number. \\\\\n")
	
	// Hiển thị bảng chuyển đổi các chữ số trong số bát phân
	result.WriteString("\\begin{center}\n")
	result.WriteString("Octal to Binary conversion table:\n")
	result.WriteString("\\begin{tabular}{|c|c|} \\hline\n")
	result.WriteString("\\text{Octal Digit} & \\text{Binary Equivalent} \\\\ \\hline\n")
	
	// Hiển thị chỉ các chữ số bát phân xuất hiện trong số đầu vào
	uniqueDigits := make(map[rune]bool)
	for _, digit := range octalStr {
		uniqueDigits[digit] = true
	}
	
	// Sắp xếp các chữ số để hiển thị theo thứ tự
	var sortedDigits []rune
	for digit := range uniqueDigits {
		sortedDigits = append(sortedDigits, digit)
	}
	sort.Slice(sortedDigits, func(i, j int) bool {
		return sortedDigits[i] < sortedDigits[j]
	})
	
	for _, digit := range sortedDigits {
		result.WriteString(fmt.Sprintf("%c & %s \\\\ \\hline\n", digit, octalToBinary[digit]))
	}
	
	result.WriteString("\\end{tabular}\n")
	result.WriteString("\\end{center}\n\n")
	
	// Chuyển đổi từng chữ số
	result.WriteString(fmt.Sprintf("\\item For octal number \\(%s_8\\):\n", octalStr))
	result.WriteString("\\begin{itemize}\n")
	
	// Hiển thị việc chuyển đổi từng chữ số
	var binaryResult strings.Builder
	for i, digit := range octalStr {
		binaryValue := octalToBinary[digit]
		binaryResult.WriteString(binaryValue)
		result.WriteString(fmt.Sprintf("    \\item %s digit \\(%c\\) = \\(%s\\)\n", 
			getOrdinalText(i+1), digit, binaryValue))
	}
	
	result.WriteString("\\end{itemize}\n")
	
	// Loại bỏ các số 0 không cần thiết ở đầu (nếu có)
	finalBinary := strings.TrimLeft(binaryResult.String(), "0")
	if finalBinary == "" {
		finalBinary = "0" // Đảm bảo ít nhất có một chữ số 0
	}
	
	// Kết quả cuối cùng
	result.WriteString("\\begin{center}\n")
	result.WriteString(fmt.Sprintf("\\textbf{Final Answer:} \\(%s_8 = %s_2\\)\n", octalStr, finalBinary))
	result.WriteString("\\end{center}\n")
	result.WriteString("\\end{enumerate}\n")
	
	return result.String()
}

// convertOctal2HexadecimalToLaTeX chuyển đổi giải thích từ bát phân sang thập lục phân sang định dạng LaTeX
func convertOctal2HexadecimalToLaTeX(input string) string {
	// Trích xuất số bát phân từ input
	var octalStr string
	parts := strings.Fields(input)
	for _, part := range parts {
		if matched, _ := regexp.MatchString(`^[0-7]+$`, part); matched {
			octalStr = part
			break
		}
	}

	if octalStr == "" {
		return convertGenericToLaTeX(input)
	}

	// Chuyển đổi số bát phân sang thập phân
	decimal, err := strconv.ParseInt(octalStr, 8, 64)
	if err != nil {
		return convertGenericToLaTeX(input)
	}

	// Chuyển số thập phân sang thập lục phân
	hexStr := strings.ToUpper(strconv.FormatInt(decimal, 16))

	// Xây dựng kết quả LaTeX
	var result strings.Builder

	// Bước 1: Chuyển từ bát phân sang thập phân
	result.WriteString("\\begin{enumerate}\n")
	result.WriteString("\\item Convert Octal to Decimal \\\\\n")
	result.WriteString("Use the formula:\n")
	result.WriteString("\\[\n")
	result.WriteString("\\text{Decimal} = d_1 \\times 8^{n-1} + d_2 \\times 8^{n-2} + \\dots + d_n \\times 8^0\n")
	result.WriteString("\\]\n")
	result.WriteString("where \\(d_i\\) is the \\(i\\)-th digit of the octal number, and \\(n\\) is the number of digits. \\\\\n")
	
	// Ứng dụng công thức cho số bát phân cụ thể
	result.WriteString(fmt.Sprintf("For \\(%s_8\\) (\\(n = %d\\) digits):\n", octalStr, len(octalStr)))
	result.WriteString("\\begin{itemize}\n")
	
	// Tính toán giá trị của từng chữ số
	var sum int64
	digits := []rune(octalStr)
	for i, digit := range digits {
		//position := len(digits) - i
		digitValue, _ := strconv.ParseInt(string(digit), 8, 64)
		
		power := len(digits) - i - 1
		powerValue := int64(math.Pow(8, float64(power)))
		value := digitValue * powerValue
		sum += value
		
		ordinal := getOrdinalText(i+1)
		result.WriteString(fmt.Sprintf("    \\item %s digit (\\(d_%d = %s\\)): \\(%s \\times 8^{%d-%d} = %s \\times 8^%d = %s \\times %d = %d\\)\n", 
			ordinal, i+1, string(digit), string(digit), len(digits), i+1, string(digit), power, string(digit), powerValue, value))
	}
	
	// Tổng kết bước 1
	result.WriteString(fmt.Sprintf("    \\item Sum: \\(%s\\)\n", formatOctalSum(digits)))
	result.WriteString("\\end{itemize}\n")
	result.WriteString(fmt.Sprintf("So, \\(%s_8 = %d_{10}\\).\n\n", octalStr, decimal))
	
	// Bước 2: Chuyển từ thập phân sang thập lục phân
	result.WriteString("\\item Convert Decimal to Hexadecimal \\\\\n")
	result.WriteString("Divide the decimal number by 16 repeatedly, record the remainders, and read the remainders from bottom to top. \\\\\n")
	
	// Giải thích giá trị của chữ số thập lục phân 10-15
	result.WriteString("Note: Hexadecimal digits 10--15 are represented as:\n")
	result.WriteString("\\begin{itemize}\n")
	result.WriteString("    \\item \\(10 = A\\), \\(11 = B\\), \\(12 = C\\), \\(13 = D\\), \\(14 = E\\), \\(15 = F\\)\n")
	result.WriteString("\\end{itemize}\n")
	
	result.WriteString(fmt.Sprintf("For \\(%d_{10}\\):\n", decimal))
	result.WriteString("\\begin{itemize}\n")
	
	// Tạo phép chia liên tiếp
	var divisions []string
	var remainders []string
	tempDecimal := decimal
	
	for tempDecimal > 0 {
		remainder := tempDecimal % 16
		quotient := tempDecimal / 16
		
		var remainderStr string
		if remainder >= 10 {
			// Chuyển đổi chữ số thập lục phân
			remainderStr = string('A' + rune(remainder - 10))
			divisions = append(divisions, fmt.Sprintf("    \\item \\(%d \\div 16 = %d\\), remainder \\(%d = %s\\)", 
				tempDecimal, quotient, remainder, remainderStr))
		} else {
			remainderStr = strconv.FormatInt(remainder, 10)
			divisions = append(divisions, fmt.Sprintf("    \\item \\(%d \\div 16 = %d\\), remainder \\(%s\\)", 
				tempDecimal, quotient, remainderStr))
		}
		
		remainders = append([]string{remainderStr}, remainders...)
		tempDecimal = quotient
	}
	
	// Hiển thị các phép chia
	for _, division := range divisions {
		result.WriteString(division + "\n")
	}
	
	result.WriteString("\\end{itemize}\n")
	
	// Đọc các số dư từ dưới lên trên
	var remainderStr strings.Builder
	for _, r := range remainders {
		remainderStr.WriteString(r)
	}
	
	result.WriteString(fmt.Sprintf("Reading the remainders from bottom to top: \\(%s_{16}\\).\n\n", remainderStr.String()))
	
	// Kết quả cuối cùng
	result.WriteString("\\begin{center}\n")
	result.WriteString(fmt.Sprintf("\\textbf{Final Answer:} \\(%s_8 = %s_{16}\\)\n", octalStr, hexStr))
	result.WriteString("\\end{center}\n")
	result.WriteString("\\end{enumerate}\n")
	
	return result.String()
}

// formatOctalSum tạo chuỗi hiển thị phép cộng của các giá trị chữ số bát phân
func formatOctalSum(digits []rune) string {
	var terms []string
	var sum int64
	
	for i, digit := range digits {
		digitValue, _ := strconv.ParseInt(string(digit), 8, 64)
		power := len(digits) - i - 1
		powerValue := int64(math.Pow(8, float64(power)))
		value := digitValue * powerValue
		
		if value > 0 {
			terms = append(terms, fmt.Sprintf("%d", value))
			sum += value
		}
	}
	
	return strings.Join(terms, " + ") + " = " + strconv.FormatInt(sum, 10)
}

// convertLineToLaTeX chuyển đổi một dòng văn bản thành biểu thức LaTeX
func convertLineToLaTeX(line string) string {
	// Thay thế các ký tự đặc biệt
	line = strings.Replace(line, "<=", "\\leq", -1)
	line = strings.Replace(line, ">=", "\\geq", -1)
	line = strings.Replace(line, "<", "\\lt", -1)
	line = strings.Replace(line, ">", "\\gt", -1)
	
	// Thay thế các dấu "x" giữa các số bằng dấu nhân
	re := regexp.MustCompile(`(\d+)\s*[xX]\s*(\d+)`)
	line = re.ReplaceAllString(line, "$1 \\times $2")
	
	// Bảo tồn khoảng trắng
	line = strings.Replace(line, " ", "~", -1)
	
	return line
} 

// formatHexSum tạo chuỗi hiển thị phép cộng của các giá trị chữ số thập lục phân
func formatHexSum(digits []rune) string {
	var terms []string
	var sum int64
	
	for i, digit := range digits {
		var digitValue int64
		
		// Xử lý chữ cái A-F
		if digit >= 'A' && digit <= 'F' {
			digitValue = int64(digit - 'A' + 10)
		} else {
			digitValue, _ = strconv.ParseInt(string(digit), 16, 64)
		}
		
		power := len(digits) - i - 1
		powerValue := int64(math.Pow(16, float64(power)))
		value := digitValue * powerValue
		
		if value > 0 {
			terms = append(terms, fmt.Sprintf("%d", value))
			sum += value
		}
	}
	
	return strings.Join(terms, " + ") + " = " + strconv.FormatInt(sum, 10)
}

// convertHexadecimal2BinaryToLaTeX chuyển đổi giải thích từ thập lục phân sang nhị phân
func convertHexadecimal2BinaryToLaTeX(input string) string {
	// Trích xuất số thập lục phân từ input
	var hexStr string
	parts := strings.Fields(input)
	for _, part := range parts {
		if matched, _ := regexp.MatchString(`^[0-9A-Fa-f]+$`, part); matched {
			hexStr = strings.ToUpper(part)
			break
		}
	}

	if hexStr == "" {
		return convertGenericToLaTeX(input)
	}

	// Bảng chuyển đổi từ thập lục phân sang nhị phân
	hexToBinary := map[rune]string{
		'0': "0000", '1': "0001", '2': "0010", '3': "0011",
		'4': "0100", '5': "0101", '6': "0110", '7': "0111",
		'8': "1000", '9': "1001", 'A': "1010", 'B': "1011",
		'C': "1100", 'D': "1101", 'E': "1110", 'F': "1111",
	}

	// Xây dựng kết quả LaTeX
	var result strings.Builder
	result.WriteString("\\begin{enumerate}\n")
	// Tiêu đề và phương pháp
	result.WriteString("\\item Method: Transform each hexadecimal digit into its corresponding 4-bit binary representation using a conversion table, then concatenate the binary sequences from left to right to form the final binary number. \\\\\n")
	
	// Bảng chuyển đổi
	
	result.WriteString("Hexadecimal to Binary conversion table:\n")
	result.WriteString("\\begin{tabular}{|c|c|} \\hline\n")
	result.WriteString("\\text{Hex Digit} & \\text{Binary Equivalent} \\\\ \\hline\n")
	
	// Hiển thị các chữ số thập lục phân cần thiết cho số này
	uniqueDigits := make(map[rune]bool)
	for _, digit := range hexStr {
		uniqueDigits[digit] = true
	}
	
	// Sắp xếp các chữ số để hiển thị theo thứ tự
	var sortedDigits []rune
	for digit := range uniqueDigits {
		sortedDigits = append(sortedDigits, digit)
	}
	sort.Slice(sortedDigits, func(i, j int) bool {
		// Chữ số thường xếp trước chữ cái
		if (sortedDigits[i] >= '0' && sortedDigits[i] <= '9') && (sortedDigits[j] >= 'A' && sortedDigits[j] <= 'F') {
			return true
		}
		if (sortedDigits[i] >= 'A' && sortedDigits[i] <= 'F') && (sortedDigits[j] >= '0' && sortedDigits[j] <= '9') {
			return false
		}
		return sortedDigits[i] < sortedDigits[j]
	})
	
	for _, digit := range sortedDigits {
		result.WriteString(fmt.Sprintf("%c & %s \\\\ \\hline\n", digit, hexToBinary[digit]))
	}
	
	result.WriteString("\\end{tabular}\n")

	
	// Chuyển đổi từng chữ số
	result.WriteString(fmt.Sprintf("\\item For hexadecimal number \\(%s_{16}\\):\n", hexStr))
	result.WriteString("\\begin{itemize}\n")
	
	for i, digit := range hexStr {
		result.WriteString(fmt.Sprintf("    \\item %s digit \\(%c\\) = \\(%s\\)\n", 
			getOrdinalText(i+1), digit, hexToBinary[digit]))
	}
	
	result.WriteString("\\end{itemize}\n")
	
	// Kết quả nhị phân
	var binaryResult strings.Builder
	for _, digit := range hexStr {
		binaryResult.WriteString(hexToBinary[digit])
	}
	
	// Loại bỏ các số 0 không cần thiết ở đầu (nếu có)
	binaryStr := strings.TrimLeft(binaryResult.String(), "0")
	if binaryStr == "" {
		binaryStr = "0" // Đảm bảo ít nhất có một chữ số 0
	}
	
	result.WriteString("\\begin{center}\n")
	result.WriteString(fmt.Sprintf("Final Answer: \\(%s_{16} = %s_{2}\\)\n", hexStr, binaryStr))
	result.WriteString("\\end{center}\n")
	result.WriteString("\\end{enumerate}\n")
	
	return result.String()
}

// convertHexadecimal2OctalToLaTeX chuyển đổi giải thích từ thập lục phân sang bát phân
func convertHexadecimal2OctalToLaTeX(input string) string {
	// Trích xuất số thập lục phân từ input
	var hexStr string
	parts := strings.Fields(input)
	for _, part := range parts {
		if matched, _ := regexp.MatchString(`^[0-9A-Fa-f]+$`, part); matched {
			hexStr = strings.ToUpper(part)
			break
		}
	}

	if hexStr == "" {
		return convertGenericToLaTeX(input)
	}

	// Chuyển đổi số thập lục phân sang thập phân
	decimal, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		return convertGenericToLaTeX(input)
	}

	// Chuyển số thập phân sang bát phân
	octalStr := strconv.FormatInt(decimal, 8)

	// Xây dựng kết quả LaTeX
	var result strings.Builder

	// Bước 1: Chuyển từ thập lục phân sang thập phân
	result.WriteString("\\begin{enumerate}\n")
	result.WriteString("\\item Convert Hexadecimal to Decimal \\\\\n")
	result.WriteString("Use the formula:\n")
	result.WriteString("\\[\n")
	result.WriteString("\\text{Decimal} = d_1 \\times 16^{n-1} + d_2 \\times 16^{n-2} + \\dots + d_n \\times 16^0\n")
	result.WriteString("\\]\n")
	result.WriteString("where \\(d_i\\) is the \\(i\\)-th digit of the hexadecimal number, and \\(n\\) is the number of digits. \\\\\n")
	
	// Giải thích giá trị của chữ cái A-F
	result.WriteString("In hexadecimal, digits range from 0 to 15, with letters A--F representing values 10--15:\n")
	result.WriteString("\\begin{itemize}\n")
	result.WriteString("\\item \\(A = 10\\), \\(B = 11\\), \\(C = 12\\), \\(D = 13\\), \\(E = 14\\), \\(F = 15\\)\n")
	result.WriteString("\\end{itemize}\n")
	
	// Ứng dụng công thức cho số thập lục phân cụ thể
	result.WriteString(fmt.Sprintf("For \\(%s_{16}\\) (\\(n = %d\\) digits):\n", hexStr, len(hexStr)))
	result.WriteString("\\begin{itemize}\n")
	
	// Tính toán giá trị của từng chữ số
	var sum int64
	digits := []rune(hexStr)
	for i, digit := range digits {
		// position := len(digits) - i
		
		var digitValue int64
		var digitStr string
		
		// Xử lý chữ cái A-F
		if digit >= 'A' && digit <= 'F' {
			digitValue = int64(digit - 'A' + 10)
			digitStr = fmt.Sprintf("%c = %d", digit, digitValue)
		} else {
			digitValue, _ = strconv.ParseInt(string(digit), 16, 64)
			digitStr = string(digit)
		}
		
		power := len(digits) - i - 1
		powerValue := int64(math.Pow(16, float64(power)))
		value := digitValue * powerValue
		sum += value
		
		// ordinal := getOrdinalText(position)
		ordinal := getOrdinalText(i+1)
		result.WriteString(fmt.Sprintf("    \\item %s digit (\\(d_%d = %s\\)): \\(%d \\times 16^{%d-%d} = %d \\times 16^%d = %d \\times %d = %d\\)\n", 
			ordinal, i+1, digitStr, digitValue, len(digits), i+1, digitValue, power, digitValue, powerValue, value))
			// ordinal, position, digitStr, digitValue, len(digits), position, digitValue, power, digitValue, powerValue, value))
	}
	
	// Tổng kết bước 1
	result.WriteString(fmt.Sprintf("    \\item Sum: \\(%s\\)\n", formatHexSum(digits)))
	result.WriteString("\\end{itemize}\n")
	result.WriteString(fmt.Sprintf("So, \\(%s_{16} = %d_{10}\\).\n\n", hexStr, decimal))
	
	// Bước 2: Chuyển từ thập phân sang bát phân
	result.WriteString("\\item Convert Decimal to Octal \\\\\n")
	result.WriteString("Divide the decimal number by 8 repeatedly, record the remainders, and read the remainders from bottom to top. \\\\\n")
	result.WriteString(fmt.Sprintf("For \\(%d_{10}\\):\n", decimal))
	result.WriteString("\\begin{itemize}\n")
	
	// Tạo phép chia liên tiếp
	var divisions []string
	var remainders []int
	tempDecimal := decimal
	
	for tempDecimal > 0 {
		remainder := tempDecimal % 8
		quotient := tempDecimal / 8
		divisions = append(divisions, fmt.Sprintf("    \\item \\(%d \\div 8 = %d\\), remainder \\(%d\\)", tempDecimal, quotient, remainder))
		remainders = append([]int{int(remainder)}, remainders...)
		tempDecimal = quotient
	}
	
	// Hiển thị các phép chia
	for _, division := range divisions {
		result.WriteString(division + "\n")
	}
	
	result.WriteString("\\end{itemize}\n")
	
	// Đọc các số dư từ dưới lên trên
	var remainderStr strings.Builder
	for _, r := range remainders {
		remainderStr.WriteString(strconv.Itoa(r))
	}
	
	result.WriteString(fmt.Sprintf("Reading the remainders from bottom to top: \\(%s_8\\).\n\n", remainderStr.String()))
	
	// Kết quả cuối cùng
	result.WriteString("\\begin{center}\n")
	result.WriteString(fmt.Sprintf("\\textbf{Final Answer:} \\(%s_{16} = %s_8\\)\n", hexStr, octalStr))
	result.WriteString("\\end{center}\n")
	result.WriteString("\\end{enumerate}\n")
	
	return result.String()
}