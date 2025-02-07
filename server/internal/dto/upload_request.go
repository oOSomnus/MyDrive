package dto

type UploadRequestEntity struct {
	FolderId    string `json:"folderId"`
	Filename    string `json:"filename"`
	Fingerprint string `json:"fingerprint"`
	Size        uint64 `json:"size"`
}
