package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	m "github.com/vaaleyard/turing-machine/machine"
	t "github.com/vaaleyard/turing-machine/tape"
	"strings"
	"time"
)

var (
	tviewApp *tview.Application
	table    *tview.Table
	machine  m.Machine
	tape     *t.Tape
)

type App struct {
	actualState *string
}

func updateTape(machineApp *App) bool {
	var done bool = false
	time.Sleep(500 * time.Millisecond)
	tviewApp.QueueUpdateDraw(func() {
		var isMachineDone bool = machine.IsDone(tape, machineApp.actualState)
		if !isMachineDone {

			var binding = machine.FindTransitionsByActualState(*machineApp.actualState,
				tape.Tape[tape.Index].Data)

			var futureTransition = binding[0]
			tape.Tape[tape.Index].Data = futureTransition.Escreve
			*machineApp.actualState = futureTransition.EstadoDestino

			tape.Move(futureTransition)

			table.SetCell(0, int(tape.Index),
				tview.NewTableCell(tape.Tape[tape.Index].Data).
					SetTextColor(tcell.ColorRed.TrueColor()).
					SetSelectable(true).
					SetAlign(tview.AlignCenter),
			)
			isMachineDone = machine.IsDone(tape, machineApp.actualState)
			done = isMachineDone
		}
	})
	//tviewApp.Draw()
	return done
}

func Ui() {
	tviewApp = tview.NewApplication()
	pages := tview.NewPages()

	table = tview.NewTable().
		SetBorders(true)
	drawEmptyTape(table)

	var inputAlphabet string
	var machineApp *App
	form := tview.NewForm()
	form.AddInputField("Input: ", "", 55, nil, nil).
		SetFieldTextColor(tcell.ColorBlack.TrueColor()).
		AddButton("Load Tape", func() {
			inputAlphabet = form.GetFormItem(0).(*tview.InputField).GetText()

			var alphabet string
			alphabet, machineApp = initMachine(inputAlphabet)
			drawInputtedTape(table, alphabet)
		}).
		AddButton("Process Machine", func() {
			go func() {
				var done bool = false
				for !done {
					done = updateTape(machineApp)
				}

				if machine.ValidateChain(*machineApp.actualState) {
					alert(pages, "alert-dialog", "Word accepted")
				} else {
					alert(pages, "alert-dialog", "Word rejected")
				}
				tviewApp.Draw()
			}()
		}).
		SetButtonsAlign(tview.AlignCenter).
		SetButtonBackgroundColor(tcell.ColorBlack).
		SetButtonTextColor(tcell.ColorYellow.TrueColor())
	form.SetBorderPadding(5, 0, 0, 0)

	// Flex
	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
				AddItem(tview.NewBox(), 0, 2, false).
				AddItem(form, 0, 2, true).
				AddItem(tview.NewBox(), 0, 2, false),
				0, 1, true).
			AddItem(table, 0, 2, false).
			AddItem(tview.NewBox().SetBorder(false), 0, 1, false), 0, 1, true)
	flex.SetTitle("Turing Machine").SetBorder(true).SetBorderAttributes(tcell.AttrDim)
	pages.AddPage("base", flex, true, true)

	if err := tviewApp.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
		panic(err)
	}
}

func initMachine(alphabet string) (string, *App) {
	machine = m.NewMachineFromFile("examples/maquina01.json")
	tape = machine.CreateTape(alphabet)
	alphabet = ""
	for i := range tape.Tape {
		alphabet += tape.Tape[i].Data
	}
	var app1 = &App{
		actualState: machine.EstadoInicial,
	}

	return alphabet, app1
}

func drawEmptyTape(table *tview.Table) {
	columns, rows := m.BLANK_CELL_QUANTITY, 1
	symbol := 0
	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			table.SetCell(row, column,
				tview.NewTableCell(machine.SimboloBranco).
					SetTextColor(tcell.ColorWhite).
					SetAlign(tview.AlignCenter))
			symbol = symbol + 1
		}
	}
	table.SetBorderPadding(5, 0, 40, 0)
}

func drawInputtedTape(table *tview.Table, tape string) {
	tapeCells := strings.Split(tape, "")
	columns, rows := t.TAPE_SIZE, 1
	symbol := 0

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			table.SetCell(row, column,
				tview.NewTableCell(tapeCells[symbol]).
					SetTextColor(tcell.ColorWhite).
					SetAlign(tview.AlignCenter))
			symbol = (symbol + 1) % len(tapeCells)
		}
	}
	table.SetBorderPadding(5, 0, 40, 0)
}

func alert(pages *tview.Pages, id string, message string) *tview.Pages {
	return pages.AddPage(
		id,
		tview.NewModal().
			SetText(message).
			AddButtons([]string{"Ok"}).SetBackgroundColor(tcell.ColorBlack).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				pages.HidePage(id).RemovePage(id)
			}).
			SetButtonBackgroundColor(tcell.ColorBlack).
			SetButtonTextColor(tcell.ColorYellow.TrueColor()),
		false,
		true,
	)
}
