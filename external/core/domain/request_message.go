package domain

type RequestMessage struct {
  Key string
  RequesterId string
  SafeKey string
  Payload interface{}
}
