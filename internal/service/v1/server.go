package v1

import (
	apiv1 "github.com/mholtzscher/formula-data/gen/api/v1/apiv1connect"
	"github.com/mholtzscher/formula-data/internal/dal"
)

func New(db *dal.Queries) *FormulaDataServer {
	return &FormulaDataServer{
		DB: db,
	}
}

type FormulaDataServer struct {
	apiv1.UnimplementedFormulaDataServiceHandler
	DB *dal.Queries
}
