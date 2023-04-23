package busi_client

import "testing"

func TestAuth(t *testing.T) {
	bo := &AuthBo{UserId: 1, Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5payIsImlkIjoxLCJwYXNzd29yZCI6ImQ0MWQ4Y2Q5OGYwMGIyMDRlOTgwMDk5OGVjZjg0MjdlIiwiZXhwIjoxNjgyMTYxMzE5LCJpc3MiOiJodHRwczovL2djaGF0LmNvbSIsInR5cGUiOjF9.zd16v9S6m6NOkA8vsp2SasRIkD_M3pBrV3KyW-I4k-U"}
	Auth(bo)
}
