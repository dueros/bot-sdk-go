package card

type LinkAccountCard struct {
	BaseCard
}

func NewLinkAccountCard() *LinkAccountCard {
	return &LinkAccountCard{BaseCard{
		Type: "LinkAccount",
	}}
}
