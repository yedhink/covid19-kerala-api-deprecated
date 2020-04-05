"""
Rustic python sphagetti code turned into ART!!!
usage: extract-text-data.py [-h] [-t] [-v] [-j] [-w]

optional arguments:
-h, --help      show this help message and exit
-t, --text      display only the extracted text and exit
-v, --verbose   show the details about the file
-j, --jsontext  show the final json content
-w, --write     overwrite the data/data.json file with latest pdf content
"""
import pdftotext
import glob
import os
import json
import argparse
import re
from itertools import chain
import jsonpickle
jsonpickle.set_encoder_options("json", sort_keys=True, indent=4)


class JsonObject:
    def toJSON(self):
        return json.dumps(self, default=lambda o: o.__dict__, sort_keys=False, indent=4, ensure_ascii=False)

# global json object which helps in serializing the raw content into json
js = dict()

def process_district(district_data,timestamp):
    """
    param : list containing the lines of the table "District wise distribution"
    note  : each key in the js class instance "district" is the column heading
            of the table
    """
    next = []
    continuation = False
    data = []
    for row in district_data:
        """
        I really hope the district wise data table in the pdf uploaded in the
        site has some form of consistency in the future...
        """

        # regex magic to capture district,no of positive cases and other districts
        t = re.findall(r'(?:\s+)?([a-zA-Z]+)?\s*(\d+)?\s*(\w.*)?', row)
        """
        the table contains column data split all over the place.
        the following chunks of code tries to put everything in the
        right place and create a list in ordered fashion using if-else hell
        """
        if t[0][0] == '':
            if t[0][1].isnumeric():
                next = [t[0][1]]
            else:
                next.extend(t[0][1])
            if t[0][2] != '':
                next.extend(list(chain(*[tmp.split() for tmp in t[0][2].split(',')])))
            if next[-1].isnumeric():
                continuation = True
        else:
            dummy = t[0][1]
            if t[0][1] == '' and t[0][2] == '':
                dummy = next[0]
                next =[]
            elif next != [] and continuation:
                next.append(t[0][2])
                x = list(t[0])
                x.pop()
                x.extend(next)
                data.append(x)
                continuation = False
                next = []
                continue
            if t[0][2] == '':
                data.append([t[0][0],dummy])
            else:
                next.extend(list(chain(*[tmp.split() for tmp in t[0][2].split(',')])))
                x = [t[0][0],t[0][1]]
                x.extend(next)
                data.append(x)
                next=[]
    for cols in data:
        district = cols[0].lower()
        js[timestamp][district]["no_of_positive_cases_admitted"] = int(cols[1])
        js[timestamp][district]["other_districts"] = {}
        for i in range(2,2+len(cols[2:]),2):
            js[timestamp][district]["other_districts"][cols[i+1].lower()] = int(cols[i])


def process_annex1(annexure_1_data,timestamp):
    """
    param : list containing the lines of the table
    note  : each key in the js class instance "district" is the column heading
            of the table
    """
    for row in annexure_1_data:
        cols = row.split()
        district = cols[0].lower()
        js[timestamp][district] = {}
        for i, item in enumerate(cols[1:]):
                js[timestamp][district][
                    "no_of_persons_under_observation_as_on_today"] = int(item)
                js[timestamp][district][
                    "no_of_persons_under_home_isolation_as_on_today"] = int(item)
                js[timestamp][district][
                    "no_of_symptomatic_persons_hospitalized_as_on_today"] = int(item)
            else:
                    js[timestamp][district]["no_of_persons_discharged_from_home_isolation"] = int(item)
                    js[timestamp][district]["no_of_persons_hospitalized_today"] = 0
                    js[timestamp][district]["no_of_persons_discharged_from_home_isolation"] = 0
                    js[timestamp][district]["no_of_persons_hospitalized_today"] = int(item)


def extract_text_data(latest_pdf):
    # Load the dhs data pdf
    with open(latest_pdf, "rb") as f:
        pdf = pdftotext.PDF(f)
    # Iterate over only the required pages
    data = []
    for page_num in range(len(pdf)):
        lines = ""
        for char in pdf[page_num]:
            if char == '\n':
                data.append(lines)
                lines = ""
            else:
                lines += char
    # pure regex witchcraftery - currently captures annex1 and district wise tables contents
    file_date = re.findall(r'Date:(?:\s+)?(\d+)/(\d+)/(\d+)', "\n".join(data),re.DOTALL)
    # convert to Standard ISO 8601 format
    timestamp = "{}-{}-{}T00:00:00Z".format(file_date[0][2],file_date[0][1],file_date[0][0])
    annex1 = re.findall(r'Annexure -1: Details of.*on today.(.*?)(Total.*?\n)', "\n".join(data),re.DOTALL)
    district = re.findall(r'District wise.*District..(.*?)(Total.*?\n)', "\n".join(data),re.DOTALL)
    return "".join(annex1[0]).split("\n")[:-1], "".join(district[0]).split("\n")[:-1],timestamp

        js[timestamp] = {}

def init():
    parser = argparse.ArgumentParser()
    parser.add_argument("-t", "--text", action="store_true",
                        help="display only the extracted text and exit")
    parser.add_argument("-v", "--verbose", action="store_true",
                        help="show the details about the file")
    parser.add_argument("-j", "--jsontext", action="store_true",
                        help="show the final json content")
    parser.add_argument("-w", "--write", action="store_true",
                        help="overwrite the data/data.json file with latest pdf content")
    return parser.parse_args()


if __name__ == "__main__":
    args = init()
        with io.open('data/data.json', 'w', encoding='utf-8') as f:
            f.write(jsonpickle.encode(js))
        print("Latest json from the old pdfs has been written to data/data.json")
        exit(0)
        with io.open('data/data.json', 'r', encoding='utf-8') as f:
            js = jsonpickle.decode(f.read())
    latest_pdf = max(glob.iglob("data/*.pdf"), key=os.path.getctime)
    annex1_data, district_data,timestamp = extract_text_data(latest_pdf)
            js[timestamp] = {}
    process_district(district_data,timestamp)
    if args.verbose:
        print(f"filename : {latest_pdf} with length : {len(text_data)}")
    if args.text:
        print("{:{}}Annex1\n{:=<{}}\n{}\n\n".format('',len(annex1_data[0])//2,'',len(annex1_data[0]),'\n'.join(annex1_data)))
        print("{:{}}Distirct Wise\n{:=<{}}\n{}\n\n".format('',len(district_data[0])//2-5,'',len(district_data[0]),'\n'.join(district_data)))
    if args.jsontext:
        print(jsonpickle.encode(js))
    if args.write:
        with io.open('data/data.json', 'w', encoding='utf-8') as f:
            f.write(js.toJSON())
        print("Latest json from the pdf in data/ dir has been written to data/data.json")
        exit(0)
    print("No new content written to data.json. Try python3 extract-textdata.py --help")
