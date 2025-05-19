package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println("hello world")
	fmt.Printf("hello\n")

	string1 := "thisisgolang"

	//cheak for 'go'
	fmt.Println(strings.Contains(string1,"go")) //true
	fmt.Println(strings.ContainsAny(string1,"g")) //true
	fmt.Println(strings.ContainsRune(string1,'x')) //false Rune cheak for any existance in string with out any condition 
	fmt.Println(strings.ContainsFunc(string1,func(r rune) bool {
		return r == 'x' || r == 's' 
	})) //cheak is x or s in the string this is like Rune but with condition 

	//now we can go and cheak for count 
	fmt.Println(strings.Count(string1,"s")) //should be 2
	fmt.Println(strings.Count(string1,"")) // ?
	//strings numer are as it said count it and +1 it so we have to -1 it 
	// 1 + unicode the number of empy "" so for the exact number we shoud (1+unicode)-1 
	number_of_strings := strings.Count(string1,"")
	number_of_strings--
	fmt.Println(number_of_strings)

	//we can make this as a function
	string2 := "hello world"
	string3 := "a,b,c,d"
	Counting_String(string2)

	//cut
	fmt.Println(strings.Cut(string2,"l")) //delete the first l word
	
	// lets cheak out for join and split
	//split cheak the existance of the string with condition
	fmt.Println(len(strings.Split(string3," "))) // -> 4
	fmt.Println(len(strings.SplitN(string3,",", 4))) /*-> a, b, c,d 4 word befor , will print */ 
	fmt.Println(strings.SplitN(string3 , ",", 4)) // -> a b c d
	
	//split ya fasele baad az shart moshakhas
	fmt.Println(strings.SplitAfter(string3 ,",")) //-> space fater  ,  a, b, c, d

	/*for clearing the string of any whitespace or tab or ... we can use this */
	letter := "this is a letter from \n \t good person 		who loves you"
	fmt.Println(strings.Fields(letter))
	

	/* 
****************************************************


split Replace any word or char into space 



***************************************************
	*/

	//now we go for join
	
	/*
	********************
	join use for replace the space 
	********************
	*/
	words:= "a b c d e r"
	fmt.Println(strings.Join([]string{words,"f"},",")) // ->add join in the end of the string
	My_Words_Arraye := strings.Join([]string{words,","}," ")
	println(My_Words_Arraye)


	//Repeat 
	fmt.Println(strings.Repeat("tehran",5)) //tehran * 5 

	//Replace && ReplaceAll
	Sentence := "This is Iran and America"
	fmt.Println(strings.Replace(Sentence,"Iran", "go", 1))
	fmt.Println(strings.ReplaceAll(Sentence,"i","x"))

	//Compare
	fmt.Println(strings.Compare("Golang","golang")) //-> -1 Case Sensitive
	//EqualFold
	fmt.Println(strings.EqualFold("golang","GOLANG")) //True none Case sensitive

	//prefix and suffix 
	fmt.Println(strings.HasPrefix("iran","ir")) //-> starts with ir  True
	fmt.Println(strings.HasSuffix("iran","n"))//->end with n	True

	//Index 
	fmt.Println(strings.Index("iran","i")) //->0
	//LastIndex
	fmt.Println(strings.LastIndex(Sentence,"r")) 
	
	//Tolower
	fmt.Println(strings.ToLower("GOLANG"))
	//Toupper
	fmt.Println(strings.ToUpper("golang"))
	//Trim	-> what to cut from left and right 
	fmt.Println(strings.Trim("   go   "," ")) //->cut the space from left and right 
	//title
	fmt.Println(strings.Title("iran is a great country"))

}




//count strings func
func Counting_String(String string){
	fmt.Printf("number of strings = %d\n",strings.Count(String,"")-1)
}
