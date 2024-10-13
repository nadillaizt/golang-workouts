.
├── **go.mod**
│   - Used by Go modules (if Go modules are enabled).
│
├── **Makefile**
│   - Used for `make` commands.
│
├── **assets/**
│   - Contains static assets like images, logos, etc.
│
├── **build/**
│   - Contains files for build and CI (Continuous Integration), such as Dockerfile, `.travis-ci.yml`, `.gitlab-ci.yml`, and files for packaging into formats like `.deb`, `.rpm`, `.pkg`.
│   ├── **ci/**
│   │   - CI-related files go here.
│   └── **package/**
│       - Files for building packages go here.
│
├── **cmd/**
│   - Contains the main application source code.
│   - In a monolith app, this includes core business logic (services, repositories).
│   - In a microservices architecture with a monorepo, this folder contains the source code divided by service.
│   ├── **your_app_1/**
│   ├── **your_app_2/**
│   └── ...
│
├── **configs/**
│   - Configuration files.
│
├── **deployments/**
│   - Orchestration, deployment, and CD-related files (e.g., `docker-compose.yml`, Kubernetes files).
│
├── **docs/**
│   - Design and documentation files.
│
├── **examples/**
│   - Example files.
│
├── **init/**
│   - System init files (e.g., systemd, upstart, sysv) and process manager configurations (e.g., runit, supervisord).
│
├── **internal/**
│   - Private app and library files. Unlike `pkg/`, packages in this folder can only be imported within this project.
│
├── **pkg/**
│   - Utility files that can be reused within this or other projects.
│   ├── **your_public_lib_1/**
│   ├── **your_public_lib_2/**
│   └── ...
│
├── **test/**
│   - Testing files. Structure is flexible. Unit tests should be placed in the same package as the files they test.
│
├── **vendor/**
│   - Contains cloned 3rd-party dependencies (used if vendor mode is enabled).
│
├── **web/**
│   - Contains web application files. For microservices, place the web app in `cmd/app`.
└── ...
