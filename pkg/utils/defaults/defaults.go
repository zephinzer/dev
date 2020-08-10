package defaults

import "io"

func GetIoReader(theDefault io.Reader, useIfAvailable ...io.Reader) io.Reader {
	useThis := theDefault
	if len(useIfAvailable) > 0 {
		useThis = useIfAvailable[0]
	}
	return useThis
}
