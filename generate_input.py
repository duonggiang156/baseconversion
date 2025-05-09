#!/usr/bin/env python3
"""
Script để tạo ngẫu nhiên các câu hỏi chuyển đổi cơ số.
Sử dụng: python generate_input.py [số lượng câu hỏi] [kiểu đầu vào] [kiểu đầu ra] [tên file đầu ra]
Các kiểu hợp lệ: binary, decimal, octal, hexadecimal, all
Ví dụ: python generate_input.py 10 hexadecimal binary output.txt
Mặc định: 10 câu hỏi, kiểu ngẫu nhiên, file output là input_v1.txt
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

def generate_conversion_problem(input_base=None, output_base=None):
    """Tạo một bài toán chuyển đổi với cơ số đầu vào và đầu ra được chỉ định"""
    valid_bases = ['binary', 'decimal', 'octal', 'hexadecimal']
    
    # Nếu không chỉ định cơ số đầu vào, chọn ngẫu nhiên
    if input_base is None or input_base == 'all':
        input_base = random.choice(valid_bases)
    elif input_base not in valid_bases:
        print(f"Lỗi: Cơ số đầu vào '{input_base}' không hợp lệ. Sử dụng giá trị mặc định.")
        input_base = random.choice(valid_bases)
    
    # Nếu không chỉ định cơ số đầu ra hoặc chọn 'all', xử lý tương ứng
    if output_base is None:
        output_bases = valid_bases.copy()
        output_bases.remove(input_base)  # Loại bỏ cơ số đầu vào ra khỏi danh sách đầu ra
        # Có thể chọn "all" với xác suất 20%
        if random.random() < 0.2:
            output_base = 'all'
        else:
            output_base = random.choice(output_bases)
    elif output_base not in valid_bases and output_base != 'all':
        print(f"Lỗi: Cơ số đầu ra '{output_base}' không hợp lệ. Sử dụng giá trị mặc định.")
        output_bases = valid_bases.copy()
        output_bases.remove(input_base)
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

def generate_unique_problems(num_problems, input_base=None, output_base=None, max_attempts=1000):
    """Tạo số lượng bài toán không trùng lặp theo yêu cầu"""
    unique_problems = set()
    attempts = 0
    
    while len(unique_problems) < num_problems and attempts < max_attempts:
        problem = generate_conversion_problem(input_base, output_base)
        unique_problems.add(problem)
        attempts += 1
        
        # Nếu không thể tạo đủ số lượng bài toán không trùng lặp sau nhiều lần thử
        if attempts >= max_attempts and len(unique_problems) < num_problems:
            print(f"Cảnh báo: Chỉ có thể tạo {len(unique_problems)} bài toán không trùng lặp sau {max_attempts} lần thử.")
            break
    
    return list(unique_problems)

def main():
    # Các giá trị mặc định
    num_problems = 10
    input_base = None  # Sẽ chọn ngẫu nhiên
    output_base = None  # Sẽ chọn ngẫu nhiên
    output_file = "2to16.txt"
    
    # python generate_input.py 256 hexadecimal binary
    # python generate_input.py 256 hexadecimal octal
    # python generate_input.py 256 hexadecimal decimal
    # python generate_input.py 256 decimal binary
    # python generate_input.py 256 decimal octal
    # python generate_input.py 256 decimal hexadecimal
    # python generate_input.py 256 octal binary
    # python generate_input.py 256 octal decimal 
    # python generate_input.py 256 octal hexadecimal
    # python generate_input.py 256 binary octal
    # python generate_input.py 256 binary decimal
    # python generate_input.py 256 binary hexadecimal
    
    # Xử lý tham số dòng lệnh
    if len(sys.argv) > 1:
        try:
            num_problems = int(sys.argv[1])
        except ValueError:
            print(f"Lỗi: '{sys.argv[1]}' không phải là số nguyên. Sử dụng giá trị mặc định 10.")
            num_problems = 10
    
    if len(sys.argv) > 2:
        input_base = sys.argv[2].lower()
    
    if len(sys.argv) > 3:
        output_base = sys.argv[3].lower()
    
    if len(sys.argv) > 4:
        output_file = sys.argv[4]
    
    # Tạo các bài toán không trùng lặp
    problems = generate_unique_problems(num_problems, input_base, output_base)
    
    # Ghi ra file
    with open(output_file, 'w') as f:
        for problem in problems:
            f.write(problem + '\n')
    
    # Hiển thị thông tin về đầu vào/đầu ra
    input_type = input_base if input_base else "ngẫu nhiên"
    output_type = output_base if output_base else "ngẫu nhiên"
    
    print(f"Đã tạo {len(problems)} bài toán chuyển đổi cơ số từ {input_type} sang {output_type}")
    print(f"Kết quả đã được lưu vào file '{output_file}'")
    print(f"Để xử lý file và xuất ra Excel với định dạng LaTeX, chạy:")
    print(f"  .\\ncalc.exe -f {output_file} -e ketqua.xlsx -l")

    # Hiển thị hướng dẫn sử dụng nếu không có tham số
    if len(sys.argv) == 1:
        print("\nHướng dẫn sử dụng:")
        print("python generate_input.py [số lượng câu hỏi] [kiểu đầu vào] [kiểu đầu ra] [tên file đầu ra]")
        print("Các kiểu hợp lệ: binary, decimal, octal, hexadecimal, all")
        print("Ví dụ: python generate_input.py 10 hexadecimal binary output.txt")

if __name__ == "__main__":
    main() 
    
