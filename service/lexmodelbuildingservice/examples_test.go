// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexmodelbuildingservice_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/Beeketing/aws-sdk-go/aws"
	"github.com/Beeketing/aws-sdk-go/aws/awserr"
	"github.com/Beeketing/aws-sdk-go/aws/session"
	"github.com/Beeketing/aws-sdk-go/service/lexmodelbuildingservice"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// To get information about a bot
//
// This example shows how to get configuration information for a bot.
func ExampleLexModelBuildingService_GetBot_shared00() {
	svc := lexmodelbuildingservice.New(session.New())
	input := &lexmodelbuildingservice.GetBotInput{
		Name:           aws.String("DocOrderPizza"),
		VersionOrAlias: aws.String("$LATEST"),
	}

	result, err := svc.GetBot(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case lexmodelbuildingservice.ErrCodeNotFoundException:
				fmt.Println(lexmodelbuildingservice.ErrCodeNotFoundException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeLimitExceededException:
				fmt.Println(lexmodelbuildingservice.ErrCodeLimitExceededException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeInternalFailureException:
				fmt.Println(lexmodelbuildingservice.ErrCodeInternalFailureException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeBadRequestException:
				fmt.Println(lexmodelbuildingservice.ErrCodeBadRequestException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To get a list of bots
//
// This example shows how to get a list of all of the bots in your account.
func ExampleLexModelBuildingService_GetBots_shared00() {
	svc := lexmodelbuildingservice.New(session.New())
	input := &lexmodelbuildingservice.GetBotsInput{
		MaxResults: aws.Int64(5),
		NextToken:  aws.String(""),
	}

	result, err := svc.GetBots(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case lexmodelbuildingservice.ErrCodeNotFoundException:
				fmt.Println(lexmodelbuildingservice.ErrCodeNotFoundException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeLimitExceededException:
				fmt.Println(lexmodelbuildingservice.ErrCodeLimitExceededException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeInternalFailureException:
				fmt.Println(lexmodelbuildingservice.ErrCodeInternalFailureException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeBadRequestException:
				fmt.Println(lexmodelbuildingservice.ErrCodeBadRequestException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To get a information about an intent
//
// This example shows how to get information about an intent.
func ExampleLexModelBuildingService_GetIntent_shared00() {
	svc := lexmodelbuildingservice.New(session.New())
	input := &lexmodelbuildingservice.GetIntentInput{
		Name:    aws.String("DocOrderPizza"),
		Version: aws.String("$LATEST"),
	}

	result, err := svc.GetIntent(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case lexmodelbuildingservice.ErrCodeNotFoundException:
				fmt.Println(lexmodelbuildingservice.ErrCodeNotFoundException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeLimitExceededException:
				fmt.Println(lexmodelbuildingservice.ErrCodeLimitExceededException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeInternalFailureException:
				fmt.Println(lexmodelbuildingservice.ErrCodeInternalFailureException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeBadRequestException:
				fmt.Println(lexmodelbuildingservice.ErrCodeBadRequestException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To get a list of intents
//
// This example shows how to get a list of all of the intents in your account.
func ExampleLexModelBuildingService_GetIntents_shared00() {
	svc := lexmodelbuildingservice.New(session.New())
	input := &lexmodelbuildingservice.GetIntentsInput{
		MaxResults: aws.Int64(10),
		NextToken:  aws.String(""),
	}

	result, err := svc.GetIntents(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case lexmodelbuildingservice.ErrCodeNotFoundException:
				fmt.Println(lexmodelbuildingservice.ErrCodeNotFoundException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeLimitExceededException:
				fmt.Println(lexmodelbuildingservice.ErrCodeLimitExceededException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeInternalFailureException:
				fmt.Println(lexmodelbuildingservice.ErrCodeInternalFailureException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeBadRequestException:
				fmt.Println(lexmodelbuildingservice.ErrCodeBadRequestException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To get information about a slot type
//
// This example shows how to get information about a slot type.
func ExampleLexModelBuildingService_GetSlotType_shared00() {
	svc := lexmodelbuildingservice.New(session.New())
	input := &lexmodelbuildingservice.GetSlotTypeInput{
		Name:    aws.String("DocPizzaCrustType"),
		Version: aws.String("$LATEST"),
	}

	result, err := svc.GetSlotType(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case lexmodelbuildingservice.ErrCodeNotFoundException:
				fmt.Println(lexmodelbuildingservice.ErrCodeNotFoundException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeLimitExceededException:
				fmt.Println(lexmodelbuildingservice.ErrCodeLimitExceededException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeInternalFailureException:
				fmt.Println(lexmodelbuildingservice.ErrCodeInternalFailureException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeBadRequestException:
				fmt.Println(lexmodelbuildingservice.ErrCodeBadRequestException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To get a list of slot types
//
// This example shows how to get a list of all of the slot types in your account.
func ExampleLexModelBuildingService_GetSlotTypes_shared00() {
	svc := lexmodelbuildingservice.New(session.New())
	input := &lexmodelbuildingservice.GetSlotTypesInput{
		MaxResults: aws.Int64(10),
		NextToken:  aws.String(""),
	}

	result, err := svc.GetSlotTypes(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case lexmodelbuildingservice.ErrCodeNotFoundException:
				fmt.Println(lexmodelbuildingservice.ErrCodeNotFoundException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeLimitExceededException:
				fmt.Println(lexmodelbuildingservice.ErrCodeLimitExceededException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeInternalFailureException:
				fmt.Println(lexmodelbuildingservice.ErrCodeInternalFailureException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeBadRequestException:
				fmt.Println(lexmodelbuildingservice.ErrCodeBadRequestException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a bot
//
// This example shows how to create a bot for ordering pizzas.
func ExampleLexModelBuildingService_PutBot_shared00() {
	svc := lexmodelbuildingservice.New(session.New())
	input := &lexmodelbuildingservice.PutBotInput{
		AbortStatement: &lexmodelbuildingservice.Statement{
			Messages: []*lexmodelbuildingservice.Message{
				{
					Content:     aws.String("I don't understand. Can you try again?"),
					ContentType: aws.String("PlainText"),
				},
				{
					Content:     aws.String("I'm sorry, I don't understand."),
					ContentType: aws.String("PlainText"),
				},
			},
		},
		ChildDirected: aws.Bool(true),
		ClarificationPrompt: &lexmodelbuildingservice.Prompt{
			MaxAttempts: aws.Int64(1),
			Messages: []*lexmodelbuildingservice.Message{
				{
					Content:     aws.String("I'm sorry, I didn't hear that. Can you repeate what you just said?"),
					ContentType: aws.String("PlainText"),
				},
				{
					Content:     aws.String("Can you say that again?"),
					ContentType: aws.String("PlainText"),
				},
			},
		},
		Description:             aws.String("Orders a pizza from a local pizzeria."),
		IdleSessionTTLInSeconds: aws.Int64(300),
		Intents: []*lexmodelbuildingservice.Intent{
			{
				IntentName:    aws.String("DocOrderPizza"),
				IntentVersion: aws.String("$LATEST"),
			},
		},
		Locale:          aws.String("en-US"),
		Name:            aws.String("DocOrderPizzaBot"),
		ProcessBehavior: aws.String("SAVE"),
	}

	result, err := svc.PutBot(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case lexmodelbuildingservice.ErrCodeConflictException:
				fmt.Println(lexmodelbuildingservice.ErrCodeConflictException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeLimitExceededException:
				fmt.Println(lexmodelbuildingservice.ErrCodeLimitExceededException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeInternalFailureException:
				fmt.Println(lexmodelbuildingservice.ErrCodeInternalFailureException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeBadRequestException:
				fmt.Println(lexmodelbuildingservice.ErrCodeBadRequestException, aerr.Error())
			case lexmodelbuildingservice.ErrCodePreconditionFailedException:
				fmt.Println(lexmodelbuildingservice.ErrCodePreconditionFailedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create an intent
//
// This example shows how to create an intent for ordering pizzas.
func ExampleLexModelBuildingService_PutIntent_shared00() {
	svc := lexmodelbuildingservice.New(session.New())
	input := &lexmodelbuildingservice.PutIntentInput{
		ConclusionStatement: &lexmodelbuildingservice.Statement{
			Messages: []*lexmodelbuildingservice.Message{
				{
					Content:     aws.String("All right, I ordered  you a {Crust} crust {Type} pizza with {Sauce} sauce."),
					ContentType: aws.String("PlainText"),
				},
				{
					Content:     aws.String("OK, your {Crust} crust {Type} pizza with {Sauce} sauce is on the way."),
					ContentType: aws.String("PlainText"),
				},
			},
			ResponseCard: aws.String("foo"),
		},
		ConfirmationPrompt: &lexmodelbuildingservice.Prompt{
			MaxAttempts: aws.Int64(1),
			Messages: []*lexmodelbuildingservice.Message{
				{
					Content:     aws.String("Should I order  your {Crust} crust {Type} pizza with {Sauce} sauce?"),
					ContentType: aws.String("PlainText"),
				},
			},
		},
		Description: aws.String("Order a pizza from a local pizzeria."),
		FulfillmentActivity: &lexmodelbuildingservice.FulfillmentActivity{
			Type: aws.String("ReturnIntent"),
		},
		Name: aws.String("DocOrderPizza"),
		RejectionStatement: &lexmodelbuildingservice.Statement{
			Messages: []*lexmodelbuildingservice.Message{
				{
					Content:     aws.String("Ok, I'll cancel your order."),
					ContentType: aws.String("PlainText"),
				},
				{
					Content:     aws.String("I cancelled your order."),
					ContentType: aws.String("PlainText"),
				},
			},
		},
		SampleUtterances: []*string{
			aws.String("Order me a pizza."),
			aws.String("Order me a {Type} pizza."),
			aws.String("I want a {Crust} crust {Type} pizza"),
			aws.String("I want a {Crust} crust {Type} pizza with {Sauce} sauce."),
		},
		Slots: []*lexmodelbuildingservice.Slot{
			{
				Description: aws.String("The type of pizza to order."),
				Name:        aws.String("Type"),
				Priority:    aws.Int64(1),
				SampleUtterances: []*string{
					aws.String("Get me a {Type} pizza."),
					aws.String("A {Type} pizza please."),
					aws.String("I'd like a {Type} pizza."),
				},
				SlotConstraint:  aws.String("Required"),
				SlotType:        aws.String("DocPizzaType"),
				SlotTypeVersion: aws.String("$LATEST"),
				ValueElicitationPrompt: &lexmodelbuildingservice.Prompt{
					MaxAttempts: aws.Int64(1),
					Messages: []*lexmodelbuildingservice.Message{
						{
							Content:     aws.String("What type of pizza would you like?"),
							ContentType: aws.String("PlainText"),
						},
						{
							Content:     aws.String("Vegie or cheese pizza?"),
							ContentType: aws.String("PlainText"),
						},
						{
							Content:     aws.String("I can get you a vegie or a cheese pizza."),
							ContentType: aws.String("PlainText"),
						},
					},
				},
			},
			{
				Description: aws.String("The type of pizza crust to order."),
				Name:        aws.String("Crust"),
				Priority:    aws.Int64(2),
				SampleUtterances: []*string{
					aws.String("Make it a {Crust} crust."),
					aws.String("I'd like a {Crust} crust."),
				},
				SlotConstraint:  aws.String("Required"),
				SlotType:        aws.String("DocPizzaCrustType"),
				SlotTypeVersion: aws.String("$LATEST"),
				ValueElicitationPrompt: &lexmodelbuildingservice.Prompt{
					MaxAttempts: aws.Int64(1),
					Messages: []*lexmodelbuildingservice.Message{
						{
							Content:     aws.String("What type of crust would you like?"),
							ContentType: aws.String("PlainText"),
						},
						{
							Content:     aws.String("Thick or thin crust?"),
							ContentType: aws.String("PlainText"),
						},
					},
				},
			},
			{
				Description: aws.String("The type of sauce to use on the pizza."),
				Name:        aws.String("Sauce"),
				Priority:    aws.Int64(3),
				SampleUtterances: []*string{
					aws.String("Make it {Sauce} sauce."),
					aws.String("I'd like {Sauce} sauce."),
				},
				SlotConstraint:  aws.String("Required"),
				SlotType:        aws.String("DocPizzaSauceType"),
				SlotTypeVersion: aws.String("$LATEST"),
				ValueElicitationPrompt: &lexmodelbuildingservice.Prompt{
					MaxAttempts: aws.Int64(1),
					Messages: []*lexmodelbuildingservice.Message{
						{
							Content:     aws.String("White or red sauce?"),
							ContentType: aws.String("PlainText"),
						},
						{
							Content:     aws.String("Garlic or tomato sauce?"),
							ContentType: aws.String("PlainText"),
						},
					},
				},
			},
		},
	}

	result, err := svc.PutIntent(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case lexmodelbuildingservice.ErrCodeConflictException:
				fmt.Println(lexmodelbuildingservice.ErrCodeConflictException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeLimitExceededException:
				fmt.Println(lexmodelbuildingservice.ErrCodeLimitExceededException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeInternalFailureException:
				fmt.Println(lexmodelbuildingservice.ErrCodeInternalFailureException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeBadRequestException:
				fmt.Println(lexmodelbuildingservice.ErrCodeBadRequestException, aerr.Error())
			case lexmodelbuildingservice.ErrCodePreconditionFailedException:
				fmt.Println(lexmodelbuildingservice.ErrCodePreconditionFailedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To Create a Slot Type
//
// This example shows how to create a slot type that describes pizza sauces.
func ExampleLexModelBuildingService_PutSlotType_shared00() {
	svc := lexmodelbuildingservice.New(session.New())
	input := &lexmodelbuildingservice.PutSlotTypeInput{
		Description: aws.String("Available pizza sauces"),
		EnumerationValues: []*lexmodelbuildingservice.EnumerationValue{
			{
				Value: aws.String("red"),
			},
			{
				Value: aws.String("white"),
			},
		},
		Name: aws.String("PizzaSauceType"),
	}

	result, err := svc.PutSlotType(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case lexmodelbuildingservice.ErrCodeConflictException:
				fmt.Println(lexmodelbuildingservice.ErrCodeConflictException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeLimitExceededException:
				fmt.Println(lexmodelbuildingservice.ErrCodeLimitExceededException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeInternalFailureException:
				fmt.Println(lexmodelbuildingservice.ErrCodeInternalFailureException, aerr.Error())
			case lexmodelbuildingservice.ErrCodeBadRequestException:
				fmt.Println(lexmodelbuildingservice.ErrCodeBadRequestException, aerr.Error())
			case lexmodelbuildingservice.ErrCodePreconditionFailedException:
				fmt.Println(lexmodelbuildingservice.ErrCodePreconditionFailedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
