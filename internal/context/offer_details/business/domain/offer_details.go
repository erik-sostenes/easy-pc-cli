package domain

type (
	Sale struct {
		Days   string
		Amount string
	}

	Rating struct {
		Rating float64
		Amount string
	}

	// OfferDetails represents an object domain
	OfferDetails struct {
		Id     string
		Image  string
		Sale   Sale
		Rating Rating
	}

	OffersDetails []OfferDetails
)
