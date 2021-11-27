package questions

type Service interface {
	getAllQnAByParticipantID(participantID int) ([]*QuestionAndAnswer, error)
	addNewQnAToParticipant(participantInformation *QuestionAndAnswer) error
	deleteQuestionByID(id int) error
}

type service struct {
	storage Storage
}

func MustNewService(storage Storage) *service {
	return &service{
		storage: storage,
	}
}

func (c *service) getAllQnAByParticipantID(participantID int) ([]*QuestionAndAnswer, error) {
	allParticipants, err := c.storage.getAllQnAByParticipantID(participantID)
	if err != nil {
		return nil, err
	}
	return allParticipants, err
}

func (c *service) addNewQnAToParticipant(qnaInformation *QuestionAndAnswer) error {
	err := c.storage.addNewQnAToParticipant(qnaInformation)
	if err != nil {
		return err
	}
	return nil
}

func (c *service) deleteQuestionByID(id int) error {
	err := c.storage.deleteQuestionByID(id)
	if err != nil {
		return err
	}
	return nil
}
