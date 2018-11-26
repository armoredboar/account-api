BUILDDIR := .
SRCDIR := ./build/package/Dockerfile

all:
	docker build -t account-api -f $(SRCDIR) $(BUILDDIR) --no-cache

run:
	docker build -t account-api -f $(SRCDIR) $(BUILDDIR) --no-cache
	docker run --publish 3500:3500 --name test --rm account-api

publish:
	docker build -t account-api -f $(SRCDIR) $(BUILDDIR) --no-cache
	docker tag account-api registry.armoredboar.com/account-api
	docker login https://registry.armoredboar.com
	docker push registry.armoredboar.com/account-api
	docker image prune -f

clean:
	docker image prune -f