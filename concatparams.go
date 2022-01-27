/* Write a function that takes the arguments received in parameters and returns them as a string.
The string is the result of all the arguments concatenated with a newline (\n) between. */

package maine

func ConcatParams(args []string) string {
	var result string = args[0]     // takse the first word of the string
	for _, word := range args[1:] { // gices the second word and all next ones
		result += "\n" + word // also can write as: reslut=result + "\n" + s
	}

	return result
}

/* func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}
*/
