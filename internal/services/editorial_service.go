package services

import (
	"lpmaarifnu-site-api/internal/models"
	"lpmaarifnu-site-api/internal/repositories"
)

type EditorialService interface {
	GetEditorialTeam() (map[string]interface{}, error)
}

type editorialService struct {
	editorialRepo repositories.EditorialRepository
}

func NewEditorialService(editorialRepo repositories.EditorialRepository) EditorialService {
	return &editorialService{editorialRepo: editorialRepo}
}

func (s *editorialService) GetEditorialTeam() (map[string]interface{}, error) {
	team, err := s.editorialRepo.FindAllTeam()
	if err != nil {
		return nil, err
	}

	council, err := s.editorialRepo.FindAllCouncil()
	if err != nil {
		return nil, err
	}

	// Group team members by role type
	var pemimpinRedaksi *models.EditorialTeam
	var wakilPemimpinRedaksi []models.EditorialTeam
	var redakturPelaksana *models.EditorialTeam
	var timRedaksi []models.EditorialTeam

	for i := range team {
		switch team[i].RoleType {
		case "pemimpin_redaksi":
			pemimpinRedaksi = &team[i]
		case "wakil_pemimpin_redaksi":
			wakilPemimpinRedaksi = append(wakilPemimpinRedaksi, team[i])
		case "redaktur_pelaksana":
			redakturPelaksana = &team[i]
		case "tim_redaksi":
			timRedaksi = append(timRedaksi, team[i])
		}
	}

	// Build response
	response := map[string]interface{}{
		"dewan_redaksi":           council,
		"tim_redaksi":             timRedaksi,
	}

	// Add pemimpin_redaksi if exists
	if pemimpinRedaksi != nil {
		response["pemimpin_redaksi"] = map[string]interface{}{
			"name":     pemimpinRedaksi.Name,
			"position": pemimpinRedaksi.Position,
			"photo":    pemimpinRedaksi.Photo,
			"bio":      pemimpinRedaksi.Bio,
			"email":    pemimpinRedaksi.Email,
			"phone":    pemimpinRedaksi.Phone,
		}
	}

	// Add wakil_pemimpin_redaksi if exists
	if len(wakilPemimpinRedaksi) > 0 {
		var wakils []map[string]interface{}
		for _, wakil := range wakilPemimpinRedaksi {
			wakils = append(wakils, map[string]interface{}{
				"name":     wakil.Name,
				"position": wakil.Position,
				"photo":    wakil.Photo,
				"bio":      wakil.Bio,
				"email":    wakil.Email,
			})
		}
		response["wakil_pemimpin_redaksi"] = wakils
	}

	// Add redaktur_pelaksana if exists
	if redakturPelaksana != nil {
		response["redaktur_pelaksana"] = map[string]interface{}{
			"name":     redakturPelaksana.Name,
			"position": redakturPelaksana.Position,
			"photo":    redakturPelaksana.Photo,
			"bio":      redakturPelaksana.Bio,
			"email":    redakturPelaksana.Email,
		}
	}

	// Add contact info (could be from settings)
	response["contact"] = map[string]interface{}{
		"email": "redaksi@lpmaarifnu.or.id",
		"phone": "021-12345678",
	}

	return response, nil
}
