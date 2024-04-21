package filter

import (
	"advertising_service/internal/models"
	"sort"
)

type filterFunc func(in []*models.Company, u *models.User) []*models.Company

var (
	filters = []filterFunc{
		filterByBrowser,
		filterByCountry,
	}
)

func MakeAuctione(in []*models.Company, u *models.User) (winner *models.Company) {
	// производим копирование для дальнейшего изменения
	company := make([]*models.Company, len(in))
	copy(company, in)

	for _, f := range filters {
		company = f(company, u)
	}

	if len(company) == 0 {
		return nil
	}

	sort.Slice(company, func(i, j int) bool {
		return company[i].Price > company[j].Price
	})

	return company[0]
}

func filterByBrowser(in []*models.Company, u *models.User) []*models.Company {
	// идем по массиву компаний и проверяим их с пользовательскими данными
	for i := len(in) - 1; i >= 0; i-- {
		if len(in[i].Targeting.Browser) == 0 {
			// если у компании нет предпочтений по бразуеру то ее не исключаем
			continue
		}

		if u.Browser == in[i].Targeting.Browser {
			continue
		}

		// если компания не подходит по таргетинкам мы ее удаляем из списка
		in[i] = in[0]
		in = in[1:]
	}

	return in
}

func filterByCountry(in []*models.Company, u *models.User) []*models.Company {

	for i := len(in) - 1; i >= 0; i-- {

		if len(in[i].Targeting.Country) == 0 {
			continue
		}

		if u.Country == in[i].Targeting.Country {
			continue
		}

		in[i] = in[0]
		in = in[1:]
	}

	return in
}

func GetStaticCampaigns() []*models.Company {
	return []*models.Company{
		{
			Price: 1,
			Targeting: models.Targeting{
				Country: "RU",
				Browser: "Chrome",
			},
			ClickURL: "https://yandex.ru",
		},
		{
			Price: 1,
			Targeting: models.Targeting{
				Country: "DE",
				Browser: "Chrome",
			},
			ClickURL: "https://google.com",
		},
		{
			Price:     1,
			Targeting: models.Targeting{},
			ClickURL:  "https://duckduckgo.com",
		},
	}
}
