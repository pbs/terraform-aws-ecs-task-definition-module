package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func testTaskDef(t *testing.T, variant string) {
	t.Parallel()

	terraformDir := fmt.Sprintf("../examples/%s", variant)

	terraformOptions := &terraform.Options{
		TerraformDir: terraformDir,
		LockTimeout:  "5m",
	}

	defer terraform.Destroy(t, terraformOptions)

	expectedName := fmt.Sprintf("tf-ecs-task-def-%s", variant)

	// This is annoying, but necessary. Log Group isn't cleaned up correctly after destroy.
	logGroupName := fmt.Sprintf("/ecs/%s", expectedName)
	deleteLogGroup(t, logGroupName)

	terraform.InitAndApply(t, terraformOptions)

	ARN := terraform.Output(t, terraformOptions, "arn")
	roleARN := terraform.Output(t, terraformOptions, "role_arn")
	containerDefinitions := terraform.Output(t, terraformOptions, "container_definitions")

	region := getAWSRegion(t)
	accountID := getAWSAccountID(t)

	partialExpectedARN := fmt.Sprintf("arn:aws:ecs:%s:%s:task-definition/%s:", region, accountID, expectedName)
	partialExpectedRoleARN := fmt.Sprintf("arn:aws:iam::%s:role/%s", accountID, expectedName)

	var expectedContainerDefinitions string
	switch variant {
	case "basic":
		expectedContainerDefinitions = fmt.Sprintf("[{\"environment\":[{\"name\":\"SSM_PATH\",\"value\":\"/sharedtools/%s/\"}],\"essential\":true,\"image\":\"nginx:alpine\",\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-group\":\"%s\",\"awslogs-region\":\"%s\",\"awslogs-stream-prefix\":\"app\"}},\"name\":\"app\",\"portMappings\":[{\"containerPort\":80}]}]", expectedName, logGroupName, region)
	case "cmd":
		expectedContainerDefinitions = fmt.Sprintf("[{\"command\":[\"hello\"],\"entrypoint\":[\"echo\"],\"environment\":[{\"name\":\"SSM_PATH\",\"value\":\"/sharedtools/%s/\"}],\"essential\":true,\"image\":\"nginx:alpine\",\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-group\":\"%s\",\"awslogs-region\":\"%s\",\"awslogs-stream-prefix\":\"app\"}},\"name\":\"app\",\"portMappings\":[{\"containerPort\":80}]}]", expectedName, logGroupName, region)
	case "efs":
		efs1ID := terraform.Output(t, terraformOptions, "efs1_id")
		efs2ID := terraform.Output(t, terraformOptions, "efs2_id")
		expectedContainerDefinitions = fmt.Sprintf("[{\"environment\":[{\"name\":\"SSM_PATH\",\"value\":\"/sharedtools/%s/\"}],\"essential\":true,\"image\":\"nginx:alpine\",\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-group\":\"%s\",\"awslogs-region\":\"%s\",\"awslogs-stream-prefix\":\"app\"}},\"mountPoints\":[{\"containerPath\":\"/mnt/efs1\",\"sourceVolume\":\"%s\"},{\"containerPath\":\"/mnt/efs2\",\"sourceVolume\":\"%s\"}],\"name\":\"app\",\"portMappings\":[{\"containerPort\":80}]}]", expectedName, logGroupName, region, efs1ID, efs2ID)
	case "newrelic":
		expectedContainerDefinitions = fmt.Sprintf("[{\"environment\":[{\"name\":\"SSM_PATH\",\"value\":\"/sharedtools/%s/\"}],\"essential\":true,\"image\":\"nginx:alpine\",\"logConfiguration\":{\"logDriver\":\"awsfirelens\",\"options\":{\"Name\":\"newrelic\"},\"secretOptions\":[{\"name\":\"apiKey\",\"valueFrom\":\"arn:aws:secretsmanager:*:*:secret:fake-newrelic-secret-arn\"}]},\"name\":\"app\",\"portMappings\":[{\"containerPort\":80}]},{\"essential\":true,\"firelensConfiguration\":{\"options\":{\"enable-ecs-log-metadata\":\"true\"},\"type\":\"fluentbit\"},\"image\":\"533243300146.dkr.ecr.%s.amazonaws.com/newrelic/logging-firelens-fluentbit\",\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-group\":\"%s\",\"awslogs-region\":\"%s\",\"awslogs-stream-prefix\":\"firelens\"}},\"name\":\"firelens\"}]", expectedName, region, logGroupName, region)
	case "virtual-gateway":
		expectedContainerDefinitions = fmt.Sprintf("[{\"environment\":[{\"name\":\"APPMESH_RESOURCE_ARN\",\"value\":\"arn:aws:appmesh:%s:%s:mesh/mesh_name/virtualGateway/virtual_gateway\"},{\"name\":\"ENABLE_ENVOY_XRAY_TRACING\",\"value\":\"1\"}],\"essential\":true,\"healthCheck\":{\"command\":[\"CMD-SHELL\",\"curl -s http://localhost:9901/server_info | grep state | grep -q LIVE\"],\"interval\":5,\"retries\":3,\"startPeriod\":10,\"timeout\":2},\"image\":\"public.ecr.aws/appmesh/aws-appmesh-envoy:v1.23.1.0-prod\",\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-group\":\"%s\",\"awslogs-region\":\"%s\",\"awslogs-stream-prefix\":\"envoy\"}},\"name\":\"envoy\",\"portMappings\":[{\"containerPort\":80,\"protocol\":\"tcp\"}],\"user\":\"1337\"},{\"cpu\":32,\"essential\":true,\"image\":\"amazon/aws-xray-daemon\",\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-group\":\"%s\",\"awslogs-region\":\"%s\",\"awslogs-stream-prefix\":\"xray\"}},\"memoryReservation\":256,\"name\":\"xray-daemon\",\"portMappings\":[{\"containerPort\":2000,\"protocol\":\"udp\"}],\"user\":\"1337\"}]", region, accountID, logGroupName, region, logGroupName, region)
	case "virtual-node":
		expectedContainerDefinitions = fmt.Sprintf("[{\"dependsOn\":[{\"condition\":\"HEALTHY\",\"containerName\":\"envoy\"}],\"environment\":[{\"name\":\"SSM_PATH\",\"value\":\"/sharedtools/%s/\"}],\"essential\":true,\"image\":\"nginx:alpine\",\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-group\":\"%s\",\"awslogs-region\":\"%s\",\"awslogs-stream-prefix\":\"app\"}},\"name\":\"app\",\"portMappings\":[{\"containerPort\":80}]},{\"environment\":[{\"name\":\"APPMESH_RESOURCE_ARN\",\"value\":\"arn:aws:appmesh:%s:%s:mesh/mesh_name/virtualNode/virtual_node\"},{\"name\":\"ENABLE_ENVOY_XRAY_TRACING\",\"value\":\"1\"}],\"essential\":true,\"healthCheck\":{\"command\":[\"CMD-SHELL\",\"curl -s http://localhost:9901/server_info | grep state | grep -q LIVE\"],\"interval\":5,\"retries\":3,\"startPeriod\":10,\"timeout\":2},\"image\":\"public.ecr.aws/appmesh/aws-appmesh-envoy:v1.23.1.0-prod\",\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-group\":\"%s\",\"awslogs-region\":\"%s\",\"awslogs-stream-prefix\":\"proxy\"}},\"memory\":500,\"name\":\"envoy\",\"user\":\"1337\"},{\"cpu\":32,\"essential\":true,\"image\":\"amazon/aws-xray-daemon\",\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-group\":\"%s\",\"awslogs-region\":\"%s\",\"awslogs-stream-prefix\":\"xray\"}},\"memoryReservation\":256,\"name\":\"xray-daemon\",\"portMappings\":[{\"containerPort\":2000,\"protocol\":\"udp\"}],\"user\":\"1337\"}]", expectedName, logGroupName, region, region, accountID, logGroupName, region, logGroupName, region)
	case "virtual-node-newrelic":
		expectedContainerDefinitions = fmt.Sprintf("[{\"environment\":[{\"name\":\"APPMESH_RESOURCE_ARN\",\"value\":\"arn:aws:appmesh:%s:%s:mesh/mesh_name/virtualNode/virtual_node\"},{\"name\":\"ENABLE_ENVOY_XRAY_TRACING\",\"value\":\"1\"}],\"essential\":true,\"healthCheck\":{\"command\":[\"CMD-SHELL\",\"curl -s http://localhost:9901/server_info | grep state | grep -q LIVE\"],\"interval\":5,\"retries\":3,\"startPeriod\":10,\"timeout\":2},\"image\":\"public.ecr.aws/appmesh/aws-appmesh-envoy:v1.23.1.0-prod\",\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-group\":\"%s\",\"awslogs-region\":\"%s\",\"awslogs-stream-prefix\":\"envoy\"}},\"name\":\"envoy\",\"portMappings\":[{\"containerPort\":80,\"protocol\":\"tcp\"}],\"user\":\"1337\"},{\"cpu\":32,\"essential\":true,\"image\":\"amazon/aws-xray-daemon\",\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-group\":\"%s\",\"awslogs-region\":\"%s\",\"awslogs-stream-prefix\":\"xray\"}},\"memoryReservation\":256,\"name\":\"xray-daemon\",\"portMappings\":[{\"containerPort\":2000,\"protocol\":\"udp\"}],\"user\":\"1337\"}]", region, accountID, logGroupName, region, logGroupName, region)
	default:
		assert.Fail(t, fmt.Sprintf("Unknown variant: %s", variant))
	}

	assert.Contains(t, ARN, partialExpectedARN)
	assert.Contains(t, roleARN, partialExpectedRoleARN)

	assert.Equal(t, expectedContainerDefinitions, containerDefinitions)
}
