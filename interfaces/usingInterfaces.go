package main

import (
	"fmt"
	"io"
)

// Exchanger is a custom interface type that specifies a single method
type Exchanger interface {
	Exchange()
}

// StringPair is a custom structure type that aggregates 2 string type fields.
type StringPair struct {
	first, second string
}

// Exchange is a method on *StringPair type
func (pair *StringPair) Exchange() {
	pair.first, pair.second = pair.second, pair.first
}

// String is a method on StringPair type
func (pair StringPair) String() string {
	return fmt.Sprintf("%q+%q", pair.first, pair.second)
}

func (pair *StringPair) Read(data []byte) (n int, err error) {
	if pair.first == "" && pair.second == "" {
		return 0, io.EOF
	}
	if pair.first != "" {
		n = copy(data, pair.first)
		pair.first = pair.first[n:]
	}
	if n < len(data) && pair.second != "" {
		m := copy(data[n:], pair.second)
		pair.second = pair.second[m:]
		n += m
	}
	return n, nil
}

// Point is a custom integer array type
type Point [2]int

// Exchange is a method on *Point type
func (point *Point) Exchange() {
	point[0], point[1] = point[1], point[0]
}

// String is a method on Point type
func (point Point) String() string {
	return fmt.Sprintf("%v, %v", point[0], point[1])
}

// ExchangeThese takes an argument of type Exchanger slice and calls the Excahange() method
// on the elements of the Exchanger slice
func ExchangeThese(exchangers []Exchanger) {
	for _, exchanger := range exchangers {
		exchanger.Exchange()
	}
}

// ToBytes takes an io.Reader value and a size limit and returns a []byte and also an error.
func ToBytes(reader io.Reader, size int) ([]byte, error) {
	data := make([]byte, size)
	n, err := reader.Read(data)
	if err != nil {
		return data, err
	}
	return data[:n], nil
}

func main() {
	var cavil Exchanger                   // variable of type Exchanger interface
	cavil = &StringPair{"henry", "cavil"} // value of the Exchanger interface type
	point := Point{3, 7}
	fmt.Printf("before Exchange call:\n %v %v \n", cavil, point)

	cavil.Exchange()
	point.Exchange()
	fmt.Printf("after Exchange call:\n %v %v \n", cavil, point)

	banner := StringPair{"bruce", "banner"}
	fmt.Printf("before ExchangeThese call:\n %v %v %v \n", cavil, banner, point)

	exchangers := []Exchanger{
		// The following pointers are Exchanger values because they have a method with signature Exchange().
		cavil, &banner, &point,
	}
	ExchangeThese(exchangers)
	fmt.Printf("after exchangeThese call:\n %v %v %v \n", cavil, banner, point)

	readers := []io.Reader{
		// The following pointers are io.Reader values because they have a method with signature Read([]byte) (int, error)
		&StringPair{"joe", "rogan"},
		&StringPair{"bobby", "lee"},
	}
	for _, reader := range readers {
		raw, err := ToBytes(reader, 16)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(raw)
		fmt.Printf("%q\n", raw)
	}
}
