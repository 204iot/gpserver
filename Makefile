.PHONY: all-in-one
all-in-one:
	docker build -f docker/all-in-one/Dockerfile . -t gpserver
