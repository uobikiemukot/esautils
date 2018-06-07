GOCMD   = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST  = $(GOCMD) test
GOGET   = $(GOCMD) get

DST = esa_attach \
	esa_get esa_getall \
	esa_new esa_update

all: $(DST)

esa_get: esa_get.go esa_common.go
	$(GOBUILD) -o $@ $^

esa_getall: esa_getall.go esa_common.go
	$(GOBUILD) -o $@ $^

esa_new: esa_new.go esa_common.go
	$(GOBUILD) -o $@ $^

esa_update: esa_update.go esa_common.go
	$(GOBUILD) -o $@ $^

esa_attach: esa_attach.go esa_common.go
	$(GOBUILD) -o $@ $^

clean:
	rm -f $(DST)
