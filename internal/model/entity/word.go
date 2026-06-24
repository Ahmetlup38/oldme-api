// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Word is the golang structure for table word.
type Word struct {
	Id      uint   `json:"id"      orm:"id"      description:""`
	Word    string `json:"word"    orm:"word"    description:""`
	Explain string `json:"explain" orm:"explain" description:""`
}
