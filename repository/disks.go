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

// Disk is an object representing the database table.
type Disk struct {
	ID         string       `boil:"id" json:"id" toml:"id" yaml:"id"`
	DiskNo     null.Int     `boil:"disk_no" json:"disk_no,omitempty" toml:"disk_no" yaml:"disk_no,omitempty"`
	Type       null.String  `boil:"type" json:"type,omitempty" toml:"type" yaml:"type,omitempty"`
	Capacity   null.Float64 `boil:"capacity" json:"capacity,omitempty" toml:"capacity" yaml:"capacity,omitempty"`
	ReadSpeed  null.Int     `boil:"read_speed" json:"read_speed,omitempty" toml:"read_speed" yaml:"read_speed,omitempty"`
	WriteSpeed null.Int     `boil:"write_speed" json:"write_speed,omitempty" toml:"write_speed" yaml:"write_speed,omitempty"`
	HostID     null.String  `boil:"host_id" json:"host_id,omitempty" toml:"host_id" yaml:"host_id,omitempty"`

	R *diskR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L diskL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DiskColumns = struct {
	ID         string
	DiskNo     string
	Type       string
	Capacity   string
	ReadSpeed  string
	WriteSpeed string
	HostID     string
}{
	ID:         "id",
	DiskNo:     "disk_no",
	Type:       "type",
	Capacity:   "capacity",
	ReadSpeed:  "read_speed",
	WriteSpeed: "write_speed",
	HostID:     "host_id",
}

var DiskTableColumns = struct {
	ID         string
	DiskNo     string
	Type       string
	Capacity   string
	ReadSpeed  string
	WriteSpeed string
	HostID     string
}{
	ID:         "disks.id",
	DiskNo:     "disks.disk_no",
	Type:       "disks.type",
	Capacity:   "disks.capacity",
	ReadSpeed:  "disks.read_speed",
	WriteSpeed: "disks.write_speed",
	HostID:     "disks.host_id",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_Int struct{ field string }

func (w whereHelpernull_Int) EQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int) NEQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int) LT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int) LTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int) GT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int) GTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Int) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpernull_Float64 struct{ field string }

