package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	assert := assert.New(t)

	assert.Nil(Validate([]string{"rds", "prod"}))
	assert.NotNil(Validate([]string{"oh", "prod"}))
}

func TestApply(t *testing.T) {
	assert := assert.New(t)

	entries := map[string]string{
		"prod/us-east-1/rds": "/home/bob/src/terraform-repo/prod/us-east-1/rds",
		"prod/us-east-2/rds": "/home/bob/src/terraform-repo/prod/us-east-2/rds",
		"prod/us-east-1/ec2": "/home/bob/src/terraform-repo/prod/us-east-1/ec2",
		"dev/us-east-1/rds":  "/home/bob/src/terraform-repo/dev/us-east-1/rds",
		"dev/eu-west-1/rds":  "/home/bob/src/terraform-repo/dev/eu-west-1/rds",
		"dev/us-east-1/ec2":  "/home/bob/src/terraform-repo/dev/us-east-1/ec2",
	}

	matches := []string{"rds", "east-1"}

	expected := []string{"dev/us-east-1/rds", "prod/us-east-1/rds"}

	assert.Equal(expected, Apply(entries, matches))
}
