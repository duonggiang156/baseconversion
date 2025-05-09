# [ncalc](https://godoc.org/github.com/clarketm/ncalc)

Command line utility for *quick* number base conversions ( **ascii** / **binary** / **octal** / **decimal** / **hexadecimal** ).

[![release-badge](https://img.shields.io/github/release/clarketm/ncalc.svg)](https://github.com/clarketm/ncalc/releases)
[![circleci-badge](https://circleci.com/gh/clarketm/ncalc.svg?style=shield)](https://circleci.com/gh/clarketm/ncalc)

```shell

Bước 1: Khởi tạo Go module: cd vào ncalc
        go mod init github.com/clarketm/ncalc
        (LẦN ĐẦU TIỀN CHẠY THÌ MỚI CHẠY BƯỚC NÀY)
Bước 2: Tải dependencies: go mod tidy
Bước 3: Build dự án: go build -o ncalc.exe
Bước 4: Chạy chương trình: .\ncalc.exe 
                            .\ncalc.exe -f input.txt -e ketqua.xlsx -l

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
    -l, --latex                 use LaTeX formatting in excel output
    -f, --file filename         read input from text file
    -e, --excel filename        export step-by-step solution to excel file
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
    ncalc -f "input.txt" -e "ketqua.xlsx" -l # read from text file, export to excel with LaTeX

```
## Installation

#### Golang
```shell
$ go get -u github.com/clarketm/ncalc
```

#### Install script
```shell
$ git clone "https://github.com/clarketm/ncalc.git"
$ cd ncalc && sudo sh install.sh
```

#### Source (Mac/Linux)
```shell
# List of builds: https://github.com/clarketm/ncalc/releases/

$ BUILD=darwin_amd64.tar.gz     # Mac (64 bit)
# BUILD=linux_amd64.tar.gz      # Linux (64 bit)

$ BIN_DIR=/usr/local/bin        # `bin` install directory
$ mkdir -p $BIN_DIR

$ curl -L https://github.com/clarketm/ncalc/releases/download/v1.1.2/$BUILD | tar xz -C $BIN_DIR        # install
```

#### Source (Windows)
* https://github.com/clarketm/ncalc/releases/download/v1.1.2/windows_amd64.zip


## Usage

#### Convert `ascii` to `decimal`
```bash
# Short form
$ ncalc -i a -o d 'w'

decimal: 119


# Long form
$ ncalc -i ascii -o decimal 'w'

decimal: 119


# Very long form
$ ncalc --input ascii --output decimal 'w'

decimal: 119


# Quite mode (-q|--quiet)
$ ncalc -q -i ascii --output decimal 'w'

119
```

#### Convert `decimal` to `binary`
```shell
$ ncalc -i decimal -o binary 170

binary: 10101010
```

#### Convert `binary` to `decimal`
```shell
$ ncalc -i b -o d 10101010

decimal: 170
```

#### Convert `ascii` to `all` formats
```shell
$ ncalc -i a 'G'

ascii: 'G'
binary: 1000111
octal: 107
decimal: 71
hexadecimal: 47
```

---

You can see the full reference documentation for the **ncalc** package at [godoc.org](https://godoc.org/github.com/clarketm/ncalc), or through go's standard documentation system:
```bash
$ godoc -http=:6060

# Open browser to: "http://localhost:6060/pkg/github.com/clarketm/ncalc"  to view godoc.
```




## Tính năng mới
    Bước 2: Tải dependencies: go mod tidy
    Bước 3: Build dự án: go build -o ncalc.exe

cách chạy 2: go run main.go -f  input_full.txt -e result_final_v5.xlsx -l
             go run main.go -f  16to10.txt -e 16to10.xlsx -l

### Xuất ra file Excel với giải pháp từng bước
```shell
$ ncalc -i d -o b -e "ketqua.xlsx" 42   # Xuất giải pháp chuyển đổi từ decimal sang binary 
$ ncalc -i d -o all -e "ketqua.xlsx" 42 # Xuất giải pháp chuyển đổi từ decimal sang tất cả các hệ số khác
```

### Định dạng LaTeX trong file Excel
```shell
$ ncalc -i d -o b -e "ketqua.xlsx" -l 42    # Xuất giải pháp với định dạng LaTeX
$ ncalc -f "input.txt" -e "ketqua.xlsx" -l  # Đọc từ file và xuất ra Excel với định dạng LaTeX
```

### Đọc đầu vào từ file
File input.txt cần có định dạng mỗi dòng một bài toán chuyển đổi, mỗi dòng gồm 3 trường:
```
<số> <cơ số đầu> <cơ số đích>
```

Ví dụ:
```
1010 binary decimal
42 decimal binary
255 decimal hexadecimal
FF hexadecimal decimal
```

### Tạo ngẫu nhiên các bài toán chuyển đổi

Sử dụng script Python để tạo ngẫu nhiên các bài toán chuyển đổi:
```shell
$ python generate_input.py          # Tạo 10 bài toán và lưu vào input.txt (mặc định)
$ python generate_input.py 20       # Tạo 20 bài toán và lưu vào input.txt
$ python generate_input.py 50 data.txt  # Tạo 50 bài toán và lưu vào data.txt
```

Sau đó, chạy ncalc với file đầu vào đã tạo:
```shell
$ ncalc -f "input.txt" -e "ketqua.xlsx" -l  # Đọc từ file và xuất ra Excel với định dạng LaTeX
```
