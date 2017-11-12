package eliza

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

var reflections map[string]string

type Response struct {
	re      *regexp.Regexp
	answers []string
}

func NewResponse(pattern string, answers []string) Response {
	response := Response{}
	re := regexp.MustCompile(pattern)
	response.re = re
	response.answers = answers
	return response
}

func buildResponseList() []Response {

	allResponses := []Response{}

	file, err := os.Open("./data/patterns.dat")
	if err != nil { // there IS an error
		panic(err) // crash the program
	}

	// file exists!
	defer file.Close() // this will be called AFTER this function.

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//fmt.Println(scanner.Text())

		patternStr := scanner.Text()
		scanner.Scan() // move onto the next line which holds the answers
		answersAsStr := scanner.Text()

		answerList := strings.Split(answersAsStr, ";")
		resp := NewResponse(patternStr, answerList)
		allResponses = append(allResponses, resp)
	}

	return allResponses
}

func getRandomAnswer(answers []string) string {
	rand.Seed(time.Now().UnixNano()) // seed to make it return different values.
	index := rand.Intn(len(answers)) // Intn generates a number between 0 and num - 1
	return answers[index]            // can be any element
}

func subWords(original string) string {
	// https://www.smallsurething.com/implementing-the-famous-eliza-chatbot-in-python/
	//reflections = readLines("file/path")// []string am:are

	if reflections == nil { // map hasn't been made yet
		reflections = map[string]string{ // will only happen once.
			"am":     "are",
			"was":    "were",
			"i":      "you",
			"i'd":    "you would",
			"i've":   "you have",
			"i'll":   "you will",
			"my":     "your",
			"are":    "am",
			"you've": "I have",
			"you'll": "I will",
			"your":   "my",
			"yours":  "mine",
			"you":    "me",
			"me":     "you",
		}
	}
	// when we're we can be sure reflectiosn map is populated.

	words := strings.Split(original, " ")

	for index, word := range words {
		// we want to change the word if it's in the map
		val, ok := reflections[word]
		if ok { // value WAS in the map
			// we want to swap with the value
			words[index] = val // eg. you -> me
		}
	}

	return strings.Join(words, " ")
}

func Ask(userInput string) string {

	// My name is bob
	responses := buildResponseList()

	for _, resp := range responses { // look at every single response/pattern/answers

		if resp.re.MatchString(userInput) {
			match := resp.re.FindStringSubmatch(userInput)
			//match[0] is full match, match[1] is the capture group
			captured := match[1]

			// remove punctuation here! <------

			captured = subWords(captured)

			formatAnswer := getRandomAnswer(resp.answers) // get random element.

			if strings.Contains(formatAnswer, "%s") { // string needs to be formatted
				formatAnswer = fmt.Sprintf(formatAnswer, captured)
			}
			return formatAnswer

		} // if

	} // for

	// if we're down here, it means there were no matches;
	return "Sorry, that's a little above my paygrade." // catch all.
	/*
		//patternStr := "name is (.*)" // Hello my name is bob
		// MustCompile, Compile to make a *regexp.Regexp struct
		//re := regexp.MustCompile(patternStr)

		if re.MatchString(userInput) {
			fmt.Println("There was a match!")
			//re.FindStringSubmatch()
			match := re.FindStringSubmatch(userInput)
			//match[0] is full match, match[1] is the capture group
			captured := match[1]
			fmt.Println(captured)

			formatString := "Hello %s, it's nice to meet" // this is the format string
			answer := fmt.Sprintf(formatString, captured)
			fmt.Println(answer)

		} else {
			fmt.Println("There was no match")
		}

		// slice / list of answers, and I return 1 at random
		// Hi bob
		// Hello bob
		// how's it hanging bob
	*/
}
