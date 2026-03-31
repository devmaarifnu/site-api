package repositories

import (
	"lpmaarifnu-site-api/internal/models"

	"gorm.io/gorm"
)

type OrganizationRepository interface {
	GetOrganizationStructure() (map[string]interface{}, error)
	FindBoardMembers(filters map[string]interface{}) ([]models.BoardMember, error)
	FindDepartments() ([]models.Department, error)
}

type organizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &organizationRepository{db: db}
}

func (r *organizationRepository) GetOrganizationStructure() (map[string]interface{}, error) {
	structure := make(map[string]interface{})

	// Get Ketua
	var ketua models.BoardMember
	err := r.db.Joins("JOIN organization_positions ON organization_positions.id = board_members.position_id").
		Where("organization_positions.position_type = ? AND organization_positions.position_name = ?", "ketua", "Ketua").
		Where("board_members.is_active = ?", true).
		First(&ketua).Error
	if err == nil {
		structure["ketua"] = ketua
	}

	// Get Wakil Ketua
	var wakilKetua []models.BoardMember
	r.db.Joins("JOIN organization_positions ON organization_positions.id = board_members.position_id").
		Where("organization_positions.position_type = ? AND organization_positions.position_name = ?", "wakil", "Wakil Ketua").
		Where("board_members.is_active = ?", true).
		Order("board_members.order_number ASC").
		Find(&wakilKetua)
	structure["wakil_ketua"] = wakilKetua

	// Get Sekretaris
	var sekretaris models.BoardMember
	err = r.db.Joins("JOIN organization_positions ON organization_positions.id = board_members.position_id").
		Where("organization_positions.position_type = ? AND organization_positions.position_name = ?", "sekretaris", "Sekretaris").
		Where("board_members.is_active = ?", true).
		First(&sekretaris).Error
	if err == nil {
		structure["sekretaris"] = sekretaris
	}

	// Get Wakil Sekretaris
	var wakilSekretaris []models.BoardMember
	r.db.Joins("JOIN organization_positions ON organization_positions.id = board_members.position_id").
		Where("organization_positions.position_type = ? AND organization_positions.position_name = ?", "sekretaris", "Wakil Sekretaris").
		Where("board_members.is_active = ?", true).
		Order("board_members.order_number ASC").
		Find(&wakilSekretaris)
	structure["wakil_sekretaris"] = wakilSekretaris

	// Get Bendahara
	var bendahara models.BoardMember
	err = r.db.Joins("JOIN organization_positions ON organization_positions.id = board_members.position_id").
		Where("organization_positions.position_type = ? AND organization_positions.position_name = ?", "bendahara", "Bendahara").
		Where("board_members.is_active = ?", true).
		First(&bendahara).Error
	if err == nil {
		structure["bendahara"] = bendahara
	}

	// Get Wakil Bendahara
	var wakilBendahara []models.BoardMember
	r.db.Joins("JOIN organization_positions ON organization_positions.id = board_members.position_id").
		Where("organization_positions.position_type = ? AND organization_positions.position_name = ?", "bendahara", "Wakil Bendahara").
		Where("board_members.is_active = ?", true).
		Order("board_members.order_number ASC").
		Find(&wakilBendahara)
	structure["wakil_bendahara"] = wakilBendahara

	// Get Departments with Members
	var departments []models.Department
	r.db.Where("is_active = ?", true).
		Order("order_number ASC").
		Preload("Members", "is_active = ? ORDER BY order_number ASC", true).
		Find(&departments)
	structure["departments"] = departments

	return structure, nil
}

func (r *organizationRepository) FindBoardMembers(filters map[string]interface{}) ([]models.BoardMember, error) {
	var members []models.BoardMember

	query := r.db.Model(&models.BoardMember{}).Preload("Position")

	if period, ok := filters["period"].(string); ok && period != "" {
		// Parse period like "2024-2029"
		query = query.Where("CONCAT(period_start, '-', period_end) = ?", period)
	}

	if active, ok := filters["active"].(bool); ok {
		query = query.Where("is_active = ?", active)
	} else {
		query = query.Where("is_active = ?", true)
	}

	err := query.Order("order_number ASC").Find(&members).Error
	return members, err
}

func (r *organizationRepository) FindDepartments() ([]models.Department, error) {
	var departments []models.Department
	err := r.db.
		Where("is_active = ?", true).
		Order("order_number ASC").
		Preload("Members", "is_active = ? ORDER BY order_number ASC", true).
		Find(&departments).Error

	return departments, err
}
