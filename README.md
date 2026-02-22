# AuditFlow - Makes Your Audit Process Flow

AuditFlow is a centralized platform designed to simplify and streamline the entire audit process. From audit planning and initialization to monitoring, managing findings, and closing audits, AuditFlow ensures your organization can conduct audits efficiently and transparently.

---

## Features

- **Centralized Audit Management** – Manage all audits for multiple organizations in one platform.  
- **Role-Based Access** – Admins, auditors, and auditees have tailored dashboards.  
- **Organization & Department Management** – Structure your organization and assign roles seamlessly.  
- **Audit Planning & Assignment** – Schedule audits, assign auditors, and define auditees for each audit.  
- **Findings & NCR Management** – Track non-conformance reports (NCRs) and corrective actions from submission to closure.  
- **File Upload & Evidence Handling** – Attach checklists, evidence, and supporting documents securely.  
- **Workflow Automation** – Propose, review, accept, or reject corrective actions within the platform.

---

## Tech Stack

- **Frontend:** SvelteKit, TypeScript, TailwindCSS  
- **Backend:** Go (Golang), REST API  
- **Database:** PostgreSQL  
- **Storage:** S3-compatible object storage  
- **Authentication & Authorization:** JWT-based RBAC system

---

## Getting Started

### Prerequisites

- [Go](https://golang.org/) >= 1.21  
- [Node.js](https://nodejs.org/) LTS  
- [pnpm](https://pnpm.io/) package manager  
- PostgreSQL database  

### Installation

1. Clone the repository:

```bash
git clone https://github.com/shojib116/auditflow.git
cd auditflow
```

2. Setup backend:

```bash
cd server
go mod tidy
# Configure .env with your database and JWT settings
go run .
```

The backend will run at http://localhost:8080/ 

3. Setup frontend:

```bash
cd client
pnpm install
pnpm dev
```

The frontend will run at http://localhost:5173/ and connect to the backend API.

### Folder Structure

```bash
auditflow/
├── client/         # SvelteKit frontend
├── server/         # Go REST backend (modular monolith)
├── docker-compose.yml
├── .env
└── README.md
```

### Contributing

We welcome contributions! Please submit pull requests or open issues for bug fixes, improvements, or feature suggestions.

### License

This project is licensed under the MIT License.

---

AuditFlow – Making Your Audit Process Flow Smoothly and Transparently
