// Copyright 2021 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package creation

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/dolthub/go-mysql-server/sql"

	"github.com/dolthub/dolt/go/libraries/doltcore/doltdb"
	"github.com/dolthub/dolt/go/libraries/doltcore/doltdb/durable"
	"github.com/dolthub/dolt/go/libraries/doltcore/schema"
	"github.com/dolthub/dolt/go/libraries/doltcore/sqle/index"
	"github.com/dolthub/dolt/go/libraries/doltcore/table/editor"
	"github.com/dolthub/dolt/go/store/prolly"
	"github.com/dolthub/dolt/go/store/prolly/tree"
	"github.com/dolthub/dolt/go/store/types"
	"github.com/dolthub/dolt/go/store/val"
)

type CreateIndexReturn struct {
	NewTable *doltdb.Table
	Sch      schema.Schema
	OldIndex schema.Index
	NewIndex schema.Index
}

// CreateIndex creates the given index on the given table with the given schema. Returns the updated table, updated schema, and created index.
func CreateIndex(
	ctx *sql.Context,
	table *doltdb.Table,
	tableName, indexName string,
	columns []string,
	prefixLengths []uint16,
	props schema.IndexProperties,
	opts editor.Options,
) (*CreateIndexReturn, error) {
	sch, err := table.GetSchema(ctx)
	if err != nil {
		return nil, err
	}

	// get the real column names as CREATE INDEX columns are case-insensitive
	var realColNames []string
	allTableCols := sch.GetAllCols()
	for _, indexCol := range columns {
		tableCol, ok := allTableCols.GetByNameCaseInsensitive(indexCol)
		if !ok {
			return nil, fmt.Errorf("column `%s` does not exist for the table", indexCol)
		}
		realColNames = append(realColNames, tableCol.Name)
	}

	if indexName == "" {
		indexName = strings.Join(realColNames, "")
		_, ok := sch.Indexes().GetByNameCaseInsensitive(indexName)
		var i int
		for ok {
			i++
			indexName = fmt.Sprintf("%s_%d", strings.Join(realColNames, ""), i)
			_, ok = sch.Indexes().GetByNameCaseInsensitive(indexName)
		}
	}
	if !doltdb.IsValidIdentifier(indexName) {
		return nil, fmt.Errorf("invalid index name `%s`", indexName)
	}

	// if an index was already created for the column set but was not generated by the user then we replace it
	existingIndex, ok := sch.Indexes().GetIndexByColumnNames(realColNames...)
	if ok && !existingIndex.IsUserDefined() {
		_, err = sch.Indexes().RemoveIndex(existingIndex.Name())
		if err != nil {
			return nil, err
		}
		table, err = table.DeleteIndexRowData(ctx, existingIndex.Name())
		if err != nil {
			return nil, err
		}
	}

	// create the index metadata, will error if index names are taken or an index with the same columns in the same order exists
	index, err := sch.Indexes().AddIndexByColNames(
		indexName,
		realColNames,
		prefixLengths,
		props,
	)
	if err != nil {
		return nil, err
	}

	// update the table schema with the new index
	newTable, err := table.UpdateSchema(ctx, sch)
	if err != nil {
		return nil, err
	}

	// TODO: in the case that we're replacing an implicit index with one the user specified, we could do this more
	//  cheaply in some cases by just renaming it, rather than building it from scratch. But that's harder to get right.
	indexRows, err := BuildSecondaryIndex(ctx, newTable, index, tableName, opts)
	if err != nil {
		return nil, err
	}

	newTable, err = newTable.SetIndexRows(ctx, index.Name(), indexRows)
	if err != nil {
		return nil, err
	}

	return &CreateIndexReturn{
		NewTable: newTable,
		Sch:      sch,
		OldIndex: existingIndex,
		NewIndex: index,
	}, nil
}

func BuildSecondaryIndex(ctx *sql.Context, tbl *doltdb.Table, idx schema.Index, tableName string, opts editor.Options) (durable.Index, error) {
	switch tbl.Format() {
	case types.Format_LD_1:
		m, err := editor.RebuildIndex(ctx, tbl, idx.Name(), opts)
		if err != nil {
			return nil, err
		}
		return durable.IndexFromNomsMap(m, tbl.ValueReadWriter(), tbl.NodeStore()), nil

	case types.Format_DOLT:
		sch, err := tbl.GetSchema(ctx)
		if err != nil {
			return nil, err
		}
		m, err := tbl.GetRowData(ctx)
		if err != nil {
			return nil, err
		}
		primary := durable.ProllyMapFromIndex(m)
		return BuildSecondaryProllyIndex(ctx, tbl.ValueReadWriter(), tbl.NodeStore(), sch, tableName, idx, primary)

	default:
		return nil, fmt.Errorf("unknown NomsBinFormat")
	}
}

// BuildSecondaryProllyIndex builds secondary index data for the given primary
// index row data |primary|. |sch| is the current schema of the table.
func BuildSecondaryProllyIndex(
	ctx *sql.Context,
	vrw types.ValueReadWriter,
	ns tree.NodeStore,
	sch schema.Schema,
	tableName string,
	idx schema.Index,
	primary prolly.Map,
) (durable.Index, error) {
	var uniqCb DupEntryCb
	if idx.IsUnique() {
		kd := idx.Schema().GetKeyDescriptor()
		uniqCb = func(ctx context.Context, existingKey, newKey val.Tuple) error {
			msg := FormatKeyForUniqKeyErr(newKey, kd)
			return sql.NewUniqueKeyErr(msg, false, nil)
		}
	}
	return BuildProllyIndexExternal(ctx, vrw, ns, sch, tableName, idx, primary, uniqCb)
}

