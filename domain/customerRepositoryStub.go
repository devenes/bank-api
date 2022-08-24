package domain

type CustomerRepositoryStub struct {
	Customers []Customer
}

func (s *CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.Customers, nil
}

func NewCustomerRepositoryStub() *CustomerRepositoryStub {
	return &CustomerRepositoryStub{
		Customers: []Customer{
			Customer{
				Id:          "1",
				Name:        "John Doe",
				City:        "New York",
				Zipcode:     "10001",
				DateOfBirth: "01/01/1990",
				Status:      "active",
			},
			Customer{
				Id:          "2",
				Name:        "Jane Doe",
				City:        "New York",
				Zipcode:     "10001",
				DateOfBirth: "01/01/1990",
				Status:      "active",
			},
		},
	}
}
