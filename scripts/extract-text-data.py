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
import argparse
import re
from itertools import chain
import io
from datetime import datetime
from utils import filters
import jsonpickle
jsonpickle.set_encoder_options("json", sort_keys=True, indent=4)

# global json object which helps in serializing the raw content into json
js = dict()
all_districts = [
    "thiruvananthapuram", "kollam", "pathanamthitta", "idukki", "kottayam",
    "alappuzha", "ernakulam", "thrissur", "palakkad", "malappuram",
    "kozhikode", "wayanad", "kannur", "kasaragod"
]


def split_joint_digit_and_dist(x):
    # for removinng inconsistency like 18Kasragod needs splitting into [18,Kasaragod]
    # this inconsistency was first noted in 06-04-2020.pdf
    match = re.findall(r'(\d+)(?:\s*)?(\w+)', x)
    # print(f"splitted digit and dist : {match}{x}")
    return match[0][0].strip(), match[0][1].strip()


def process_district(district_data, timestamp):
    """
    params
    ------
    lines of the table "District wise distribution" - list
    timestamp - string
    note  : each key in the js class instance "district" is the column heading
            of the table
    I really hope the district wise data table in the pdf uploaded in the
    site has some form of consistency in the future...
    """
    next = []
    continuation = False
    data = []
    for row in district_data:
        # regex magic to capture district,no of positive cases and other districts
        t = re.findall(r'(?:\s+)?([a-zA-Z]+)?\s*(\d+)?\s*(\w.*)?', row)
        # print(t)
        """
        the table contains column data split all over the place.
        the following chunks of code tries to put everything in the
        right place and create a list in ordered fashion using if-else hell
        for eg, sometimes the pdf table extracted looks like this:-
                                            1malappuram, 1
            Thriuvananthapuram      8       Kollam
        if it was a single line and if there was a whitespace btw '1malappuram1',
        then how much easier life wouldve been.
        then sometimes the data very inconsistent like this:-
                                    114     1 kannur, 2 kozhikode
        Kasaragod
        it shouldve been "Kasaragod  114  1 kannur, 2 kozhikode
        """
        if t[0][0] == '':
            if t[0][1].isnumeric():
                next = [t[0][1]]
            else:
                next.extend(t[0][1])
            if t[0][2] != '':
                next.extend(
                    list(
                        chain(*[
                            split_joint_digit_and_dist(tmp)
                            for tmp in t[0][2].split(',')
                        ])))
            if next[-1].isnumeric():
                continuation = True
        else:
            dummy = t[0][1]
            if t[0][1] == '' and t[0][2] == '':
                dummy = next[0]
                next = []
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
                data.append([t[0][0], dummy])
            else:
                next.extend(
                    list(
                        chain(*[
                            split_joint_digit_and_dist(tmp)
                            for tmp in t[0][2].split(',')
                        ])))
                x = [t[0][0], t[0][1]]
                x.extend(next)
                data.append(x)
                next = []
    flag = all_districts[:]
    for cols in data:
        district = cols[0].lower()
        if district != "total":
            flag.remove(district)
        js[timestamp][district]["no_of_positive_cases_admitted"] = int(cols[1])
        js[timestamp][district]["other_districts"] = {}
        for i in range(2, 2 + len(cols[2:]), 2):
            js[timestamp][district]["other_districts"][cols[i +
                                                            1].lower()] = int(
                                                                cols[i])
    # not all district tables have full districts mentioned
    # therefore provide default values to unmentioned
    for i in flag:
        js[timestamp][i]["no_of_positive_cases_admitted"] = 0
        js[timestamp][i]["other_districts"] = {}


def process_annex1(annexure_1_data, timestamp, other):
    """
    params
    ------
    lines of the annex-1 table - list
    timestamp of the file - string
    other denotes whether default value for district table is required in json - bool
    note  : each key in the js object instance "district" is the column heading
            of the table
    """
    file_date = datetime.strptime(timestamp, '%Y-%m-%dT%H:%M:%SZ').date()
    with_discharged_from_home_date = datetime.strptime(
        "2020-03-08T00:00:00Z", '%Y-%m-%dT%H:%M:%SZ').date()
    # some of the old pdf datasets(<=08-03-2020) has annex-1 table with last column as persons_discharged_from_home_isolation
    persons_discharged_from_home_isolation = False
    if file_date <= with_discharged_from_home_date:
        persons_discharged_from_home_isolation = True
    for row in annexure_1_data:
        cols = row.split()
        district = cols[0].lower()
        # probably an extra word from regex match. but needs to be avoided
        if district == "today":
            continue
        js[timestamp][district] = {}
        for i, item in enumerate(cols[1:]):
            if i == 0:
                js[timestamp][district][
                    "no_of_persons_under_observation_as_on_today"] = int(item)
            elif i == 1:
                js[timestamp][district][
                    "no_of_persons_under_home_isolation_as_on_today"] = int(
                        item)
            elif i == 2:
                js[timestamp][district][
                    "no_of_symptomatic_persons_hospitalized_as_on_today"] = int(
                        item)
                # even if no district table exist, we provide default values
                if other:
                    js[timestamp][district][
                        "no_of_positive_cases_admitted"] = 0
                    js[timestamp][district]["other_districts"] = {}
            else:
                if persons_discharged_from_home_isolation:
                    js[timestamp][district][
                        "no_of_persons_discharged_from_home_isolation"] = int(
                            item)
                    js[timestamp][district][
                        "no_of_persons_hospitalized_today"] = 0
                else:
                    js[timestamp][district][
                        "no_of_persons_discharged_from_home_isolation"] = 0
                    js[timestamp][district][
                        "no_of_persons_hospitalized_today"] = int(item)
    return persons_discharged_from_home_isolation


