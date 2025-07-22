package security

import (
	"encoding/json"
	"errors"
	"trust-credit-back/models"

	"github.com/skip2/go-qrcode"
)

func GenerateQRCode(ref models.RegForm) ([]byte, error) {
	data, err := json.Marshal(ref)
	if err != nil {
		return nil, errors.New("failed to serialize JSON form")
	}
	
	json_str := string(data)

	qr_code, err := qrcode.Encode(json_str, qrcode.Medium, 250)
	if err != nil {
		return nil, errors.New("coudn't generate the qr code")
	}

	return qr_code, nil
}