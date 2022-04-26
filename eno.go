package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/TwiN/go-color"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	//	"github.com/charmbracelet/lipgloss"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/qeesung/image2ascii/convert"
)

const (
	padding  = 2
	maxWidth = 100
)

var (
	subtle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}

	InfoStyle = lipgloss.NewStyle().
			Padding(0, 0, 1, 2)
	BorderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.AdaptiveColor{Light: "#2B53AF", Dark: "#37B9FF"}).
			Width(80).
			Height(80).
			Padding(0, 1, 1, 2)
)

type model struct {
	progress     progress.Model
	fileContent  []string
	altScreen    bool
	selectedText string
	layoutStyle  lipgloss.Style
	borderStyle  lipgloss.Style
}

type tickMsg time.Time

func main() {
	args := os.Args[1]
	switch args {

	case "idk":
		help_menu()
	case "help":
		m := model{
			progress: progress.New(progress.WithDefaultGradient()),
			layoutStyle: InfoStyle,
			borderStyle: BorderStyle,
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
	//	f, err := os.ReadFile("strategies.txt")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	file, _ := os.Open("strategies.txt")
	//if err != nil {
	//	return nil, err
	//}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	m.fileContent = lines
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

		//		r := rand.Intn(len(m.fileContent))
		//		m.selectedText = m.fileContent[r]
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

	return
}

func (e model) View() string {
	file, _ := os.Open("strategies.txt")
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if !e.altScreen {
		return loading(e)
	} else {
		view := lines[rand.Intn(len(lines))]
		return e.layoutStyle.Render(e.borderStyle.Render(view))
		//Outline of card with clickable option
	}
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
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
