<div style="display:flex;align-items:center;flex-direction:column;">
<h1>COVID19 Kerala API</h1>
<h3><a href="https://covid19-kerala-api.herokuapp.com">https://covid19-kerala-api.herokuapp.com</a><h3>
</div>

## Why?
Manually collecting and updating the data from the pdf sources is time consuming and energy draining! Make use of this API to automatically retrieve the latest as well some of the old COVID19 data specific to Kerala in JSON format, from your applications with ease.

---

## Table of Contents (Optional)

- [Source](#source)
- [Usage](#usage)
- [API](#api)
- [Contributing](#contributing)
- [Team](#team)
- [FAQ](#faq)
- [Support](#support)
- [License](#license)

---

## Source
Currently the API auto collects the data from <a href="http://dhs.kerala.gov.in/">http://dhs.kerala.gov.in/</a>. This site provides reliable data in an unreliable format. Thus some of the data from certain dates are still missing in the dataset. We are actively trying to find a solution to extract data from some of the inconsistent data PDF's.

## Usage
> All you have to do is make a simple GET request to get the indented JSON

---

## API Details
The API currently contains three endpoints `/api`, `/timeline` and `/location` at the moment.

The data can be viewed from the browser by visiting say `https://covid19-kerala-api.herokuapp.com/api` or just use `curl` magic sugar coated by `jq` to view neat response:
```bash
curl "https://covid19-kerala-api.herokuapp.com/api" | jq
```
**Note that all the timestamps in results follow <a href="https://www.w3.org/TR/NOTE-datetime">ISO 8601</a>**

### API Endpoint
The `/api` endpoint serves the whole available data in the following JSON format:
```json
{
    "oldest-timestamp" : {
        "district1": {
            "cases-deaths-etc": int(cases),
            .
            .
            .
            "other_districts": {
                "district" : number_of_persons
            }
        },
        "district2":{...}
    },
    .
    .
    .
    "latest-timestamp": {
        similar-to-above-entry-but-diff-data-values
    }
}
```

#### Example
```bash
curl "https://covid19-kerala-api.herokuapp.com/api" | jq
```

### Location Endpoint
The `/api/location` endpoint can serve a variety of data based on the query parameters that the user provides.
The default response is an array of the districts:
```bash
curl "https://covid19-kerala-api.herokuapp.com/api/location" | jq
```

The parameters that are currently supported include `loc`(specify location) and `date`(specify date of data).

#### Example

---
##### Loc
We can specify an array of locations to be filtered out:
```bash
curl "https://covid19-kerala-api.herokuapp.com/api/location?loc=kasaragod&loc=ernakulam" | jq
```
The above request provides the data of Kasaragod and Ernakulam for all the dates from start till end.

---
##### Date
We can also filter using `date={dd-mm-yyyy|dd/mm/yyyy}` formatted parameter. Date supports inclusion of `<` and `>` characters in the query and even a keyword `latest` to get the latest data only.


Retrieve the data of all locations for the date 1st April 2020:
```bash
curl "https://covid19-kerala-api.herokuapp.com/api/location?date=01-04-2020" | jq
```

Retrieve the data of all locations for all dates greater than 1st April 2020 till last updated date:
```bash
curl "https://covid19-kerala-api.herokuapp.com/api/location?date=>01-04-2020" | jq
```
---
##### Combination
We can also combine these parameters for querying specific entries:
```bash
curl "https://covid19-kerala-api.herokuapp.com/api/location?date=>04-04-2020&loc=ernakulam&loc=kannur" | jq
```

The following query will retrieve the data of Ernakulam and Kannur districts for all dates after 4th April 2020 till latest available date.

### Timeline Endpoint
The `/timeline` endpoint serves the timeline of the number of cases in each district[WIP]. An example response format:
```bash
curl "https://covid19-kerala-api.herokuapp.com/api/timeline" | jq
```
```json
{
    "total_no_of_positive_cases_admitted": {
        "latest": 256,
        "timeline": {
            "2020-02-28T00:00:00Z": 0,
            "...": 1,
            .
            .
            .
            "latest-timestamp": 256
        }
    }
}
```

# Contributing

## Libraries
### Golang
- gin - highly performant web framework
- favicon - middleware for gin
- cron - scheduler
- jsoniter - just faster json encode/decode
- soup - scraper(a bs4 clone)
- color - rainbow puke
### Python
- pdftotext - not the best, but still provides a layout
- jsonpickle - encode/decode objects into json
---

## Project Structure
```
├── bin-------------------------------->covid19keralaapi executable
│
├── cmd
│   └── covid19keralaapi
│       └── main.go-------------------->Entry point and initialization of all pkgs
├── data
│   ├── 05-04-2020.pdf----------------->Latest pdf data collected
│   └── data.json---------------------->Latest json data extracted from the above pdf
├── go.mod---------------------------| 
├── go.sum---------------------------|->Go Modules tracker
├── internal--------------------------->Pkgs for internal use only
│   ├── controller
│   │   └── controller.go-------------->Deserialization, Timeline Generation, Location Array Gen
│   ├── logger
│   │   └── logger.go------------------>Custom logging for all pkgs
│   ├── model
│   │   └── model.go------------------->Primarily for all json unmarshalling
│   ├── scheduler
│   │   └── scheduler.go--------------->Scheduling of scraper,downloader and exec python script
│   ├── scraper
│   │   └── scraper.go----------------->Interface to scrape any website with limited attrs
│   ├── server
│   │   ├── api_handler.go----------|-->'/api' serves the server.JsonData.All.Data
│   │   ├── api_location_handler.go-|-->'/api/location' filters based on loc and date params
│   │   ├── api_timeline_handler.go-|-->'/api/timeline' serves the TimeLine struct
│   │   ├── root_handler.go---------|-->'/' endpoint renders the html frontpage
│   │   └── server.go---------------|-->Server running, alloting handlers to url
│   ├── storage
│   │   └── storage.go----------------->PDF,json filenames, deletion of old pdf file
│   └── website
│       ├── error.go------------------->Custom error handler while scraping the netzz
│       └── website.go----------------->Implements scraper functions and downloads latest data
├── Makefile--------------------------->For easier building and running
├── Procfile
├── readme.md
├── scripts
│   ├── extract-text-data.py----------->Messy script - converts pdf to json
│   ├── Pipfile
│   └── Pipfile.lock
└── web
    └── index.html--------------------->Frontpage
    └── assets/------------------------>No css yet. Just favicons
```

That's a general idea about the structure I have used. I'll happily accept new contributions and ideas. Make sure you check out the issues, or raise one and follow the contribution guidelines, and make your PR(**raise issue before PR or claim already existing one**).

# License
Use this repo in the name of *Freeeeeedommmmmmmm!!* and open source. or this would do [license]()