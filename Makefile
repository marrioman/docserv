local-pg:
	docker run -d --name test-user-postgres -p 5444:5432 -e POSTGRES_PASSWORD=password -e POSTGRES_DB=userdb postgres:11-alpine
	docker run -d --name test-docs-postgres -p 5432:5432 -e POSTGRES_PASSWORD=password -e POSTGRES_DB=docsdb postgres:11-alpine

clean-local-pg:
	docker rm -f test-user-postgres
	docker rm -f test-docs-postgres