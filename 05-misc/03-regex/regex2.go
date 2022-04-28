package main

import (
	"fmt"
	"regexp"
)

func main() {
	file := `<books>
    <book>
      <author>Shakespeare</author>
      <title>Romeo and Juliet</title>
      <pages>400</pages>
      <type>paperback</type>
      <cost>17</cost>
    </book>
    <book>
      <author>Shakespeare</author>
      <title>Hamlet</title>
      <pages>270</pages>
      <type>paperback</type>
      <cost>15</cost>
    </book>
</books>`

	r := regexp.MustCompile(`(\<\/?)(title)(\>)`)
	s := r.ReplaceAllString(file, "${1}name${3}")
	fmt.Println(s)

	r = regexp.MustCompile(`(\<\/?)(cost)(\>)`)
	s = r.ReplaceAllString(s, "${1}price${3}")
	fmt.Println(s)
}
