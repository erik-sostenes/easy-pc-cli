package domain

type (
	Sale struct {
		Days   string
		Amount string
	}

	Category struct {
		Id   string
		Name string
	}

	Rating struct {
		Rating float32
		Amount float32
	}

	RatingOpinion struct {
		Rating string
		Amount float64
	}

	// Offer represents an object domain
	Offer struct {
		Title              string
		OriginalPrice      float64
		DiscountPrice      float64
		DiscountPercentage string
		OfferUrl           string
		IsOfferDay         bool
		IsAvailable        bool
		DeliveryIsFree     string
		Sale               Sale
		Rating             Rating
		RatingOpinions     []RatingOpinion
		Opinions           []string
		Category           Category
	}
	Offers []Offer
)
