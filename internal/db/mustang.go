package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/caring/go-packages/pkg/errors"
	"github.com/google/uuid"

	"github.com/caring/ford-mustang/pb"
)



// mustangService provides an API for interacting with the mustangs table
type mustangService struct {
	db    *sql.DB
	stmts map[string]*sql.Stmt
}

// Mustang is a struct representation of a row in the mustangs table
type Mustang struct {
	ID  	uuid.UUID
	Name  string
}

// protoMustang is an interface that most proto mustang objects will satisfy
type protoMustang interface {
	GetName() string
}

// NewMustang is a convenience helper cast a proto mustang to it's DB layer struct
func NewMustang(ID string, proto protoMustang) (*Mustang, error) {
	mID, err := ParseUUID(ID)
	if err != nil {
		return nil, err
	}

	return &Mustang{
		ID:  	mID,
		Name: proto.GetName(),
	}, nil
}

// ToProto casts a db mustang into a proto response object
func (m *Mustang) ToProto() *pb.MustangResponse {
	return &pb.MustangResponse{
		Id:  				m.ID.String(),
		Name:       m.Name,
	}
}

// Get fetches a single mustang from the db
func (svc *mustangService) Get(ctx context.Context, ID uuid.UUID) (*Mustang, error) {
	return svc.get(ctx, false, ID)
}

// GetTx fetches a single mustang from the db inside of a tx from ctx
func (svc *mustangService) GetTx(ctx context.Context, ID uuid.UUID) (*Mustang, error) {
	return svc.get(ctx, true, ID)
}

// get fetches a single mustang from the db
func (svc *mustangService) get(ctx context.Context, useTx bool, ID uuid.UUID) (*Mustang, error) {
	errMsg := func() string { return "Error executing get mustang - " + fmt.Sprint(ID) }

	var (
		stmt *sql.Stmt
		err  error
		tx   *sql.Tx
	)

	if useTx {

		if tx, err = FromCtx(ctx); err != nil {
			return nil, err
		}

		stmt = tx.Stmt(svc.stmts["get-mustang"])
	} else {
		stmt = svc.stmts["get-mustang"]
	}

	p := Mustang{}

	err = stmt.QueryRowContext(ctx, ID).
		Scan(&m.MustangID, &m.Name)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(ErrNotFound, errMsg())
		}

		return nil, errors.Wrap(err, errMsg())
	}

	return &p, nil
}

// Create a new mustang
func (svc *mustangService) Create(ctx context.Context, input *Mustang) error {
	return svc.create(ctx, false, input)
}

// CreateTx creates a new mustang withing a tx from ctx
func (svc *mustangService) CreateTx(ctx context.Context, input *Mustang) error {
	return svc.create(ctx, true, input)
}

// create a new mustang. if useTx = true then it will attempt to create the mustang within a transaction
// from context.
func (svc *mustangService) create(ctx context.Context, useTx bool, input *Mustang) error {
	errMsg := func() string { return "Error executing create mustang - " + fmt.Sprint(input) }

	var (
		stmt *sql.Stmt
		err  error
		tx   *sql.Tx
	)

	if useTx {

		if tx, err = FromCtx(ctx); err != nil {
			return err
		}

		stmt = tx.Stmt(svc.stmts["create-mustang"])
	} else {
		stmt = svc.stmts["create-mustang"]
	}

	result, err := stmt.ExecContext(ctx, input.MustangID, input.Name)
	if err != nil {
		return errors.Wrap(err, errMsg())
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, errMsg())
	}

	if rowCount == 0 {
		return errors.Wrap(ErrNotCreated, errMsg())
	}

	return nil
}

// Update updates a single mustang row in the DB
func (svc *mustangService) Update(ctx context.Context, input *Mustang) error {
	return svc.update(ctx, false, input)
}

// UpdateTx updates a single mustang row in the DB within a tx from ctx
func (svc *mustangService) UpdateTx(ctx context.Context, input *Mustang) error {
	return svc.update(ctx, true, input)
}

// update a mustang. if useTx = true then it will attempt to update the mustang within a transaction
// from context.
func (svc *mustangService) update(ctx context.Context, useTx bool, input *Mustang) error {
	errMsg := func() string { return "Error executing update mustang - " + fmt.Sprint(input) }

	var (
		stmt *sql.Stmt
		err  error
		tx   *sql.Tx
	)

	if useTx {

		if tx, err = FromCtx(ctx); err != nil {
			return err
		}

		stmt = tx.Stmt(svc.stmts["update-mustang"])
	} else {
		stmt = svc.stmts["update-mustang"]
	}

	result, err := stmt.ExecContext(ctx, input.Name, input.MustangID)
	if err != nil {
		return errors.Wrap(err, errMsg())
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, errMsg())
	}

	if rowCount == 0 {
		return errors.Wrap(ErrNoRowsAffected, errMsg())
	}

	return nil
}

// Delete sets deleted_at for a single mustangs row
func (svc *mustangService) Delete(ctx context.Context, ID uuid.UUID) error {
	return svc.delete(ctx, false, ID)
}

// DeleteTx sets deleted_at for a single mustangs row within a tx from ctx
func (svc *mustangService) DeleteTx(ctx context.Context, ID uuid.UUID) error {
	return svc.delete(ctx, true, ID)
}

// delete a mustang by setting deleted at. if useTx = true then it will attempt to delete the mustang within a transaction
// from context.
func (svc *mustangService) delete(ctx context.Context, useTx bool, ID uuid.UUID) error {
	errMsg := func() string { return "Error executing delete mustang - " + ID.String() }

	var (
		stmt *sql.Stmt
		err  error
		tx   *sql.Tx
	)

	if useTx {

		if tx, err = FromCtx(ctx); err != nil {
			return err
		}

		stmt = tx.Stmt(svc.stmts["delete-mustang"])
	} else {
		stmt = svc.stmts["delete-mustang"]
	}

	result, err := stmt.ExecContext(ctx, ID)
	if err != nil {
		return errors.Wrap(err, errMsg())
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, errMsg())
	}

	if rowCount == 0 {
		return errors.Wrap(ErrNotFound, errMsg())
	}

	return nil
}

