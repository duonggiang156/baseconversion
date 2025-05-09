/*

Copyright 2018 Travis Clarke. All rights reserved.
Use of this source code is governed by a Apache-2.0
license that can be found in the LICENSE file.

NAME:
    ncalc – number base converter.

SYNOPSIS:
    ncalc [ opts... ] [ number|character ]

OPTIONS:
    -h, --help                  print usage.
    -i, --input format          input format. see FORMATS. (default: decimal|ascii)
    -o, --output format         output format. see FORMATS. (default: all)
    -q, --quiet                 suppress printing of output format type(s)
    -s, --steps                 show step-by-step solution
    -e, --excel filename        export step-by-step solution to excel file
    -l, --latex                 use LaTeX formatting in excel output
    -f, --file filename         read input from text file
    -v, --version               print version number.

FORMATS:
    (a)scii                     character
    (b)inary                    base 2
    (o)ctal                     base 8
    (d)ecimal                   base 10
    (h)exadecimal               base 16

EXAMPLES:
    ncalc "6"                               # output `decimal` number `6` in `all` formats
    ncalc "G"                               # output `ascii` character `G` in `all` formats
    ncalc -i a "f"                          # output `ascii` character `f` in `all` formats
    ncalc -i decimal -o ascii "15"          # output `decimal` number `15` as `ascii`
    ncalc --input h --output o "ff"         # output `hexadecimal` number `ff` as `octal`
    ncalc -i d -o b -s "15"                 # convert decimal 15 to binary with steps
    ncalc -i d -o all -e "result.xlsx" "42" # export conversions to Excel file
    ncalc -f "input.txt" -e "result.xlsx" -l # read from text file, export to excel with LaTeX

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	flag "github.com/clarketm/pflag"

	"github.com/clarketm/ncalc/ascii"
	"github.com/clarketm/ncalc/binary"
	"github.com/clarketm/ncalc/decimal"
	"github.com/clarketm/ncalc/hexadecimal"
	"github.com/clarketm/ncalc/octal"
	"github.com/clarketm/ncalc/stepbystep"
	"github.com/clarketm/ncalc/utils"
	"github.com/fatih/color"
)

// VERSION - current version number
const VERSION = "v1.2.3"

type inputFlag []string

func (i *inputFlag) String() string {
	return "decimal|ascii"
}

func (i *inputFlag) Type() string {
	return "string"
}

func (i *inputFlag) Set(value string) error {
	*i = getFormat(value)
	return nil
}

type outputFlag []string

func (o *outputFlag) String() string {
	return "all"
}
func (o *outputFlag) Type() string {
	return "string"
}

func (o *outputFlag) Set(value string) error {
	*o = getFormat(value)
	return nil
}

// Flags
var version bool
var quiet bool
var showSteps bool
var excelFile string
var inputFile string
var useLaTeX bool

var inputFormat inputFlag
var outputFormat outputFlag = utils.ALL

// Globals
var statusCode int
var bold = color.New(color.Bold).SprintFunc()
var funcMap = map[string]interface{}{
	"ascii|ascii":             ascii.String,
	"ascii|binary":            ascii.Ascii2Binary,
	"ascii|octal":             ascii.Ascii2Octal,
	"ascii|decimal":           ascii.Ascii2Decimal,
	"ascii|hexadecimal":       ascii.Ascii2Hexadecimal,
	"binary|ascii":            binary.Binary2Ascii,
	"binary|binary":           binary.String,
	"binary|octal":            binary.Binary2Octal,
	"binary|decimal":          binary.Binary2Decimal,
	"binary|hexadecimal":      binary.Binary2Hexadecimal,
	"octal|ascii":             octal.Octal2Ascii,
	"octal|binary":            octal.Octal2Binary,
	"octal|octal":             octal.String,
	"octal|decimal":           octal.Octal2Decimal,
	"octal|hexadecimal":       octal.Octal2Hexadecimal,
	"decimal|ascii":           decimal.Decimal2Ascii,
	"decimal|binary":          decimal.Decimal2Binary,
	"decimal|octal":           decimal.Decimal2Octal,
	"decimal|decimal":         decimal.String,
	"decimal|hexadecimal":     decimal.Decimal2Hexadecimal,
	"hexadecimal|ascii":       hexadecimal.Hexadecimal2Ascii,
	"hexadecimal|binary":      hexadecimal.Hexadecimal2Binary,
	"hexadecimal|octal":       hexadecimal.Hexadecimal2Octal,
	"hexadecimal|decimal":     hexadecimal.Hexadecimal2Decimal,
	"hexadecimal|hexadecimal": hexadecimal.String,
}

// Map cho các hàm giải từng bước
var stepsFuncMap = map[string]func(string) *stepbystep.StepByStepResult{
	"binary|decimal":          stepbystep.Binary2DecimalSteps,
	"binary|octal":            stepbystep.Binary2OctalSteps,
	"binary|hexadecimal":      stepbystep.Binary2HexadecimalSteps,
	"octal|decimal":           stepbystep.Octal2DecimalSteps,
	"octal|binary":            stepbystep.Octal2BinarySteps,
	"octal|hexadecimal":       stepbystep.Octal2HexadecimalSteps,
	"decimal|binary":          stepbystep.Decimal2BinarySteps,
	"decimal|octal":           stepbystep.Decimal2OctalSteps,
	"decimal|hexadecimal":     stepbystep.Decimal2HexadecimalSteps,
	"hexadecimal|decimal":     stepbystep.Hexadecimal2DecimalSteps,
	"hexadecimal|binary":      stepbystep.Hexadecimal2BinarySteps,
	"hexadecimal|octal":       stepbystep.Hexadecimal2OctalSteps,
}

// init () - initialize command-line flags
func init() {
	// -q, --quiet
	flag.BoolVarP(&quiet, "quiet", "q", false, "suppress printing of output format type(s)")
	
	// -s, --steps
	flag.BoolVarP(&showSteps, "steps", "s", false, "show step-by-step solution")
	
	// -e, --excel
	flag.StringVarP(&excelFile, "excel", "e", "", "export step-by-step solution to excel file")
	
	// -l, --latex
	flag.BoolVarP(&useLaTeX, "latex", "l", false, "use LaTeX formatting in excel output")
	
	// -f, --file
	flag.StringVarP(&inputFile, "file", "f", "", "read input from text file")

	// -i, --input
	flag.VarP(&inputFormat, "input", "i", "input `format`: see FORMATS.")

	// -o, --output
	flag.VarP(&outputFormat, "output", "o", "output `format`: see FORMATS.")

	// -v, --version
	flag.BoolVarP(&version, "version", "v", false, "print version number")

	// Usage
	flag.Usage = func() {
		println()
		fmt.Printf("NAME:\n")
		fmt.Printf("\tncalc – number base converter.\n")
		println()
		fmt.Printf("SYNOPSIS:\n")
		fmt.Printf("\t%v [ opts... ] [ number|character ]\n", bold("ncalc"))
		println()
		fmt.Printf("OPTIONS:\n")
		flag.PrintDefaults()
		println()
		fmt.Printf("FORMATS:\n")
		fmt.Printf("\t(a)scii      \tcharacter\n")
		fmt.Printf("\t(b)inary     \tbase 2\n")
		fmt.Printf("\t(o)ctal      \tbase 8\n")
		fmt.Printf("\t(d)ecimal    \tbase 10\n")
		fmt.Printf("\t(h)exadecimal\tbase 16\n")
		println()
		os.Exit(statusCode)
	}
}

func printVersion() {
	fmt.Printf("\n%s %v\n", bold("Version:"), VERSION)
	os.Exit(0)
}

func getFormat(format string) []string {
	o := make([]string, 1)

	switch format {
	case "all":
		return utils.ALL
	case utils.ASCII, string(utils.ASCII[0]):
		o = []string{utils.ASCII}
	case utils.BINARY, string(utils.BINARY[0]):
		o = []string{utils.BINARY}
	case utils.OCTAL, string(utils.OCTAL[0]):
		o = []string{utils.OCTAL}
	case utils.DECIMAL, string(utils.DECIMAL[0]):
		o = []string{utils.DECIMAL}
	case utils.HEXADECIMAL, string(utils.HEXADECIMAL[0]):
		o = []string{utils.HEXADECIMAL}
	default:
		fmt.Fprintln(os.Stderr, "Unkown format", format)
		os.Exit(1)
	}
	return o
}

func setDefaultInputFormat(v interface{}) {
	if utils.IsDecimal(v.(string)) {
		inputFormat = []string{utils.DECIMAL}
	} else {
		inputFormat = []string{utils.ASCII}
	}
}

// baseNameToFormat chuyển đổi tên cơ số thành định dạng
func baseNameToFormat(baseName string) string {
	baseName = strings.ToLower(baseName)
	switch baseName {
	case "2", "binary", "b":
		return utils.BINARY
	case "8", "octal", "o":
		return utils.OCTAL
	case "10", "decimal", "d":
		return utils.DECIMAL
	case "16", "hexadecimal", "h", "hex":
		return utils.HEXADECIMAL
	case "ascii", "a":
		return utils.ASCII
	case "all":
		return "all"
	default:
		fmt.Fprintf(os.Stderr, "Không hỗ trợ cơ số %s\n", baseName)
		os.Exit(1)
		return ""
	}
}

// main ()
func main() {
	flag.Parse()

	if version {
		printVersion() // version and EXIT
	}

	// Kiểm tra nếu có file đầu vào
	if inputFile != "" {
		processInputFile()
		return
	}

	if len(flag.Args()) < 1 {
		flag.Usage() // usage and EXIT
	}

	arg := flag.Args()[0] // extract arg

	if len(inputFormat) < 1 {
		setDefaultInputFormat(arg)
	}

	// Kiểm tra xem có cần xuất ra file Excel không
	if excelFile != "" {
		exportToExcel(arg)
		return
	}

	// Xử lý hiển thị từng bước 
	if showSteps {
		showStepByStep(arg)
		return
	}

	buffer := bufio.NewWriter(os.Stdout)
	defer buffer.Flush()
	for _, o := range outputFormat {
		fn := funcMap[string(inputFormat[0])+"|"+string(o)]
		result := utils.Invoke(fn, arg)
		if !quiet {
			fmt.Fprintf(buffer, "%v: %v\n", bold(o), result)
		} else {
			fmt.Fprintf(buffer, "%v\n", result)
		}
	}
}

// processInputFile xử lý dữ liệu từ file đầu vào
func processInputFile() {
	// Đọc dữ liệu từ file
	inputs, err := stepbystep.ReadInputFromTxt(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Lỗi khi đọc file %s: %v\n", inputFile, err)
		os.Exit(1)
	}
	
	if len(inputs) == 0 {
		fmt.Fprintf(os.Stderr, "Không có dữ liệu đầu vào từ file %s\n", inputFile)
		os.Exit(1)
	}
	
	// Xử lý từng dòng dữ liệu
	var results []*stepbystep.StepByStepResult
	
	for _, input := range inputs {
		fromBase := baseNameToFormat(input.FromBase)
		toBase := baseNameToFormat(input.ToBase)
		
		if toBase == "all" {
			// Nếu đầu ra là "all", thực hiện chuyển đổi sang tất cả các cơ số khác
			possibleOutputs := []string{utils.BINARY, utils.OCTAL, utils.DECIMAL, utils.HEXADECIMAL}
			for _, outBase := range possibleOutputs {
				if outBase != fromBase && outBase != utils.ASCII {
					key := fromBase + "|" + outBase
					stepsFunc, exists := stepsFuncMap[key]
					if exists {
						result := stepsFunc(input.Input)
						results = append(results, result)
					}
				}
			}
		} else {
			// Nếu đầu ra là một cơ số cụ thể
			key := fromBase + "|" + toBase
			stepsFunc, exists := stepsFuncMap[key]
			
			if exists {
				result := stepsFunc(input.Input)
				results = append(results, result)
			} else {
				fmt.Fprintf(os.Stderr, "Không hỗ trợ chuyển đổi từ %s sang %s\n", 
					input.FromBase, input.ToBase)
			}
		}
	}
	
	// Xuất kết quả
	if len(results) > 0 {
		if excelFile != "" {
			// Xuất ra file Excel
			var err error
			if useLaTeX {
				err = stepbystep.ExportToExcelWithLaTeX(results, excelFile)
			} else {
				err = stepbystep.ExportToExcel(results, excelFile)
			}
			
			if err != nil {
				fmt.Fprintf(os.Stderr, "Lỗi khi xuất file Excel: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Đã xuất kết quả ra file Excel: %s\n", excelFile)
		} else {
			// Hiển thị trên màn hình
			for _, result := range results {
				fmt.Printf("Chuyển đổi %s (cơ số %s) sang %s (cơ số %s):\n",
					result.Input, stepbystep.FormatBaseName(result.InputBase),
					result.Output, stepbystep.FormatBaseName(result.OutputBase))
				
				for _, step := range result.Steps {
					fmt.Println(step)
				}
				fmt.Println()
			}
		}
	} else {
		fmt.Fprintln(os.Stderr, "Không có kết quả nào để xuất")
	}
}

// showStepByStep hiển thị giải pháp từng bước
func showStepByStep(arg string) {
	// Nếu định dạng đầu ra là "all", thực hiện tất cả các chuyển đổi
	if len(outputFormat) == len(utils.ALL) {
		for _, o := range outputFormat {
			if o == inputFormat[0] || o == utils.ASCII {
				continue // Bỏ qua chuyển đổi cùng định dạng hoặc ASCII
			}
			key := string(inputFormat[0]) + "|" + string(o)
			stepsFunc, exists := stepsFuncMap[key]
			if exists {
				result := stepsFunc(arg)
				fmt.Println(bold("Chuyển đổi từ " + inputFormat[0] + " sang " + o))
				for _, step := range result.Steps {
					fmt.Println(step)
				}
				fmt.Println()
			}
		}
		return
	}

	// Thực hiện chuyển đổi cụ thể
	for _, o := range outputFormat {
		key := string(inputFormat[0]) + "|" + string(o)
		stepsFunc, exists := stepsFuncMap[key]
		if exists {
			result := stepsFunc(arg)
			for _, step := range result.Steps {
				fmt.Println(step)
			}
		} else {
			fmt.Fprintf(os.Stderr, "Không có giải pháp từng bước cho chuyển đổi từ %s sang %s\n", 
				inputFormat[0], o)
		}
	}
}

// exportToExcel xuất kết quả giải pháp từng bước ra file Excel
func exportToExcel(arg string) {
	var results []*stepbystep.StepByStepResult

	// Nếu định dạng đầu ra là "all", thực hiện tất cả các chuyển đổi
	if len(outputFormat) == len(utils.ALL) {
		for _, o := range outputFormat {
			if o == inputFormat[0] || o == utils.ASCII {
				continue // Bỏ qua chuyển đổi cùng định dạng hoặc ASCII
			}
			key := string(inputFormat[0]) + "|" + string(o)
			stepsFunc, exists := stepsFuncMap[key]
			if exists {
				result := stepsFunc(arg)
				results = append(results, result)
			}
		}
	} else {
		// Thực hiện chuyển đổi cụ thể
		for _, o := range outputFormat {
			key := string(inputFormat[0]) + "|" + string(o)
			stepsFunc, exists := stepsFuncMap[key]
			if exists {
				result := stepsFunc(arg)
				results = append(results, result)
			}
		}
	}
	
	// Xuất ra file Excel
	if len(results) > 0 {
		var err error
		if useLaTeX {
			err = stepbystep.ExportToExcelWithLaTeX(results, excelFile)
		} else {
			err = stepbystep.ExportToExcel(results, excelFile)
		}
		
		if err != nil {
			fmt.Fprintf(os.Stderr, "Lỗi khi xuất file Excel: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Đã xuất giải pháp ra file Excel: %s\n", excelFile)
	} else {
		fmt.Fprintf(os.Stderr, "Không có giải pháp nào để xuất ra Excel\n")
	}
}
