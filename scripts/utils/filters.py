import re


def extract_date(data):
    for line in data.split('\n'):
        if line.strip().startswith("Date") or line.strip().startswith("DATE"):
            return re.findall(r'(\d+)/(\d+)/(\d+)', line.strip())
    print("Wasn't able to extract Date from pdf")
    return ""


def extract_annex1(data):
    table = []
    flag = False
    for line in data.split('\n'):
        line = line.strip()
        if 'on today ' in line:
            flag = True
            continue
        if flag and re.match(
                r'^Thir|^Koll|^Path|^Iduk|^Kott|^Alap|^Erna|^Thri|^Pala|^Mala|^Kozh|^Waya|^Kann|^Kasa',
                line):
            table.append(line)
        elif flag and re.match(r'^[T|t]otal', line):
            table.append(line)
            return table
    return []


def extract_district(data):
    table = []
    captured = re.findall(
        r'District(?:\s+)(?:[N|n]o\..*?[\n\r])(.*Total.*?\n)', data, re.DOTALL)
    for line in captured[0].split('\n'):
        line = line.strip()
        if re.match(
                r'^Thir|^Koll|^Path|^Iduk|^Kott|^Alap|^Erna|^Thri|^Pala|^Mala|^Kozh|^Waya|^Kann|^Kasa|^\d+\s',
                line):
            table.append(line)
        elif re.match(r'^[T|t]otal', line):
            table.append(line)
            return table
    return table
