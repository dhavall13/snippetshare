package main

import "github.com/dhvll/snippetshare/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
