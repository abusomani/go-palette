package main

import (
	pal "github.com/abusomani/go-palette/palette"
)

func main() {
	p := pal.New()
	p.SetOptions(pal.WithForeground(56)).Println("\n	    =============================================================================================")
	p.SetOptions(pal.WithForeground(56)).Print("	    #	")
	p.SetOptions(pal.WithForeground(207)).Print(" @@@@@                          @@@@@@                                                ")
	p.SetOptions(pal.WithForeground(56)).Println("	#")
	p.SetOptions(pal.WithForeground(56)).Print("	    #	")
	p.SetOptions(pal.WithForeground(171)).Print("@     @   @@@@                  @     @    @@    @       @@@@@@  @@@@@  @@@@@  @@@@@@ ")
	p.SetOptions(pal.WithForeground(56)).Println("	#")
	p.SetOptions(pal.WithForeground(56)).Print("	    #	")
	p.SetOptions(pal.WithForeground(135)).Print("@        @    @                 @     @   @  @   @       @         @      @    @      ")
	p.SetOptions(pal.WithForeground(56)).Println("	#")
	p.SetOptions(pal.WithForeground(56)).Print("	    #	")
	p.SetOptions(pal.WithForeground(99)).Print("@  @@@@  @    @      @@@@@      @@@@@@   @    @  @       @@@@@     @      @    @@@@@  ")
	p.SetOptions(pal.WithForeground(56)).Println("	#")
	p.SetOptions(pal.WithForeground(56)).Print("	    #	")
	p.SetOptions(pal.WithForeground(134)).Print("@     @  @    @                 @        @@@@@@  @       @         @      @    @      ")
	p.SetOptions(pal.WithForeground(56)).Println("	#")
	p.SetOptions(pal.WithForeground(56)).Print("	    #	")
	p.SetOptions(pal.WithForeground(169)).Print("@     @  @    @                 @        @    @  @       @         @      @    @      ")
	p.SetOptions(pal.WithForeground(56)).Println("	#")
	p.SetOptions(pal.WithForeground(56)).Print("	    #	")
	p.SetOptions(pal.WithForeground(206)).Print(" @@@@@    @@@@                  @        @    @  @@@@@@  @@@@@@    @      @    @@@@@@ ")
	p.SetOptions(pal.WithForeground(56)).Println("	#")
	p.SetOptions(pal.WithForeground(56)).Println("	    =============================================================================================\n\n")

	p.Flush().Print("	")
	p.SetOptions(pal.WithForeground(255), pal.WithBackground(207)).Print(" Go Palette ")
	p.Flush().Print("										    ")
	p.SetOptions(pal.WithForeground(255), pal.WithBackground(207)).Println(" Go Palette ")

	p.Flush().Print("	")
	p.Flush().Print("  ")
	p.SetOptions(pal.WithForeground(255), pal.WithBackground(206)).Print(" Go Palette ")
	p.Flush().Print("		Now paint your terminals with colors using the Go Palette	  ")
	p.SetOptions(pal.WithForeground(255), pal.WithBackground(206)).Println(" Go Palette ")

	p.Flush().Print("	")
	p.Flush().Print("    ")
	p.SetOptions(pal.WithForeground(255), pal.WithBackground(205)).Print(" Go Palette ")
	p.Flush().Print("	=========================================================	")
	p.SetOptions(pal.WithForeground(255), pal.WithBackground(205)).Println(" Go Palette ")

	p.Flush().Print("	")
	p.Flush().Print("  ")
	p.SetOptions(pal.WithForeground(255), pal.WithBackground(204)).Print(" Go Palette ")
	p.Flush().Print("		Abhishek Somani - ")
	p.SetOptions(pal.WithForeground(pal.Green)).Print("https://github.com/abusomani/go-palette	  ")
	p.SetOptions(pal.WithForeground(255), pal.WithBackground(204)).Println(" Go Palette ")

	p.Flush().Print("	")
	p.SetOptions(pal.WithForeground(255), pal.WithBackground(203)).Print(" Go Palette ")
	p.Flush().Print("										    ")
	p.SetOptions(pal.WithForeground(255), pal.WithBackground(203)).Print(" Go Palette ")
	p.Flush().Println("			\n")
}
