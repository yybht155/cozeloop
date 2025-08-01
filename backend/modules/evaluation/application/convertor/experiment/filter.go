// Copyright (c) 2025 coze-dev Authors
// SPDX-License-Identifier: Apache-2.0

package experiment

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/bytedance/gg/gslice"

	domain_expt "github.com/coze-dev/cozeloop/backend/kitex_gen/coze/loop/evaluation/domain/expt"
	"github.com/coze-dev/cozeloop/backend/modules/evaluation/domain/entity"
	"github.com/coze-dev/cozeloop/backend/modules/evaluation/domain/service"
	"github.com/coze-dev/cozeloop/backend/pkg/errorx"
	"github.com/coze-dev/cozeloop/backend/pkg/json"
	"github.com/coze-dev/cozeloop/backend/pkg/logs"
)

func NewExptFilterConvertor(evalTargetService service.IEvalTargetService) *ExptFilterConvertor {
	return &ExptFilterConvertor{
		evalTargetService: evalTargetService,
	}
}

type ExptFilterConvertor struct {
	evalTargetService service.IEvalTargetService
}

func (e *ExptFilterConvertor) Convert(ctx context.Context, efo *domain_expt.ExptFilterOption, spaceID int64) (*entity.ExptListFilter, error) {
	if efo == nil {
		return nil, nil
	}

	filters, err := e.ConvertFilters(ctx, efo.GetFilters(), spaceID)
	if err != nil {
		return nil, err
	}

	filters.FuzzyName = efo.GetFuzzyName()

	return filters, nil
}

