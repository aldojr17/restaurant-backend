package domain

type Response struct {
	Data interface{}
	Code int
	Err  error
}
