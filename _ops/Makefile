BACKUP ?= undefined
DEPLOYMENT ?= default
HELM ?= helm
PLAYER ?= undefined
ZONEIP ?= 127.0.0.1

SQL_POD ?= $(shell kubectl get pods --no-headers -o custom-columns=":metadata.name" | grep xi-$(DEPLOYMENT)-sql)

docker:
	TAG=$(TAG) docker compose build

install:
	$(HELM) install $(DEPLOYMENT) -f "`pwd`/helm/xi/$(DEPLOYMENT).yaml" helm/xi

upgrade:
	$(HELM) upgrade $(DEPLOYMENT) -f "`pwd`/helm/xi/$(DEPLOYMENT).yaml" helm/xi

template:
	$(HELM) template $(DEPLOYMENT) -f "`pwd`/helm/xi/$(DEPLOYMENT).yaml" helm/xi --debug

uninstall:
	$(HELM) uninstall $(DEPLOYMENT)

db-backup:
	kubectl exec -ti -c sql $(SQL_POD) -- python3 /opt/xi/tools/dbtool.py backup
	kubectl cp -c sql $(SQL_POD):/opt/xi/sql/backups ./backups

db-restore:
	kubectl cp -c sql $(BACKUP) $(SQL_POD):/opt/xi/sql/backups

db-admin:
	kubectl exec -ti -c sql $(SQL_POD) -- python3 /opt/xi/tools/dbtool.py