func (e *ExptFilterConvertor) ConvertFilters(ctx context.Context, filters *domain_expt.Filters, spaceID int64) (*entity.ExptListFilter, error) {
	efo := &entity.ExptListFilter{
		Includes: &entity.ExptFilterFields{},
		Excludes: &entity.ExptFilterFields{},
	}

	if filters == nil {
		return efo, nil
	}

	if filters.GetLogicOp() != domain_expt.FilterLogicOp_And {
		return nil, fmt.Errorf("ConvertFilters fail, opertaor type must be and, got: %v", filters.GetLogicOp())
	}

	ffieldsFn := func(operatorType domain_expt.FilterOperatorType) *entity.ExptFilterFields {
		switch operatorType {
		case domain_expt.FilterOperatorType_In, domain_expt.FilterOperatorType_Equal:
			return efo.Includes
		case domain_expt.FilterOperatorType_NotIn, domain_expt.FilterOperatorType_NotEqual:
			return efo.Excludes
		default:
			return &entity.ExptFilterFields{}
		}
	}

	setDefaultExptTypeFlag := true
	for _, cond := range filters.GetFilterConditions() {
		if cond.GetField() == nil {
			continue
		}
		ff := ffieldsFn(cond.GetOperator())
		switch cond.GetField().GetFieldType() {
		case domain_expt.FieldType_CreatorBy:
			if len(cond.GetValue()) == 0 {
				continue
			}
			ff.CreatedBy = intersectIgnoreNull(ff.CreatedBy, []string{cond.GetValue()})
		case domain_expt.FieldType_ExptStatus:
			if len(cond.GetValue()) == 0 {
				continue
			}
			status, err := parseIntList(cond.GetValue())
			if err != nil {
				return nil, errorx.Wrapf(err, "string to int64 assert fail, str: %v", cond.GetValue())
			}
			if gslice.Contains(status, int64(domain_expt.ExptStatus_Processing)) {
				status = append(status, int64(domain_expt.ExptStatus_Draining))
			}
			ff.Status = intersectIgnoreNull(ff.Status, status)
		case domain_expt.FieldType_EvalSetID:
			if len(cond.GetValue()) == 0 {
				continue
			}
			ids, err := parseIntList(cond.GetValue())
			if err != nil {
				return nil, err
			}
			ff.EvalSetIDs = intersectIgnoreNull(ff.EvalSetIDs, ids)
		case domain_expt.FieldType_TargetID:
			if len(cond.GetValue()) == 0 {
				continue
			}
			ids, err := parseIntList(cond.GetValue())
			if err != nil {
				return nil, err
			}
			ff.TargetIDs = intersectIgnoreNull(ff.TargetIDs, ids)
		case domain_expt.FieldType_EvaluatorID:
			if len(cond.GetValue()) == 0 {
				continue
			}
			ids, err := parseIntList(cond.GetValue())
			if err != nil {
				return nil, err
			}
			ff.EvaluatorIDs = intersectIgnoreNull(ff.EvaluatorIDs, ids)
		case domain_expt.FieldType_TargetType:
			if len(cond.GetValue()) == 0 {
				continue
			}
			ty, err := strconv.ParseInt(cond.GetValue(), 10, 64)
			if err != nil {
				return nil, errorx.Wrapf(err, "string to int64 assert fail, str: %v", cond.GetValue())
			}
			ff.TargetType = intersectIgnoreNull(ff.TargetType, []int64{ty})
		case domain_expt.FieldType_SourceTarget:
			if cond.GetSourceTarget() == nil || len(cond.GetSourceTarget().GetSourceTargetIds()) == 0 {
				continue
			}
			param := &entity.BatchGetEvalTargetBySourceParam{
				SpaceID:        spaceID,
				SourceTargetID: cond.GetSourceTarget().GetSourceTargetIds(),
				TargetType:     entity.EvalTargetType(cond.GetSourceTarget().GetEvalTargetType()),
			}
			targets, err := e.evalTargetService.BatchGetEvalTargetBySource(ctx, param)
			// targets, err := e.evalCall.BatchGetEvalTargetBySource(ctx, cond.GetSourceTarget().GetSourceTargetIds(), 0, spaceID)
			if err != nil {
				return nil, err
			}
			if len(cond.GetSourceTarget().GetSourceTargetIds()) == 1 && len(targets) == 0 {
				ff.TargetIDs = append(ff.TargetIDs, -1) // 无效查询，返回空结果
				break
			}
			targetIDs := make([]int64, 0, len(targets))
			for _, target := range targets {
				targetIDs = append(targetIDs, target.ID)
			}
			ff.TargetIDs = intersectIgnoreNull(ff.TargetIDs, targetIDs)
		case domain_expt.FieldType_ExptType:
			setDefaultExptTypeFlag = false
			types, err := parseIntList(cond.GetValue())
			if err != nil {
				return nil, err
			}
			ff.ExptType = intersectIgnoreNull(ff.ExptType, types)
		case domain_expt.FieldType_SourceType:
			if len(cond.GetValue()) == 0 {
				continue
			}
			types, err := parseIntList(cond.GetValue())
			if err != nil {
				return nil, err
			}
			ff.SourceType = intersectIgnoreNull(ff.SourceType, types)
		case domain_expt.FieldType_SourceID:
			if len(cond.GetValue()) == 0 {
				continue
			}
			sourceIDs := parseStringList(cond.GetValue())
			ff.SourceID = intersectIgnoreNull(ff.SourceID, sourceIDs)
		default:
			logs.CtxWarn(ctx, "ConvertFilters with unsupport condition: %v", json.Jsonify(cond))
		}
	}
	if setDefaultExptTypeFlag {
		efo.Includes.ExptType = intersectIgnoreNull(efo.Includes.ExptType, []int64{int64(domain_expt.ExptType_Offline)})
	}

	return efo, nil
}

func intersectIgnoreNull[T comparable](s1, s2 []T) []T {
	if len(s1) == 0 {
		return s2
	}
	if len(s2) == 0 {
		return s1
	}
	var res []T
	memo := gslice.ToMap(s1, func(t T) (T, bool) { return t, true })
	for _, item := range s2 {
		if memo[item] {
			res = append(res, item)
		}
	}
	return res
}

func parseIntList(str string) ([]int64, error) {
	split := strings.Split(str, ",")
	res := make([]int64, 0, len(split))
	for _, s := range split {
		val, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, errorx.Wrapf(err, "string to int64 assert fail, str: %s", str)
		}
		res = append(res, val)
	}
	return res, nil
}

