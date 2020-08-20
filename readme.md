<h1 align="center">COVID19 Kerala API</h1>
<h3 align="center"><a href="https://covid19-kerala-api.herokuapp.com">https://covid19-kerala-api.herokuapp.com</a><h3>

## Note
The project is currently unmaintained since the source of the covid data, in pdf form, had got highly inconsistent. Thus the parsing of such data pdfs, which changes layout on a daily basis, requires a lot of modifications to the existing scripts and a lot of time investment. 

## Why?
Manually collecting and updating the data from the pdf sources is time consuming and energy draining! Make use of this API to **automatically** retrieve the **latest** as well some of the old COVID19 data specific to Kerala in JSON format, into your applications with ease.

---

## Table of Contents

* [Source](#source)
* [Usage](#usage)
* [API Details](#api-details)
    * [Notes](#notes)
    * [API Endpoint](#api-endpoint)
        * [Example](#example)
    * [Location Endpoint](#location-endpoint)
        * [Example](#example-1)
            * [Loc](#loc)
            * [Date](#date)
            * [Combination](#combination)
    * [Timeline Endpoint](#timeline-endpoint)
* [Contributing](#contributing)
    * [Libraries](#libraries)
        * [Golang](#golang)
        * [Python](#python)
    * [Running](#running)
    * [Project Structure](#project-structure)
* [License](#license)

---

## Source
Currently the API auto collects the data from <a href="http://dhs.kerala.gov.in/">http://dhs.kerala.gov.in/</a>. This site provides reliable data in a very unreliable and inconsistent format. Thus some of the data from certain dates are still missing in the dataset. Currently trying to find a solution to extract data from some of the inconsistent data PDF's.

## Usage
> All you have to do is make a simple GET request to get the indented JSON

---

## API Details
The API currently contains three endpoints `/api`, `/timeline` and `/location` at the moment.

The data can be viewed from the browser by visiting say `https://covid19-kerala-api.herokuapp.com/api` or just use `curl` magic, sugar coated by `jq` to view a neat response:
```bash
curl "https://covid19-kerala-api.herokuapp.com/api" | jq
```

### Notes
* **All json responses consist of a `success` key, which denotes whether the user request was valid or whether it retrieved any data**
* **All the timestamps in results follow <a href="https://www.w3.org/TR/NOTE-datetime">ISO 8601</a><br>**
* **Sometimes the response might be slow - because heroku shuts down it's dynos after a certain interval of inactivity and it has to restart when a request is made in such a state**
* **The JSON responses are not indented for performance sake. All examples are shown in indented format for readability only**

### API Endpoint
The `/api` endpoint serves the whole available data in the following JSON format(this is a rough format):
```json
{
	"success": true | false, // whether the request is valid or not
    {oldest-timestamp} : {
        {district1}: {
            {cases-deaths-etc}: {int(cases)},
            ...,
            ...,
            ...,
            "other_districts": {
                {district} : {number_of_persons}
            }
        },
        {district2}:{...}
        ...,
        "total":{...}
    },
    ...,
    ...,
    ...,
    {latest-timestamp}: {
        {similar-to-above-entry-but-data-values-corresponds-to-timestamp}
    }
}
```

#### Example
```bash
curl "https://covid19-kerala-api.herokuapp.com/api" | jq
```

### Location Endpoint
The `/api/location` endpoint can serve a variety of data based on the query parameters that the user provides.
The default response is an array of the possible location values acceptable by `loc` parameter:
```bash
curl "https://covid19-kerala-api.herokuapp.com/api/location" | jq
```

The parameters that are currently supported include `loc`(specify location) and `date`(specify date/timestamp).

#### Example

---
##### Loc
We can specify an array of locations to be filtered out:
```bash
curl "https://covid19-kerala-api.herokuapp.com/api/location?loc=kasaragod&loc=ernakulam" | jq
```
The above request provides the data pertaining to Kasaragod and Ernakulam districts from the oldest timestamp till latest.

---
##### Date
We can also filter using `date={dd-mm-yyyy|dd/mm/yyyy}` formatted parameter. Here the `date` supports inclusion of `<` and `>` characters in the query and even a keyword `latest` to get the latest data only.


Retrieve the data of all locations for the date 1st April 2020:
```bash
curl "https://covid19-kerala-api.herokuapp.com/api/location?date=01-04-2020" | jq
```

Retrieve the data from all locations with dates(timestamp) greater than 1st April 2020 till the last updated date:
```bash
curl "https://covid19-kerala-api.herokuapp.com/api/location?date=>07-04-2020" | jq
```
---
##### Combination
We can also combine these parameters for querying specific entries:

Getting the total summary from the latest data:-
```bash
curl "https://covid19-kerala-api.herokuapp.com/api/location?loc=total&date=latest" | jq
```

Retrieving the data of Ernakulam and Kannur districts for all dates after 4th April 2020 till latest timestamp.
```bash
curl "https://covid19-kerala-api.herokuapp.com/api/location?date=>04-04-2020&loc=ernakulam&loc=kannur" | jq
```

### Timeline Endpoint
The `/timeline` endpoint serves the timeline of the number of cases in each district[WIP]. An example response format:
```bash
curl "https://covid19-kerala-api.herokuapp.com/api/timeline" | jq
```
```json
{
	"success": true,
    "total_no_of_positive_cases_admitted": {
        "latest": 256,
        "timeline": {
            "2020-02-28T00:00:00Z": 0,
            "...": 1,
            ...,
            ...,
            ...,
            {latest-timestamp}: 256
        }
    }
}
```

# Contributing
This is a general idea about the structure I have used. I'll happily accept new contributions and ideas. Make sure you check out the [issues](https://github.com/yedhink/covid19-kerala-api/issues), or raise one and follow the [contribution guidelines](https://github.com/yedhink/covid19-kerala-api/blob/master/CONTRIBUTING.md), and make your PR(**raise issue before PR or claim already existing issue**).

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

## Running
> *i 'gnu' that `make` is gods own creation, the moment i laid my hands on it*

Start off by installing the go and python packages - only needs to be done the first time:-
```bash
make init
```

The python script is invoked from within the gin-server. Therefore activate the pipenv shell first:-
```bash
pipenv shell
```
Then run the server(note that the executable will be stored in bin/):-  
To run in production mode
```bash
make build
```
To run in development mode (i.e. gin server will be in development mode)
```bash
make run
```
Once everything is setup, essentially running `make build` or `make run` (as per requirement) from project root can restart the server everytime.

---

## Project Structure
```
├── bin-------------------------------->covid19keralaapi executable
├── cmd
│   └── covid19keralaapi
│       └── main.go-------------------->Entry point and initialization of all pkgs
├── data
│   ├── 09-04-2020.pdf----------------->Latest pdf data collected
│   └── data.json---------------------->Latest json data extracted from the above pdf
├── go.mod---------------------------| 
├── go.sum---------------------------|->Go Modules tracker
├── internal--------------------------->Pkgs for internal use only
│   ├── date
│   │   └── date.go-------------------->All date related functions and validations
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
│   │   ├── no_route_handler.go-----|-->handles invalid routes
│   │   ├── api_handler.go----------|-->'/api' serves the server.JsonData.All.Data
│   │   ├── api_location_handler.go-|-->'/api/location' filters based on loc and date params
│   │   ├── api_timeline_handler.go-|-->'/api/timeline' serves the TimeLine struct
│   │   ├── root_handler.go---------|-->'/' endpoint renders the html frontpage
│   │   └── server.go---------------|-->Server running, allotting handlers to url
│   ├── storage
│   │   └── storage.go----------------->PDF,json filenames, deletion of old pdf file
│   └── website
│       ├── error.go------------------->Custom error handler while scraping the netzz
│       └── website.go----------------->Implements scraper functions and downloads latest data
├── Makefile--------------------------->For easier building and running
├── readme.md
├── scripts
│   ├── extract-text-data.py----------->Messy script - converts pdf to json
│   ├── Pipfile
│   └── Pipfile.lock
└── web
    └── index.html--------------------->Frontpage
    └── assets/------------------------>No css yet. Just favicons
```

# License
Use this repo in the name of *Freeeeeedommmmmmmm!!* and open source. or this would do - [license](https://github.com/yedhink/covid19-kerala-api/blob/master/LICENSE.txt)
