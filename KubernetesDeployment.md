# Kubernetes Deployment Documentation

## Access Links

- **Backend Ping Endpoint:** [http://103.192.198.15:30008/api/ping](http://103.192.198.15:30008/api/ping)
- **Frontend Website/App:** [http://103.192.198.15:30006/](http://103.192.198.15:30006/)

## Overview
This document describes the deployment of the Expense Tracker backend application on a Kubernetes cluster. The deployment includes the backend service, integrated PostgreSQL database, and supporting infrastructure.

---

## Deployment Details

- **Kubernetes Namespace:** `go-app`
- **Pods Running:**
  - `expense-tracker-frontend-7cf6bcc764-8hlkf` (Frontend)
  - `go-app-deployment-54b5d9fd5d-7f7f7` (Backend)
  - `go-app-deployment-54b5d9fd5d-867v5` (Backend)
  - `go-app-deployment-54b5d9fd5d-fpszt` (Backend)
  - `go-app-deployment-54b5d9fd5d-r227n` (Backend)
  - `postgres-deployment-6d47c57b46-7f5zt` (PostgreSQL)

- **Pod Status:** All pods are in `Running` state with 0 restarts.

---

## Integrated Services

- **Backend API:** Go (Golang)
- **Frontend:** (See frontend documentation)
- **Database:** PostgreSQL (deployed as a pod)

---

## Technologies & Tools Used

- **Programming Language:** Go (Golang)
- **Web Framework:** Custom/Standard Go HTTP
- **Database:** PostgreSQL
- **Containerization:** Docker
- **Orchestration:** Kubernetes
- **CI/CD:** Jenkins (see `jenkins.jenkins`)
- **API Testing:** Postman (see `expense.postman_collection.json`)
- **Logging:** Custom logging (see `Loging/logger.go`)
- **Authentication:** JWT-based authentication
- **Configuration:** YAML files (e.g., `go-app-deployment.yml`)
- **Cloud Native Networking:** Calico (as seen in `calico.yaml`)

---

## Deployment Files
- `go-app-deployment.yml`: Kubernetes deployment manifest for backend
- `app-config.yml`: Application configuration
- `calico.yaml`: Network policy configuration

---

## How to Check Pod Status
Run the following command on your Kubernetes master node:
```sh
kubectl get pods -n go-app
```

---

## Example Output
```
NAME                                        READY   STATUS    RESTARTS   AGE
expense-tracker-frontend-7cf6bcc764-8hlkf   1/1     Running   0          18m
go-app-deployment-54b5d9fd5d-7f7f7          1/1     Running   0          119m
go-app-deployment-54b5d9fd5d-867v5          1/1     Running   0          119m
go-app-deployment-54b5d9fd5d-fpszt          1/1     Running   0          119m
go-app-deployment-54b5d9fd5d-r227n          1/1     Running   0          119m
postgres-deployment-6d47c57b46-7f5zt        1/1     Running   0          25h
```

---

## Notes
- PostgreSQL is fully integrated and managed as a Kubernetes pod.
- All services are containerized and orchestrated via Kubernetes for scalability and reliability.
- For more details, refer to the respective YAML and configuration files in the repository. 