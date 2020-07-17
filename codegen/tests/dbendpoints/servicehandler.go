// Code generated by sysl DO NOT EDIT.
package dbendpoints

import (
	"database/sql"
	"net/http"

	"github.com/anz-bank/sysl-go/common"
	"github.com/anz-bank/sysl-go/convert"
	"github.com/anz-bank/sysl-go/core"
	"github.com/anz-bank/sysl-go/database"
	"github.com/anz-bank/sysl-go/restlib"
	"github.com/anz-bank/sysl-go/validator"
)

// Handler interface for DbEndpoints
type Handler interface {
	GetCompanyLocationListHandler(w http.ResponseWriter, r *http.Request)
}

// ServiceHandler for DbEndpoints API
type ServiceHandler struct {
	genCallback      core.RestGenCallback
	serviceInterface *ServiceInterface
	DB               *sql.DB
}

// NewServiceHandler for DbEndpoints
func NewServiceHandler(genCallback core.RestGenCallback, serviceInterface *ServiceInterface) *ServiceHandler {
	db, err := database.GetDBHandle()
	if err != nil {
		return nil
	}

	return &ServiceHandler{genCallback, serviceInterface, db}
}

// GetCompanyLocationListHandler ...
func (s *ServiceHandler) GetCompanyLocationListHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.GetCompanyLocationList == nil {
		common.HandleError(r.Context(), w, common.InternalError, "not implemented", nil, s.genCallback.MapError)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req GetCompanyLocationListRequest

	req.DeptLoc = restlib.GetQueryParam(r, "deptLoc")

	var CompanyNameParam string

	var convErr error
	CompanyNameParam = restlib.GetQueryParam(r, "companyName")
	req.CompanyName, convErr = convert.StringToStringPtr(ctx, CompanyNameParam)
	if convErr != nil {
		common.HandleError(ctx, w, common.BadRequestError, "Invalid request", convErr, s.genCallback.MapError)
		return
	}

	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	valErr := validator.Validate(&req)
	if valErr != nil {
		common.HandleError(ctx, w, common.BadRequestError, "Invalid request", valErr, s.genCallback.MapError)
		return
	}

	conn, err := s.DB.Conn(ctx)
	if err != nil {
		common.HandleError(ctx, w, common.InternalError, "Database connection could not be retrieved", err, s.genCallback.MapError)
		return
	}

	defer conn.Close()
	retrievebycompanyandlocationStmt, err_retrievebycompanyandlocation := conn.PrepareContext(ctx, "select company.abnnumber, company.companyname, company.companycountry, department.deptid, department.deptname, department.deptloc from company JOIN department ON company.abnnumber=department.abn WHERE department.deptloc=? and company.companyname=? order by company.abnnumber;")
	if err_retrievebycompanyandlocation != nil {
		common.HandleError(ctx, w, common.InternalError, "could not parse the sql query with the name sql_retrieveByCompanyAndLocation", err_retrievebycompanyandlocation, s.genCallback.MapError)
		return
	}

	tx, err := conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		common.HandleError(ctx, w, common.DownstreamUnavailableError, "DB Transaction could not be created", err, s.genCallback.MapError)
		return
	}

	client := GetCompanyLocationListClient{
		conn:                         conn,
		retrievebycompanyandlocation: retrievebycompanyandlocationStmt,
	}

	getcompanylocationresponse, err := s.serviceInterface.GetCompanyLocationList(ctx, &req, client)
	if err != nil {
		tx.Rollback()
		common.HandleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err, s.genCallback.MapError)
		return
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		common.HandleError(ctx, w, common.InternalError, "Failed to commit the transaction", commitErr, s.genCallback.MapError)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	if headermap.Get("Content-Type") == "" {
		headermap.Set("Content-Type", "application/json")
	}
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, getcompanylocationresponse)
}
