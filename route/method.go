package route

type Method int

const (
	Get Method = iota
	Post
	Put
	Delete
)

func (m Method) String() string {
	switch m {
	case Get:
		return "GET"
	case Post:
		return "POST"
	case Put:
		return "PUT"
	case Delete:
		return "DELETE"
	}
	return "unknown"
}