// FormatKeyForUniqKeyErr formats the given tuple |key| using |d|. The resulting
// string is suitable for use in a sql.UniqueKeyError
// This is copied from the writer package to avoid pulling in that dependency and prevent cycles
func FormatKeyForUniqKeyErr(key val.Tuple, d val.TupleDesc) string {
	var sb strings.Builder
	sb.WriteString("[")
	seenOne := false
	for i := range d.Types {
		if seenOne {
			sb.WriteString(",")
		}
		seenOne = true
		sb.WriteString(d.FormatValue(i, key.GetField(i)))
	}
	sb.WriteString("]")
	return sb.String()
}

// DupEntryCb receives duplicate unique index entries.
type DupEntryCb func(ctx context.Context, existingKey, newKey val.Tuple) error

// BuildUniqueProllyIndex builds a unique index based on the given |primary| row
// data. If any duplicate entries are found, they are passed to |cb|. If |cb|
// returns a non-nil error then the process is stopped.
func BuildUniqueProllyIndex(
	ctx *sql.Context,
	vrw types.ValueReadWriter,
	ns tree.NodeStore,
	sch schema.Schema,
	tableName string,
	idx schema.Index,
	primary prolly.Map,
	cb DupEntryCb,
) (durable.Index, error) {
	empty, err := durable.NewEmptyIndex(ctx, vrw, ns, idx.Schema())
	if err != nil {
		return nil, err
	}
	secondary := durable.ProllyMapFromIndex(empty)
	if schema.IsKeyless(sch) {
		secondary = prolly.ConvertToSecondaryKeylessIndex(secondary)
	}

	iter, err := primary.IterAll(ctx)
	if err != nil {
		return nil, err
	}
	p := primary.Pool()

	prefixDesc := secondary.KeyDesc().PrefixDesc(idx.Count())
	secondaryBld, err := index.NewSecondaryKeyBuilder(ctx, tableName, sch, idx, secondary.KeyDesc(), p, secondary.NodeStore())
	if err != nil {
		return nil, err
	}

	mut := secondary.Mutate()
	for {
		var k, v val.Tuple
		k, v, err = iter.Next(ctx)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		idxKey, err := secondaryBld.SecondaryKeyFromRow(ctx, k, v)
		if err != nil {
			return nil, err
		}

		if prefixDesc.HasNulls(idxKey) {
			continue
		}

		err = mut.GetPrefix(ctx, idxKey, prefixDesc, func(existingKey, _ val.Tuple) error {
			// register a constraint violation if |idxKey| collides with |existingKey|
			if existingKey != nil {
				return cb(ctx, existingKey, idxKey)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}

		if err = mut.Put(ctx, idxKey, val.EmptyTuple); err != nil {
			return nil, err
		}
	}

	secondary, err = mut.Map(ctx)
	if err != nil {
		return nil, err
	}
	return durable.IndexFromProllyMap(secondary), nil
}

// PrefixItr iterates all keys of a given prefix |p| and its descriptor |d| in map |m|.
// todo(andy): move to pkg prolly
type PrefixItr struct {
	itr prolly.MapIter
	p   val.Tuple
	d   val.TupleDesc
}

func NewPrefixItr(ctx context.Context, p val.Tuple, d val.TupleDesc, m rangeIterator) (PrefixItr, error) {
	rng := prolly.PrefixRange(p, d)
	itr, err := m.IterRange(ctx, rng)
	if err != nil {
		return PrefixItr{}, err
	}
	return PrefixItr{p: p, d: d, itr: itr}, nil
}

func (itr PrefixItr) Next(ctx context.Context) (k, v val.Tuple, err error) {
OUTER:
	for {
		k, v, err = itr.itr.Next(ctx)
		if err != nil {
			return nil, nil, err
		}

		// check if p is a prefix of k
		// range iteration currently can return keys not in the range
		for i := 0; i < itr.p.Count(); i++ {
			f1 := itr.p.GetField(i)
			f2 := k.GetField(i)
			if bytes.Compare(f1, f2) != 0 {
				// if a field in the prefix does not match |k|, go to the next row
				continue OUTER
			}
		}

		return k, v, nil
	}
}

type rangeIterator interface {
	IterRange(ctx context.Context, rng prolly.Range) (prolly.MapIter, error)
}

var _ error = (*prollyUniqueKeyErr)(nil)

// prollyUniqueKeyErr is an error that is returned when a unique constraint has been violated. It contains the index key
// (which is the full row).
type prollyUniqueKeyErr struct {
	k         val.Tuple
	kd        val.TupleDesc
	IndexName string
}

// Error implements the error interface.
func (u *prollyUniqueKeyErr) Error() string {
	keyStr, _ := formatKey(u.k, u.kd)
	return fmt.Sprintf("duplicate unique key given: %s", keyStr)
}

// formatKey returns a comma-separated string representation of the key given
// that matches the output of the old format.
func formatKey(key val.Tuple, td val.TupleDesc) (string, error) {
	vals := make([]string, td.Count())
	for i := 0; i < td.Count(); i++ {
		vals[i] = td.FormatValue(i, key.GetField(i))
	}

	return fmt.Sprintf("[%s]", strings.Join(vals, ",")), nil
}