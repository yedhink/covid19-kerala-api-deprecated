import pdftotext
import glob
import os
import json

class JsonObject:
    def toJSON(self):
        return json.dumps(self, default=lambda o: o.__dict__,
            sort_keys=True, indent=4)

js = JsonObject()

def process_annex1(data):
    annexure_1 = data[16:31]
    for row in annexure_1:
        cols = row.split()
        district = cols[0].lower()
        setattr(js,district,{})
        for i,item in enumerate(cols[1:]):
            if i == 1:
                getattr(js,district)["no_of_persons_under_observation_as_on_today"]=item
            elif i == 2:
                getattr(js,district)["no_of_persons_under_home_isolation_as_on_today"]=item
            elif i == 3:
                getattr(js,district)["no_of_symptomatic_persons_hospitalized_as_on_today"]=item
            else:
                getattr(js,district)["no_of_persons_hospitalized_today"]=item

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
    process_annex1(text_data)
    print(js.toJSON())