func (w whereHelpernull_Float64) EQ(x null.Float64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Float64) NEQ(x null.Float64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Float64) LT(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Float64) LTE(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Float64) GT(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Float64) GTE(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Float64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Float64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var DiskWhere = struct {
	ID         whereHelperstring
	DiskNo     whereHelpernull_Int
	Type       whereHelpernull_String
	Capacity   whereHelpernull_Float64
	ReadSpeed  whereHelpernull_Int
	WriteSpeed whereHelpernull_Int
	HostID     whereHelpernull_String
}{
	ID:         whereHelperstring{field: "\"disks\".\"id\""},
	DiskNo:     whereHelpernull_Int{field: "\"disks\".\"disk_no\""},
	Type:       whereHelpernull_String{field: "\"disks\".\"type\""},
	Capacity:   whereHelpernull_Float64{field: "\"disks\".\"capacity\""},
	ReadSpeed:  whereHelpernull_Int{field: "\"disks\".\"read_speed\""},
	WriteSpeed: whereHelpernull_Int{field: "\"disks\".\"write_speed\""},
	HostID:     whereHelpernull_String{field: "\"disks\".\"host_id\""},
}

// DiskRels is where relationship names are stored.
var DiskRels = struct {
	Host string
}{
	Host: "Host",
}

// diskR is where relationships are stored.
type diskR struct {
	Host *Host `boil:"Host" json:"Host" toml:"Host" yaml:"Host"`
}

// NewStruct creates a new relationship struct
func (*diskR) NewStruct() *diskR {
	return &diskR{}
}

// diskL is where Load methods for each relationship are stored.
type diskL struct{}

var (
	diskAllColumns            = []string{"id", "disk_no", "type", "capacity", "read_speed", "write_speed", "host_id"}
	diskColumnsWithoutDefault = []string{}
	diskColumnsWithDefault    = []string{"id", "disk_no", "type", "capacity", "read_speed", "write_speed", "host_id"}
	diskPrimaryKeyColumns     = []string{"id"}
	diskGeneratedColumns      = []string{}
)

type (
	// DiskSlice is an alias for a slice of pointers to Disk.
	// This should almost always be used instead of []Disk.
	DiskSlice []*Disk
	// DiskHook is the signature for custom Disk hook methods
	DiskHook func(context.Context, boil.ContextExecutor, *Disk) error

	diskQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	diskType                 = reflect.TypeOf(&Disk{})
	diskMapping              = queries.MakeStructMapping(diskType)
	diskPrimaryKeyMapping, _ = queries.BindMapping(diskType, diskMapping, diskPrimaryKeyColumns)
	diskInsertCacheMut       sync.RWMutex
	diskInsertCache          = make(map[string]insertCache)
	diskUpdateCacheMut       sync.RWMutex
	diskUpdateCache          = make(map[string]updateCache)
	diskUpsertCacheMut       sync.RWMutex
	diskUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var diskAfterSelectHooks []DiskHook

var diskBeforeInsertHooks []DiskHook
var diskAfterInsertHooks []DiskHook

var diskBeforeUpdateHooks []DiskHook
var diskAfterUpdateHooks []DiskHook

var diskBeforeDeleteHooks []DiskHook
var diskAfterDeleteHooks []DiskHook

var diskBeforeUpsertHooks []DiskHook
var diskAfterUpsertHooks []DiskHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Disk) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range diskAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Disk) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range diskBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Disk) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range diskAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Disk) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range diskBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Disk) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range diskAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Disk) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range diskBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Disk) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range diskAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Disk) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range diskBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Disk) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range diskAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDiskHook registers your hook function for all future operations.
func AddDiskHook(hookPoint boil.HookPoint, diskHook DiskHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		diskAfterSelectHooks = append(diskAfterSelectHooks, diskHook)
	case boil.BeforeInsertHook:
		diskBeforeInsertHooks = append(diskBeforeInsertHooks, diskHook)
	case boil.AfterInsertHook:
		diskAfterInsertHooks = append(diskAfterInsertHooks, diskHook)
	case boil.BeforeUpdateHook:
		diskBeforeUpdateHooks = append(diskBeforeUpdateHooks, diskHook)
	case boil.AfterUpdateHook:
		diskAfterUpdateHooks = append(diskAfterUpdateHooks, diskHook)
	case boil.BeforeDeleteHook:
		diskBeforeDeleteHooks = append(diskBeforeDeleteHooks, diskHook)
	case boil.AfterDeleteHook:
		diskAfterDeleteHooks = append(diskAfterDeleteHooks, diskHook)
	case boil.BeforeUpsertHook:
		diskBeforeUpsertHooks = append(diskBeforeUpsertHooks, diskHook)
	case boil.AfterUpsertHook:
		diskAfterUpsertHooks = append(diskAfterUpsertHooks, diskHook)
	}
}

// One returns a single disk record from the query.
func (q diskQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Disk, error) {
	o := &Disk{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for disks")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Disk records from the query.
func (q diskQuery) All(ctx context.Context, exec boil.ContextExecutor) (DiskSlice, error) {
	var o []*Disk

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Disk slice")
	}

	if len(diskAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Disk records in the query.
func (q diskQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count disks rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q diskQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if disks exists")
	}

	return count > 0, nil
}

// Host pointed to by the foreign key.
func (o *Disk) Host(mods ...qm.QueryMod) hostQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.HostID),
	}

	queryMods = append(queryMods, mods...)

	query := Hosts(queryMods...)
	queries.SetFrom(query.Query, "\"hosts\"")

	return query
}

// LoadHost allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (diskL) LoadHost(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDisk interface{}, mods queries.Applicator) error {
	var slice []*Disk
	var object *Disk

	if singular {
		object = maybeDisk.(*Disk)
	} else {
		slice = *maybeDisk.(*[]*Disk)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &diskR{}
		}
		if !queries.IsNil(object.HostID) {
			args = append(args, object.HostID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &diskR{}
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

	if len(diskAfterSelectHooks) != 0 {
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
		foreign.R.Disks = append(foreign.R.Disks, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.HostID, foreign.ID) {
				local.R.Host = foreign
				if foreign.R == nil {
					foreign.R = &hostR{}
				}
				foreign.R.Disks = append(foreign.R.Disks, local)
				break
			}
		}
	}

	return nil
}

