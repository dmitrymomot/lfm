# The Leads Form Manager
Collect leads via a form on a website and manage them by simple admin panel.

[![Deploy to DO](https://www.deploytodo.com/do-btn-blue.svg)](https://cloud.digitalocean.com/apps/new?repo=github.com/dmitrymomot/lfm/tree/master)


## Environment Variables

| Name             | Is Required | Default Value      |
| ---------------- | ----------- | ------------------ |
| APP_PORT         | no          | 8000               |
| APP_BASE_URL     | no          | http://localhost   |
| DEBUG_MODE       | no          | false              |
| HEALTH_ENDPOINT  | no          | /health            |
| TEMPLATE_DIR     | no          | ./src/views        |
| FORM_CONFIG_PATH | no          | ./form.config.yaml |
| STATIC_FILES_DIR | no          | ./src/assets       |