func parseStringList(str string) []string {
	return strings.Split(str, ",")
}

func parseOperator(operatorType domain_expt.FilterOperatorType) (string, error) {
	var operator string
	switch operatorType {
	case domain_expt.FilterOperatorType_Equal:
		operator = "="
	case domain_expt.FilterOperatorType_NotEqual:
		operator = "!="
	case domain_expt.FilterOperatorType_Greater:
		operator = ">"
	case domain_expt.FilterOperatorType_GreaterOrEqual:
		operator = ">="
	case domain_expt.FilterOperatorType_Less:
		operator = "<"
	case domain_expt.FilterOperatorType_LessOrEqual:
		operator = "<="
	case domain_expt.FilterOperatorType_In:
		operator = "IN"
	case domain_expt.FilterOperatorType_NotIn:
		operator = "NOT IN"
	default:
		return "", fmt.Errorf("invalid operator")
	}

	return operator, nil
}

func ConvertExptTurnResultFilter(filters *domain_expt.Filters) (*entity.ExptTurnResultFilter, error) {
	trunRunStateFilters := make([]*entity.TurnRunStateFilter, 0)
	scoreFilters := make([]*entity.ScoreFilter, 0)
	if filters != nil && len(filters.FilterConditions) > 0 {
		if filters.GetLogicOp() != domain_expt.FilterLogicOp_And {
			return nil, fmt.Errorf("invalid logic op")
		}

		for _, filterCondition := range filters.GetFilterConditions() {
			if filterCondition == nil {
				continue
			}
			err := checkFilterCondition(*filterCondition)
			if err != nil {
				return nil, err
			}

			operator, err := parseOperator(filterCondition.GetOperator())
			if err != nil {
				return nil, err
			}

			switch filterCondition.GetField().GetFieldType() {
			case domain_expt.FieldType_TurnRunState:
				turnRunStates, err := parseTurnRunState(*filterCondition)
				if err != nil {
					return nil, err
				}
				turnRunStateFilter := &entity.TurnRunStateFilter{
					Status:   turnRunStates,
					Operator: operator,
				}
				trunRunStateFilters = append(trunRunStateFilters, turnRunStateFilter)
			case domain_expt.FieldType_EvaluatorScore:
				score, err := strconv.ParseFloat(filterCondition.GetValue(), 64)
				if err != nil {
					return nil, err
				}
				evaluatorVersionID, err := strconv.ParseInt(filterCondition.GetField().GetFieldKey(), 10, 64)
				if err != nil {
					return nil, err
				}
				scoreFilter := &entity.ScoreFilter{
					Score:              score,
					Operator:           operator,
					EvaluatorVersionID: evaluatorVersionID,
				}
				scoreFilters = append(scoreFilters, scoreFilter)
			default:
				return nil, fmt.Errorf("invalid field type")
			}
		}
	}

	return &entity.ExptTurnResultFilter{
		TrunRunStateFilters: trunRunStateFilters,
		ScoreFilters:        scoreFilters,
	}, nil
}

func parseTurnRunState(filterCondition domain_expt.FilterCondition) ([]entity.TurnRunState, error) {
	// 使用“,”分割
	strStates := strings.Split(filterCondition.GetValue(), ",")

	// 解析为TurnRunState
	states := make([]entity.TurnRunState, 0, len(strStates))
	for _, strState := range strStates {
		if strState == "" { //	兜底：前端取消筛选后TurnRunState可能会传空字符串
			continue
		}
		turnRunState, err := strconv.ParseInt(strState, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid turn run state")
		}

		states = append(states, entity.TurnRunState(turnRunState))
	}

	return states, nil
}

func checkFilterCondition(filterCondition domain_expt.FilterCondition) error {
	switch filterCondition.GetField().GetFieldType() {
	case domain_expt.FieldType_TurnRunState:
		if filterCondition.GetOperator() != domain_expt.FilterOperatorType_In &&
			filterCondition.GetOperator() != domain_expt.FilterOperatorType_NotIn {
			return fmt.Errorf("invalid operator")
		}
	}
	return nil
}
