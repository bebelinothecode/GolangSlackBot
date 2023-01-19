package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <- chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events");
		fmt.Println(event.Timestamp);
		fmt.Println(event.Command);
		fmt.Println(event.Parameters);
		fmt.Println(event.Event)
	}
}



func main() {
	err := godotenv.Load();

    if err != nil {
	log.Fatal("Error loading .env file")
    }
	const elevyPercentage = 0.01;

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"),os.Getenv("SLACK_APP_TOKEN"));

	go printCommandEvents(bot.CommandEvents());

	bot.Command("My amount is {amount}",&slacker.CommandDefinition{
		Description: "E-Levy Calculator",
		Examples: []string{"Your e-levy is \"gh\" 10"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			amount := request.FloatParam("amount",0.0);
			var elevy float64;
			// var res string;
			if amount >= 100 {
				elevy = amount * elevyPercentage;
				r := fmt.Sprintf("your e-levy is %.2f",elevy);
			    response.Reply(r);
			} else {
				// elevy = amount;
			    response.Reply("You wont pay elevy");
			}
			
		},
	});

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	error := bot.Listen(ctx)
	if error != nil {
		log.Fatal(error)
	}
}