package model

type ReadingStatus uint

var StatusText = map[uint]string{
	10: "弃读",
	15: "特殊书类",
	21: "粗读完结",
	29: "正常完结",
	95: "在读",
}
