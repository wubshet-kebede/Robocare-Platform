# Nuxt Minimal Starter

Look at the [Nuxt documentation](https://nuxt.com/docs/getting-started/introduction) to learn more.

## Setup

Make sure to install dependencies:

```bash
# npm
npm install

# pnpm
pnpm install

# yarn
yarn install

# bun
bun install
```

## Development Server

Start the development server on `http://localhost:3000`:

```bash
# npm
npm run dev

# pnpm
pnpm dev

# yarn
yarn dev

# bun
bun run dev
```

## Production

Build the application for production:

```bash
# npm
npm run build

# pnpm
pnpm build

# yarn
yarn build

# bun
bun run build
```

Locally preview production build:

```bash
# npm
npm run preview

# pnpm
pnpm preview

# yarn
yarn preview

# bun
bun run preview
```

Check out the [deployment documentation](https://nuxt.com/docs/getting-started/deployment) for more information.
# RoboCare

## Overview

RoboCare is an IoT-integrated autonomous mobile robot platform designed for remote patient monitoring and telepresence in healthcare environments. The system combines robotics, real-time communication, and web technologies to enable healthcare professionals to remotely monitor patient conditions, track robot status, and interact through a centralized dashboard.

---

## Key Features

- Remote patient monitoring dashboard
- Real-time robot telemetry and status tracking
- Secure JWT-based authentication and authorization
- Real-time updates using WebSocket communication
- MQTT-based IoT communication architecture
- WebRTC-powered telepresence support
- RESTful API architecture
- PostgreSQL-backed data management
- Responsive and modern web interface

---

## Technology Stack

### Frontend
- Nuxt 4
- Vue 3
- TypeScript
- Tailwind CSS

### Backend
- Golang (`net/http`)
- GORM
- JWT Authentication
- WebSocket
- MQTT

### Database
- PostgreSQL

### Robotics & Simulation
- ROS 2
- Gazebo Harmonic
- Nav2
- SLAM Toolbox

---

## Project Highlights

- Designed and developed as a final-year Computer Engineering thesis project.
- Successfully integrated robotics, IoT, backend services, databases, and real-time communication into a unified healthcare platform.
- Demonstrated autonomous navigation, patient monitoring, and telepresence capabilities within a simulated hospital environment.
- Combined multiple technologies across robotics, distributed systems, web development, and healthcare-focused software engineering.

---

## Architecture

```text
Robot Layer
    │
    ├── MQTT Telemetry
    │
Backend Services (Golang)
    │
    ├── REST APIs
    ├── WebSocket Updates
    ├── Authentication
    │
PostgreSQL Database
    │
Frontend Dashboard (Nuxt 4)
```

---

## Screenshots

### Login Page
<img width="1844" height="917" alt="image" src="https://github.com/user-attachments/assets/f09d44b3-7ef2-4c04-a862-a4d3ce74ea54" />


### Dashboard
_Add screenshot here_

### Patient Monitoring
_Add screenshot here_

### Robot Monitoring
_Add screenshot here_

---

## Demo

- Demo Video: _Add LinkedIn or YouTube demo link here_
- Source Code: _Add repository link here_

---

## Research Context

This project was developed as a final-year Computer Engineering thesis project titled:

**"RoboCare: An IoT-Integrated Autonomous Mobile Robot Framework for Remote Patient Monitoring and Telepresence"**

The project explored the integration of autonomous robotics, IoT communication, real-time web technologies, and healthcare monitoring systems to improve remote healthcare assistance and patient interaction.

---

## Author

**Wubshet Ayellew**

Computer Engineer | Full Stack Developer

- GitHub: https://github.com/wubshet-kebede
- LinkedIn: https://www.linkedin.com/in/wubshet-kebede
