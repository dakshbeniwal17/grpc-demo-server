// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// HostGpu is an object representing the database table.
type HostGpu struct {
	ID        string       `boil:"id" json:"id" toml:"id" yaml:"id"`
	HostID    null.String  `boil:"host_id" json:"host_id,omitempty" toml:"host_id" yaml:"host_id,omitempty"`
	Updated   null.Time    `boil:"updated" json:"updated,omitempty" toml:"updated" yaml:"updated,omitempty"`
	Gpu       null.String  `boil:"gpu" json:"gpu,omitempty" toml:"gpu" yaml:"gpu,omitempty"`
	GpuNo     null.Int     `boil:"gpu_no" json:"gpu_no,omitempty" toml:"gpu_no" yaml:"gpu_no,omitempty"`
	SlotNo    null.Int     `boil:"slot_no" json:"slot_no,omitempty" toml:"slot_no" yaml:"slot_no,omitempty"`
	Available null.Int     `boil:"available" json:"available,omitempty" toml:"available" yaml:"available,omitempty"`
	Vram      null.Float64 `boil:"vram" json:"vram,omitempty" toml:"vram" yaml:"vram,omitempty"`
	VramFree  null.Float64 `boil:"vram_free" json:"vram_free,omitempty" toml:"vram_free" yaml:"vram_free,omitempty"`

	R *hostGpuR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L hostGpuL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var HostGpuColumns = struct {
	ID        string
	HostID    string
	Updated   string
	Gpu       string
	GpuNo     string
	SlotNo    string
	Available string
	Vram      string
	VramFree  string
}{
	ID:        "id",
	HostID:    "host_id",
	Updated:   "updated",
	Gpu:       "gpu",
	GpuNo:     "gpu_no",
	SlotNo:    "slot_no",
	Available: "available",
	Vram:      "vram",
	VramFree:  "vram_free",
}

var HostGpuTableColumns = struct {
	ID        string
	HostID    string
	Updated   string
	Gpu       string
	GpuNo     string
	SlotNo    string
	Available string
	Vram      string
	VramFree  string
}{
	ID:        "host_gpus.id",
	HostID:    "host_gpus.host_id",
	Updated:   "host_gpus.updated",
	Gpu:       "host_gpus.gpu",
	GpuNo:     "host_gpus.gpu_no",
	SlotNo:    "host_gpus.slot_no",
	Available: "host_gpus.available",
	Vram:      "host_gpus.vram",
	VramFree:  "host_gpus.vram_free",
}

// Generated where

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var HostGpuWhere = struct {
	ID        whereHelperstring
	HostID    whereHelpernull_String
	Updated   whereHelpernull_Time
	Gpu       whereHelpernull_String
	GpuNo     whereHelpernull_Int
	SlotNo    whereHelpernull_Int
	Available whereHelpernull_Int
	Vram      whereHelpernull_Float64
	VramFree  whereHelpernull_Float64
}{
	ID:        whereHelperstring{field: "\"host_gpus\".\"id\""},
	HostID:    whereHelpernull_String{field: "\"host_gpus\".\"host_id\""},
	Updated:   whereHelpernull_Time{field: "\"host_gpus\".\"updated\""},
	Gpu:       whereHelpernull_String{field: "\"host_gpus\".\"gpu\""},
	GpuNo:     whereHelpernull_Int{field: "\"host_gpus\".\"gpu_no\""},
	SlotNo:    whereHelpernull_Int{field: "\"host_gpus\".\"slot_no\""},
	Available: whereHelpernull_Int{field: "\"host_gpus\".\"available\""},
	Vram:      whereHelpernull_Float64{field: "\"host_gpus\".\"vram\""},
	VramFree:  whereHelpernull_Float64{field: "\"host_gpus\".\"vram_free\""},
}

// HostGpuRels is where relationship names are stored.
var HostGpuRels = struct {
	Host        string
	GpuGpuModel string
}{
	Host:        "Host",
	GpuGpuModel: "GpuGpuModel",
}

// hostGpuR is where relationships are stored.
type hostGpuR struct {
	Host        *Host     `boil:"Host" json:"Host" toml:"Host" yaml:"Host"`
	GpuGpuModel *GpuModel `boil:"GpuGpuModel" json:"GpuGpuModel" toml:"GpuGpuModel" yaml:"GpuGpuModel"`
}

// NewStruct creates a new relationship struct
func (*hostGpuR) NewStruct() *hostGpuR {
	return &hostGpuR{}
}

// hostGpuL is where Load methods for each relationship are stored.
type hostGpuL struct{}

var (
	hostGpuAllColumns            = []string{"id", "host_id", "updated", "gpu", "gpu_no", "slot_no", "available", "vram", "vram_free"}
	hostGpuColumnsWithoutDefault = []string{}
	hostGpuColumnsWithDefault    = []string{"id", "host_id", "updated", "gpu", "gpu_no", "slot_no", "available", "vram", "vram_free"}
	hostGpuPrimaryKeyColumns     = []string{"id"}
	hostGpuGeneratedColumns      = []string{}
)

type (
	// HostGpuSlice is an alias for a slice of pointers to HostGpu.
	// This should almost always be used instead of []HostGpu.
	HostGpuSlice []*HostGpu
	// HostGpuHook is the signature for custom HostGpu hook methods
	HostGpuHook func(context.Context, boil.ContextExecutor, *HostGpu) error

	hostGpuQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	hostGpuType                 = reflect.TypeOf(&HostGpu{})
	hostGpuMapping              = queries.MakeStructMapping(hostGpuType)
	hostGpuPrimaryKeyMapping, _ = queries.BindMapping(hostGpuType, hostGpuMapping, hostGpuPrimaryKeyColumns)
	hostGpuInsertCacheMut       sync.RWMutex
	hostGpuInsertCache          = make(map[string]insertCache)
	hostGpuUpdateCacheMut       sync.RWMutex
	hostGpuUpdateCache          = make(map[string]updateCache)
	hostGpuUpsertCacheMut       sync.RWMutex
	hostGpuUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var hostGpuAfterSelectHooks []HostGpuHook

var hostGpuBeforeInsertHooks []HostGpuHook
var hostGpuAfterInsertHooks []HostGpuHook

var hostGpuBeforeUpdateHooks []HostGpuHook
var hostGpuAfterUpdateHooks []HostGpuHook

var hostGpuBeforeDeleteHooks []HostGpuHook
var hostGpuAfterDeleteHooks []HostGpuHook

var hostGpuBeforeUpsertHooks []HostGpuHook
var hostGpuAfterUpsertHooks []HostGpuHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *HostGpu) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostGpuAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *HostGpu) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostGpuBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *HostGpu) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostGpuAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *HostGpu) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostGpuBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *HostGpu) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostGpuAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *HostGpu) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostGpuBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *HostGpu) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostGpuAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *HostGpu) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostGpuBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *HostGpu) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostGpuAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddHostGpuHook registers your hook function for all future operations.
func AddHostGpuHook(hookPoint boil.HookPoint, hostGpuHook HostGpuHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		hostGpuAfterSelectHooks = append(hostGpuAfterSelectHooks, hostGpuHook)
	case boil.BeforeInsertHook:
		hostGpuBeforeInsertHooks = append(hostGpuBeforeInsertHooks, hostGpuHook)
	case boil.AfterInsertHook:
		hostGpuAfterInsertHooks = append(hostGpuAfterInsertHooks, hostGpuHook)
	case boil.BeforeUpdateHook:
		hostGpuBeforeUpdateHooks = append(hostGpuBeforeUpdateHooks, hostGpuHook)
	case boil.AfterUpdateHook:
		hostGpuAfterUpdateHooks = append(hostGpuAfterUpdateHooks, hostGpuHook)
	case boil.BeforeDeleteHook:
		hostGpuBeforeDeleteHooks = append(hostGpuBeforeDeleteHooks, hostGpuHook)
	case boil.AfterDeleteHook:
		hostGpuAfterDeleteHooks = append(hostGpuAfterDeleteHooks, hostGpuHook)
	case boil.BeforeUpsertHook:
		hostGpuBeforeUpsertHooks = append(hostGpuBeforeUpsertHooks, hostGpuHook)
	case boil.AfterUpsertHook:
		hostGpuAfterUpsertHooks = append(hostGpuAfterUpsertHooks, hostGpuHook)
	}
}

