package codes

import (
	"net/http"

	"github.com/moshenahmias/failure"
)

const StatusLabel = "status"

var (
	ErrArticleNotFound = failure.Build("article not found").WithField("status", http.StatusNotFound).Done()

	// validations
	ErrNumberOfImagesExceeded   = failure.Build("number of images exceeded").WithField(StatusLabel, http.StatusBadRequest).Done()
	ErrImageSizeExceeded        = failure.Build("image size exceeded").WithField(StatusLabel, http.StatusBadRequest).Done()
	ErrInvalidImageType         = failure.Build("invalid image type").WithField(StatusLabel, http.StatusBadRequest).Done()
	ErrInvalidExpirationDate    = failure.Build("expiration date needs to be in the future").WithField(StatusLabel, http.StatusBadRequest).Done()
	ErrTitleIsRequired          = failure.Build("title is a required field").WithField(StatusLabel, http.StatusBadRequest).Done()
	ErrPathIsRequired           = failure.Build("path is a required field").WithField(StatusLabel, http.StatusBadRequest).Done()
	ErrDescriptionIsRequired    = failure.Build("title is a required field").WithField(StatusLabel, http.StatusBadRequest).Done()
	ErrDescriptionIsTooLong     = failure.Build("description is too long. max: 4000 chars").WithField(StatusLabel, http.StatusBadRequest).Done()
	ErrExpirationDateIsRequired = failure.Build("expirationDate is a required field").WithField(StatusLabel, http.StatusBadRequest).Done()
	ErrInvalidArticleID         = failure.Build("invalid article id").WithField(StatusLabel, http.StatusBadRequest).Done()
)
