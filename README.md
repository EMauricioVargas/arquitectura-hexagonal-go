mkdir -p cmd/api internal/{adapters,application,domain,ports,config,utils}
mkdir -p internal/adapters/{db,http,cache,queue,storage}
mkdir -p internal/domain/{entities,repositories,events}
mkdir -p internal/ports/{in,out}
mkdir -p api docs scripts test
touch cmd/api/main.go go.mod go.sum README.md .gitignore Makefile