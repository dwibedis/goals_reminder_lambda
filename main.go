package goals_reminder_lambda

import (
	"dwibedis/goal_reminder_lambda/internal/controllers"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(controllers.HandleURLShortenRequest)
}