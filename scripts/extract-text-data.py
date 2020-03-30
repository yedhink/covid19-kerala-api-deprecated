import pdftotext
import glob
import os

def process_data(data):
    for index,line in enumerate(data):
        print(line)

def extract_text_data(latest_pdf):
    # Load the dhs data pdf
    with open(latest_pdf, "rb") as f:
        pdf = pdftotext.PDF(f)

    # Iterate over only the required pages
    data = []
    for page_num in [2,3,4]:
        lines = ""
        for char in pdf[page_num]:
            if char == '\n':
                data.append(lines)
                lines = ""
            else:
                lines += char
    return data

if __name__ == "__main__":
    latest_pdf = max(glob.iglob("../data/*.pdf"),key=os.path.getctime)
    text_data = extract_text_data(latest_pdf)
    process_data(text_data)
