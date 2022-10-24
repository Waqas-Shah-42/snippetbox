package main

import "github.com/Waqas-Shah-42/snippetbox/internal/models"

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}
