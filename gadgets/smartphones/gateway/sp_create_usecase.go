package gateway

import "phones-review/gadgets/smartphones/models"

func (s *SmartphoneCreateGtw) Create(cmd *models.CreateSmartphoneCMD) (*models.Smartphone, error) {
	return s.create(cmd)
}
