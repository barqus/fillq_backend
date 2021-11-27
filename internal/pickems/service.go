package pickems

import "sort"

type Service interface {
	getUsersPickems(id int) ([]*PickEm, error)
	createUsersPickems(usersPickems []*PickEm,userId int) error
	deleteUsersAllPickems(id int) error
}

type service struct {
	storage Storage
}

func MustNewService(storage Storage) *service {
	return &service{
		storage: storage,
	}
}

func (c *service) getUsersPickems(id int) ([]*PickEm, error) {
	usersPickems, err := c.storage.getUsersPickems(id)

	sort.Slice(usersPickems, func(i, j int) bool {
		return usersPickems[i].Position < usersPickems[j].Position
	})



	if err != nil {
		return nil, err
	}
	return usersPickems, err
}

func (c *service) createUsersPickems(usersPickems []*PickEm,userId int) error{
	err := c.storage.createUsersPickems(usersPickems)
	if err != nil {
		return err
	}
	return nil
}

func (c *service) deleteUsersAllPickems(userID int) error {
	err := c.storage.deleteUsersAllPickems(userID)
	if err != nil {
		return err
	}
	return nil
}