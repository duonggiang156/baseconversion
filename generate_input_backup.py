#!/usr/bin/env python3
"""
Script để tạo ngẫu nhiên các câu hỏi chuyển đổi cơ số.
Sử dụng: python generate_input.py [số lượng câu hỏi] [tên file đầu ra]
Mặc định: 10 câu hỏi, file output là input.txt
"""

import random
import sys

def generate_binary_number(min_len=3, max_len=8):
    """Tạo ngẫu nhiên một số nhị phân với độ dài từ min_len đến max_len"""
    length = random.randint(min_len, max_len)
    # Đảm bảo bit đầu tiên là 1 để không có số 0 ở đầu
    binary = '1' + ''.join(random.choice('01') for _ in range(length-1))
    return binary

def generate_decimal_number(min_val=1, max_val=500):
    """Tạo ngẫu nhiên một số thập phân từ min_val đến max_val"""
    return str(random.randint(min_val, max_val))

def generate_octal_number(min_len=2, max_len=4):
    """Tạo ngẫu nhiên một số bát phân với độ dài từ min_len đến max_len"""
    length = random.randint(min_len, max_len)
    # Đảm bảo chữ số đầu tiên không phải là 0
    octal = random.choice('1234567') + ''.join(random.choice('01234567') for _ in range(length-1))
    return octal

def generate_hexadecimal_number(min_len=2, max_len=3):
    """Tạo ngẫu nhiên một số thập lục phân với độ dài từ min_len đến max_len"""
    length = random.randint(min_len, max_len)
    # Đảm bảo chữ số đầu tiên không phải là 0
    hex_digits = '123456789ABCDEF'
    hexadecimal = random.choice(hex_digits) + ''.join(random.choice('0123456789ABCDEF') for _ in range(length-1))
    return hexadecimal

def generate_conversion_problem():
    """Tạo một bài toán chuyển đổi ngẫu nhiên"""
    # Chọn ngẫu nhiên cơ số đầu vào
    input_base = random.choice(['binary', 'decimal', 'octal', 'hexadecimal'])
    
    # Chọn ngẫu nhiên cơ số đầu ra
    output_bases = ['binary', 'decimal', 'octal', 'hexadecimal']
    output_bases.remove(input_base)  # Loại bỏ cơ số đầu vào ra khỏi danh sách đầu ra
    
    # Có thể chọn "all" với xác suất 20%
    if random.random() < 0.2:
        output_base = 'all'
    else:
        output_base = random.choice(output_bases)
    
    # Tạo số ngẫu nhiên dựa vào cơ số đầu vào
    if input_base == 'binary':
        number = generate_binary_number()
    elif input_base == 'decimal':
        number = generate_decimal_number()
    elif input_base == 'octal':
        number = generate_octal_number()
    else:  # hexadecimal
        number = generate_hexadecimal_number()
    
    return f"{number} {input_base} {output_base}"

def main():
    # Lấy tham số từ dòng lệnh
    num_problems = 10  # mặc định
    output_file = "input_v2.txt"  # mặc định
    
    if len(sys.argv) > 1:
        try:
            num_problems = int(sys.argv[1])
        except ValueError:
            print(f"Lỗi: '{sys.argv[1]}' không phải là số nguyên.")
            sys.exit(1)
    
    if len(sys.argv) > 2:
        output_file = sys.argv[2]
    
    # Tạo các bài toán
    problems = [generate_conversion_problem() for _ in range(num_problems)]
    
    # Ghi ra file
    with open(output_file, 'w') as f:
        for problem in problems:
            f.write(problem + '\n')
    
    print(f"Đã tạo {num_problems} bài toán chuyển đổi cơ số và lưu vào file '{output_file}'")
    print(f"Để xử lý file và xuất ra Excel với định dạng LaTeX, chạy:")
    print(f"  .\\ncalc.exe -f {output_file} -e ketqua.xlsx -l")

if __name__ == "__main__":
    main() 