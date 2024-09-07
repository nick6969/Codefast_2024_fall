package app

import (
	"codefast_2024/pages"
	"embed"
)

type App struct {
	// Embed Pages Folder FileSystems
	PageFS embed.FS
}

func NewApp() *App {
	return &App{
		PageFS: pages.FS,
	}
}
