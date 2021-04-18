.PHONY: new-app

new-app:
	doctl --context do-startups-makers apps create --spec ./.do/app.yaml