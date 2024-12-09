package gui

import (
	"strconv"
	"strings"

	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
)

func Get(convRate float64) declarative.MainWindow {
	var inTE, outTE *walk.TextEdit

	return declarative.MainWindow{
		Title:  "UY to USD",
		Size:   declarative.Size{Width: 300, Height: 200},
		Layout: declarative.VBox{},
		Children: []declarative.Widget{
			declarative.Label{
				Text: strconv.FormatFloat(convRate, 'f', -1, 64),
			},
			declarative.HSplitter{
				Children: []declarative.Widget{
					declarative.TextEdit{AssignTo: &inTE},
					declarative.TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			declarative.PushButton{
				Text: "Convert",
				OnClicked: func() {
					input, err := strconv.ParseFloat(strings.TrimSpace(inTE.Text()), 64)
					if err != nil {
						outTE.SetText("error")
						return
					}
					output := input / convRate
					outTE.SetText(strconv.FormatFloat(output, 'f', -1, 64))
				},
			},
		},
	}
}
