package main
//go:generate sqlc generate -f .\internal\pgstore\sqlc.yaml
//go:generate tern migrate -m .\internal\pgstore\migrations\ -c .\internal\pgstore\migrations\tern.conf
//go:generate goapi-gen --package=spec --out .\internal\api\spec\journey.gen.spec.go .\internal\api\spec\journey.spec.json

