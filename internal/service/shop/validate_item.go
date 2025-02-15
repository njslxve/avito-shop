package shop

import "log/slog"

func (ss *ShopService) ValidateItem(itemname string) bool {
	item, err := ss.repo.Item.FindItem(itemname)
	if err != nil {
		ss.logger.Error("failed to find item",
			slog.String("item", item.Type),
			slog.String("error", err.Error()),
		)

		return false
	}

	return true
}
