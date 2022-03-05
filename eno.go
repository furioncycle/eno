package main 

import (
	"fmt";
	"log";
	"time";
	"os";
	"github.com/TwiN/go-color";
	"github.com/qeesung/image2ascii/convert";
	_ "image/jpeg";
	_ "image/png";
	"github.com/common-nighthawk/go-figure";
	tea "github.com/charmbracelet/bubbletea"
)

type model int
type tickMsg time.Time
func main(){
	
	
	args := os.Args[1]
	switch args {
	
		case "idk": 
			help_menu()
		case "help":
			p := tea.NewProgram(model(5), tea.WithAltScreen())
		    if err := p.Start(); err != nil {
				log.Fatal(err)
			}
	    default:
			fmt.Println("Error no such arg")
	}

}

func (m model) Init() tea.Cmd {
	return tea.Batch(tick(), tea.EnterAltScreen)
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type){
		case tea.KeyMsg: 
			switch msg.String() {
				case "q", "esc", "ctrl+c":
					return m, tea.Quit
			}
		case tickMsg: 
			m -= 1
			if m <= 0 {
				return m, tea.Quit
			}
			return m, tick()
	}
	
	return m, nil
}

func (m model) View() string {

	convertOptions := convert.DefaultOptions;
	convertOptions.FixedWidth = 100
	convertOptions.FixedHeight = 40
	
	converter := convert.NewImageConverter()	
	
	return fmt.Sprintf(converter.ImageFile2ASCIIString("brian.jpg",&convertOptions))
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg{
			return tickMsg(t)
	})
}
//CLI for eno
//eno idk - Shows help menu for a new person or if someone was forgetful

//eno help -- shows an oblique strategy on the screen




//Help menu displayed when ran with eno idk
func help_menu(){
	figure.NewColorFigure("Eno","","cyan",true).Print()
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
