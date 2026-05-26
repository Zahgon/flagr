package handler

import (
	"github.com/openflagr/flagr/pkg/entity"
	"github.com/openflagr/flagr/swagger_gen/restapi/operations"
)

var getDB = entity.GetDB

// Setup initialize all the handler functions
func Setup(api *operations.FlagrAPI) { _ = "STUB: not implemented"; return }

func setupCRUD(api *operations.FlagrAPI) { _ = "STUB: not implemented"; return }

func setupEvaluation(api *operations.FlagrAPI) { _ = "STUB: not implemented"; return }

func setupHealth(api *operations.FlagrAPI) { _ = "STUB: not implemented"; return }

func setupExport(api *operations.FlagrAPI) { _ = "STUB: not implemented"; return }
