package oriflame

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	strip "github.com/grokify/html-strip-tags-go"
)

// This struct represents a product
type Product struct {
	Code                string  `json:"code"`
	Name                string  `json:"name"`
	BrandName           string  `json:"brandName"`
	BrandURL            string  `json:"brandUrl"`
	Rating              float64 `json:"rating"`
	TotalReviews        int     `json:"totalReviews"`
	IsRatingEnabled     bool    `json:"isRatingEnabled"`
	ShowSubscribeButton bool    `json:"showSubscribeButton"`
	Description         string  `json:"description"`
	IsMultiProduct      bool    `json:"isMultiProduct"`
	Products            []struct {
		Code              string      `json:"code"`
		ShadeName         interface{} `json:"shadeName"`
		Size              string      `json:"size"`
		CurrentPrice      string      `json:"currentPrice"`
		CurrentPriceValue float64     `json:"currentPriceValue"`
		OldPrice          string      `json:"oldPrice"`
		BusinessPoints    interface{} `json:"businessPoints"`
		Images            []struct {
			Sizes []struct {
				URL   string `json:"url"`
				Width int    `json:"width"`
			} `json:"sizes"`
		} `json:"images"`
		VideoURL           string        `json:"videoUrl"`
		LabelText          interface{}   `json:"labelText"`
		LabelCSSClass      interface{}   `json:"labelCssClass"`
		DealLabelText      interface{}   `json:"dealLabelText"`
		DealLabelCSSClass  interface{}   `json:"dealLabelCssClass"`
		HasReplacements    bool          `json:"hasReplacements"`
		ColorHexCodes      []interface{} `json:"colorHexCodes"`
		ColorImageURL      string        `json:"colorImageUrl"`
		IsOutOfStock       bool          `json:"isOutOfStock"`
		IsAvailable        bool          `json:"isAvailable"`
		CanBeAddedToBasket bool          `json:"canBeAddedToBasket"`
		CanBeReserved      bool          `json:"canBeReserved"`
		ProductInfo        string        `json:"productInfo"`
		Benefits           []string      `json:"benefits"`
		Ingredients        []struct {
			Name          string `json:"name"`
			IngredientURL string `json:"ingredientUrl"`
			Image         struct {
				Sizes []struct {
					URL   string `json:"url"`
					Width int    `json:"width"`
				} `json:"sizes"`
			} `json:"image"`
			BenefitSummary string `json:"benefitSummary"`
		} `json:"ingredients"`
		Sample                       interface{}   `json:"sample"`
		FullProductUrls              []interface{} `json:"fullProductUrls"`
		RelatedSet                   interface{}   `json:"relatedSet"`
		OlapicTags                   string        `json:"olapicTags"`
		Barcode                      string        `json:"barcode"`
		ShowNotifyMeSubscription     bool          `json:"showNotifyMeSubscription"`
		NextAvailableDateText        string        `json:"nextAvailableDateText"`
		NextAvailableDateTooltipText string        `json:"nextAvailableDateTooltipText"`
	} `json:"products"`
	HowToUse struct {
		Label string `json:"label"`
		Text  string `json:"text"`
	} `json:"howToUse"`
	About struct {
		Label string `json:"label"`
		Text  string `json:"text"`
	} `json:"about"`
	Ingredients struct {
		MainDescription struct {
			Label string `json:"label"`
			Text  string `json:"text"`
		} `json:"mainDescription"`
		DisplayWithAdditionalDescription bool `json:"displayWithAdditionalDescription"`
		AdditionalDescription            struct {
			Label string `json:"label"`
			Text  string `json:"text"`
		} `json:"additionalDescription"`
	} `json:"ingredients"`
	AdditionalInformation interface{} `json:"additionalInformation"`
}

// A method that returns a product.
func (client *Client) GetProduct(code string) (*Product, error) {
	url := fmt.Sprintf("https://id.oriflame.com/system/ajax/pdp/concept?code=%s", code)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	bodyString := strip.StripTags(string(body))

	dec := json.NewDecoder(strings.NewReader(bodyString))
	var product *Product
	if err := dec.Decode(&product); err != nil {
		return nil, err
	}

	return product, nil
}
