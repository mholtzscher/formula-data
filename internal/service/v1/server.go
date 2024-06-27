package v1

import (
	apiv1 "github.com/mholtzscher/formula-data/gen/api/v1/apiv1connect"
)

type FormulaDataServer struct {
	apiv1.UnimplementedFormulaDataServiceHandler
}
