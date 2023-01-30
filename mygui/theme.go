package gui

// import (
// 	"image/color"

// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/theme"
// )

// type MyTheme struct{}

// var _ fyne.Theme = (*MyTheme)(nil)

// func (m MyTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
// 	if name == theme.ColorNameBackground {
// 		if variant == theme.VariantDark {
// 			return color.RGBA{255, 48, 109, 100}
// 		}
// 		return color.Black
// 	} else if name == theme.ColorNameInputBackground {
// 		return color.RGBA{255, 48, 225, 255}
// 	}

// 	return theme.DefaultTheme().Color(name, variant)
// }
// func (m MyTheme) Font(style fyne.TextStyle) fyne.Resource {
// 	return theme.DefaultTheme().Font(style)
// }

// func (m MyTheme) Size(name fyne.ThemeSizeName) float32 {
// 	return theme.DefaultTheme().Size(name)
// }
// func (m MyTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
// 	// if name == theme.IconNameHome {
// 	// 	fyne.NewStaticResource("myHome", homeBytes)
// 	// }

// 	return theme.DefaultTheme().Icon(name)
// }
