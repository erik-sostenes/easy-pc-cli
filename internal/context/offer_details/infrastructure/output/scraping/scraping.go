package scraping

import (
	"github.com/erik-sostenes/easy-pc-cli/internal/context/offer_details/business/domain"
	"github.com/erik-sostenes/easy-pc-cli/internal/context/offer_details/business/domain/ports"
	"github.com/gocolly/colly"
	"strconv"
)

// offerDetailsScraper represents an output adapter in charge of scraping offer details
type offerDetailsScraper struct {
	colly.Collector
}

// NewOfferDetailsScraper returns an instances of ports.OfferDetailsScraper interface
func NewOfferDetailsScraper(collector colly.Collector) ports.OfferDetailsScraper {
	return &offerDetailsScraper{
		Collector: collector,
	}
}

// Scraping method that searches for an offer details by scraping
func (o *offerDetailsScraper) Scraping(url, offerDetailsId string) (offerDetails domain.OfferDetails, err error) {
	offerDetails.Id = offerDetailsId

	o.OnHTML("ul[class=ui-pdp-seller__list-description] li:first-child", func(element *colly.HTMLElement) {
		offerDetails.Sale.Days = element.ChildText(".ui-pdp-seller__text-description")
		offerDetails.Sale.Amount = element.ChildText(".ui-pdp-seller__sales-description")
	})

	o.OnHTML("div[class=ui-review-capability__rating]", func(element *colly.HTMLElement) {
		offerDetails.Rating.Rating, _ = strconv.ParseFloat(element.ChildText(".ui-review-capability__rating__average--desktop"), 64)
		offerDetails.Rating.Amount = element.ChildText("p[class=ui-review-capability__rating__label]")
	})

	o.OnHTML("div[class=ui-pdp-gallery] > .ui-pdp-gallery__column", func(h *colly.HTMLElement) {
		offerDetails.Image = h.ChildAttr(".ui-pdp-gallery__wrapper > figure > img", "src")
	})

	for _, url := range []string{url} {
		_ = o.Visit(url)
	}

	o.Wait()

	return offerDetails, nil
}
