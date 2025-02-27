// ---------------------------------------------------------
// EXERCISE: Unique Words 2
//
//  Use your solution from the previous "Unique Words"
//  exercise.
//
//  Before adding the words to your map, remove the
//  punctuation characters and numbers from them.
//
//
// BE CAREFUL
//
//  Now the shakespeare.txt contains upper and lower
//  case letters too.
//
//
// EXPECTED OUTPUT
//
//  go run main.go < shakespeare.txt
//
//   There are 100 words, 69 of them are unique.
//
// ---------------------------------------------------------
package main
import(
"fmt"
"os"
"bufio"
"regexp"
"strings"
)
func main() {

	total, words := 0, make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	re:=regexp.MustCompile('[^a-zA-Z\\s]+')
	for input.Scan() {
		text:=input.Text()
		text=strings.ToLower(text)
		text=re.ReplaceAllString(text,"")
		total++
		words[text]++
	}
	unique := len(words)
	fmt.Printf("There are %d , %d are them unique", total, unique)

}
