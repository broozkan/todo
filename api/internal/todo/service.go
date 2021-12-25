package todo

type (
	Service struct {
		client SwapiClient
	}

	SwapiClient interface {
		GetStarshipDetails(starshipName string) (string, error)
	}
)

func NewService() {

}
