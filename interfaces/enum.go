package interfaces

type TracingName string

const (
	Error    TracingName = "Error"
	Response TracingName = "Response"
	Request  TracingName = "Request"
)
