package eliza

// https://www.smallsurething.com/implementing-the-famous-eliza-chatbot-in-python/

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

//map to link the Regular expressions with the answers
type Response struct {
	re      *regexp.Regexp
	answers []string
}

func NewResponse(pattern string, answers []string) Response { //get eliza's answer
	response := Response{}
	re := regexp.MustCompile(pattern) //fail if there is a mistake in the regular expressions.
	response.re = re
	response.answers = answers
	return response
}

func buildResponseList() []Response { //make a list to store the responses

	allResponses := []Response{}

	file, err := os.Open("./data/patterns.dat") //responses are in the patterns file. RegExp
	if err != nil {                             // there IS an error
		panic(err) // crash the program
	}

	// file exists!
	defer file.Close() // this will be called AFTER this function to close the file.

	scanner := bufio.NewScanner(file)

	for scanner.Scan() { //scan text file to match with the RegExp

		patternStr := scanner.Text()
		scanner.Scan() // move onto the next line which holds the answers
		answersAsStr := scanner.Text()

		answerList := strings.Split(answersAsStr, ";") //answers are split in patterns.dat by ;
		resp := NewResponse(patternStr, answerList)
		allResponses = append(allResponses, resp)
	}

	return allResponses
}

func getRandomAnswer(answers []string) string { //pick out a random answer from possible answer's in the RegExp file
	rand.Seed(time.Now().UnixNano()) // seed to make it return different values.
	index := rand.Intn(len(answers)) // Intn generates a number between 0 and num - 1
	return answers[index]            // can be any element
}

func subWords(original string) string { //substitute words from the user so the sentence is from Eliza's point of view and not the users

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
	// the reflections map is populated.

	words := strings.Split(original, " ")

	for index, word := range words {
		// change the word if it's in the map
		val, ok := reflections[word]
		if ok { // value WAS in the map
			// swap with the value
			words[index] = val // for example, you changes to me, my to your etc.
		}
	}
	return strings.Join(words, " ")
}

func Ask(userInput string) string { //

	responses := buildResponseList()

	for _, resp := range responses { // look at every single response/pattern/answers

		if resp.re.MatchString(userInput) {
			match := resp.re.FindStringSubmatch(userInput)
			//match[0] is full match, match[1] is the capture group
			captured := match[1]
			captured = subWords(captured)

			formatAnswer := getRandomAnswer(resp.answers) // get random element.

			if strings.Contains(formatAnswer, "%s") { // string needs to be formatted
				formatAnswer = fmt.Sprintf(formatAnswer, captured)
			}
			return formatAnswer

		} //end if

	} // end for

	// if we're down here, it means there were no matches;
	return "Sorry, that's a little above my paygrade. Let's move onto something else..." // catch all.
}