// One returns a single hostGpu record from the query.
func (q hostGpuQuery) One(ctx context.Context, exec boil.ContextExecutor) (*HostGpu, error) {
	o := &HostGpu{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for host_gpus")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all HostGpu records from the query.
func (q hostGpuQuery) All(ctx context.Context, exec boil.ContextExecutor) (HostGpuSlice, error) {
	var o []*HostGpu

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to HostGpu slice")
	}

	if len(hostGpuAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all HostGpu records in the query.
func (q hostGpuQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count host_gpus rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q hostGpuQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if host_gpus exists")
	}

	return count > 0, nil
}

// Host pointed to by the foreign key.
func (o *HostGpu) Host(mods ...qm.QueryMod) hostQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.HostID),
	}

	queryMods = append(queryMods, mods...)

	query := Hosts(queryMods...)
	queries.SetFrom(query.Query, "\"hosts\"")

	return query
}

// GpuGpuModel pointed to by the foreign key.
func (o *HostGpu) GpuGpuModel(mods ...qm.QueryMod) gpuModelQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.Gpu),
	}

	queryMods = append(queryMods, mods...)

	query := GpuModels(queryMods...)
	queries.SetFrom(query.Query, "\"gpu_models\"")

	return query
}

// LoadHost allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (hostGpuL) LoadHost(ctx context.Context, e boil.ContextExecutor, singular bool, maybeHostGpu interface{}, mods queries.Applicator) error {
	var slice []*HostGpu
	var object *HostGpu

	if singular {
		object = maybeHostGpu.(*HostGpu)
	} else {
		slice = *maybeHostGpu.(*[]*HostGpu)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &hostGpuR{}
		}
		if !queries.IsNil(object.HostID) {
			args = append(args, object.HostID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &hostGpuR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.HostID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.HostID) {
				args = append(args, obj.HostID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`hosts`),
		qm.WhereIn(`hosts.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Host")
	}

	var resultSlice []*Host
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Host")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for hosts")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for hosts")
	}

	if len(hostGpuAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Host = foreign
		if foreign.R == nil {
			foreign.R = &hostR{}
		}
		foreign.R.HostGpus = append(foreign.R.HostGpus, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.HostID, foreign.ID) {
				local.R.Host = foreign
				if foreign.R == nil {
					foreign.R = &hostR{}
				}
				foreign.R.HostGpus = append(foreign.R.HostGpus, local)
				break
			}
		}
	}

	return nil
}

// LoadGpuGpuModel allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (hostGpuL) LoadGpuGpuModel(ctx context.Context, e boil.ContextExecutor, singular bool, maybeHostGpu interface{}, mods queries.Applicator) error {
	var slice []*HostGpu
	var object *HostGpu

	if singular {
		object = maybeHostGpu.(*HostGpu)
	} else {
		slice = *maybeHostGpu.(*[]*HostGpu)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &hostGpuR{}
		}
		if !queries.IsNil(object.Gpu) {
			args = append(args, object.Gpu)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &hostGpuR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.Gpu) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.Gpu) {
				args = append(args, obj.Gpu)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`gpu_models`),
		qm.WhereIn(`gpu_models.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load GpuModel")
	}

	var resultSlice []*GpuModel
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice GpuModel")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for gpu_models")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for gpu_models")
	}

	if len(hostGpuAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.GpuGpuModel = foreign
		if foreign.R == nil {
			foreign.R = &gpuModelR{}
		}
		foreign.R.GpuHostGpus = append(foreign.R.GpuHostGpus, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.Gpu, foreign.ID) {
				local.R.GpuGpuModel = foreign
				if foreign.R == nil {
					foreign.R = &gpuModelR{}
				}
				foreign.R.GpuHostGpus = append(foreign.R.GpuHostGpus, local)
				break
			}
		}
	}

	return nil
}

