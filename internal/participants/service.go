package participants

type Service interface {
	getAllParticipants() ([]*Participant, error)
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
	return c.storage.getAllParticipants()
}
