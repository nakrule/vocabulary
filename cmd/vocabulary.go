package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

const buttonWidth = 130 // Some buttons have a fixed width

func main() {

	// needs to be called once before you can start using the QWidgets
	app := widgets.NewQApplication(len(os.Args), os.Args)
	app.SetApplicationName("Vocabulary")

	// create a window
	// with a minimum size of 250*200
	// and sets the title to "Hello Widgets Example"
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(250, 200)
	window.SetWindowTitle("Vocabulary")

	menuBar()

	// create a regular widget
	// give it a QVBoxLayout
	// and make it the central widget of the window
	mainVbox := widgets.NewQWidget(nil, 0)
	mainVbox.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(mainVbox)

	mainVbox.Layout().AddWidget(createLearnButton())
	mainVbox.Layout().AddWidget(createWordTableContainer())
	mainVbox.Layout().AddWidget(createBottomButtons())

	// make the window visible
	window.Show()

	// start the main Qt event loop
	// and block until app.Exit() is called
	// or the window is closed by the user
	app.Exec()
}

// Create a table to enter words and definitions
func createWordTableContainer() widgets.QWidget_ITF {

	rowsNumber := 5
	columnsNumber := 2

	table := widgets.NewQTableWidget(nil)
	table.SetRowCount(rowsNumber)
	table.SetColumnCount(columnsNumber)
	table.SetAlternatingRowColors(true)

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