package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

const buttonWidth = 130 // Some buttons have a fixed width

var table *widgets.QTableWidget

func main() {

	// needs to be called once before you can start using the QWidgets
	app := widgets.NewQApplication(len(os.Args), os.Args)
	app.SetApplicationName("Vocabulary")

	window := setupUi()

	// make the window visible
	window.Show()

	// start the main Qt event loop
	// and block until app.Exit() is called
	// or the window is closed by the user
	app.Exec()
}

// Create a table to enter words and definitions
func createWordTableContainer() *widgets.QTableWidget {

	rowsNumber := 10
	columnsNumber := 2

	table := widgets.NewQTableWidget(nil)
	table.SetRowCount(rowsNumber)
	table.SetColumnCount(columnsNumber)
	table.SetAlternatingRowColors(true)
	table.SetShowGrid(false)
	table.SetSelectionBehavior(widgets.QAbstractItemView__SelectRows) // Click a cell select the whole row
	table.SetCornerButtonEnabled(false)                               // Disable the top left corner button to select all cells

	headers := []string{"Terms", "Definitions"}
	table.SetHorizontalHeaderLabels(headers)

	// Set the table to horizontally stretch to the window
	table.QTableView.HorizontalHeader().SetSectionResizeMode(widgets.QHeaderView__Stretch)

	return table
}

// Create a right aligned "Learn" button.
func createLearnButton() *widgets.QWidget {
	hbox := widgets.NewQWidget(nil, 0)
	hbox.SetLayout(widgets.NewQHBoxLayout())

	button := widgets.NewQPushButton2("Learn", nil)
	button.SetFixedWidth(buttonWidth)

	hbox.Layout().AddWidget(button)
	hbox.Layout().SetAlignment(button, core.Qt__AlignRight) // Align the button on the right side
	return hbox
}

// Create the two bottom buttons: "delete" and "add row"
func createBottomButtons() *widgets.QWidget {
	hbox := widgets.NewQWidget(nil, 0)
	hbox.SetLayout(widgets.NewQHBoxLayout())

	deleteButton := widgets.NewQPushButton2("Delete", nil)
	newRowButton := widgets.NewQPushButton2("Add row", nil)
	deleteButton.SetMaximumWidth(buttonWidth)
	deleteButton.SetFixedWidth(buttonWidth)
	newRowButton.SetFixedWidth(buttonWidth)

	newRowButton.ConnectClicked(func(bool) {
		table.InsertRow(table.RowCount())
	})
	deleteButton.ConnectClicked(func(bool) {
		model := table.SelectionModel()
		if model.HasSelection() {
			rowIndex := model.CurrentIndex().Row()
			table.RemoveRow(rowIndex)
			// select next row after deleting the old one
			if rowIndex > 0 {
				table.SelectRow(rowIndex - 1)
			} else {
				table.SelectRow(0)
			}
		}
	})

	hbox.Layout().AddWidget(deleteButton)
	hbox.Layout().AddWidget(newRowButton)
	hbox.Layout().SetAlignment(newRowButton, core.Qt__AlignRight) // Align the button on the right side
	return hbox
}

// Create to top menu bar with File and Edit menu
func menuBar() {

	menuBar := widgets.NewQMenuBar(nil)

	fileMenu := menuBar.AddMenu2("File")
	fileMenu.AddAction("About")
	fileMenu.AddSeparator()
	fileMenu.AddAction("Open file")
	fileMenu.AddAction("Save")
	fileMenu.AddAction("Save as")
	fileMenu.AddAction("Print")
	fileMenu.AddAction("Exit")

	editMenu := menuBar.AddMenu2("Edit")
	editMenu.AddAction("Undo")
	editMenu.AddAction("Redo")
}

func setupUi() *widgets.QMainWindow {

	// create a window
	window := widgets.NewQMainWindow(nil, 0)
	window.SetBaseSize2(450, 500)
	window.SetWindowTitle("Vocabulary")

	menuBar()

	// create a regular widget
	// give it a QVBoxLayout
	// and make it the central widget of the window
	mainVbox := widgets.NewQWidget(nil, 0)
	mainVbox.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(mainVbox)

	table = createWordTableContainer()
	mainVbox.Layout().AddWidget(createLearnButton())
	mainVbox.Layout().AddWidget(table)
	mainVbox.Layout().AddWidget(createBottomButtons())

	return window
}
