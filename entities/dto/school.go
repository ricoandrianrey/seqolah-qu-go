package dto

import "seqolah-qu/types"

type School struct {
	ID            uint        `json:"id"`
	Address       string      `json:"address"`
	AddressDetail types.JSONB `json:"address_detail"`
	Logo          string      `json:"logo"`
	Email         string      `json:"email"`
	Accreditation string      `json:"accreditation"`
	Phone         string      `json:"phone"`
	Curriculum    string      `json:"curriculum"`
}
