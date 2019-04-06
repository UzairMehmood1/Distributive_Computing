//I150035_Assignment3_Section(A)
package Assignment_3

import (
	"strings"
	"github.com/ehteshamz/greetings"
)
//Function concatenates byte slice in string seperated by -
func ConcatSlice(sliceToConcat []byte) string {
	var ret string = ""
	for _, val := range sliceToConcat {
			ret += string(val)+"-"
	}
	//Removing last dash -
	ret = strings.TrimRight(ret, "-")
	return ret
}
//Function performs Ceaser Cypher Encryption.
func Encrypt(sliceToEncrypt []byte, ceaserCount int) []byte {
	for i := 0; i < len(sliceToEncrypt); i++{
		//If it lies betweem A-Z
		if sliceToEncrypt[i] >= byte('A') && sliceToEncrypt[i] <= byte('Z') {
			//Adding Ceaser Count to each character
			sliceToEncrypt[i] = sliceToEncrypt[i] + byte(ceaserCount)
			//If, after adding, it exceeds Z, cycle it back to A and so on.
			if sliceToEncrypt[i] >= byte('Z') {
				sliceToEncrypt[i] = (sliceToEncrypt[i] % byte('Z') + (byte('A') - 1))
			}
			//If it lies betweem a-z
		} else if sliceToEncrypt[i] >= byte('a') && sliceToEncrypt[i] <= byte('z') {
			//Adding Ceaser Count to each character
			sliceToEncrypt[i] = sliceToEncrypt[i] + byte(ceaserCount)
			//If, after adding, it exceeds z, cycle it back to a and so on.
			if sliceToEncrypt[i] >= byte('z') {
				sliceToEncrypt[i] = (sliceToEncrypt[i] % byte('z') + (byte('a') - 1))
			}
		}
	}
	return sliceToEncrypt
}
//Function calls the function of another package "github.com/ehteshamz/greetings"
func EZGreetings(name string) string  {
	str := greetings.PrintGreetings(name)
	return str
}
