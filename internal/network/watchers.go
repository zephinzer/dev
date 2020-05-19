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

var (
	isOnline  = map[string]pkgnetwork.Network{}
	isOffline = map[string]pkgnetwork.Network{}
)

type NetworkConnectionStatus struct {
	Online  bool
	Network pkgnetwork.Network
}

func WatchConnections(
	networks []pkgnetwork.Network,
	updateInterval time.Duration,
	stop chan struct{},
) chan types.Notification {
	notificationsChannel := make(chan types.Notification, 16)
	statusChannel := make(chan NetworkConnectionStatus, 8)
	isInitialRun := true
	go func(tick <-chan time.Time) {
		for {
			select {
			case <-stop:
				return
			case statusUpdate := <-statusChannel:
				nwName := statusUpdate.Network.Name
				if statusUpdate.Online {
					isOnline[nwName] = statusUpdate.Network
					if _, ok := isOffline[nwName]; ok {
						delete(isOffline, nwName)
						if !isInitialRun {
							notificationsChannel <- notifications.New(
								"Network change detected",
								fmt.Sprintf("Network [ %s ] is now ONLINE", nwName),
							)
						}
					}
				} else {
					isOffline[nwName] = statusUpdate.Network
					if _, ok := isOnline[nwName]; ok {
						delete(isOnline, nwName)
						if !isInitialRun {
							notificationsChannel <- notifications.New(
								"Network change detected",
								fmt.Sprintf("Network [ %s ] is now OFFLINE", nwName),
							)
						}
					}
				}
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
						defer waiter.Done()
						if runError := nw.Check.Run(); runError != nil {
							log.Warnf("unable to run network check '%s': %s", networkName, runError)
							statusChannel <- NetworkConnectionStatus{false, nw}
						} else if verificationError := nw.Check.Verify(); verificationError != nil {
							log.Warnf("unable to verify network check '%s' to %s '%s': %s",
								networkName,
								networkCheckMethod,
								networkCheckURL,
								verificationError,
							)
							statusChannel <- NetworkConnectionStatus{false, nw}
						} else {
							statusChannel <- NetworkConnectionStatus{true, nw}
						}
					}(nw)
				}
				waiter.Wait()
				isInitialRun = false
			}
		}
	}(time.Tick(updateInterval))
	return notificationsChannel
}
