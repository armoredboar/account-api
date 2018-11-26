BUILDDIR := .
SRCDIR := ./build/package/Dockerfile

# Display usage information.
all:
	@echo ""
	@echo "Usage: make [OPTION]"
	@echo ""
	@echo "OPTIONS:"
	@echo ""
	@echo "  build          - Build the docker image."
	@echo "  run            - Build and run the docker image."
	@echo "  publish        - Build, create a tag and pushes the image to the registry."
	@echo "  clean          - Clean unused images."
	@echo ""

# Build the docker image.
build:
	docker build -t account-api -f $(SRCDIR) $(BUILDDIR) --no-cache

# Build and run the docker image.
run:
	docker build -t account-api -f $(SRCDIR) $(BUILDDIR) --no-cache
	docker run --publish 3500:3500 --name test --rm account-api

# Build, create a tag and pushes the image to the registry.
publish:
	docker build -t account-api -f $(SRCDIR) $(BUILDDIR) --no-cache
	docker tag account-api registry.armoredboar.com/account-api
	docker login https://registry.armoredboar.com
	docker push registry.armoredboar.com/account-api
	docker image prune -f

# Clean unused images.
clean:
	docker image prune -f