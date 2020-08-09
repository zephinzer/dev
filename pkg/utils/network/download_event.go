package network

import "fmt"

type DownloadEvent struct {
	State        DownloadState   `json:"download_state"`
	URL          string          `json:"url"`
	FilePath     string          `json:"file_path"`
	TempFilePath string          `json:"temp_file_path"`
	Status       *DownloadStatus `json:"status"`
}

func (de DownloadEvent) String() string {
	return fmt.Sprintf("[%s] downloading from '%s' into '%s' via '%s' [%v%% done]", de.State, de.URL, de.FilePath, de.TempFilePath, de.Status.GetPercentage())
}
