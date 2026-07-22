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
<img width="1815" height="918" alt="image" src="https://github.com/user-attachments/assets/6a01495c-936a-4c8a-9c57-0163c695aa60" />


### Telepresence page
<img width="1815" height="918" alt="image" src="https://github.com/user-attachments/assets/2837b36a-220a-4e1b-aca3-f2df297dc98a" />



### Staff Management page
<img width="1815" height="918" alt="image" src="https://github.com/user-attachments/assets/c834ad18-fc38-4e77-92f5-72336bf7a83d" />


---

## Demo
The web platform (frontend  and backend services) has been deployed and is available online.

The following video demonstrates the autonomous robot simulation inside the hospital environment, including robot navigation and telepresence features. This recording was captured locally during development using ROS 2 and Gazebo before deployment of the web platform.

- Demo Video:
   <iframe src="https://www.linkedin.com/embed/feed/update/urn:li:ugcPost:7476887508044247040?compact=1" height="399" width="504" frameborder="0" allowfullscreen="" title="Embedded post"></iframe>
- Source Code:
  https://github.com/wubshet-kebede/Robocare-Platform
---

## Research Context

This project was developed as a final-year Computer Engineering thesis project titled:

**"RoboCare: An IoT-Integrated Autonomous Mobile Robot Framework for Remote Patient Monitoring and Telepresence"**

The project explored the integration of autonomous robotics, IoT communication, real-time web technologies, and healthcare monitoring systems to improve remote healthcare assistance and patient interaction.

---
## Deployment

The platform is deployed using:

- Frontend: Vercel (Nuxt.js)
- Backend: Render (Golang REST API)
- Database: PostgreSQL
- Robot Simulation: ROS 2 Jazzy + Gazebo Harmonic (local simulation environment)
## Project Status

🚧 **Under Development**

RoboCare Platform is currently under active development. The web platform (frontend and backend services) has been deployed, while the autonomous robot system and additional hardware integrations are still being improved and tested.

The current version demonstrates the complete software architecture, including:
- Hospital management dashboard
- Patient monitoring platform
- Authentication and backend services
- Robot simulation environment
- Autonomous navigation and telepresence capabilities

Future updates will include further integration, testing, and improvements to the robotic system.

## Author

**Wubshet Ayellew**

Computer Engineer | Full Stack Developer

- GitHub: https://github.com/wubshet-kebede
- LinkedIn: https://www.linkedin.com/in/wubshet-kebede
