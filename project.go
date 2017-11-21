package main

import (
	"fmt"
	"net/http"
	"math/rand"
	"regexp"
	"time"
	
)

func userinputhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w ,elizaResponse(r.URL.Query().Get("value"))) 
	
}

func elizaResponse(input string )string {

	//Hello
	if matched, _ := regexp.MatchString(`(?i).*\b[hH]ello\b.*`, input); matched{
		responses := []string {
		"How do you do.  Please state your problem.",
		"Hi.  What seems to be your problem ?",
		}
		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}

	//sorry
	if matched, _ := regexp.MatchString(`(?i).*\b([aA]pologise|[sS]orry)\b.*`, input); matched{
		responses := []string {
			"Please don't apologise.",
			"Apologies are not necessary.",
			"I've told you that apologies are not required.",
			"It did not bother me.  Please continue.",
		}
		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}

	//i remember
	if matched, _ := regexp.MatchString(`(?i).*\bi [rR]emember\b.*`, input); matched{
		responses := []string {
			"Do you often think of this ?",
			"Does thinking of this bring anything else to mind ?",
			"What else do you recollect ?",
			"Why do you remember this just now ?",
			"What in the present situation reminds you of this ?",
			"What is the connection between me and this ?",
			"What else does this remind you of ?",
		}
		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}

	//do you remember
	if matched, _ := regexp.MatchString(`(?i).*\b[dD]o you remember\b.*`, input); matched{
		responses := []string {
			"Did you think I would forget that ?",
			"Why do you think I should recall that now ?",
			"What about that?",
		}
		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}



	//i forget
	if matched, _ := regexp.MatchString(`(?i).*\bi[fF]orget\b.*`, input); matched{
		responses := []string {
			"Can you think of why you might forget that ?",
			"Why can't you remember that ?",
			"How often do you think of that?",
			"Does it bother you to forget that ?",
			"Could it be a mental block ?",
			"Are you generally forgetful ?",
			"Do you think you are suppressing thoughts about that ?",
		}
		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}

	//did you forget
	if matched, _ := regexp.MatchString(`(?i).*\b[dD]id you forget\b.*`, input); matched{
		responses := []string {
			"Why do you ask ?",
			"Are you sure you told me ?",
			"Would it bother you if I forgot that ?",
			"Why should I recall that just now ?",
			"Tell me more about that.",
		}
		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}

	//different
	if matched, _ := regexp.MatchString(`(?i).*\b[dD]ifferent\b.*`, input); matched{
		responses := []string {
			"How is it different ?",
			"What differences do you see ?",
			"What does that difference suggest to you ?",
			"What other distinctions do you see ?",
			"What do you suppose that disparity means ?",
			"Could there be some connection, do you suppose ?",
			"How ?",
		}
		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}
	
	//alike
	if matched, _ := regexp.MatchString(`(?i).*\b[aA]like\b.*`, input); matched{
		responses := []string {
			"In what way ?",
			"What resemblence do you see ?",
			"What does that similarity suggest to you ?",
			"What other connections do you see ?",
			"What do you suppose that resemblence means ?",
			"What is the connection, do you suppose ?",
			"Could there really be some connection ?",
			"How ?",
		}
		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}
	
	//always
	if matched, _ := regexp.MatchString(`(?i).*\b[aA]lways\b.*`, input); matched{
		responses := []string {
		"Can you think of a specific example ?",
		"When ?",
		"What incident are you thinking of ?",
		"Really, always ?",
		}
		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}
	
	//everyone/everybody
	if matched, _ := regexp.MatchString(`(?i).*\b([eE]veryone|[eE]verybody)\b.*`, input); matched{
		responses := []string {
		"Really, everyone ?",
		"Surely not everybody.",
		"Can you think of anyone in particular ?",
		"Who, for example?",
		"Are you thinking of a very special person ?",
		"Who, may I ask ?",
		"Someone special perhaps ?",
		"You have a particular person in mind, don't you ?",
		"Who do you think you're talking about ?",
		}
		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}
	
	//why dont you
	if matched, _ := regexp.MatchString(`(?i).*\b[wW]hy dont you\b.*`, input); matched{
		responses := []string {
		"Do you believe I don't?",
		"Perhaps I will do that in good time.",
		"Should you do that to yourself ?",
		"You want me to do that?",

		}
		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}
	
	//why cant i
	if matched, _ := regexp.MatchString(`(?i).*\b[wW]hy can't i\b.*`, input); matched{
		responses := []string {
			"Do you think you should be able to do that ?",
			"Do you want to be able to do that ?",
			"Do you believe this will help you to do that ?",
			"Have you any idea why you can't do that ?",
		}
		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}
	
	//because
	if matched, _ := regexp.MatchString(`(?i).*\b[bB]ecause\b.*`, input); matched{
		responses := []string {
		"Is that the real reason ?",
		"Don't any other reasons come to mind ?",
		"Does that reason seem to explain anything else ?",
		"What other reasons might there be ?",
		}
		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}
	
	//why/what
	if matched, _ := regexp.MatchString(`(?i).*\b([wW]hy|[wW]hat)\b.*`, input); matched{
		responses := []string {
			"Why do you ask ?",
			"Does that question interest you ?",
			"What is it you really want to know ?",
			"Are such questions much on your mind ?",
			"What answer would please you most ?",
			"What do you think ?",
			"What comes to mind when you ask that ?",
			"Have you asked such questions before ?",
			"Have you asked anyone else ?",
		}
		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}
	
	//who
	if matched, _ := regexp.MatchString(`(?i).*\b[wW]ho\b.*`, input); matched{
		responses := []string {
			"Are you sure I know them ?",
			"Why do you want to know who that is?",
			"If i knew who that was I would tell you?",
		}
		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}
	
	//can you
	if matched, _ := regexp.MatchString(`(?i).*\b[cC]an you\b.*`, input); matched{
		responses := []string {
		"You believe I can do that don't you ?",
		"You want me to be able to do that.",
		"Perhaps you would like to be able to do that yourself.",
		}
		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}
	
	//can i
	if matched, _ := regexp.MatchString(`(?i).*\b[cC]an i\b.*`, input); matched{
		responses := []string {
		"Whether or not you can do that depends on you more than on me.",
		"Do you want to be able to do that ?",
		"Perhaps you don't want to do that.",
		}
		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}
	
	//computer
	if matched, _ := regexp.MatchString(`(?i).*\b[cC]omputer\b.*`, input); matched{
		responses := []string {
			"Do computers worry you ?",
			"Why do you mention computers ?",
			"What do you think machines have to do with your problem ?",
			"Don't you think computers can help people ?",
			"What about machines worries you ?",
			"What do you think about machines ?",
			"You don't think I am a computer program, do you ?",

		}
		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}
	
	//if 
	if matched, _ := regexp.MatchString(`(?i).*\bif\b.*`, input); matched{
		responses := []string {
			"Do you think that's likely ?",
			"Do you wish that to happen ?",
			"What do you know about that?",
			"What would you do if that happened ?",
			"But what are the chances of that happening ?",
			"What does this speculation lead to ?",
		}
		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}
	
	//dreamed
	if matched, _ := regexp.MatchString(`(?i).*\bdreamed\b.*`, input); matched{
		responses := []string {
			"Really, that's what you dreamt ?",
			"Have you ever fantasized about that while you were awake ?",
			"Have you ever dreamed about that before ?",
		}
		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}
	
	//dream
	if matched, _ := regexp.MatchString(`(?i).*\bdream\b.*`, input); matched{
		responses := []string {
			"What does that dream suggest to you ?",
			"Do you dream often ?",
			"What persons appear in your dreams ?",
			"Do you believe that dreams have something to do with your problem ?",
		}
		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}
	
	//perhaps
	if matched, _ := regexp.MatchString(`(?i).*\b[pP]erhaps\b.*`, input); matched{
		responses := []string {
		"You don't seem quite certain.",
		"Why the uncertain tone ?",
		"Can't you be more positive ?",
		"You aren't sure ?",
		"Don't you know ?",
		"How likely, would you estimate ?",
		}
		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}
		
	//name
	if matched, _ := regexp.MatchString(`(?i).*\bname\b.*`, input); matched{
		responses := []string {
		"I am not interested in names.",
		"I've told you before, I don't care about names -- please continue.",
		}
		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}
	
	//perhaps
	if matched, _ := regexp.MatchString(`(?i).*\b[pP]erhaps\b.*`, input); matched{
		responses := []string {
		"You don't seem quite certain.",
		"Why the uncertain tone ?",
		"Can't you be more positive ?",
		"You aren't sure ?",
		"Don't you know ?",
		"How likely, would you estimate ?",
		}
		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}
	
	//perhaps
	if matched, _ := regexp.MatchString(`(?i).*\b[pP]erhaps\b.*`, input); matched{
		responses := []string {
		"You don't seem quite certain.",
		"Why the uncertain tone ?",
		"Can't you be more positive ?",
		"You aren't sure ?",
		"Don't you know ?",
		"How likely, would you estimate ?",
		}
		randIndex := rand.Intn(len(responses))
		return responses[randIndex]
	}




	














































	*/

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