package main

import (
	"github.com/casbin/casbin"
	"testing"
)

func TestEnforce(t *testing.T) {
	e := casbin.NewEnforcer("cb/basic_model.conf", "cb/basic_policy.csv")

	e.EnableEnforce(false)
	testEnforce(t, e, "alice", "data1", "read", true)
	testEnforce(t, e, "alice", "data1", "write", true)
	testEnforce(t, e, "alice", "data2", "read", true)
	testEnforce(t, e, "alice", "data2", "write", true)
	testEnforce(t, e, "bob", "data1", "read", true)
	testEnforce(t, e, "bob", "data1", "write", true)
	testEnforce(t, e, "bob", "data2", "read", true)
	testEnforce(t, e, "bob", "data2", "write", true)

	e.EnableEnforce(true)
	testEnforce(t, e, "alice", "data1", "GET", true)
	testEnforce(t, e, "alice", "data1", "write", false)
	testEnforce(t, e, "alice", "data2", "read", false)
	testEnforce(t, e, "alice", "data2", "write", false)
	testEnforce(t, e, "bob", "data1", "read", false)
	testEnforce(t, e, "bob", "data1", "write", false)
	testEnforce(t, e, "bob", "data2", "read", false)
	testEnforce(t, e, "bob", "data2", "POST", true)
}

func TestSavePolicy(t *testing.T) {
	e := casbin.NewEnforcer("cb/basic_model.conf", "cb/basic_policy.csv")
	e.SavePolicy()
}

func testEnforce(t *testing.T, e *casbin.Enforcer, sub string, obj interface{}, act string, res bool) {
	t.Helper()
	if e.Enforce(sub, obj, act) != res {
		t.Errorf("%s, %v, %s: %t, supposed to be %t", sub, obj, act, !res, res)
	}
}