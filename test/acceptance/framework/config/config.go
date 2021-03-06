package config

// TestConfig holds configuration for the test suite.
type TestConfig struct {
	NoCleanupOnFailure bool
	ECSClusterARN      string      `json:"ecs_cluster_arn"`
	PrivateSubnets     interface{} `json:"private_subnets"`
	PublicSubnets      interface{} `json:"public_subnets"`
	Region             string      `json:"region"`
	LogGroupName       string      `json:"log_group_name"`
	VPCID              string      `json:"vpc_id"`
	ECRImageURI        string      `json:"ecr_image_uri"`
	Suffix             string      `json:"suffix"`
}

func (t TestConfig) TFVars(ignoreVars ...string) map[string]interface{} {
	vars := map[string]interface{}{
		"ecs_cluster_arn": t.ECSClusterARN,
		"private_subnets": t.PrivateSubnets,
		"public_subnets":  t.PublicSubnets,
		"region":          t.Region,
		"log_group_name":  t.LogGroupName,
		"vpc_id":          t.VPCID,
		"ecr_image_uri":   t.ECRImageURI,
		"suffix":          t.Suffix,
	}

	for _, v := range ignoreVars {
		delete(vars, v)
	}
	return vars
}
