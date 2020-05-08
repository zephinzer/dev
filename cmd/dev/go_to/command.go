package go_to

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/config"
	"github.com/usvc/dev/internal/constants"
	"github.com/usvc/dev/internal/log"
	"github.com/usvc/dev/pkg/utils"
	"github.com/usvc/dev/pkg/validator"

	"github.com/sahilm/fuzzy"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.GotoCanonicalVerb,
		Aliases: constants.GotoAliases,
		Short:   "Goes to the specified argument URI",
		Run:     Run,
	}
	return &cmd
}

func Run(command *cobra.Command, args []string) {
	if len(args) == 0 {
		// links reference
		log.Debug("no arguments received, opening links...")
		startFuzzySearchInterface()
		return
	}
	argument := strings.Join(args, " ")
	log.Infof("received argument: %s", argument)
	switch true {
	case validator.IsGitHTTPUrl(argument):
		log.Debug("this should be a git http url")
	case validator.IsGitSSHUrl(argument):
		log.Debug("this should be a git ssh url")
	}
	command.Help()
}

var g *gocui.Gui
var err error
var searchTerms []string
var searchResults fuzzy.Matches

var selectionIndex int = -1

func startFuzzySearchInterface() {
	links := config.Global.Links
	for _, link := range links {
		searchTerms = append(searchTerms, fmt.Sprintf("%s [%s] (%s)", link.Label, strings.Join(link.Categories, ", "), link.URL))
	}
	g, err = gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer g.Close()

	g.Cursor = true
	g.Mouse = false

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return gocui.ErrQuit
	}); err != nil {
		log.Error(err)
	}

	if err := g.SetKeybinding("finder", gocui.KeyArrowRight, gocui.ModNone, func(g *gocui.Gui, view *gocui.View) error {
		switch view.Name() {
		case "finder":
			curX, curY := view.Cursor()
			if curX < len(view.ViewBuffer())-1 {
				view.SetCursor(curX+1, curY)
			}
		}
		return nil
	}); err != nil {
		log.Error(err)
	}

	if err := g.SetKeybinding("finder", gocui.KeyArrowLeft, gocui.ModNone, func(g *gocui.Gui, view *gocui.View) error {
		switch view.Name() {
		case "finder":
			curX, curY := view.Cursor()
			if curX > 0 {
				view.SetCursor(curX-1, curY)
			}
		}
		return nil
	}); err != nil {
		log.Error(err)
	}

	if err := g.SetKeybinding("finder", gocui.KeyArrowDown, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		if selectionIndex == -1 {
			selectionIndex = 0
			if _, err := g.SetCurrentView("results"); err != nil {
				return err
			}
			resultsView, err := g.View("results")
			if err != nil {

			}
			resultsView.SetCursor(0, 0)
		}
		return nil
	}); err != nil {
		log.Error(err)
	}

	if err := g.SetKeybinding("results", gocui.KeyArrowDown, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		selectionIndex++
		if selectionIndex >= len(searchResults) {
			selectionIndex--
			return nil
		}
		v.SetCursor(0, selectionIndex)
		return nil
	}); err != nil {
		log.Error(err)
	}

	if err := g.SetKeybinding("results", gocui.KeyArrowUp, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		selectionIndex--
		if selectionIndex == -1 {
			if _, err := g.SetCurrentView("finder"); err != nil {
				return err
			}
		}
		v.SetCursor(0, selectionIndex)
		return nil
	}); err != nil {
		log.Error(err)
	}

	if err := g.SetKeybinding("results", gocui.KeyBackspace2, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		selectionIndex = -1
		if _, err := g.SetCurrentView("finder"); err != nil {
			return err
		}
		finderView, err := g.View("finder")
		if err != nil {
		}
		v.MoveCursor(len(finderView.ViewBuffer())-1, 0, true)
		return nil
	}); err != nil {
		log.Error(err)
	}

	if err := g.SetKeybinding("results", gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		targetString := strings.Split(searchResults[selectionIndex].Str, " ")
		targetURI := strings.Trim(targetString[len(targetString)-1], "()")
		utils.OpenURIWithDefaultApplication(targetURI)
		return gocui.ErrQuit
	}); err != nil {
		log.Error(err)
	}

	if err := g.SetKeybinding("finder", gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		selectionIndex = 0
		if _, err := g.SetCurrentView("results"); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Error(err)
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		results, err := g.View("results")
		if err != nil {
		}
		results.Title = fmt.Sprintf("Search results (selectionIndex: %v)", selectionIndex)
		return nil
	}); err != nil {
		log.Error(err)
	}

	if err := g.SetKeybinding("", gocui.KeyEsc, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return gocui.ErrQuit
	}); err != nil {
		log.Error(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Error(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("finder", 0, 0, maxX, 10); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.Editable = true
		v.Frame = true
		v.Title = "Search for a link"
		if _, err := g.SetCurrentView("finder"); err != nil {
			return err
		}
		v.Editor = gocui.EditorFunc(finder)
	}

	if v, err := g.SetView("results", 0, 3, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		v.Wrap = true
		v.Frame = true
		v.Title = "Search results"
	}

	return nil
}

func recalculateResults(gui *gocui.Gui) error {
	finderView, err := g.View("finder")
	if err != nil {
		// handle error
	}
	results, err := g.View("results")
	if err != nil {
		// handle error
	}
	results.Clear()
	t := time.Now()
	searchResults = fuzzy.Find(strings.TrimSpace(finderView.ViewBuffer()), searchTerms)
	elapsed := time.Since(t)
	if len(searchResults) > 0 {
		results.Title = fmt.Sprintf("Search results (%v matches in %v)", len(searchResults), elapsed)
	} else {
		results.Title = "Search results"
	}
	for _, match := range searchResults {
		for i := 0; i < len(match.Str); i++ {
			if contains(i, match.MatchedIndexes) {
				fmt.Fprintf(results, fmt.Sprintf("\033[1m%s\033[0m", string(match.Str[i])))
			} else {
				fmt.Fprintf(results, string(match.Str[i]))
			}
		}
		fmt.Fprintln(results, "")
	}
	return nil
}

func finder(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch {
	case ch != 0 && mod == 0:
		v.EditWrite(ch)
		g.Update(recalculateResults)
	case key == gocui.KeySpace:
		v.EditWrite(' ')
	case key == gocui.KeyBackspace || key == gocui.KeyBackspace2:
		v.EditDelete(true)
		g.Update(recalculateResults)
	case key == gocui.KeyInsert:
		v.Overwrite = !v.Overwrite
	}
}

func contains(needle int, haystack []int) bool {
	for _, i := range haystack {
		if needle == i {
			return true
		}
	}
	return false
}
