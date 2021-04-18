.PHONY: help, new-app, build-assets

%:
	@:

help: ## Show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

new-app: ## Create new application in DO App Platform
	doctl apps create --spec ./.do/app.yaml

build-assets: ## Rebuild CSS assets
	cd src && npm run build-prod && cd ../ \
	&& git add src/assets/* && git commit -m "Rebuilt CSS-assets for production environment at $(date)"