package main

import (
	"fmt"
	"net/http"
	"math/rand"
	"regexp"
	"time"
	
)

func userinputhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w ,elizaResponse(r.URL.Query().Get("value"))) //.Path[1:])
	
}

func elizaResponse(input string )string {

	if matched, _ := regexp.MatchString(`(?i).*\b[Ii]reland\b.*`, input); matched{
		return "\n Ireland will 1-0. Shane Duffy will score a header"
	}


	if matched, _ := regexp.MatchString(`(?i).*\b[fF]ootball\b.*`, input); matched{
		responses := []string {
		"Do you often think of this ?",
		"Does thinking of (i) bring anything else to mind ?",
		"What else do you recollect ?",
		"Why do you remember (i) just now ?",
		"What in the present situation reminds you of (i) ?",
		"What is the connection between me and (i) ?",
		"What else does (i) remind you of ?",
	}

		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}

	responses := []string {
		"I'm not sure I understand you fully.",
		"Please go on.",
		"What does that suggest to you ?",
		"Do you feel strongly about discussing such things ?",
		"That is interesting.  Please continue.",
		"Tell me more about that.",
		"Does talking about this bother you ?",
	}

	randIndex := rand.Intn(len(responses))
	return responses[randIndex]
}

func main() {
	rand.Seed(time.Now().Unix()) 

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/user-input", userinputhandler)
	http.ListenAndServe(":8080", nil)

}