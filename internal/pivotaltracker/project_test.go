package pivotaltracker

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ProjectTests struct {
	suite.Suite
}

func TestProject(t *testing.T) {
	suite.Run(t, &ProjectTests{})
}

func (s ProjectTests) Test_GetSanitized_Project() {
	proj := Project{
		Name:        "__name",
		Description: "__description",
		AccessToken: "__access_token",
		ProjectID:   "__project_id",
		Public:      true,
	}
	sanitized := proj.GetSanitized()
	s.NotEqual("__access_token", sanitized)
	s.Equal(proj.Name, sanitized.Name)
	s.Equal(proj.Description, sanitized.Description)
	s.Equal(proj.ProjectID, sanitized.ProjectID)
	s.Equal(proj.Public, sanitized.Public)
}

func (s ProjectTests) Test_GetSanitized_Projects() {
	projs := Projects{
		{
			Name:        "__name_0",
			Description: "__description_0",
			AccessToken: "__access_token_0",
			ProjectID:   "__project_id_0",
			Public:      true,
		},
		{
			Name:        "__name_1",
			Description: "__description_1",
			AccessToken: "__access_token_1",
			ProjectID:   "__project_id_1",
			Public:      true,
		},
	}
	sanitizeds := projs.GetSanitized()
	for index, sanitized := range sanitizeds {
		s.NotEqual(fmt.Sprintf("__access_token_%v", index), sanitized)
		s.Equal(projs[index].Name, sanitized.Name)
		s.Equal(projs[index].Description, sanitized.Description)
		s.Equal(projs[index].ProjectID, sanitized.ProjectID)
		s.Equal(projs[index].Public, sanitized.Public)
	}
}
