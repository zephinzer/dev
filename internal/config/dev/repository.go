package dev

import (
	"fmt"

	"github.com/zephinzer/dev/internal/config/dev/repository"
)

// Repository holds client configuration for operations related to repositories
type Repository struct {
	Templates []repository.Template `json:"templates" yaml:"templates,omitempty"`
}

func (dr *Repository) MergeWith(o Repository) []error {
	warnings := []error{}
	seen := map[string]bool{}
	for _, t := range dr.Templates {
		seen[t.GetKey()] = true
	}
	for _, T := range o.Templates {
		if exists, ok := seen[T.GetKey()]; exists && ok {
			warnings = append(warnings, fmt.Errorf("template repository with key '%s' already exists", T.GetKey()))
			continue
		}
		dr.Templates = append(dr.Templates, T)
		seen[T.GetKey()] = true
	}
	return warnings
}
