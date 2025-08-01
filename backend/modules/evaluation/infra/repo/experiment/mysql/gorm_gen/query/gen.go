// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                         db,
		Experiment:                 newExperiment(db, opts...),
		ExptAggrResult:             newExptAggrResult(db, opts...),
		ExptEvaluatorRef:           newExptEvaluatorRef(db, opts...),
		ExptItemResult:             newExptItemResult(db, opts...),
		ExptItemResultRunLog:       newExptItemResultRunLog(db, opts...),
		ExptRunLog:                 newExptRunLog(db, opts...),
		ExptStats:                  newExptStats(db, opts...),
		ExptTurnEvaluatorResultRef: newExptTurnEvaluatorResultRef(db, opts...),
		ExptTurnResult:             newExptTurnResult(db, opts...),
		ExptTurnResultRunLog:       newExptTurnResultRunLog(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Experiment                 experiment
	ExptAggrResult             exptAggrResult
	ExptEvaluatorRef           exptEvaluatorRef
	ExptItemResult             exptItemResult
	ExptItemResultRunLog       exptItemResultRunLog
	ExptRunLog                 exptRunLog
	ExptStats                  exptStats
	ExptTurnEvaluatorResultRef exptTurnEvaluatorResultRef
	ExptTurnResult             exptTurnResult
	ExptTurnResultRunLog       exptTurnResultRunLog
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                         db,
		Experiment:                 q.Experiment.clone(db),
		ExptAggrResult:             q.ExptAggrResult.clone(db),
		ExptEvaluatorRef:           q.ExptEvaluatorRef.clone(db),
		ExptItemResult:             q.ExptItemResult.clone(db),
		ExptItemResultRunLog:       q.ExptItemResultRunLog.clone(db),
		ExptRunLog:                 q.ExptRunLog.clone(db),
		ExptStats:                  q.ExptStats.clone(db),
		ExptTurnEvaluatorResultRef: q.ExptTurnEvaluatorResultRef.clone(db),
		ExptTurnResult:             q.ExptTurnResult.clone(db),
		ExptTurnResultRunLog:       q.ExptTurnResultRunLog.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                         db,
		Experiment:                 q.Experiment.replaceDB(db),
		ExptAggrResult:             q.ExptAggrResult.replaceDB(db),
		ExptEvaluatorRef:           q.ExptEvaluatorRef.replaceDB(db),
		ExptItemResult:             q.ExptItemResult.replaceDB(db),
		ExptItemResultRunLog:       q.ExptItemResultRunLog.replaceDB(db),
		ExptRunLog:                 q.ExptRunLog.replaceDB(db),
		ExptStats:                  q.ExptStats.replaceDB(db),
		ExptTurnEvaluatorResultRef: q.ExptTurnEvaluatorResultRef.replaceDB(db),
		ExptTurnResult:             q.ExptTurnResult.replaceDB(db),
		ExptTurnResultRunLog:       q.ExptTurnResultRunLog.replaceDB(db),
	}
}

type queryCtx struct {
	Experiment                 *experimentDo
	ExptAggrResult             *exptAggrResultDo
	ExptEvaluatorRef           *exptEvaluatorRefDo
	ExptItemResult             *exptItemResultDo
	ExptItemResultRunLog       *exptItemResultRunLogDo
	ExptRunLog                 *exptRunLogDo
	ExptStats                  *exptStatsDo
	ExptTurnEvaluatorResultRef *exptTurnEvaluatorResultRefDo
	ExptTurnResult             *exptTurnResultDo
	ExptTurnResultRunLog       *exptTurnResultRunLogDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Experiment:                 q.Experiment.WithContext(ctx),
		ExptAggrResult:             q.ExptAggrResult.WithContext(ctx),
		ExptEvaluatorRef:           q.ExptEvaluatorRef.WithContext(ctx),
		ExptItemResult:             q.ExptItemResult.WithContext(ctx),
		ExptItemResultRunLog:       q.ExptItemResultRunLog.WithContext(ctx),
		ExptRunLog:                 q.ExptRunLog.WithContext(ctx),
		ExptStats:                  q.ExptStats.WithContext(ctx),
		ExptTurnEvaluatorResultRef: q.ExptTurnEvaluatorResultRef.WithContext(ctx),
		ExptTurnResult:             q.ExptTurnResult.WithContext(ctx),
		ExptTurnResultRunLog:       q.ExptTurnResultRunLog.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
