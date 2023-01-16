package gui

import (
	"hodei/web1/db"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
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

	text := widget.NewEntry()
	text2 := widget.NewTextGrid()
	text2.SetRowStyle(0, &widget.CustomTextGridStyle{color.White, color.White})
	text2.ExtendBaseWidget(text)

	return text
}
func MyUpdateTable(i widget.TableCellID, o fyne.CanvasObject) {
	o.(*widget.Entry).SetText(tracksData[i.Row][i.Col])
}

func AlbumsPanel(create func() fyne.Container) fyne.Container {
	return create()
}
func TrackModal(w fyne.Window, track db.DbTrack) dialog.Dialog {
	return dialog.NewCustom(track.Title, "Ados", container.NewVBox(widget.NewLabel(track.Key),
		widget.NewLabel(track.Title),
		widget.NewLabel(track.Description),
		widget.NewLabel(track.Date),
		widget.NewLabel(track.FileRef)), w)
}
