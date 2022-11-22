package handlers

import (
	"http-server/storage"
)

//Handler...
type Handler struct{
	Stg storage.StorageI
}