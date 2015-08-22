/*
 *Slice uses a underlying array. Use a slice may cause cannot release the array memory.
 */

var digitRegexp = regexp.MustCompile("[0-9]+")

// cause can not release the file bytes
func FindDigits(filename string) []byte {
    b, _ := ioutil.ReadFile(filename)
    return digitRegexp.Find(b)
}

//copy the result
func CopyDigits(filename string) []byte {
    b, _ := ioutil.ReadFile(filename)
    b = digitRegexp.Find(b)
    c := make([]byte, len(b))
    copy(c, b)
    return c
}

// or use append
func CopyDigits(filename string) []byte {
    b, _ := ioutil.ReadFile(filename)
    b = digitRegexp.Find(b)
	var c []byte
	append(c, b)
    return c
}
