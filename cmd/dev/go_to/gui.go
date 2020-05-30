package go_to

import (
	"fmt"
	"os"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/sahilm/fuzzy"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/pkg/utils"
)

const (
	finderTitleInactive = "Going somewhere? (hit enter to open link, ctrl+c to quit)"
	finderTitleActive   = "Going somewhere? (type to filter items, ctrl+c to quit)"
	noResultsTitle      = "All items"
	resultsTitle        = "Search results"
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

func handleFinderNavigateRight(g *gocui.Gui, view *gocui.View) error {
	switch view.Name() {
	case "finder":
		curX, _ := view.Cursor()
		if curX < len(view.ViewBuffer())-1 {
			view.MoveCursor(1, 0, true)
		}
	}
	return nil
}

func handleFinderNavigateLeft(g *gocui.Gui, view *gocui.View) error {
	switch view.Name() {
	case "finder":
		curX, _ := view.Cursor()
		if curX > 0 {
			view.MoveCursor(-1, 0, true)
		}
	}
	return nil
}

func handleFinderDone(g *gocui.Gui, v *gocui.View) error {
	if selectionIndex == -1 {
		selectionIndex = 0
		if _, err := g.SetCurrentView("results"); err != nil {
			return err
		}
		g.Cursor = false
		resultsView, err := g.View("results")
		if err != nil {
			return err
		}
		g.Update(recalculateResults)
		v.Title = finderTitleInactive
		resultsView.SetCursor(0, 0)
	}
	return nil
}

func handleResultsNavigateDown(g *gocui.Gui, v *gocui.View) error {
	selectionIndex++
	if selectionIndex >= len(searchResults) {
		selectionIndex--
		return nil
	}
	g.Update(recalculateResults)
	v.MoveCursor(0, 1, true)
	return nil
}

func handleResultsNavigateUp(g *gocui.Gui, v *gocui.View) error {
	selectionIndex--
	if selectionIndex == -1 {
		if _, err := g.SetCurrentView("finder"); err != nil {
			return err
		}
		g.Cursor = true
		finderView, err := g.View("finder")
		if err != nil {
			return err
		}
		finderView.Title = finderTitleActive
	}
	g.Update(recalculateResults)
	v.MoveCursor(0, -1, true)
	return nil
}

func handleResultsDone(g *gocui.Gui, v *gocui.View) error {
	selectionIndex = -1
	if _, err := g.SetCurrentView("finder"); err != nil {
		return err
	}
	g.Cursor = true
	finderView, err := g.View("finder")
	if err != nil {
	}
	if len(finderView.ViewBuffer()) > 0 {
		v.MoveCursor(len(finderView.ViewBuffer())-1, 0, true)
	}
	finderView.EditDelete(true)
	g.Update(recalculateResults)
	return nil
}

func handleResultsFound(g *gocui.Gui, v *gocui.View) error {
	targetString := strings.Split(searchResults[selectionIndex].Str, " ")
	targetURI := targetString[len(targetString)-1]
	utils.OpenURIWithDefaultApplication(targetURI)
	return gocui.ErrQuit
}

func startFuzzySearchInterface() {
	links := config.Global.Links
	repositories := config.Global.Repositories
	for _, link := range links {
		searchTerms = append(searchTerms, fmt.Sprintf("%s [%s] @ %s", link.Label, strings.Join(link.Categories, ", "), link.URL))
	}
	for _, repo := range repositories {
		repoURL, getWebsiteError := repo.GetWebsiteURL()
		if getWebsiteError == nil {
			searchTerms = append(searchTerms, fmt.Sprintf("repo['%s'] (%s) - %s @ %s", repo.Name, strings.Join(repo.Workspaces, ", "), repo.Description, repoURL))
		}
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

	if err := g.SetKeybinding("finder", gocui.KeyArrowRight, gocui.ModNone, handleFinderNavigateRight); err != nil {
		log.Error(err)
	}
	if err := g.SetKeybinding("finder", gocui.KeyArrowLeft, gocui.ModNone, handleFinderNavigateLeft); err != nil {
		log.Error(err)
	}
	if err := g.SetKeybinding("finder", gocui.KeyArrowDown, gocui.ModNone, handleFinderDone); err != nil {
		log.Error(err)
	}
	if err := g.SetKeybinding("finder", gocui.KeyEnter, gocui.ModNone, handleFinderDone); err != nil {
		log.Error(err)
	}
	if err := g.SetKeybinding("results", gocui.KeyArrowDown, gocui.ModNone, handleResultsNavigateDown); err != nil {
		log.Error(err)
	}
	if err := g.SetKeybinding("results", gocui.KeyArrowUp, gocui.ModNone, handleResultsNavigateUp); err != nil {
		log.Error(err)
	}
	if err := g.SetKeybinding("results", gocui.KeyBackspace2, gocui.ModNone, handleResultsDone); err != nil {
		log.Error(err)
	}
	if err := g.SetKeybinding("results", gocui.KeyEnter, gocui.ModNone, handleResultsFound); err != nil {
		log.Error(err)
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		results, err := g.View("results")
		if err != nil {
		}
		results.Title = fmt.Sprintf("%s (selectionIndex: %v)", resultsTitle, selectionIndex)
		return nil
	}); err != nil {
		log.Error(err)
	}

	if err := g.SetKeybinding("", gocui.KeyEsc, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return gocui.ErrQuit
	}); err != nil {
		log.Error(err)
	}

	g.Update(recalculateResults)
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
		finderView.Title = finderTitleActive
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
	searchResults = fuzzy.Find(strings.TrimSpace(finderView.ViewBuffer()), searchTerms)
	results.Title = fmt.Sprintf("%s (%v matches)", resultsTitle, len(searchResults))
	if len(searchResults) == 0 {
		results.Title = noResultsTitle
		for _, searchTerm := range searchTerms {
			searchResults = append(searchResults, fuzzy.Match{Str: searchTerm})
		}
	}

	for index, match := range searchResults {
		if selectionIndex == index {
			fmt.Fprintf(results, fmt.Sprintf("\033[1m> %s\033[0m", string(match.Str)))
		} else {
			fmt.Fprintf(results, "  ")
			for i := 0; i < len(match.Str); i++ {
				if utils.ContainsInt(i, match.MatchedIndexes) {
					fmt.Fprintf(results, fmt.Sprintf("\033[1m%s\033[0m", string(match.Str[i])))
				} else {
					fmt.Fprintf(results, string(match.Str[i]))
				}
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
