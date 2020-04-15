package main

var chars = []byte("123456789bcdfghjkmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ")

func generateShortUrl(Id int) (string) {
	var code []byte;
	id := Id
	for id > len(chars) - 1 {
		var remainder = id % len(chars)
		code = append(code, chars[remainder])
		id = id / len(chars)
	}
	code = append(code, chars[id])
	result := string(code)
	return result
}