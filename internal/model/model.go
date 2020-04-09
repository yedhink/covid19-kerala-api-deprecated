package model

type DataSet struct{
    TimeLineData TimeLine
    All Data
    Districts Location
}

type TimeLine struct {
    TimeLine map[string]interface{}
}

type Location struct{
    Loc map[string]interface{}
}

type Data struct {
    Data map[string]interface{}
}

func NewData() Data {
    var d Data
    d.Data = make(map[string]interface{})
	d.Data["success"]=false
    d.Data = make(map[string]interface{})
    return d
}

func NewTimeLine() TimeLine{
    var t TimeLine
    t.TimeLine = make(map[string]interface{})
    return t
}

func NewLocation() Location {
    var l Location
    l.Loc = make(map[string]interface{})
    return l
}