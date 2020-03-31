import pdftotext
import glob
import os
import json
import argparse
import re
from itertools import chain


class JsonObject:
    def toJSON(self):
        return json.dumps(self, default=lambda o: o.__dict__, sort_keys=False, indent=4)


js = JsonObject()

def process_district(district_data):
    next = []
    continuation = False
    data = []
    for row in district_data:
        t = re.findall(r'(?:\s+)([a-zA-Z]+)?\s+(\d+)?\s*(\w.*)?', row)
        if t[0][0] == '':
            next.extend(t[0][1])
            next.extend(list(chain(*[tmp.split() for tmp in t[0][2].split(',')])))
            if next[-1].isnumeric():
                continuation = True
        else:
            if next != [] and continuation:
                next.append(t[0][2])
                x = list(t[0])
                x.pop()
                x.extend(next)
                data.append(x)
                continuation = False
                next = []
                continue
            if t[0][2] == '':
                data.append([t[0][0],t[0][1]])
            else:
                next.extend(list(chain(*[tmp.split() for tmp in t[0][2].split(',')])))
                x = [t[0][0],t[0][1]]
                x.extend(next)
                data.append(x)
                next=[]
    for cols in data:
        district = cols[0].lower()
        getattr(js, district)["no_of_positive_cases_admitted"] = int(cols[1])
        getattr(js, district)["other_districts"] = {}
        for i in range(2,2+len(cols[2:]),2):
            getattr(js, district)["other_districts"][cols[i+1].lower()] = int(cols[i])


def process_annex1(annexure_1_data):
    for row in annexure_1_data:
        cols = row.split()
        district = cols[0].lower()
        setattr(js, district, {})
        for i, item in enumerate(cols[1:]):
            if i == 1:
                getattr(js, district)[
                    "no_of_persons_under_observation_as_on_today"] = int(item)
            elif i == 2:
                getattr(js, district)[
                    "no_of_persons_under_home_isolation_as_on_today"] = int(item)
            elif i == 3:
                getattr(js, district)[
                    "no_of_symptomatic_persons_hospitalized_as_on_today"] = int(item)
            else:
                getattr(js, district)[
                    "no_of_persons_hospitalized_today"] = int(item)


def extract_text_data(latest_pdf):
    # Load the dhs data pdf
    with open(latest_pdf, "rb") as f:
        pdf = pdftotext.PDF(f)
    # Iterate over only the required pages
    data = []
    for page_num in [2, 3, 4]:
        lines = ""
        for char in pdf[page_num]:
            if char == '\n':
                data.append(lines)
                lines = ""
            else:
                lines += char
    # pure regex witchcraftery - currently captures annex1 and district wise tables
    annex1 = re.findall(r'Annexure -1.*on today.(.*?)(Total.*?\n)', "\n".join(data),re.DOTALL)
    district = re.findall(r'District wise.*District..(.*?)(Total.*?\n)', "\n".join(data),re.DOTALL)
    return "".join(annex1[0]).split("\n")[:-1], "".join(district[0]).split("\n")[:-1]


def init():
    parser = argparse.ArgumentParser()
    parser.add_argument("-t", "--text", action="store_true",
                        help="display only the extracted text and exit")
    parser.add_argument("-v", "--verbose", action="store_true",
                        help="show the details about the file")
    return parser.parse_args()


if __name__ == "__main__":
    args = init()
    latest_pdf = max(glob.iglob("../data/*.pdf"), key=os.path.getctime)
    annex1_data, district_data = extract_text_data(latest_pdf)
    if args.verbose:
        print(f"filename : {latest_pdf} with length : {len(text_data)}")
    if args.text:
        print("\n".join(annex1_data))
        print("\n".join(district_data))
        exit(0)
    process_annex1(annex1_data)
    process_district(district_data)
    # print(js.toJSON())
