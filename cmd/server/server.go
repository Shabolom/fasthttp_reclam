package adr

import (
	"advertising_service/internal/service"
	"github.com/ferluci/fast-realip"
	"github.com/mssola/user_agent"
	"github.com/oschwald/geoip2-golang"
	"github.com/valyala/fasthttp"
	"net/http"
)

type Server struct {
	geoip *geoip2.Reader
}

func NewServer(reader *geoip2.Reader) *Server {
	return &Server{geoip: reader}
}

var ser = service.NewService()

func (s *Server) Listen() error {
	return fasthttp.ListenAndServe(":8080", s.handleHttp)
}

func (s *Server) handleHttp(ctx *fasthttp.RequestCtx) {
	// получаем публичный ip адрес клиента
	clientIP := realip.FromRequest(ctx)
	// получаем строку данные из заголовка userAgent и переводим их из байт в строку они нужны для парсинга и получения данных о пользователе
	ua := string(ctx.Request.Header.UserAgent())

	// парси полученные данные
	parsed := user_agent.New(ua)

	winner := ser.FindWinner(parsed, clientIP, s.geoip)

	if winner == nil {
		ctx.Redirect("https://example.com", http.StatusSeeOther)
		return
	}
	ctx.Redirect(winner.ClickURL, http.StatusSeeOther)

	//ctx.WriteString(fmt.Sprintf("Name Browser: %s v.%s \n", browserName, browserVersion))
	//ctx.WriteString(fmt.Sprintf("OS: %s v.%s \n", parsed.OS(), parsed.OSInfo().Version))
	//ctx.WriteString(fmt.Sprintf("Platform %s \n", parsed.Platform()))
	//ctx.WriteString(fmt.Sprintf("IP %s, client country: %s", clientIP, country.Country.IsoCode))
}
