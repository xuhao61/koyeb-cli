package koyeb

import (
	"fmt"
	"strings"

	"github.com/koyeb/koyeb-api-client-go/api/v1/koyeb"
	"github.com/koyeb/koyeb-cli/pkg/koyeb/errors"
	"github.com/koyeb/koyeb-cli/pkg/koyeb/idmapper"
	"github.com/koyeb/koyeb-cli/pkg/koyeb/renderer"
	"github.com/spf13/cobra"
)

func (h *DeploymentHandler) Get(ctx *CLIContext, cmd *cobra.Command, args []string) error {
	deployment, err := h.ResolveDeploymentArgs(ctx, args[0])
	if err != nil {
		return err
	}

	res, resp, err := ctx.Client.DeploymentsApi.GetDeployment(ctx.Context, deployment).Execute()
	if err != nil {
		return errors.NewCLIErrorFromAPIError(
			fmt.Sprintf("Error while retrieving the deployment `%s`", args[0]),
			err,
			resp,
		)
	}

	full := GetBoolFlags(cmd, "full")
	getDeploymentsReply := NewGetDeploymentReply(ctx.Mapper, res, full)
	ctx.Renderer.Render(getDeploymentsReply)
	return nil
}

type GetDeploymentReply struct {
	mapper *idmapper.Mapper
	value  *koyeb.GetDeploymentReply
	full   bool
}

func NewGetDeploymentReply(mapper *idmapper.Mapper, value *koyeb.GetDeploymentReply, full bool) *GetDeploymentReply {
	return &GetDeploymentReply{
		mapper: mapper,
		value:  value,
		full:   full,
	}
}

func (GetDeploymentReply) Title() string {
	return "Deployment"
}

func (r *GetDeploymentReply) MarshalBinary() ([]byte, error) {
	return r.value.GetDeployment().MarshalJSON()
}

func (r *GetDeploymentReply) Headers() []string {
	return []string{"id", "service", "type", "status", "messages", "regions", "created_at"}
}

func (r *GetDeploymentReply) Fields() []map[string]string {
	item := r.value.GetDeployment()
	fields := map[string]string{
		"id":         renderer.FormatID(item.GetId(), r.full),
		"service":    renderer.FormatServiceSlug(r.mapper, item.GetServiceId(), r.full),
		"type":       formatDeploymentType(item.Definition.GetType()),
		"status":     formatDeploymentStatus(item.GetStatus()),
		"messages":   formatDeploymentMessages(item.GetMessages(), 0),
		"regions":    renderRegions(item.Definition.Regions),
		"created_at": renderer.FormatTime(item.GetCreatedAt()),
	}

	resp := []map[string]string{fields}
	return resp
}

func formatDeploymentType(dt koyeb.DeploymentDefinitionType) string {
	return string(dt)
}

func formatDeploymentStatus(ds koyeb.DeploymentStatus) string {
	return string(ds)
}

func formatDeploymentMessages(messages []string, max int) string {
	concat := strings.Join(messages, " ")
	if max == 0 || len(concat) < max {
		return concat
	}
	return fmt.Sprint(concat[:max], "...")
}

func renderRegions(regions []string) string {
	if len(regions) == 0 {
		return "-"
	}

	return strings.Join(regions, ",")
}
