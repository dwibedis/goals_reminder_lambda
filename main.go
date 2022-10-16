package goals_reminder_lambda

import (
	"dwibedis/goal_reminder_lambda/internal/services"
	"fmt"
)

func main() {
	g := services.GoalsCalculator{}
	fmt.Println(g.GetGoal())
}