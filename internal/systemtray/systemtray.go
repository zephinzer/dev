package systemtray

import (
	"github.com/getlantern/systray"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
)

type MenuItemType string

const (
	TypeItem      MenuItemType = "item"
	TypeFolder    MenuItemType = "folder"
	TypeSeparator MenuItemType = "separator"
)

type Menu []MenuItem

type MenuItem struct {
	Type    MenuItemType
	Label   string
	Tooltip string
	Handler func()
	Menu    *Menu
}

func Start(menu Menu, stopped chan struct{}) {
	systray.SetTitle("dev")
	systray.SetTooltip("dev cli")
	systray.SetIcon(constants.DefaultSystrayIcon)
	systray.Run(func() {
		for _, item := range menu {
			switch item.Type {
			case TypeFolder:
				log.Debugf("adding menu folder to main system tray: %s (%s)", item.Label, item.Tooltip)
				addedItem := systray.AddMenuItem(item.Label, item.Tooltip)
				addNestedMenuItems(addedItem, &item)
			case TypeSeparator:
				log.Debug("adding separator to main system tray")
				systray.AddSeparator()
			default:
				log.Debugf("adding menu item to main system tray: %s (%s)", item.Label, item.Tooltip)
				addedItem := systray.AddMenuItem(item.Label, item.Tooltip)
				go func(handler func()) {
					for {
						select {
						case _, ok := <-addedItem.ClickedCh:
							if !ok {
								return
							}
							handler()
						}
					}
				}(item.Handler)
			}
		}
	}, func() {
		stopped <- struct{}{}
	})
}

func addNestedMenuItems(to *systray.MenuItem, item *MenuItem) {
	switch item.Type {
	case TypeFolder:
		for _, subItem := range *item.Menu {
			switch subItem.Type {
			case TypeSeparator:
			case TypeFolder:
				log.Debugf("adding sub-menu folder to menu item '%s': %s (%s)", item.Label, subItem.Label, subItem.Tooltip)
				addedItem := to.AddSubMenuItem(subItem.Label, subItem.Tooltip)
				addNestedMenuItems(addedItem, &subItem)
			default:
				log.Debugf("adding sub-menu item to menu item '%s': %s (%s)", item.Label, subItem.Label, subItem.Tooltip)
				addedItem := to.AddSubMenuItem(subItem.Label, subItem.Tooltip)
				go func(handler func()) {
					for {
						select {
						case _, ok := <-addedItem.ClickedCh:
							if !ok {
								return
							}
							handler()
						}
					}
				}(subItem.Handler)
			}
		}
	}
	return
}
