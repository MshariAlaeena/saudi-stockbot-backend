package handler

import (
	"mime/multipart"
)

type HandlerResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewResponse(data interface{}, message string) HandlerResponse {
	return HandlerResponse{
		Data:    data,
		Message: message,
	}
}

type ChatResponseDTO struct {
	Answer string `json:"answer"`
}

type UploadRequestDTO struct {
	File *multipart.FileHeader `form:"file"  binding:"required"`
}

type UploadResponseDTO struct {
	Answer string `json:"answer"`
}

type Extension string

const (
	ExtensionPDF  Extension = ".pdf"
	ExtensionDOC  Extension = ".doc"
	ExtensionJPG  Extension = ".jpg"
	ExtensionDOCX Extension = ".docx"
	ExtensionTXT  Extension = ".txt"
	ExtensionCSV  Extension = ".csv"
	ExtensionXLSX Extension = ".xlsx"
	ExtensionXLS  Extension = ".xls"
	ExtensionPPTX Extension = ".pptx"
	ExtensionPPT  Extension = ".ppt"
	ExtensionJPEG Extension = ".jpeg"
	ExtensionPNG  Extension = ".png"
)

type PaginationRequest struct {
	Page     int `form:"page,default=1"    binding:"min=1"`
	PageSize int `form:"page_size,default=10" binding:"min=1,max=50"`
}

type GetDocumentsResponseDTO struct {
	Documents []Documents `json:"documents"`
	PageSize  int         `json:"page_size"`
	Page      int         `json:"page"`
	Total     int         `json:"total"`
}

type Documents struct {
	DocumentID        string             `json:"document_id"`
	DocumentName      string             `json:"document_name"`
	DocumentExtension Extension          `json:"document_extension"`
	ExtractedContent  []ExtractedContent `json:"extracted_content"`
	UploadedAt        string             `json:"uploaded_at"`
}

type ExtractedContent struct {
	ContentID string `json:"content_id"`
	Content   string `json:"content"`
}
