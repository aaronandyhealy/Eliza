**Eliza Go Project**
--------------------

My name is Aaron Healy. This is my repository for my chatbot based on Eliza.

This chatbot was made as a 3rd Year Project at GMIT.

**Setup Instructions**

To run this chatbot, you'll have to install Go. To do this visit [here](https://golang.org/dl/) . Continue by cloning this repository by opening up a command line typing this command.

git clone https://github.com/aaronandyhealy/Eliza.git

You need to first compile the code using:
go build .\project.go

and  by simply running the command:
go run .\project.go

To start talking to Eliza and using the chatbot you need to open a browser and navigate to:
http://localhost:8080/

Just enter "Hello" in the textbox and press enter to start a converstaion....


**Design Decisions**
Firstly used a userInputHandler to recieve the users message and pass it to my elizaResponse function before writing it back out to the web server page.

In the elizaResponse function I used regular expressions to capture user input and output an answer that would make the user believe Eliza can understand what the conversation is about. Eliza picks from a certain group of answers for each different expression. 

I used jQuery and CSS styling to make the chatbox appear as if Eliza was replying the same way you would recieve a message from a friend on social media platforms. 

**Problems That I Had.**

I initially wanted to keep the responses in a separate file but could not get them to output correctly when trying to do it this way. I had to settle for one main Go file which would have a huge amount of lines of code due to the number of response options.

I could not find a succesful way of being able to store the word that matched the regular expression pattern and be able to out put it in the answer e.g. Input: "I like football"
								Output: "Why do you like football?"
My Eliza will say something like "Why do you like that?"


**Languages**

Go programming language was used to make the web server and Eliza. HTML, JavaScript and jQuery were used in the creation of the web server page.

Bootstrap was also used in the web server page.
