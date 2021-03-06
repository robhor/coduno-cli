package main

import (
	"fmt"
	"os"
	"os/exec"
)

var cmdChallenge = &Command{
	Run:       getChallenge,
	UsageLine: "challenge <challenge-name>",
	Short:     "fetch challenge from server",
	Long: `
Fetch challenge <challenge-name> from server.

All needed files will be saved in a directory named <challende-name>.
	`,
}

func getChallenge(cmd *Command, args []string) {
	if len(args) < 1 {
		// No argument given, list all available challenges
		listChallenges(cmd, args)
	} else if len(args) == 1 {
		// challenge name given, clone into directory
		cloneChallenge(cmd, args)
	} else {
		// Invalid number of arguments, show usage
		fmt.Fprintln(os.Stderr, "Invalid number of arguments")
		os.Exit(2)
	}
}

func listChallenges(cmd *Command, args []string) {
	fmt.Printf("%s:\t%s\n", "tictactoe", "Write a bot that plays Tic Tac Toe")
}

func cloneChallenge(cmd *Command, args []string) {
	repo := "git@git.cod.uno:" + args[0]
	c := exec.Command("git", "clone", repo)
	outstr, err := c.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to fetch challenge <"+args[0]+">.")
		fmt.Fprintln(os.Stderr, "Reason: "+string(outstr)) // TODO: Remove message from production tool
		os.Exit(2)
	} else {
		fmt.Println(os.Stdout, "Successfully fetched into directory "+args[0])
	}
}
