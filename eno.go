package main

import (
	"fmt"
	"strings"
	"time"
	"github.com/TwiN/go-color"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/viewport"
	//"github.com/charmbracelet/glamour"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/common-nighthawk/go-figure"
	"github.com/qeesung/image2ascii/convert"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

type state int

const (
	index state = iota
	detail
)

type indexModel struct {
	progress progress.Model	
	fileContent []string
}

type detailModel struct {
	viewport viewport.Model
}

type model struct {
	state state
	index indexModel
	detail detailModel
}



type tickMsg time.Time

const (
	padding  = 2
	maxWidth = 100
)

func main() {
	args := os.Args[1]
	switch args {

	case "idk":
		help_menu()
	case "help":
		m := model{}
		p := tea.NewProgram(m, tea.WithAltScreen())
		if err := p.Start(); err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Println("Error no such arg")
	}

}

func newIndexModel() indexModel {
	f, err := os.ReadFile("strategies.txt")
	if err != nil {
		log.Fatal(err)
	}
//	m.fileContent = strings.Split(string(f),"\n")}
	return indexModel{
		progress: progress.New(progress.WithDefaultGradient()),
		fileContent: strings.Split(string(f),"\n"),
	}
}

func newDetailModel() detailModel {

	vp := viewport.New(78,20)
	vp.Style = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62")).
			PaddingRight(2)
	return detailModel{
		viewport: vp,		
	}
}

func (m model) Init() tea.Cmd {
	
	 m = model{
		state: index,
		index: newIndexModel(),
		detail: newDetailModel(),
	}
	
	return nil
}

func indexUpdate(message tea.Msg, m model) (indexModel, tea.Cmd) {
	switch msg := message.(type) {	
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m.index, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.index.progress.Width = msg.Width - padding*2 - 4
		if m.index.progress.Width > maxWidth {
			m.index.progress.Width = maxWidth
		}
		return m.index, nil
	case tickMsg:
		if m.index.progress.Percent() == 1.0 {
			m.state = detail
		}
		
		cmd := m.index.progress.IncrPercent(0.25)
		return m.index, tea.Batch(tickCmd(), cmd)
	case progress.FrameMsg:
		progressModel, cmd := m.index.progress.Update(msg)
		m.index.progress = progressModel.(progress.Model)
		return m.index, cmd
	}
	return m.index, nil
}

func detailUpdate(message tea.Msg, m model) (detailModel, tea.Cmd) {
	//viewport and randomize string 	
	switch msg := message.(type) {	
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m.detail, tea.Quit
		}
	}
	return m.detail, nil
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg:= message.(type) {	
		case tea.KeyMsg:
			switch msg.String() {
			case "q", "esc", "ctrl+c":
				return m, tea.Quit
		}
	}	
	
	switch m.state {
		case index:
			indexModel,cmd := indexUpdate(message,m)
			m.index = indexModel
			return m, cmd
		case detail:
			detailModel,cmd := detailUpdate(message,m)
			m.detail = detailModel
			return m, cmd
	}
	return m, nil
}

func indexView(e indexModel) (s string) {
	
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 100
	convertOptions.FixedHeight = 40

	converter := convert.NewImageConverter()
	pad := strings.Repeat(" ", padding)
	s += fmt.Sprintf(converter.ImageFile2ASCIIString("brian.jpg", &convertOptions))
	s += fmt.Sprintf("\n                         Honor your mistake as a hidden intention\n")
	s += pad + e.progress.View()

	for _, s := range e.fileContent {
		fmt.Println(s)
	}
	return 
}

func detailView(m detailModel)(s string) {
	return
}

func (m model) View() (s string) {
	switch m.state {
		case index:
			indexView(m.index)
		case detail:
			detailView(m.detail)	
	}
	return
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

//Help menu displayed when ran with eno idk
func help_menu() {
	figure.NewColorFigure("Eno", "", "cyan", true).Print()
	fmt.Println(color.InCyan("========================"))
	fmt.Println(color.InCyan("Your own creative helper"))
	fmt.Println(color.InWhite("Usage:"))
	fmt.Println(color.InWhite("        $ eno <command>\n"))
	fmt.Println(color.InYellow("Commands:"))
	fmt.Println(color.InYellow("  idk     Show this help menu"))
	fmt.Println(color.InYellow("  help    Will randomly show an oblique strategy for your creative output\n"))
	fmt.Println(color.InGreen("Examples:"))
	fmt.Println(color.InGreen("        $ eno idk"))
	fmt.Println(color.InGreen("        $ eno help"))
}
