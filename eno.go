package main 

import (
	"fmt";
	"os";
	"github.com/TwiN/go-color";
	"github.com/qeesung/image2ascii/convert";
	_ "image/jpeg";
	_ "image/png";
	"github.com/common-nighthawk/go-figure"
)

func main(){
	
	
	args := os.Args[1]
	switch args {
	
		case "idk": 
			help_menu()
		case "help":
			fmt.Println("TODO")
	    default:
			fmt.Println("Error no such arg")
	}

}

//CLI for eno
//eno idk - Shows help menu for a new person or if someone was forgetful

//eno help -- shows an oblique strategy on the screen




//Help menu displayed when ran with eno idk
func help_menu(){

	convertOptions := convert.DefaultOptions;
	convertOptions.FixedWidth = 100
	convertOptions.FixedHeight = 40
	
	converter := convert.NewImageConverter()	
	
	figure.NewColorFigure("                    Eno","","cyan",true).Print()
	fmt.Print(converter.ImageFile2ASCIIString("brian.jpg",&convertOptions))	
	fmt.Println(color.InCyan("                        Your own creative helper"))
	
	fmt.Println(color.InWhite("Usage:"))
	fmt.Println(color.InWhite("        $ eno <command>\n"))
	fmt.Println(color.InYellow("Commands:"))
	fmt.Println(color.InYellow("  idk     Show this help menu"))
	fmt.Println(color.InYellow("  help    Will randomly show an oblique strategy for your creative output\n"))
	fmt.Println(color.InGreen("Examples:"))
	fmt.Println(color.InGreen("        $ eno idk"))
	fmt.Println(color.InGreen("        $ eno help"))
}
