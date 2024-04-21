package service

import (
	"advertising_service/internal/filter"
	"advertising_service/internal/models"
	"github.com/mssola/user_agent"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (se *Service) FindWinner(pars *user_agent.UserAgent, clientIP string, s *geoip2.Reader) (winner *models.Company) {

	// получаем информацию о браузере и его версию из спаршенного заголовка
	browserName, _ := pars.Browser()

	// получаем данные о местоположении пользователя исходя из публичного ip
	country, err := s.Country(net.ParseIP(clientIP))
	if err != nil {
		log.Fatal(err)
	}

	userEntity := models.User{
		Country: country.Country.IsoCode,
		Browser: browserName,
	}

	comps := filter.GetStaticCampaigns()

	winner = filter.MakeAuctione(comps, &userEntity)

	return winner
}
