package nd

type Service interface {
	Open(string, int) error
	Dispense(qty byte) error
	Status() (string, error)
}

func NewService(repo Repository) (Service, error) {
	s := service{repo}
	return &s, nil
}

type service struct {
	hw Repository
}

func (s *service) Open(port string, speed int) error {
	err := s.hw.Connect(port, speed)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Dispense(qty byte) error {
	err := s.hw.PayNote(qty)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Status() (string, error) {
	id, err := s.hw.GetStatus()
	if err != nil {
		return "", err
	}
	return id, nil
}
