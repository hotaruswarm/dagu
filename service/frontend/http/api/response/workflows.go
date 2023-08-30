package response

import (
	"github.com/samber/lo"
	"github.com/yohamta/dagu/internal/controller"
	"github.com/yohamta/dagu/internal/dag"
	domain "github.com/yohamta/dagu/internal/models"
	"github.com/yohamta/dagu/service/frontend/models"
)

func ToListWorkflowResponse(
	workflows []*controller.DAGStatus,
	errs []string,
	hasError bool,
) *models.ListWorkflowsResponse {
	return &models.ListWorkflowsResponse{
		DAGs: lo.Map(workflows, func(item *controller.DAGStatus, _ int) *models.WorkflowListItem {
			return ToWorkflowListItem(item)
		}),
		Errors:   errs,
		HasError: lo.ToPtr(hasError),
	}
}

func ToWorkflowListItem(s *controller.DAGStatus) *models.WorkflowListItem {
	return &models.WorkflowListItem{
		Dir:       lo.ToPtr(s.Dir),
		Error:     lo.ToPtr(toErrorText(s.Error)),
		ErrorT:    s.ErrorT,
		File:      lo.ToPtr(s.File),
		Status:    ToWorkflowStatus(s.Status),
		Suspended: lo.ToPtr(s.Suspended),
		DAG:       ToDAG(s.DAG),
	}
}

func ToDAG(d *dag.DAG) *models.Workflow {
	return &models.Workflow{
		Name:          lo.ToPtr(d.Name),
		Group:         lo.ToPtr(d.Group),
		Description:   lo.ToPtr(d.Description),
		Params:        d.Params,
		DefaultParams: lo.ToPtr(d.DefaultParams),
		Tags:          d.Tags,
		Schedule: lo.Map(d.Schedule, func(item *dag.Schedule, _ int) *models.Schedule {
			return ToSchedule(item)
		}),
	}
}

func ToSchedule(s *dag.Schedule) *models.Schedule {
	return &models.Schedule{
		Expression: lo.ToPtr(s.Expression),
	}
}

func ToWorkflowStatus(s *domain.Status) *models.WorkflowStatus {
	return &models.WorkflowStatus{
		Log:        lo.ToPtr(s.Log),
		Name:       lo.ToPtr(s.Name),
		Params:     lo.ToPtr(s.Params),
		Pid:        lo.ToPtr(int64(s.Pid)),
		RequestID:  lo.ToPtr(s.RequestId),
		StartedAt:  lo.ToPtr(s.StartedAt),
		FinishedAt: lo.ToPtr(s.FinishedAt),
		Status:     lo.ToPtr(int64(s.Status)),
		StatusText: lo.ToPtr(s.StatusText),
	}
}