def extract_text_data(latest_pdf):
    """
    params
    ------
    latest-pdf file - string
    it's very important to note that some of the pdf files
    are not convertable to text using pdftotext library
    """
    # Load the dhs data pdf
    with open(latest_pdf, "rb") as f:
        pdf = pdftotext.PDF(f)
    data = []
    # Iterate over only the required pages by providing range
    for page_num in range(len(pdf)):
        lines = ""
        for char in pdf[page_num]:
            if char == '\n':
                data.append(lines)
                lines = ""
            else:
                lines += char
    data = "\n".join(data)
    file_date = filters.extract_date(data)
    annex1_table = filters.extract_annex1(data)
    district_table = filters.extract_district(data)
    try:
        # convert to Standard ISO 8601 format
        timestamp = "{}-{}-{}T00:00:00Z".format(file_date[0][2],
                                                file_date[0][1],
                                                file_date[0][0])
        # print(f"current file : {latest_pdf}")
    except IndexError:
        # if the pdf file cant be read as text
        return "", "", ""
    if len(district_table) == 0:
        district_table = ""
    return annex1_table, district_table, timestamp


def process_old_data(folder):
    """
    params
    -----
    folder which contains all the old pdfs - string
    """
    for f in glob.iglob(f"{folder}/*.pdf"):
        d = False
        annex1_data, district_data, timestamp = extract_text_data(f)
        if annex1_data == "" and district_data == "" and timestamp == "":
            continue
        js[timestamp] = {}
        # condition is true only if a district wise table exists
        if district_data == "":
            d = True
        if not process_annex1(annex1_data, timestamp, d):
            process_district(district_data, timestamp)


def init():
    parser = argparse.ArgumentParser()
    parser.add_argument(
        "-t",
        "--text",
        action="store_true",
        help=
        "display only the extracted text from both annex-1 and district wise tables"
    )
    parser.add_argument("-v",
                        "--verbose",
                        action="store_true",
                        help="show the details about the file")
    parser.add_argument("-j",
                        "--jsontext",
                        action="store_true",
                        help="show the final json content")
    parser.add_argument(
        "-w",
        "--write",
        action="store_true",
        help="overwrite the data/data.json file with latest pdf content")
    parser.add_argument(
        "-o",
        "--old",
        action='store',
        default=False,
        dest='path_to_old_data_dir',
        nargs=1,
        type=str,
        help="parse the old pdf datasets - provide the directory as arg")
    return parser.parse_args()


if __name__ == "__main__":
    args = init()
    if args.path_to_old_data_dir:
        """
        this is just a one time thing. we need to get the data from dataset starting
        back from 01-31-2020 and dump it into our data.json.
        """
        process_old_data(args.path_to_old_data_dir[0])
        with io.open('data/data.json', 'w', encoding='utf-8') as f:
            f.write(jsonpickle.encode(js))
        print(
            "Latest json from the old pdfs has been written to data/data.json")
        exit(0)
    else:
        """
        parse the latest data pdf by intially creating a json obj
        from existing json data
        """
        with io.open('data/data.json', 'r', encoding='utf-8') as f:
            js = jsonpickle.decode(f.read())
        latest_pdf = max(glob.iglob("data/*.pdf"), key=os.path.getctime)
        annex1_data, district_data, timestamp = extract_text_data(latest_pdf)
        try:
            x = js[timestamp]
            if x is not None:
                print(
                    "the file {} has already been extracted and added to data/data.json"
                )
                exit(0)
        except KeyError:
            js[timestamp] = {}
        if not process_annex1(annex1_data, timestamp, False):
            process_district(district_data, timestamp)
    if args.verbose:
        print(f"filename : {latest_pdf} with length : {len(text_data)}")
    if args.text:
        # print just the tables
        print("{:{}}Annex1\n{:=<{}}\n{}\n\n".format('',
                                                    len(annex1_data[0]) // 2,
                                                    '', len(annex1_data[0]),
                                                    '\n'.join(annex1_data)))
        print("{:{}}Distirct Wise\n{:=<{}}\n{}\n\n".format(
            '',
            len(district_data[0]) // 2 - 5, '', len(district_data[0]),
            '\n'.join(district_data)))
    if args.jsontext:
        # print the json output without writing to file
        print(jsonpickle.encode(js))
    if args.write:
        file = "data/data.json"
        with io.open(file, 'w', encoding='utf-8') as f:
            f.write(jsonpickle.encode(js))
        print(
            "Python Script : Latest json from the pdf in data/ dir has been written to "
            + file)
        exit(0)

    print(
        "No new content written to data.json. Try python3 extract-textdata.py --help"
    )
