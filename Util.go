package main

import (
	"strings"
	"bytes"
	"unicode"
)

const (
	A_ADDRESS = "127.0.0.1:6379"
)


func CapitalliseFirstLetterofEveryWord(title string) string {
	var buffer bytes.Buffer
	words := strings.Fields(title)
	for _, s := range words {
		buffer.WriteString(Caps(s) + " ")
	}
	return buffer.String()
}

func Caps(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i + 1:]
	}
	return ""
}

//func RemoveStopWordsTest(input []string) string {
//	var buffer bytes.Buffer
//	if err != nil {
//		log.Fatal(err)
//	}
//	c.Send("MULTI")
//	c.Send("DEL", "inputWords")
//	c.Send("SADD", redis.Args{}.Add("inputWords").AddFlat(input)...)
//	c.Send("SDIFF", "inputWords", "stopwords")
//	reply, err := c.Do("EXEC")
//	if err != nil {
//		fmt.Println("Error Executing Commands", err)
//	}
//	values, _ := redis.Values(reply, nil)
//	fliteredWords, err := redis.Strings(values[2], nil)
//	if err != nil {
//		fmt.Println("Wrong Type Received", err)
//	}
//	if (len(fliteredWords)) > 0 {
//		for _, v := range fliteredWords {
//			buffer.WriteString(v + " ")
//		}
//	} else {
//		fmt.Println(">>Nothing found")
//		for _, v := range input {
//			buffer.WriteString(v + " ")
//		}
//	}
//	return buffer.String()
//}

func GetTermsTEST(sentence string) [] string {
	terms := strings.Split(string(sentence), " ")
	return terms;
}






