package data

//用于处理JSON数据
type JsonHandler struct {
}

type DataIO interface {
	writeData(DataList)
	readData()
}
