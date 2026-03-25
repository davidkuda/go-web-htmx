.PHONY: bundle
bundle: bundle/check bundle/js bundle/css

# curly braces: "command group", shell syntax to group multiple commands into a single unit.
# the OR operator expects a single command (unit) afterwards.
bundle/check:
	@command -v esbuild >/dev/null 2>&1 || { \
	echo "esbuild not installed"; \
	exit 1; \
	}

bundle/js:
	@esbuild \
	--bundle \
	--minify \
	./ui/static/bundle.js \
	--outfile=./ui/static/dist/app.js

bundle/js/watch:
	@esbuild \
	--bundle \
	./ui/static/bundle.js \
	--outfile=./ui/static/dist/app.js \
	--watch


bundle/css:
	@esbuild \
	--bundle \
	--minify \
	./ui/static/bundle.css \
	--outfile=./ui/static/dist/styles.css

bundle/css/watch:
	@esbuild \
	--bundle \
	./ui/static/bundle.css \
	--outfile=./ui/static/dist/styles.css \
	--watch
