package usecase

import "log/slog"

func (u *Usecase) ValidateItem(itemname string) bool {
	item, err := u.repo.Item.FindItem(itemname)
	if err != nil {
		u.logger.Error("failed to find item",
			slog.String("item", item.Type),
			slog.String("error", err.Error()),
		)

		return false
	}

	return true
}
