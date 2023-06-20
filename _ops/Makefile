BACKUP ?= undefined
DEPLOYMENT ?= default
HELM ?= helm
PLAYER ?= undefined
REPOSITORY ?= undefined
REVISION ?= undefined

POD ?= $(shell kubectl get pods --no-headers -o custom-columns=":metadata.name" | grep xi-$(DEPLOYMENT))

SET_REPOSITORY ?=
ifneq ($(REPOSITORY), undefined)
SET_REPOSITORY = --set repository=$(REPOSITORY)
endif

SET_REVISION ?=
ifneq ($(REVISION), undefined)
SET_REVISION = --set revision=$(REVISION)
endif

docker:
	REPOSITORY=$(REPOSITORY) REVISION=$(REVISION) docker compose build

publish:
	REPOSITORY=$(REPOSITORY) REVISION=$(REVISION) docker compose push

install:
	$(HELM) install $(DEPLOYMENT) -f "`pwd`/helm/xi/$(DEPLOYMENT).yaml" helm/xi $(SET_REPOSITORY) $(SET_REVISION)

upgrade:
	$(HELM) upgrade $(DEPLOYMENT) -f "`pwd`/helm/xi/$(DEPLOYMENT).yaml" helm/xi $(SET_REPOSITORY) $(SET_REVISION)

template:
	$(HELM) template $(DEPLOYMENT) -f "`pwd`/helm/xi/$(DEPLOYMENT).yaml" helm/xi --debug $(SET_REPOSITORY) $(SET_REVISION)

uninstall:
	$(HELM) uninstall $(DEPLOYMENT)

db-backup:
	kubectl exec -ti -c sql $(POD) -- python3 /opt/xi/tools/dbtool.py backup
	kubectl cp -c sql $(POD):/opt/xi/sql/backups ./backups

db-restore:
	kubectl cp -c sql $(BACKUP) $(POD):/opt/xi/sql/backups

db-admin:
	kubectl exec -ti -c sql $(POD) -- python3 /opt/xi/tools/dbtool.py

db-map-partitions:
	kubectl exec -ti -c sql $(POD) -- python3 /opt/xi/tools/map-partitions.py
