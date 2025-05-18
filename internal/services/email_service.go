package services

import (
	"fmt"
	"net/smtp"

	"github.com/ihorlenko/weather_notifier/internal/config"
)

type EmailService struct {
	from     string
	password string
	smtpHost string
	smtpPort string
	baseURL  string
}

func NewEmailService(cfg *config.Config) *EmailService {
	return &EmailService{
		from:     cfg.EmailConfig.From,
		password: cfg.EmailConfig.Password,
		smtpHost: cfg.EmailConfig.SMTPHost,
		smtpPort: cfg.EmailConfig.SMTPPort,
		baseURL:  cfg.AppConfig.BaseURL,
	}
}

func (s *EmailService) SendConfirmationEmail(email, city, token string) error {
	subject := "Confirming subscription to Weather Notifier"
	confirmURL := fmt.Sprintf("%s/api/confirm/%s", s.baseURL, token)

	body := fmt.Sprintf(`
Hello!

You have subscribed to the weather notifications for location: %s.

To confirm your subscription visit the link:
%s

If it wasn't you, just ignore this letter.

Sincerely yours,
Weather Notifier Team
`, city, confirmURL)

	return s.sendEmail(email, subject, body)
}

func (s *EmailService) SendWeatherUpdate(email, city string, weather *WeatherData, unsubscribeToken string) error {
	subject := fmt.Sprintf("Weather in %s", city)
	unsubscribeURL := fmt.Sprintf("%s/api/unsubscribe/%s", s.baseURL, unsubscribeToken)

	body := fmt.Sprintf(`
Hello!

Current weather in the city %s:
Temperature: %.1f°C
Humidity: %.2f%%
Description: %s

To unsubscribe from notificetions visit the link:
%s

Kind regards,
Weather Notifier Team
`, city, weather.Temperature, weather.Humidity, weather.Description, unsubscribeURL)

	return s.sendEmail(email, subject, body)
}

func (s *EmailService) sendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", s.from, s.password, s.smtpHost)

	msg := fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", to, subject, body)

	addr := fmt.Sprintf("%s:%s", s.smtpHost, s.smtpPort)

	return smtp.SendMail(addr, auth, s.from, []string{to}, []byte(msg))
}
