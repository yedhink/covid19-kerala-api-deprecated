# TODO
## Python
* allow querying keys with say TVm or 01 as in KL 01 for thiruvananthapuram
* import only necessary items
* redo the verbose flag such that printing in loops etc are possible
* change loop in extract_text_data func into for page in pdf:
* extract the no of persons recovered from annex-2
* add last modified field to json
## GO
* in server currently we compare if the file is latest from remote by checking the filename vs title. but there might
    occur a problem when more than one file is uploaded in same day
* server.go line 71 no need for iteration. simply files[0] would do

# Features
* auto parsing the data from dhs. this feature is very much in the **experimental** stage, since conversion of the pdf to text may not be that consistent.

# Breif the idea
My idea is to purely technical. I propose to create an API for tracking the coronavirus (COVID-19) outbreak specific to KERALA. For instance: you can get data per location by using this example URL say :-  https://covid19-kerala-api/api/locations or say getting latest amount of total confirmed cases, deaths, and recovered https://covid19-kerala-api/api/latest. The data can be retrieved in JSON format as of now.

# Use Case
The primary use case is to make easier accessibility and availability of the real time data to developers such that they could easily query and make of this API in their applications. Also this API can enable people to develop applications which can show the realtime "Covid-19 Kerala Infection Map" or "Covid19 Kerala Tracker". This will be very useful for people who have to go out for work or even police officials. With such an application people could easily lookup a place and understand where they can safely go and not go. Currently i propose to make the api, then I can move on to developing the above mentioned application if you provide support.

# web scraper
* the dhs kerala website is too slow to load up. thus querying time using scraper too increased