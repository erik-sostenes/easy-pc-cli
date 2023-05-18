package domain

type (
	Category struct {
		Id   string
		Name string
	}

	// Offer represents an object domain
	Offer struct {
		Website            string
		Id                 string
		Title              string
		OriginalPrice      float64
		DiscountPrice      float64
		DiscountPercentage string
		OfferUrl           string
		IsOfferDay         bool
		IsAvailable        bool
		DeliveryIsFree     string
		Opinions           []string
		Category           Category
	}
	Offers []Offer
)
