package simplenet

func WriteString(c *netSocket, s string) (n int, err error) {
	return c.Write([]byte(s))
}