// SetHost of the disk to the related item.
// Sets o.R.Host to related.
// Adds o to related.R.Disks.
func (o *Disk) SetHost(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Host) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"disks\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"host_id"}),
		strmangle.WhereClause("\"", "\"", 2, diskPrimaryKeyColumns),
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
		o.R = &diskR{
			Host: related,
		}
	} else {
		o.R.Host = related
	}

	if related.R == nil {
		related.R = &hostR{
			Disks: DiskSlice{o},
		}
	} else {
		related.R.Disks = append(related.R.Disks, o)
	}

	return nil
}

// RemoveHost relationship.
// Sets o.R.Host to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Disk) RemoveHost(ctx context.Context, exec boil.ContextExecutor, related *Host) error {
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

	for i, ri := range related.R.Disks {
		if queries.Equal(o.HostID, ri.HostID) {
			continue
		}

		ln := len(related.R.Disks)
		if ln > 1 && i < ln-1 {
			related.R.Disks[i] = related.R.Disks[ln-1]
		}
		related.R.Disks = related.R.Disks[:ln-1]
		break
	}
	return nil
}

// Disks retrieves all the records using an executor.
func Disks(mods ...qm.QueryMod) diskQuery {
	mods = append(mods, qm.From("\"disks\""))
	return diskQuery{NewQuery(mods...)}
}

// FindDisk retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDisk(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Disk, error) {
	diskObj := &Disk{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"disks\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, diskObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from disks")
	}

	if err = diskObj.doAfterSelectHooks(ctx, exec); err != nil {
		return diskObj, err
	}

	return diskObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Disk) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no disks provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(diskColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	diskInsertCacheMut.RLock()
	cache, cached := diskInsertCache[key]
	diskInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			diskAllColumns,
			diskColumnsWithDefault,
			diskColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(diskType, diskMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(diskType, diskMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"disks\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"disks\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into disks")
	}

	if !cached {
		diskInsertCacheMut.Lock()
		diskInsertCache[key] = cache
		diskInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Disk.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Disk) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	diskUpdateCacheMut.RLock()
	cache, cached := diskUpdateCache[key]
	diskUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			diskAllColumns,
			diskPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update disks, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"disks\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, diskPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(diskType, diskMapping, append(wl, diskPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update disks row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for disks")
	}

	if !cached {
		diskUpdateCacheMut.Lock()
		diskUpdateCache[key] = cache
		diskUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q diskQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for disks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for disks")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DiskSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), diskPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"disks\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, diskPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in disk slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all disk")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Disk) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no disks provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(diskColumnsWithDefault, o)

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

	diskUpsertCacheMut.RLock()
	cache, cached := diskUpsertCache[key]
	diskUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			diskAllColumns,
			diskColumnsWithDefault,
			diskColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			diskAllColumns,
			diskPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert disks, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(diskPrimaryKeyColumns))
			copy(conflict, diskPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"disks\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(diskType, diskMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(diskType, diskMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert disks")
	}

	if !cached {
		diskUpsertCacheMut.Lock()
		diskUpsertCache[key] = cache
		diskUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Disk record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Disk) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Disk provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), diskPrimaryKeyMapping)
	sql := "DELETE FROM \"disks\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from disks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for disks")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q diskQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no diskQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from disks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for disks")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DiskSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(diskBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), diskPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"disks\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, diskPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from disk slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for disks")
	}

	if len(diskAfterDeleteHooks) != 0 {
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
func (o *Disk) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDisk(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiskSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DiskSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), diskPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"disks\".* FROM \"disks\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, diskPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DiskSlice")
	}

	*o = slice

	return nil
}

// DiskExists checks if the Disk row exists.
func DiskExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"disks\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if disks exists")
	}

	return exists, nil
}