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

func MyTableLength(data [][]string) (int, int) {
	return len(data), len(data[0])
}
func MyCreateTable() fyne.CanvasObject {
	entry := widget.NewEntry()
	// entry.SetText("123456789ABCD")
	// entry.SetPlaceHolder("aaaaaaaaaaaaaaaaaaaaaaaa")
	return entry
}
func MyUpdateTable(i widget.TableCellID, o fyne.CanvasObject) {
	o.(*widget.Entry).Resize(fyne.NewSize(234, 30))
	o.(*widget.Entry).SetText(tracksData[i.Row][i.Col])
}

func AlbumsPanel(create func() fyne.Container) fyne.Container {
	return create()
}
