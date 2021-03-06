package repository

import (
	"fmt"
	"github.com/hellyab/techreview/answer"
	"github.com/hellyab/techreview/entities"
	"github.com/jinzhu/gorm"
)

//AnswerGormRepo implements answer.AnswerRepository interface
type AnswerGormRepo struct {
	conn *gorm.DB
}

//NewAnswerGormRepo returns new object of AnswerGormRepo
func NewAnswerGormRepo(db *gorm.DB) answer.AnswerRepository{
	return &AnswerGormRepo{conn: db}
}

//Answers returns all user answers stored in the database
func (ansRepo *AnswerGormRepo) Answers() ([]entities.Answer, []error) {
	ans := []entities.Answer{}
	errs := ansRepo.conn.Find(&ans).GetErrors()
	if len(errs) > 0 {
		
		return nil, errs
	}
	return ans, errs
}

//Answer returns a user answer stored in the database which has the given id
func (ansRepo *AnswerGormRepo) Answer(id string) (*entities.Answer, []error) {
	qstn := entities.Answer{}
	errs := ansRepo.conn.Where("id = ?", id).First(&qstn).GetErrors()
	// errs := ansRepo.conn.First(&qstn, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &qstn, errs
}

//UpdateAnswer updates a given answer in the database
func (ansRepo *AnswerGormRepo) UpdateAnswer(answer *entities.Answer) (*entities.Answer, []error) {
	qstn := answer
	errs := ansRepo.conn.Save(qstn).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return qstn, errs
}

//DeleteAnswer deletes a answer with a given id from the database
func (ansRepo *AnswerGormRepo) DeleteAnswer(id string) (*entities.Answer, []error) {
	qstn, errs := ansRepo.Answer(id)
	if len(errs) > 0 {
		return nil, errs
	}
	// errs := ansRepo.conn.Where("id = ?", id).First(&qstn).GetErrors()
	errs = ansRepo.conn.Delete(qstn).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return qstn, errs
}

//StoreAnswer stores a given answer in the database
func (ansRepo *AnswerGormRepo) StoreAnswer(answer *entities.Answer) (*entities.Answer, []error) {
	qstn := answer
	errs := ansRepo.conn.Create(qstn).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return qstn, errs
}

func (ansRepo *AnswerGormRepo) AnswersByQuestionId(questionId string) ([]entities.AnswersByQuesId, []error){
	answs := []entities.Answer{}
	user := entities.User{}
	ansByQ := entities.AnswersByQuesId{}

	ansByQuestionArray := []entities.AnswersByQuesId{}

	errs := ansRepo.conn.Where("question_id = ?", questionId). Find(&answs).GetErrors()
	if len(errs) > 0 {
		fmt.Println("errors fetching answer")
		return nil, errs
	}

	for _, ans := range answs{
		errsForUser := ansRepo.conn.Where("id = ?", ans.ReplierID). First(&user).GetErrors()

		if len(errsForUser) > 0 {
			fmt.Println("errors fetching the user")
			return nil, errsForUser
		}
		// test
		ansByQ.AnsweredByFirstName = user.FirstName
		ansByQ.AnsweredByUserName = user.Username
		ansByQ.AnsweredByLastName = user.LastName
		ansByQ.Votes = int(ans.Votes)
		ansByQ.Answer = ans.Answer
		ansByQ.AnswerID = ans.ID

		ansByQuestionArray = append(ansByQuestionArray, ansByQ)
	}


	return ansByQuestionArray, nil

}

func (ansRepo *AnswerGormRepo) UpVoteAnswer(answerUpvote *entities.AnswerUpvote) {

	errs := ansRepo.conn.Where("answer_id = ? AND user_id= ?", answerUpvote.AnswerID, answerUpvote.UserID).First(&answerUpvote).GetErrors()
	//fmt.Println("the found answerUpvote form", ansUpvote)
	if len(errs) > 0 {

		errs := ansRepo.conn.Create(answerUpvote).GetErrors()
		if len(errs) > 0 {
			fmt.Println("error storing the answer upvote")
			return
		}
		fmt.Println("stored  answer upvote")
		fmt.Println("query worked")
		return

	} else{
		errs := ansRepo.conn.Delete(answerUpvote).GetErrors()
		if len(errs) > 0 {
			fmt.Println("error while deleting answer upvote", errs)
			return
		}
		fmt.Println("deleted answer upvote succuessfully")
		return
	}
}

func (ansRepo *AnswerGormRepo) UpVoteCount(answerId string) int {
	var count int
	var questionFollows entities.AnswerUpvote
	//questionFollows := entities.QuestionFollow{}
	ansRepo.conn.Model(&questionFollows).Where("answer_id = ?", answerId).Count(&count)
	// update the count value in answers table also
	ansRepo.conn.Model(&entities.Answer{}).Where("id = ?", answerId).UpdateColumn("votes", count)

	fmt.Println("counted succesfully, with value: ", count)

	return count
}



