package data

type DataList struct {
	list []Data
}

type Data struct {
	uuid            string
	event           string
	importanceLevel int
}

type DataHandler interface {
	findData(string) Data
	deleteData(string) error
	addData(Data) error
	updateData(former Data, Later Data) error
}
