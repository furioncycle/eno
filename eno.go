package main

import (
	"fmt"
	"strings"
	"time"
	"github.com/TwiN/go-color"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
//	"github.com/charmbracelet/lipgloss"
	"github.com/common-nighthawk/go-figure"
	"github.com/qeesung/image2ascii/convert"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

type model struct {
	progress progress.Model
	fileContent []string
	altScreen bool
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
		m := model{
			progress: progress.New(progress.WithDefaultGradient()),
		}

		p := tea.NewProgram(m, tea.WithAltScreen())
		if err := p.Start(); err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Println("Error no such arg")
	}

}

func (m model) Init() tea.Cmd {
	f, err := os.ReadFile("strategies.txt")
	if err != nil {
		log.Fatal(err)
	}
	m.fileContent = strings.Split(string(f),"\n")
	return tea.Batch(tickCmd(), tea.EnterAltScreen)
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil
	case tickMsg:
		if m.progress.Percent() == 1.0 {
			m.altScreen = true
			return m, tea.EnterAltScreen
		}
		
		cmd := m.progress.IncrPercent(0.25)
		return m, tea.Batch(tickCmd(), cmd)
	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd
	}

	return m, nil
}

func loading(e model) (s string) {
	
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

func (e model) View() (s string) {
    if !e.altScreen {
		s = loading(e)
	}else{
		//Outline of card with clickable option
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
