package main

import (
    "fmt"
    "os"
    "log"
    "context"
    "strings"
    "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
    msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
    "github.com/microsoftgraph/msgraph-sdk-go/models"
)

func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func checkEmpty(envVar string) {
    if len(envVar) == 0 { 
        log.Fatal("env var missing, ARM_TENANT_ID, ARM_CLIENT_ID, ARM_CLIENT_SECRET must be set for this binary")
    }
}
func main() {
    if len(os.Args) < 2 {
        log.Fatal("Must supply Azure Group ID as command-line argument")
    }
    groupId := os.Args[1]
    armTenantId := os.Getenv("ARM_TENANT_ID")
    checkEmpty(armTenantId)
    armClientId := os.Getenv("ARM_CLIENT_ID")
    checkEmpty(armClientId)
    armClientSecret := os.Getenv("ARM_CLIENT_SECRET")
    checkEmpty(armClientSecret)
    cred, err := azidentity.NewClientSecretCredential(armTenantId, armClientId, armClientSecret, nil)
    check(err)
    graphId := "00000003-0000-0000-c000-000000000000"
    client , err  := msgraphsdk.NewGraphServiceClientWithCredentials(cred, []string{fmt.Sprintf("%s/.default", graphId)})
    check(err)

    result, err := client.GroupsById(groupId).TransitiveMembers().Get(context.Background(), nil)
    check(err)
    var users []string

    val := result.GetValue()
    for _, v := range val {
        user, ok := v.(models.Userable)
        if ok {
            users = append(users, *user.GetUserPrincipalName())
        }
    }
    fmt.Printf("{\"value\":\"%s\"}",strings.Join(users, ","))
}
