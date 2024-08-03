package house

import "avito-test/internal/models"

const (
	onModerate = iota
	created
	declined
	approved
)

func codeFromFlatStatus(status models.ModerationStatus) byte {
	switch status {
	case models.OnModerate:
		return onModerate
	case models.Created:
		return created
	case models.Declined:
		return declined
	case models.Approved:
		return approved
	}
	panic("unimplemented status")
}

func flatStatusFromCode(status byte) models.ModerationStatus {
	switch status {
	case onModerate:
		return models.OnModerate
	case created:
		return models.Created
	case declined:
		return models.Declined
	case approved:
		return models.Approved
	}
	panic("unimplemented status")
}
