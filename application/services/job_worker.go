package services

import (
	"github.com/Andreis3/encoder-video-golang/domain"
	"github.com/Andreis3/encoder-video-golang/framework/utils"
	"github.com/streadway/amqp"
)

type JobWorkerResult struct {
	Job     domain.Job
	Message *amqp.Delivery
	Error   error
}

func JobWorker(messageChanel chan amqp.Delivery, returnChan chan JobWorkerResult, jobService JobService, workerID int) {
	// pega msg de body do json
	// validar se o json é um json. Se é valido o json
	// validar o video
	// inserir no banco de dados
	// start
	for message := range messageChanel {
		err := utils.IsJson(string(message.Body))
		if err != nil {
			returnChan <- returnJobResult(domain.Job{}, message, err)
		}
	}
}

func returnJobResult(job domain.Job, message amqp.Delivery, err error) JobWorkerResult {
	result := JobWorkerResult{
		Job:     job,
		Message: &message,
		Error:   err,
	}

	return result
}
