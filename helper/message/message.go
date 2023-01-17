package message

type Message struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var SuccessMsg = Message{Code: 2000, Message: "Success"}
var ErrNotFound = Message{Code: 5000, Message: "Data Not Found"}
var ErrSaveData = Message{Code: 7001, Message: "Data cannot be saved, please check your request"}
var ErrDeleteData = Message{Code: 7002, Message: "Data cannot be deleted, please check your request"}
var ErrServerTimeout = Message{Code: 5005, Message: "Service Timeout"}
var ErrValidation = Message{Code: 4002, Message: "Validation Failed"}
