package go_to

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/sahilm/fuzzy"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/pkg/utils"
)

var g *gocui.Gui
var err error
var searchTerms []string
var searchResults fuzzy.Matches
var selectionIndex int = -1

func bindSigintHandler() {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, func(_ *gocui.Gui, _ *gocui.View) error {
		return gocui.ErrQuit
	}); err != nil {
		log.Error(err)
	}
}

func startFuzzySearchInterface() {
	links := config.Global.Links
	for _, link := range links {
		searchTerms = append(searchTerms, fmt.Sprintf("%s [%s] @ %s", link.Label, strings.Join(link.Categories, ", "), link.URL))
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
	bindSigintHandler()

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
		targetURI := targetString[len(targetString)-1]
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

func layout(terminalGUI *gocui.Gui) error {
	maxX, maxY := terminalGUI.Size()
	if finderView, setViewError := terminalGUI.SetView("finder", 0, 0, maxX, 10); setViewError != nil {
		if setViewError != gocui.ErrUnknownView {
			return setViewError
		}
		finderView.Wrap = true
		finderView.Editable = true
		finderView.Frame = true
		finderView.Title = "Going somewhere?"
		if _, setCurrentViewError := terminalGUI.SetCurrentView("finder"); setCurrentViewError != nil {
			return setCurrentViewError
		}
		finderView.Editor = gocui.EditorFunc(finderController)
	}

	if resultsView, setViewError := terminalGUI.SetView("results", 0, 3, maxX, maxY); setViewError != nil {
		if setViewError != gocui.ErrUnknownView {
			return setViewError
		}
		resultsView.Editable = false
		resultsView.Wrap = true
		resultsView.Frame = true
		resultsView.Title = "Search results"
	}
	return nil
}

func recalculateResults(terminalGUI *gocui.Gui) error {
	finderView, err := terminalGUI.View("finder")
	if err != nil {
		return err
	}
	results, err := terminalGUI.View("results")
	if err != nil {
		return err
	}
	results.Clear()
	startingSearchAt := time.Now()
	searchResults = fuzzy.Find(strings.TrimSpace(finderView.ViewBuffer()), searchTerms)
	elapsed := time.Since(startingSearchAt)
	if len(searchResults) > 0 {
		results.Title = fmt.Sprintf("Search results (%v matches in %v)", len(searchResults), elapsed)
	} else {
		results.Title = "Search results"
	}
	for _, match := range searchResults {
		for i := 0; i < len(match.Str); i++ {
			if utils.ContainsInt(i, match.MatchedIndexes) {
				fmt.Fprintf(results, fmt.Sprintf("\033[1m%s\033[0m", string(match.Str[i])))
			} else {
				fmt.Fprintf(results, string(match.Str[i]))
			}
		}
		fmt.Fprintln(results, "")
	}
	return nil
}

func finderController(view *gocui.View, key gocui.Key, input rune, modifier gocui.Modifier) {
	switch {
	case input != 0 && modifier == 0:
		view.EditWrite(input)
		g.Update(recalculateResults)
	case key == gocui.KeySpace:
		view.EditWrite(' ')
	case key == gocui.KeyBackspace || key == gocui.KeyBackspace2:
		view.EditDelete(true)
		g.Update(recalculateResults)
	case key == gocui.KeyInsert:
		view.Overwrite = !view.Overwrite
	}
}
