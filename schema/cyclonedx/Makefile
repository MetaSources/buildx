.DEFAULT_GOAL := validate-schema
.PHONY: validate-schema
validate-schema:
	go run ../../cmd/buildx/main.go ubuntu:latest -vv -o cyclonedx > bom.xml
	xmllint --noout --schema ./cyclonedx.xsd bom.xml
	go run ../../cmd/buildx/main.go ubuntu:latest -vv -o cyclonedx-json > bom.json
	../../.tmp/yajsv -s cyclonedx.json bom.json
