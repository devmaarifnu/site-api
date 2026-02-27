package services

import (
	"lpmaarifnu-site-api/internal/repositories"
)

type SettingService interface {
	GetPublicSettings() (map[string]interface{}, error)
}

type settingService struct {
	settingRepo repositories.SettingRepository
}

func NewSettingService(settingRepo repositories.SettingRepository) SettingService {
	return &settingService{settingRepo: settingRepo}
}

func (s *settingService) GetPublicSettings() (map[string]interface{}, error) {
	settings, err := s.settingRepo.FindPublicSettings()
	if err != nil {
		return nil, err
	}

	// Convert to map structure
	result := make(map[string]interface{})
	result["site_name"] = ""
	result["site_description"] = ""
	result["logo"] = ""
	result["favicon"] = ""

	contact := make(map[string]string)
	socialMedia := make(map[string]string)

	for _, setting := range settings {
		switch setting.SettingKey {
		case "site_name":
			result["site_name"] = setting.SettingValue
		case "site_description":
			result["site_description"] = setting.SettingValue
		case "contact_email":
			contact["email"] = setting.SettingValue
		case "contact_phone":
			contact["phone"] = setting.SettingValue
		case "contact_address":
			contact["address"] = setting.SettingValue
		case "social_facebook":
			socialMedia["facebook"] = setting.SettingValue
		case "social_twitter":
			socialMedia["twitter"] = setting.SettingValue
		case "social_instagram":
			socialMedia["instagram"] = setting.SettingValue
		case "social_youtube":
			socialMedia["youtube"] = setting.SettingValue
		case "site_logo":
			result["logo"] = setting.SettingValue
		case "site_favicon":
			result["favicon"] = setting.SettingValue
		}
	}

	result["contact"] = contact
	result["social_media"] = socialMedia

	return result, nil
}
