package app

import (
	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/application/commands"
	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/application/queris"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	ApproveTrainingReschedule commands.ApproveTrainingRescheduleHandler
	CancelTraining            commands.CancelTrainingHandler
	RejectTrainingReschedule  commands.RejectTrainingRescheduleHandler
	RescheduleTraining        commands.RescheduleTrainingHandler
	RequestTrainingReschedule commands.RequestTrainingRescheduleHandler
	ScheduleTraining          commands.ScheduleTrainingHandler
}

type Queries struct {
	AllTrainings     queris.AllTrainingsHandler
	TrainingsForUser queris.TrainingsForUserHandler
}
