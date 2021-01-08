package util

type Method string

const (
	GET    Method = "GET"
	PUT    Method = "PUT"
	POST   Method = "POST"
	DELETE Method = "DELETE"
)

func (e Method) String() string {
	extensions := [...]string{"JPG", "PNG", "GIF", "BMP"}

	x := string(e)
	for _, v := range extensions {
		if v == x {
			return x
		}
	}

	return ""
}
