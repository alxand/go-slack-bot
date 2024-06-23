package main

import (
	"context"
	"log"
	"os"

	"github.com/slack-io/slacker"
)

//	func printCommand(analyticsChannel <-chan *slacker.CommandEvents) {
//		for event := range analyticsChannel {
//			fmt.Println("Command events")
//			fmt.Println(event.Timestamp)
//			fmt.Println(event.Command)
//			fmt.Println(event.Parameters)
//			fmt.Println(event.Event)
//		}
//	}
func main() {
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	// go printCommand(bot.CommandEvents())
	bot.AddCommand(&slacker.CommandDefinition{
		Command: "ping",
		Handler: func(ctx *slacker.CommandContext) {
			// year := request.Param("year")
			// yob, err := strconv.Atoi(year)
			// if err != nil {
			// 	fmt.Println("error occured")
			// }
			// age := 2024 - yob
			// r := fmt.Sprintf("age is %d", age)
			// response.Reply(r)
			ctx.Response().Reply("Oh okk")
		},
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
