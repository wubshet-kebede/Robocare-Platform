<template>
  <div
    class="w-full max-w-5xl mx-auto bg-slate-950 p-6 sm:p-8 rounded-2xl border border-slate-800/80 shadow-2xl space-y-8 select-none text-slate-100"
  >
    <!-- Header -->
    <div class="text-center space-y-1">
      <h2 class="text-xs font-bold tracking-widest text-indigo-400 uppercase">
        System Topology Map
      </h2>
      <p class="text-xl font-bold text-white tracking-tight">
        IoT-Integrated Healthcare Robotic Framework
      </p>
      <p class="text-xs text-slate-400 max-w-xl mx-auto">
        Interactive blueprint displaying structural cross-layer data flows,
        security boundaries, and telemetry loops.
      </p>
    </div>

    <!-- LAYER 1: FRONTEND -->
    <div
      class="border border-emerald-500/20 bg-emerald-950/5 p-5 rounded-xl transition-all duration-300 hover:border-emerald-500/40"
    >
      <div class="flex items-center justify-between mb-3">
        <div
          class="text-xs font-bold text-emerald-400 tracking-wider uppercase"
        >
          1. Frontend Application Layer
        </div>
        <span class="h-2 w-2 rounded-full bg-emerald-400 animate-ping"></span>
      </div>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div
          class="bg-slate-900/90 border border-slate-800 p-4 rounded-lg flex items-center justify-between shadow-md"
        >
          <div>
            <div class="font-semibold text-sm text-slate-200">
              Nuxt.js Web Healthcare Dashboard
            </div>
            <div class="text-xs text-slate-400 mt-0.5">
              Hospital registration, user authentication, & task mgmt
            </div>
          </div>
          <span
            class="bg-emerald-500/10 text-emerald-400 text-[10px] font-mono font-bold px-2 py-0.5 rounded border border-emerald-500/20"
            >Nuxt 3</span
          >
        </div>
        <div
          class="bg-slate-900/90 border border-slate-800 p-4 rounded-lg flex items-center justify-between shadow-md"
        >
          <div>
            <div class="font-semibold text-sm text-slate-200">
              WebRTC Telepresence / Viewport
            </div>
            <div class="text-xs text-slate-400 mt-0.5">
              Low-latency real-time video consultation monitor
            </div>
          </div>
          <span
            class="bg-emerald-500/10 text-emerald-400 text-[10px] font-mono font-bold px-2 py-0.5 rounded border border-emerald-500/20"
            >MediaStream</span
          >
        </div>
      </div>
    </div>

    <!-- LINK BLOCK 1: FRONTEND TO BACKEND -->
    <div class="flex flex-col items-center justify-center -my-4 relative h-10">
      <div
        class="h-full w-0.5 bg-gradient-to-b from-emerald-500 to-orange-500 relative"
      >
        <div
          class="absolute w-1.5 h-1.5 bg-emerald-400 rounded-full left-1/2 -translate-x-1/2 animate-bounce-down"
        ></div>
        <div
          class="absolute w-1.5 h-1.5 bg-orange-400 rounded-full left-1/2 -translate-x-1/2 bottom-0 animate-bounce-up"
        ></div>
      </div>
      <div
        class="absolute text-[9px] font-mono text-slate-400 bg-slate-950 px-2.5 py-0.5 border border-slate-800 rounded-full shadow-sm"
      >
        RESTful APIs & WebSockets
      </div>
    </div>

    <!-- LAYER 2: BACKEND -->
    <div
      class="border border-orange-500/20 bg-orange-950/5 p-5 rounded-xl transition-all duration-300 hover:border-orange-500/40"
    >
      <div
        class="text-xs font-bold text-orange-400 tracking-wider uppercase mb-3"
      >
        2. Backend Infrastructure Layer
      </div>
      <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
        <div
          class="md:col-span-2 bg-slate-900/90 border border-slate-800 p-4 rounded-lg flex items-center justify-between shadow-md"
        >
          <div>
            <div class="font-semibold text-sm text-orange-300">
              Golang Core Service Engine
            </div>
            <div class="text-xs text-slate-400 mt-0.5">
              Authentication processing, coordination services & MQTT/WebRTC
              modules
            </div>
          </div>
          <span
            class="bg-orange-500/10 text-orange-400 text-[10px] font-mono font-bold px-2 py-0.5 rounded border border-orange-500/20"
            >Go Lang</span
          >
        </div>
        <div
          class="bg-slate-900/90 border border-slate-800 p-4 rounded-lg flex flex-col justify-between shadow-md"
        >
          <div class="font-semibold text-sm text-slate-200">
            PostgreSQL Database
          </div>
          <div
            class="text-[10px] text-slate-400 font-mono mt-2 bg-slate-950 p-1 rounded text-center border border-slate-800"
          >
            via GORM ORM
          </div>
        </div>
        <div
          class="bg-slate-900/90 border border-slate-800 p-4 rounded-lg flex flex-col justify-between shadow-md"
        >
          <div class="font-semibold text-sm text-slate-200">
            WebRTC Signaling
          </div>
          <div
            class="text-[10px] text-slate-400 font-mono mt-2 bg-slate-950 p-1 rounded text-center border border-slate-800"
          >
            ICE / SDP Session
          </div>
        </div>
      </div>
    </div>

    <!-- LINK BLOCK 2: BACKEND TO MQTT -->
    <div class="flex flex-col items-center justify-center -my-4 relative h-8">
      <div
        class="h-full w-0.5 bg-gradient-to-b from-orange-500 to-purple-500"
      ></div>
    </div>

    <!-- LAYER 3: MQTT LAYER -->
    <div
      class="border border-purple-500/20 bg-purple-950/5 p-4 rounded-xl flex flex-col sm:flex-row sm:items-center sm:justify-between gap-3 transition-all duration-300 hover:border-purple-500/40"
    >
      <div>
        <div class="text-xs font-bold text-purple-400 tracking-wider uppercase">
          3. MQTT Communication Layer
        </div>
        <div class="text-sm font-semibold text-slate-200 mt-0.5">
          Central Distributed MQTT Message Broker
        </div>
        <div class="text-xs text-slate-400 mt-1">
          Topics Matrix:
          <span
            class="text-purple-300 font-mono bg-slate-900 px-1 py-0.5 rounded border border-slate-800"
            >/nav_goal</span
          >
          <span
            class="text-purple-300 font-mono bg-slate-900 px-1 py-0.5 rounded border border-slate-800 ml-1"
            >/vitals</span
          >
        </div>
      </div>
      <div
        class="self-start sm:self-center bg-slate-900/90 border border-purple-500/30 px-3 py-1.5 rounded-lg text-xs font-mono text-purple-300 flex items-center gap-1.5 shadow-inner"
      >
        <span>🔒</span>
        <span class="font-bold uppercase tracking-wider text-[10px]"
          >HMAC Authenticated Pub/Sub</span
        >
      </div>
    </div>

    <!-- LINK BLOCK 3: MQTT TO ROBOT -->
    <div class="flex flex-col items-center justify-center -my-4 relative h-8">
      <div
        class="h-full w-0.5 bg-gradient-to-b from-purple-500 to-cyan-500"
      ></div>
    </div>

    <!-- LAYER 4: ROBOTIC SUBSYSTEM -->
    <div
      class="border border-cyan-500/20 bg-cyan-950/5 p-5 rounded-xl transition-all duration-300 hover:border-cyan-500/40"
    >
      <div
        class="text-xs font-bold text-cyan-400 tracking-wider uppercase mb-3"
      >
        4. Robotic Subsystem Layer (Simulated Edge Node)
      </div>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
        <div
          class="bg-slate-900/90 border border-cyan-500/20 p-4 rounded-lg relative overflow-hidden shadow-md"
        >
          <div class="font-semibold text-sm text-cyan-300">
            ROS-MQTT Bridge Node
          </div>
          <div class="text-xs text-slate-400 mt-1">
            Translates broker payloads to internal ROS topics. Intercepts
            commands & validates HMAC tokens.
          </div>
        </div>
        <div
          class="bg-slate-900/90 border border-slate-800 p-4 rounded-lg shadow-md"
        >
          <div class="font-semibold text-sm text-slate-200">
            Nav2 & SLAM Toolbox
          </div>
          <div class="text-xs text-slate-400 mt-1">
            Handles path planning, occupancy grid parsing, obstacle avoidance &
            goal execution via twist_mux.
          </div>
        </div>
        <div
          class="bg-slate-900/90 border border-slate-800 p-4 rounded-lg shadow-md"
        >
          <div class="font-semibold text-sm text-slate-200">
            WebRTC Media Pipeline
          </div>
          <div class="text-xs text-slate-400 mt-1">
            Captures camera frames from ROS topics, compresses stream with
            low-latency jitter buffers.
          </div>
        </div>
      </div>
      <div
        class="mt-4 text-center text-[10px] font-mono text-cyan-400/80 bg-cyan-950/30 py-1.5 rounded-lg border border-cyan-500/10"
      >
        Execution Environment Base: ROS 2 Jazzy & Gazebo Harmonic Simulation
        Runtime
      </div>
    </div>

    <!-- LINK BLOCK 4: ROBOT TO SENSORS -->
    <div class="flex flex-col items-center justify-center -my-4 relative h-8">
      <div
        class="h-full w-0.5 bg-gradient-to-b from-cyan-500 to-amber-500"
      ></div>
    </div>

    <!-- LAYER 5: SENSORS -->
    <div
      class="border border-amber-500/20 bg-amber-950/5 p-5 rounded-xl transition-all duration-300 hover:border-amber-500/40"
    >
      <div
        class="text-xs font-bold text-amber-400 tracking-wider uppercase mb-3"
      >
        5. Hardware Sensor Layer (Simulated Components)
      </div>
      <div class="grid grid-cols-1 sm:grid-cols-3 gap-4 text-center">
        <div
          class="bg-slate-900/90 border border-slate-800 p-3 rounded-lg shadow-sm"
        >
          <div class="font-semibold text-sm text-slate-200">LiDAR Sensor</div>
          <div class="text-[10px] text-slate-400 mt-0.5">
            Laser Scan Sweeps for SLAM Maps
          </div>
        </div>
        <div
          class="bg-slate-900/90 border border-slate-800 p-3 rounded-lg shadow-sm"
        >
          <div class="font-semibold text-sm text-slate-200">RGB Camera</div>
          <div class="text-[10px] text-slate-400 mt-0.5">
            Visual Video Stream Ingestion Feed
          </div>
        </div>
        <div
          class="bg-slate-900/90 border border-slate-800 p-3 rounded-lg shadow-sm"
        >
          <div class="font-semibold text-sm text-slate-200">IMU Sensor</div>
          <div class="text-[10px] text-slate-400 mt-0.5">
            Odometry & Balance Calculations
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
@keyframes bounceDown {
  0%,
  100% {
    transform: translate(-50%, 0);
    opacity: 0.3;
  }
  50% {
    transform: translate(-50%, 12px);
    opacity: 1;
  }
}
@keyframes bounceUp {
  0%,
  100% {
    transform: translate(-50%, 0);
    opacity: 0.3;
  }
  50% {
    transform: translate(-50%, -12px);
    opacity: 1;
  }
}

.animate-bounce-down {
  animation: bounceDown 2.5s infinite ease-in-out;
}
.animate-bounce-up {
  animation: bounceUp 2.5s infinite ease-in-out;
}
</style>
