package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ChooseDirectory(w fyne.Window, h *widget.Label) {
	dialog.ShowFolderOpen(func(dir fyne.ListableURI, err error) {
		save_dir := "NoPathYet!"
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		if dir != nil {

			save_dir = dir.Path()
		}

		h.SetText(save_dir)
	}, w)
}
func ChooseFile(w fyne.Window, h *widget.Label) {
	dialog.ShowFileOpen(func(dir fyne.URIReadCloser, err error) {
		save_file := "NoFile"
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		if dir != nil {

			save_file = dir.URI().Path()
		}

		h.SetText(save_file)
	}, w)
}
