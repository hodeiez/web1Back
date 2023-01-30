package gui

// import (
// 	"hodei/web1/db"
// 	"hodei/web1/service"

// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/data/binding"
// 	"fyne.io/fyne/v2/layout"
// 	"fyne.io/fyne/v2/widget"
// )

// func TrackPanel(w fyne.Window) *fyne.Container {
// 	titleInp := widget.NewEntry()
// 	label1 := widget.NewLabel("Title")
// 	descrInp := widget.NewEntry()
// 	label2 := widget.NewLabel("Description")
// 	dateInp := widget.NewEntry()
// 	label3 := widget.NewLabel("Date")
// 	label4 := widget.NewLabel("fileRef")
// 	label5 := widget.NewLabel("")
// 	data := binding.NewStringList()

// 	but := widget.NewButton("Save Track", func() {
// 		data.Set([]string{titleInp.Text, descrInp.Text, dateInp.Text, label4.Text})
// 		setData, _ := data.Get()
// 		saved := saveTrack(setData[0], setData[1], setData[2], setData[3])
// 		data.Set([]string{saved.Key, saved.Title, saved.Description, saved.Date, saved.FileRef})

// 	})
// 	mockBut := widget.NewButton("Select file", func() { ChooseFile(w, label4) })
// 	result := widget.NewListWithData(data,
// 		func() fyne.CanvasObject {
// 			return widget.NewLabel("Response")
// 		},
// 		func(i binding.DataItem, o fyne.CanvasObject) {
// 			o.(*widget.Label).Bind(i.(binding.String))
// 		})
// 	return container.New(layout.NewGridLayout(2), container.New(layout.NewFormLayout(), label1, titleInp, label2, descrInp, label3, dateInp, label4, mockBut, label5, but), result)
// }
// func saveTrack(title, descr, dateInp, fileRef string) db.DbTrack {
// 	return service.AddTrack(db.DbTrack{Title: title, Description: descr, Date: dateInp, FileRef: fileRef})
// }
