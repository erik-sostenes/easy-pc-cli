package scraping

import (
	"github.com/erik-sostenes/easy-pc-cli/internal/context/offer/business/domain"
	"github.com/erik-sostenes/easy-pc-cli/internal/context/offer/business/domain/ports"
	"github.com/gocolly/colly"
	"github.com/google/uuid"
	"strconv"
	"strings"
)

var _ ports.OfferScraper = &offerScraper{}

// offerScraper represents an output adapter in charge of scraping offers
type offerScraper struct {
	colly.Collector
}

// NewOfferScraper returns an instances of ports.OfferScraper interface
func NewOfferScraper(collector colly.Collector) ports.OfferScraper {
	return &offerScraper{
		Collector: collector,
	}
}

// Scraping method that searches for an offer by scraping
func (o *offerScraper) Scraping(values map[string]string, urls []string) (offers domain.Offers, err error) {
	o.OnHTML(".items_container", func(htmlElement *colly.HTMLElement) {
		htmlElement.ForEach("li", func(_ int, h *colly.HTMLElement) {
			originalPrice, _ := strconv.ParseFloat(h.ChildText(".andes-money-amount-combo__previous-value > .andes-money-amount__fraction"), 64)
			discountPrice, _ := strconv.ParseFloat(h.ChildText(".andes-money-amount--cents-superscript > .andes-money-amount__fraction"), 64)

			var isOfferDay bool
			if strings.TrimSpace(h.ChildText(".promotion-item__today-offer-text")) != "" {
				isOfferDay = true
			}

			var isAvailable bool
			if strings.TrimSpace(h.ChildText(".promotion-item__item-lightning-status > span")) == "" {
				isAvailable = true
			}
			u, _ := uuid.NewRandom()
			offer := domain.Offer{
				Website:            values["website"],
				Id:                 u.String(),
				Title:              h.ChildText(".promotion-item__title"),
				OriginalPrice:      originalPrice,
				DiscountPrice:      discountPrice,
				DiscountPercentage: h.ChildText("p[class=andes-money-amount-combo] > span[class=andes-money-amount-amount__discount]"),
				OfferUrl:           h.ChildAttr(".promotion-item__link-container", "href"),
				IsOfferDay:         isOfferDay,
				IsAvailable:        isAvailable,
				DeliveryIsFree:     h.ChildText(".promotion-item__next-day-text"),
				Category: domain.Category{
					Id:   values["category_id"],
					Name: values["category"],
				},
			}

			offers = append(offers, offer)
		})
	})

	for _, url := range urls {
		_ = o.Visit(url)
	}

	o.Wait()

	return
}
