package service

import (
	"github.com/hellyab/techreview/answer"
	"github.com/hellyab/techreview/entities"
)

//AnswerService implements answer.Service
type AnswerService struct {
	answerRepo answer.AnswerRepository
}

// NewAnswerService returns a new AnswerService object
func NewAnswerService(ansRepo answer.AnswerRepository) answer.AnswerService {
	return &AnswerService{answerRepo: ansRepo}
}

//Answers returns all stored Answers
func (as *AnswerService) Answers() ([]entities.Answer, []error) {
	qstns, errs := as.answerRepo.Answers()
	if len(errs) > 0 {
		return nil, errs
	}
	return qstns, errs
}

//Answer retrieves stored answer given its id
func (as *AnswerService) Answer(id string) (*entities.Answer, []error) {
	qstn, errs := as.answerRepo.Answer(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return qstn, errs
}

//UpdateAnswer updates a given answer
func (as *AnswerService) UpdateAnswer(answer *entities.Answer) (*entities.Answer, []error) {
	qstn, errs := as.answerRepo.UpdateAnswer(answer)
	if len(errs) > 0 {
		return nil, errs
	}
	return qstn, errs
}

//DeleteAnswer deletes a given answer
func (as *AnswerService) DeleteAnswer(id string) (*entities.Answer, []error) {
	qstn, errs := as.answerRepo.DeleteAnswer(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return qstn, errs
}

//StoreAnswer stores a given answer
func (as *AnswerService) StoreAnswer(answer *entities.Answer) (*entities.Answer, []error) {
	qstn, errs := as.answerRepo.StoreAnswer(answer)
	if len(errs) > 0 {
		return nil, errs
	}
	//changed
	return qstn, nil
}

func (as *AnswerService) AnswersByQuestionId(questionId string) ([]entities.AnswersByQuesId, []error) {
	answers, errs := as.answerRepo.AnswersByQuestionId(questionId)
	if len(errs) > 0 {
		return nil, errs
	}
	return answers, nil
}


func (as *AnswerService) UpVoteAnswer(answerUpvote *entities.AnswerUpvote){
	as.answerRepo.UpVoteAnswer(answerUpvote)
}

func (as *AnswerService) UpVoteCount(answerId string) int {
	return as.answerRepo.UpVoteCount(answerId)
}