package participants

type Service interface {
	getAllParticipants() ([]*Participant, error)
	addParticipants(participantInformation *Participant) error
	deleteParticipant(id int) error
}

type service struct {
	storage Storage
}

func MustNewService(storage Storage) *service {
	return &service{
		storage: storage,
	}
}

func (c *service) getAllParticipants() ([]*Participant, error) {
	allParticipants, err := c.storage.GetAllParticipants()
	if err != nil {
		return nil, err
	}
	return allParticipants, err
}

func (c *service) addParticipants(participantInformation *Participant) error {
	err := c.storage.AddParticipant(participantInformation)
	if err != nil {
		return err
	}
	return nil
}

func (c *service) deleteParticipant(id int) error {
	err := c.storage.deleteParticipant(id)
	if err != nil {
		return err
	}
	return nil
}
