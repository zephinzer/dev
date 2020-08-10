package network

import (
	"fmt"
	"os"
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

type ConnectionParams struct {
	Hostname             string
	NotificationsChannel chan types.Notification
	StatusChannel        chan NetworkConnectionStatus
	IsInitialRun         bool
}

func WatchConnections(
	networks []pkgnetwork.Network,
	updateInterval time.Duration,
	stop chan struct{},
) chan types.Notification {
	// assign defaults
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown-host"
	}

	// intiialise
	connectionParams := &ConnectionParams{
		Hostname:             hostname,
		NotificationsChannel: make(chan types.Notification, 16),
		StatusChannel:        make(chan NetworkConnectionStatus, 8),
		IsInitialRun:         true,
	}

	// lets do dis
	ticker := time.NewTicker(updateInterval)
	go func(tick <-chan time.Time) {
		for {
			select {
			case <-stop:
				return
			case statusUpdate := <-connectionParams.StatusChannel:
				HandleStatusUpdate(statusUpdate, connectionParams)
			case <-tick:
				HandleCheckTrigger(networks, connectionParams)
			}
		}
	}(ticker.C)

	return connectionParams.NotificationsChannel
}

func HandleCheckTrigger(networksToCheck []pkgnetwork.Network, connectionParams *ConnectionParams) {
	log.Trace("network connections watcher triggered")
	var waiter sync.WaitGroup
	for _, nw := range networksToCheck {
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
				connectionParams.StatusChannel <- NetworkConnectionStatus{false, nw}
			} else if verificationError := nw.Check.Verify(); verificationError != nil {
				log.Warnf("unable to verify network check '%s' to %s '%s': %s",
					networkName,
					networkCheckMethod,
					networkCheckURL,
					verificationError,
				)
				connectionParams.StatusChannel <- NetworkConnectionStatus{false, nw}
			} else {
				connectionParams.StatusChannel <- NetworkConnectionStatus{true, nw}
			}
		}(nw)
	}
	waiter.Wait()
	connectionParams.IsInitialRun = false
}

func HandleStatusUpdate(statusUpdate NetworkConnectionStatus, connectionParams *ConnectionParams) {
	nwName := statusUpdate.Network.Name
	if statusUpdate.Online {
		isOnline[nwName] = statusUpdate.Network
		if _, ok := isOffline[nwName]; ok {
			delete(isOffline, nwName)
			if !connectionParams.IsInitialRun {
				connectionParams.NotificationsChannel <- notifications.New(
					"✅ Network Online",
					fmt.Sprintf("[`%s`@`%s`] is now ONLINE", nwName, connectionParams.Hostname),
				)
			}
		}
	} else {
		isOffline[nwName] = statusUpdate.Network
		if _, ok := isOnline[nwName]; ok {
			delete(isOnline, nwName)
			if !connectionParams.IsInitialRun {
				connectionParams.NotificationsChannel <- notifications.New(
					"⛔️ Network Offline",
					fmt.Sprintf("[`%s`@`%s`] is now OFFLINE", nwName, connectionParams.Hostname),
				)
			}
		}
	}
}
