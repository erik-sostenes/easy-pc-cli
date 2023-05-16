package scraping

import (
	"github.com/erik-sostenes/easy-pc-cli/internal/context/offer/business/domain"
	"github.com/erik-sostenes/easy-pc-cli/internal/context/offer/business/domain/ports"
	"github.com/gocolly/colly"
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
func (o *offerScraper) Scraping(query map[string]string, urls []string) (offers domain.Offers, err error) {
	o.OnHTML(query["container"], func(htmlElement *colly.HTMLElement) {
		htmlElement.ForEach(query["item"], func(_ int, h *colly.HTMLElement) {
			originalPrice, _ := strconv.ParseFloat(h.ChildText(query["original_price"]), 64)
			discountPrice, _ := strconv.ParseFloat(h.ChildText(query["discounted"]), 64)

			var isOfferDay bool
			if strings.TrimSpace(h.ChildText(query["offer_day"])) != "" {
				isOfferDay = true
			}

			var isAvailable bool
			if strings.TrimSpace(h.ChildText(query["available"])) == "" {
				isAvailable = true
			}

			offer := domain.Offer{
				Title:              h.ChildText(query["title"]),
				OriginalPrice:      originalPrice,
				DiscountPrice:      discountPrice,
				DiscountPercentage: h.ChildText(query["percentage"]),
				OfferUrl:           h.ChildAttr(query["offer_url"], "href"),
				IsOfferDay:         isOfferDay,
				IsAvailable:        isAvailable,
				DeliveryIsFree:     h.ChildText(query["delivery"]),
				Category: domain.Category{
					Id:   query["category_id"],
					Name: query["category"],
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
