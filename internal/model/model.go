package model

type DataSet struct{
    TimeData TimeLine
    All Data
    Districts Location
}

type TimeLine struct {
    TimeLine map[string]interface{}
}

type Location struct{
    Loc map[string][]string
}

type Data struct {
    Data map[string]interface{}
}

func NewTimeLine() TimeLine{
    var t TimeLine
    t.TimeLine = make(map[string]interface{})
    return t
}

func NewLocation() Location {
    var l Location
    l.Loc = make(map[string][]string)
    return l
}