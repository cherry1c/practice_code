package testDemo

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"testing"
	"time"
)

const rbac_models = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
`

func TestManyRoles(t *testing.T) {
	m, _ := model.NewModelFromString(rbac_models)
	e, _ := casbin.NewEnforcer(m, false)

	roles := []string{"admin", "manager", "developer", "tester"}

	// 2500 projects
	for nbPrj := 1; nbPrj < 2500; nbPrj++ {
		// 4 objects and 1 role per object (so 4 roles)
		for _, role := range roles {
			roleDB := fmt.Sprintf("%s_project:%d", role, nbPrj)
			objectDB := fmt.Sprintf("/projects/%d", nbPrj)
			e.AddPolicy(roleDB, objectDB, "GET")
		}
		jasmineRole := fmt.Sprintf("%s_project:%d", roles[1], nbPrj)
		e.AddGroupingPolicy("jasmine", jasmineRole)
	}

	e.AddGroupingPolicy("abu", "manager_project:1")
	e.AddGroupingPolicy("abu", "manager_project:2499")

	// With same number of policies
	//  User 'abu' has only two roles
	//  User 'jasmine' has many roles (1 role per policy, here 2500 roles)

	request := func(subject, object, action string) {
		t0 := time.Now()
		resp, _ := e.Enforce(subject, object, action)
		tElapse := time.Since(t0)
		t.Logf("RESPONSE %-10s %s\t %s : %5v IN: %+v", subject, object, action, resp, tElapse)
		if tElapse > time.Millisecond*100 {
			t.Errorf("More than 100 milliseconds for %s %s %s : %+v", subject, object, action, tElapse)
		}
	}

	request("abu", "/projects/1", "GET")     // really fast because only 2 roles in all policies and at the beginning of the casbin_rule table
	request("abu", "/projects/2499", "GET")  // fast because only 2 roles in all policies
	request("jasmine", "/projects/1", "GET") // really fast at the beginning of the casbin_rule table

	request("jasmine", "/projects/2499", "GET") // slow and fail the only 1st time   <<<< pb here
	request("jasmine", "/projects/2499", "GET") // fast maybe due to internal cache mechanism

	//same issue with non-existing roles
	request("jasmine", "/projects/999999", "GET") // slow fail the only 1st time   <<<< pb here
	request("jasmine", "/projects/999999", "GET") // fast maybe due to internal cache mechanism
}
