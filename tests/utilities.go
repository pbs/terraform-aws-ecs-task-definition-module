package test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/sts"
)

func deleteLogGroup(t *testing.T, logGroupName string) {
	session, err := session.NewSession()
	if err != nil {
		t.Fatalf("Failed to create AWS session: %v", err)
	}
	svc := cloudwatchlogs.New(session)
	input := cloudwatchlogs.DeleteLogGroupInput{
		LogGroupName: &logGroupName,
	}
	_, err = svc.DeleteLogGroup(&input)
	if err != nil {
		t.Logf("Failed to delete log group: %v.\nThis is probably OK, as we're just making sure it's not there.", err)
	}
}

func getAWSAccountID(t *testing.T) string {
	session, err := session.NewSession()
	if err != nil {
		t.Fatalf("Failed to create AWS session: %v", err)
		return ""
	}
	svc := sts.New(session)
	result, err := svc.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		t.Fatalf("Failed to get AWS Account ID: %v", err)
		return ""
	}
	return *result.Account
}

func getAWSRegion(t *testing.T) string {
	session, err := session.NewSession()
	if err != nil {
		t.Fatalf("Failed to create AWS session: %v", err)
		return ""
	}
	return *session.Config.Region
}
