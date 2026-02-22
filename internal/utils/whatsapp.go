package utils

import (
	"fmt"
	"net/url"
)

// GenerateWhatsAppLink génère un lien WhatsApp avec message pré-rempli
func GenerateWhatsAppLink(phoneNumber, productName string, price float64) string {
	// Format du numéro: +33612345678 (sans espaces ni tirets)
	// Message pré-rempli
	message := fmt.Sprintf("Bonjour, je suis intéressé par %s au prix de %.2f€", productName, price)
	encodedMessage := url.QueryEscape(message)

	// Format WhatsApp: https://wa.me/NUMERO?text=MESSAGE
	return fmt.Sprintf("https://wa.me/%s?text=%s", phoneNumber, encodedMessage)
}
