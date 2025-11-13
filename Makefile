# ---------------------------------------------------------
#  esbuild:

.PHONY: bundle
bundle: bundle/js bundle/css

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


# ---------------------------------------------------------
#  fmt:

fmt/ui:
	./node_modules/.bin/prettier --write ./ui

