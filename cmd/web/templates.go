// Rendering multiple pieces of data on the same page

package main

import "github.com/TomiwaAribisala-git/snippetbox.git/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
