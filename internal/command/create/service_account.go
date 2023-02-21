package create

import (
	"fmt"
	"github.com/cirruslabs/orchard/pkg/client"
	v1 "github.com/cirruslabs/orchard/pkg/resource/v1"
	"github.com/spf13/cobra"
	"strings"
)

var token string
var roles []string

func newCreateServiceAccount() *cobra.Command {
	command := &cobra.Command{
		Use:  "service-account",
		RunE: runCreateServiceAccount,
		Args: cobra.ExactArgs(1),
	}

	command.PersistentFlags().StringVar(&token, "token", "",
		"token to use for this service account (autogenerated by the API server if left empty)")

	var serviceAccountRoleList []string
	for _, role := range v1.AllServiceAccountRoles() {
		serviceAccountRoleList = append(serviceAccountRoleList, string(role))
	}
	command.PersistentFlags().StringArrayVar(&roles, "roles", []string{},
		fmt.Sprintf("roles to grant to this service account (supported roles: %s)",
			strings.Join(serviceAccountRoleList, ", ")))

	return command
}

func runCreateServiceAccount(cmd *cobra.Command, args []string) error {
	name := args[0]

	client, err := client.New()
	if err != nil {
		return err
	}

	var serviceAccountRoles []v1.ServiceAccountRole

	for _, role := range roles {
		// Don't bother checking if the role name is valid
		// since this will be checked by the API server anyway
		serviceAccountRoles = append(serviceAccountRoles, v1.ServiceAccountRole(role))
	}

	return client.ServiceAccounts().Create(cmd.Context(), &v1.ServiceAccount{
		Meta: v1.Meta{
			Name: name,
		},
		Token: token,
		Roles: serviceAccountRoles,
	})
}