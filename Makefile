WORKSPACE ?= dev
PROJECT ?= go-api-355418
REGION ?= us-central1
REBUILD ?=

define header
  $(info $(START)▶▶▶ $(1)$(END))
endef

set-project:
	$(call header, "Setting project to $(PROJECT)...")
	gcloud config set project $(PROJECT)

tf-init:
	$(call header, Initializing terraform for deploy in $(WORKSPACE)...)
	(cd infrastructure && terraform workspace select $(WORKSPACE) && terraform init)

tf-plan:
	$(call header, Creating plan for deploy $(WORKSPACE)...)
	(cd infrastructure && terraform workspace select $(WORKSPACE) && terraform plan)

tf-apply:
	$(call header, Applying plan for deploy in $(WORKSPACE)...)
	(cd infrastructure && terraform apply)

# Deployment steps
deploy-api: set-project
	$(call header, Deploy api for project $(PROJECT) in workspace $(WORKSPACE)...)
	docker build . -t gcr.io/$(PROJECT)/$(WORKSPACE)-api
	docker push gcr.io/$(PROJECT)/$(WORKSPACE)-api && \
	gcloud run deploy $(WORKSPACE)-api --image=gcr.io/$(PROJECT)/$(WORKSPACE)-api --region $(REGION) --platform managed --allow-unauthenticated