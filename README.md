# SWACD Local KCP Demo

This repository contains a **fully working local KCP (Kubernetes Control Plane)** setup for the **SWACD (Secure Web Application Control Delivery)** project.  
It demonstrates the end-to-end flow of defining Custom Resources (CRDs), applying sample CRs, and running a lightweight controller that reconciles them.

---

## 🧱 Project Structure

| Folder | Description |
|---------|--------------|
| **kcp/** | Local KCP binary and workspace setup (root + swacd workspace). |
| **swacd-controller/** | Minimal Go-based controller using a dynamic client to detect CR changes and simulate reconciliation. |
| **swacd-yamls/** | All CRDs and CRs used for this demo — Tenant, EdgeRoute, OriginService, and Provider definitions. |

---

## 🚀 How to Run the Demo

### 1️⃣ Start KCP locally
```bash
cd kcp
./bin/kcp start --root-directory=.kcp --run-controllers=true
