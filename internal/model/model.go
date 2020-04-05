package model

type DataSet struct{
    TimeData TimeLine
    All Data
}

type TimeLine struct {
    TimeLine map[string]interface{}
}

func NewTimeLine() TimeLine{
    var t TimeLine
    t.TimeLine = make(map[string]interface{})
    return t
}

type Data struct {
    Data map[string]interface{}
}