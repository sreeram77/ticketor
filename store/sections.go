package store

import (
	"ticketor/errors"
	"ticketor/models"
)

type sections struct {
	store map[string]models.Section
}

func NewSections() Sections {
	store := make(map[string]models.Section)

	// Populate data.
	store["1"] = models.Section{ID: "1", Number: 1}
	store["2"] = models.Section{ID: "2", Number: 2}

	return &sections{
		store: store,
	}
}

// Get fetches a section by ID.
func (s sections) Get(id string) (models.Section, error) {
	section, exists := s.store[id]
	if !exists {
		return models.Section{}, errors.ErrNotFound
	}

	return section, nil
}
