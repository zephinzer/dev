package network

// DownloadStatus contains progress information on a Download
type DownloadStatus struct {
	TotalBytes     uint64 `json:"total_bytes"`
	ProcessedBytes uint64 `json:"processed_bytes"`
}

// GetPercentage returns a floating point value representing the progress
func (ds DownloadStatus) GetPercentage() float64 {
	return float64(ds.ProcessedBytes) / float64(ds.TotalBytes) * 100
}

// Write implements io.Writer and logically updates the progress
func (ds *DownloadStatus) Write(content []byte) (int, error) {
	contentLength := len(content)
	ds.ProcessedBytes += uint64(contentLength)
	return contentLength, nil
}