// SetHost of the hostGpu to the related item.
// Sets o.R.Host to related.
// Adds o to related.R.HostGpus.
func (o *HostGpu) SetHost(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Host) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"host_gpus\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"host_id"}),
		strmangle.WhereClause("\"", "\"", 2, hostGpuPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.HostID, related.ID)
	if o.R == nil {
		o.R = &hostGpuR{
			Host: related,
		}
	} else {
		o.R.Host = related
	}

	if related.R == nil {
		related.R = &hostR{
			HostGpus: HostGpuSlice{o},
		}
	} else {
		related.R.HostGpus = append(related.R.HostGpus, o)
	}

	return nil
}

// RemoveHost relationship.
// Sets o.R.Host to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *HostGpu) RemoveHost(ctx context.Context, exec boil.ContextExecutor, related *Host) error {
	var err error

	queries.SetScanner(&o.HostID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("host_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Host = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.HostGpus {
		if queries.Equal(o.HostID, ri.HostID) {
			continue
		}

		ln := len(related.R.HostGpus)
		if ln > 1 && i < ln-1 {
			related.R.HostGpus[i] = related.R.HostGpus[ln-1]
		}
		related.R.HostGpus = related.R.HostGpus[:ln-1]
		break
	}
	return nil
}

// SetGpuGpuModel of the hostGpu to the related item.
// Sets o.R.GpuGpuModel to related.
// Adds o to related.R.GpuHostGpus.
func (o *HostGpu) SetGpuGpuModel(ctx context.Context, exec boil.ContextExecutor, insert bool, related *GpuModel) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"host_gpus\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"gpu"}),
		strmangle.WhereClause("\"", "\"", 2, hostGpuPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.Gpu, related.ID)
	if o.R == nil {
		o.R = &hostGpuR{
			GpuGpuModel: related,
		}
	} else {
		o.R.GpuGpuModel = related
	}

	if related.R == nil {
		related.R = &gpuModelR{
			GpuHostGpus: HostGpuSlice{o},
		}
	} else {
		related.R.GpuHostGpus = append(related.R.GpuHostGpus, o)
	}

	return nil
}

