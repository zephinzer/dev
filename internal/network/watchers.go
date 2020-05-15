package network

import (
	"fmt"
	"sync"
	"time"

	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/internal/notifications"
	"github.com/zephinzer/dev/internal/types"
	pkgnetwork "github.com/zephinzer/dev/pkg/network"
)

func WatchConnections(
	networks []pkgnetwork.Network,
	updateInterval time.Duration,
	stop chan struct{},
) chan types.Notification {
	notificationsChannel := make(chan types.Notification, 16)
	isOnline := map[string]pkgnetwork.Network{}
	isOffline := map[string]pkgnetwork.Network{}
	isInitialRun := true
	go func(tick <-chan time.Time) {
		for {
			select {
			case <-stop:
				return
			case <-tick:
				log.Trace("network connections watcher triggered")
				var waiter sync.WaitGroup
				for _, nw := range networks {
					networkName := nw.Name
					if len(networkName) == 0 {
						networkName = "unnamed"
					}

					networkCheckURL := nw.Check.URL
					networkCheckMethod := nw.Check.Method
					if len(networkCheckMethod) == 0 {
						networkCheckMethod = constants.DefaultNetworkCheckMethod
					}
					waiter.Add(1)
					go func(nw pkgnetwork.Network) {
						stateChanged := false
						if runError := nw.Check.Run(); runError != nil {
							log.Warnf("unable to run network check '%s': %s",
								networkName,
								runError,
							)
							isOffline[networkName] = nw
							if _, ok := isOnline[networkName]; ok {
								delete(isOnline, networkName)
								stateChanged = true
							}
						} else if verificationError := nw.Check.Verify(); verificationError != nil {
							log.Warnf("unable to verify network check '%s' to %s '%s': %s",
								networkName,
								networkCheckMethod,
								networkCheckURL,
								verificationError,
							)
							isOffline[networkName] = nw
							if _, ok := isOnline[networkName]; ok {
								delete(isOnline, networkName)
								stateChanged = true
							}
						} else {
							isOnline[networkName] = nw
							if _, ok := isOffline[networkName]; ok {
								delete(isOffline, networkName)
								stateChanged = true
							}
						}
						if stateChanged && !isInitialRun {
							if _, ok := isOnline[networkName]; ok {
								notificationsChannel <- notifications.New(
									"Network change detected",
									fmt.Sprintf("Network [ %s ] is now ONLINE", networkName),
								)
							} else if _, ok := isOffline[networkName]; ok {
								notificationsChannel <- notifications.New(
									"Network change detected",
									fmt.Sprintf("Network [ %s ] is now OFFLINE", networkName),
								)
							}
						}
						waiter.Done()
					}(nw)
				}
				waiter.Wait()
				isInitialRun = false
			}
		}
	}(time.Tick(updateInterval))
	return notificationsChannel
}
