package repository

import "fmt"

func (s *Repository) FindItem(item string) error {
	const op = "storage.FindItem"

	if _, ok := s.wh.Items[item]; !ok {
		return fmt.Errorf("%s: item not found", op)
	}

	return nil
}
