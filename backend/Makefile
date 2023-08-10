# Makefile
# 参考 https://zenn.dev/tarotaro0/articles/1d5bf3e32d5ef2

# formatter, linter, language_serverなど, gqlgen
# 参考 https://zenn.dev/bun913/articles/f0a6c6177a4716
TOOLS=\
	golang.org/x/tools/cmd/goimports@latest \
	golang.org/x/lint/golint@latest \
	honnef.co/go/tools/cmd/staticcheck@latest \
	golang.org/x/tools/gopls@latest \
	github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest \
	github.com/ramya-rao-a/go-outline@latest \
	github.com/nsf/gocode@latest \
	github.com/acroca/go-symbols@latest \
	github.com/fatih/gomodifytags@latest \
	github.com/josharian/impl@latest \
	github.com/haya14busa/goplay/cmd/goplay@latest \
	github.com/go-delve/delve/cmd/dlv@latest \
	github.com/99designs/gqlgen@latest \
	golang.org/x/tools/cmd/godoc@latest 

install-tools:
	@for tool in $(TOOLS) ; do \
		go install $$tool; \
	done

.PHONY: build-docs

build-docs:
	wget -np -k -p -q -r --directory-prefix docs --no-host-directories -E http://localhost:6060/pkg/github.com/catechlounge/tf63_server/ ;


