<template>
  <section
    class="relative min-h-[90vh] flex items-center overflow-hidden bg-slate-950 border-b border-slate-900"
  >
    <!-- Technical Grid Background -->
    <div
      class="absolute inset-0 opacity-10 bg-[linear-gradient(to_right,#1e293b_1px,transparent_1px),linear-gradient(to_bottom,#1e293b_1px,transparent_1px)] bg-[size:4rem_4rem]"
    ></div>
    <div
      class="absolute inset-0 bg-gradient-to-b from-cyan-500/5 via-transparent to-slate-950"
    ></div>

    <div class="container mx-auto px-6 relative z-10">
      <div class="grid lg:grid-cols-2 gap-12 items-center">
        <!-- Text Content -->
        <div class="space-y-8 text-center lg:text-left">
          <div
            class="inline-flex items-center gap-2 px-3 py-1 rounded-full border border-cyan-500/20 bg-cyan-500/5 backdrop-blur-sm"
          >
            <span class="flex h-2 w-2">
              <span
                class="animate-ping absolute inline-flex h-full w-full rounded-full bg-cyan-400 opacity-75"
              ></span>
              <span
                class="relative inline-flex rounded-full h-2 w-2 bg-cyan-500"
              ></span>
            </span>
            <span
              class="text-cyan-400 text-[10px] font-black uppercase tracking-[0.3em]"
              >Project Status: Simulation v1.0</span
            >
          </div>

          <h1
            class="text-5xl lg:text-7xl font-black text-white leading-[1.05] tracking-tighter"
          >
            Bridging
            <span
              class="text-transparent bg-clip-text bg-gradient-to-r from-cyan-400 to-emerald-400"
              >Healthcare</span
            >
            <br />
            with Robotics
          </h1>

          <p
            class="text-slate-400 text-lg leading-relaxed max-w-xl mx-auto lg:mx-0 font-medium"
          >
            An IoT-integrated SLAM framework for remote patient monitoring.
            Designed for the specific infrastructural needs of
            <span class="text-white border-b border-cyan-500/30"
              >Ethiopian healthcare facilities</span
            >.
          </p>

          <div
            class="flex flex-col sm:flex-row items-center justify-center lg:justify-start gap-5"
          >
            <NuxtLink
              to="/login"
              class="px-8 py-4 bg-cyan-500 hover:bg-cyan-400 text-slate-950 font-black rounded-xl transition-all shadow-[0_20px_40px_rgba(6,182,212,0.2)] hover:-translate-y-1"
            >
              LAUNCH DASHBOARD
            </NuxtLink>
            <NuxtLink
              to="#about"
              class="px-8 py-4 border border-slate-800 text-white font-bold rounded-xl hover:bg-slate-900 transition-all"
            >
              VIEW FRAMEWORK
            </NuxtLink>
          </div>

          <!-- Component Specs -->
          <div
            class="flex flex-wrap justify-center lg:justify-start gap-10 pt-10 border-t border-slate-900"
          >
            <div v-for="spec in specs" :key="spec.label">
              <div class="text-white font-bold text-xl">{{ spec.value }}</div>
              <div
                class="text-slate-500 text-[9px] font-black uppercase tracking-widest mt-1"
              >
                {{ spec.label }}
              </div>
            </div>
          </div>
        </div>

        <!-- Robot Visual / Gazebo Frame -->
        <div class="relative group hidden lg:block">
          <div
            class="absolute -inset-1 bg-gradient-to-r from-cyan-500 to-emerald-500 rounded-[2rem] blur opacity-20 group-hover:opacity-40 transition duration-1000"
          ></div>
          <div
            class="relative bg-slate-900 border border-slate-800 rounded-[2rem] p-3 overflow-hidden shadow-2xl"
          >
            <!-- Simulated Gazebo View -->
            <div
              class="aspect-video bg-slate-950 rounded-2xl overflow-hidden relative border border-slate-800"
            >
              <img
                src="/assets/images/gazebo.png"
                alt="Gazebo View"
                class="w-full h-full object-cover opacity-80"
              />
              <!-- Telemetry Overlay -->
              <div class="absolute top-4 right-4 flex gap-2">
                <div
                  class="px-2 py-1 bg-slate-900/90 rounded text-[9px] font-mono text-cyan-400 border border-cyan-500/20"
                >
                  LIDAR: ACTIVE
                </div>
                <div
                  class="px-2 py-1 bg-slate-900/90 rounded text-[9px] font-mono text-emerald-400 border border-emerald-500/20"
                >
                  SLAM: MAP GENERATED
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
const specs = [
  { label: "Platform", value: "ROS2 Jazzy" },
  { label: "Sensing", value: "Multi-Modal" },
  { label: "Protocol", value: "MQTT / micro-ROS" },
];
</script>