// RemoveGpuGpuModel relationship.
// Sets o.R.GpuGpuModel to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *HostGpu) RemoveGpuGpuModel(ctx context.Context, exec boil.ContextExecutor, related *GpuModel) error {
	var err error

	queries.SetScanner(&o.Gpu, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("gpu")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.GpuGpuModel = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.GpuHostGpus {
		if queries.Equal(o.Gpu, ri.Gpu) {
			continue
		}

		ln := len(related.R.GpuHostGpus)
		if ln > 1 && i < ln-1 {
			related.R.GpuHostGpus[i] = related.R.GpuHostGpus[ln-1]
		}
		related.R.GpuHostGpus = related.R.GpuHostGpus[:ln-1]
		break
	}
	return nil
}

// HostGpus retrieves all the records using an executor.
func HostGpus(mods ...qm.QueryMod) hostGpuQuery {
	mods = append(mods, qm.From("\"host_gpus\""))
	return hostGpuQuery{NewQuery(mods...)}
}

// FindHostGpu retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindHostGpu(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*HostGpu, error) {
	hostGpuObj := &HostGpu{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"host_gpus\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, hostGpuObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from host_gpus")
	}

	if err = hostGpuObj.doAfterSelectHooks(ctx, exec); err != nil {
		return hostGpuObj, err
	}

	return hostGpuObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *HostGpu) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no host_gpus provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(hostGpuColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	hostGpuInsertCacheMut.RLock()
	cache, cached := hostGpuInsertCache[key]
	hostGpuInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			hostGpuAllColumns,
			hostGpuColumnsWithDefault,
			hostGpuColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(hostGpuType, hostGpuMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(hostGpuType, hostGpuMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"host_gpus\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"host_gpus\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into host_gpus")
	}

	if !cached {
		hostGpuInsertCacheMut.Lock()
		hostGpuInsertCache[key] = cache
		hostGpuInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the HostGpu.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *HostGpu) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	hostGpuUpdateCacheMut.RLock()
	cache, cached := hostGpuUpdateCache[key]
	hostGpuUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			hostGpuAllColumns,
			hostGpuPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update host_gpus, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"host_gpus\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, hostGpuPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(hostGpuType, hostGpuMapping, append(wl, hostGpuPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update host_gpus row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for host_gpus")
	}

	if !cached {
		hostGpuUpdateCacheMut.Lock()
		hostGpuUpdateCache[key] = cache
		hostGpuUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q hostGpuQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for host_gpus")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for host_gpus")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o HostGpuSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), hostGpuPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"host_gpus\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, hostGpuPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in hostGpu slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all hostGpu")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *HostGpu) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no host_gpus provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(hostGpuColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	hostGpuUpsertCacheMut.RLock()
	cache, cached := hostGpuUpsertCache[key]
	hostGpuUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			hostGpuAllColumns,
			hostGpuColumnsWithDefault,
			hostGpuColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			hostGpuAllColumns,
			hostGpuPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert host_gpus, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(hostGpuPrimaryKeyColumns))
			copy(conflict, hostGpuPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"host_gpus\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(hostGpuType, hostGpuMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(hostGpuType, hostGpuMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert host_gpus")
	}

	if !cached {
		hostGpuUpsertCacheMut.Lock()
		hostGpuUpsertCache[key] = cache
		hostGpuUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single HostGpu record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *HostGpu) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no HostGpu provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), hostGpuPrimaryKeyMapping)
	sql := "DELETE FROM \"host_gpus\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from host_gpus")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for host_gpus")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q hostGpuQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no hostGpuQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from host_gpus")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for host_gpus")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o HostGpuSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(hostGpuBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), hostGpuPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"host_gpus\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, hostGpuPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from hostGpu slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for host_gpus")
	}

	if len(hostGpuAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *HostGpu) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindHostGpu(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *HostGpuSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := HostGpuSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), hostGpuPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"host_gpus\".* FROM \"host_gpus\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, hostGpuPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in HostGpuSlice")
	}

	*o = slice

	return nil
}

// HostGpuExists checks if the HostGpu row exists.
func HostGpuExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"host_gpus\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if host_gpus exists")
	}

	return exists, nil
}