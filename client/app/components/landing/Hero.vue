<template>
  <section
    class="relative min-h-[90vh] flex items-center overflow-hidden bg-white border-b border-slate-200"
  >
    <!-- Subtle Grid Background -->
    <div
      class="absolute inset-0 opacity-[0.04] bg-[linear-gradient(to_right,#0f172a_1px,transparent_1px),linear-gradient(to_bottom,#0f172a_1px,transparent_1px)] bg-[size:4rem_4rem]"
    ></div>

    <!-- Soft Gradient Glow -->
    <div
      class="absolute inset-0 bg-gradient-to-b from-cyan-500/5 via-transparent to-white"
    ></div>

    <div class="container mx-auto px-8 relative z-10">
      <div class="grid lg:grid-cols-2 gap-14 items-center">
        <!-- Text Content -->
        <div class="space-y-8 text-center lg:text-left">
          <!-- Status Badge -->
          <div
            class="inline-flex items-center gap-2 px-4 py-1.5 rounded-full border border-cyan-200 bg-cyan-50"
          >
            <span class="flex h-2 w-2 relative">
              <span
                class="animate-ping absolute inline-flex h-full w-full rounded-full bg-cyan-400 opacity-60"
              ></span>
              <span
                class="relative inline-flex rounded-full h-2 w-2 bg-cyan-500"
              ></span>
            </span>
            <span
              class="text-cyan-700 text-[10px] font-semibold uppercase tracking-[0.25em]"
            >
              Project Status: Simulation v1.0
            </span>
          </div>

          <!-- Heading -->
          <h1
            class="text-5xl lg:text-7xl font-extrabold text-slate-900 leading-[1.05] tracking-tight"
          >
            Bridging
            <span
              class="text-transparent bg-clip-text bg-gradient-to-r from-cyan-600 to-emerald-500"
            >
              Healthcare
            </span>
            <br />
            with Robotics
          </h1>

          <!-- Description -->
          <p
            class="text-slate-600 text-lg leading-relaxed max-w-xl mx-auto lg:mx-0"
          >
            An IoT-integrated SLAM framework for remote patient monitoring.
            Designed for the specific infrastructural needs of
            <span
              class="text-slate-900 font-medium border-b border-cyan-400/40"
            >
              Ethiopian healthcare facilities </span
            >.
          </p>

          <!-- Actions -->
          <div
            class="flex flex-col sm:flex-row items-center justify-center lg:justify-start gap-5"
          >
            <NuxtLink
              to="/login"
              class="px-8 py-4 bg-gradient-to-r from-cyan-500 to-emerald-500 hover:from-cyan-600 hover:to-emerald-600 text-white font-semibold rounded-xl transition-all shadow-md hover:shadow-lg hover:-translate-y-0.5"
            >
              Launch Dashboard
            </NuxtLink>

            <NuxtLink
              to="#about"
              class="px-8 py-4 border border-slate-300 text-slate-700 font-semibold rounded-xl hover:bg-slate-100 transition-all"
            >
              View Framework
            </NuxtLink>
          </div>

          <!-- Specs -->
          <div
            class="flex flex-wrap justify-center lg:justify-start gap-10 pt-10 border-t border-slate-200"
          >
            <div v-for="spec in specs" :key="spec.label">
              <div class="text-slate-900 font-semibold text-xl">
                {{ spec.value }}
              </div>
              <div
                class="text-slate-400 text-[10px] font-semibold uppercase tracking-widest mt-1"
              >
                {{ spec.label }}
              </div>
            </div>
          </div>
        </div>

        <!-- Visual -->
        <div class="relative group hidden lg:block">
          <div
            class="absolute -inset-1 bg-gradient-to-r from-cyan-400 to-emerald-400 rounded-[2rem] blur opacity-20 group-hover:opacity-30 transition duration-700"
          ></div>

          <div
            class="relative bg-white border border-slate-200 rounded-[2rem] p-3 overflow-hidden shadow-xl"
          >
            <!-- Frame -->
            <div
              class="aspect-video bg-slate-100 rounded-2xl overflow-hidden relative border border-slate-200"
            >
              <img
                src="/assets/images/gazebo.png"
                alt="Gazebo View"
                class="w-full h-full object-cover"
              />

              <!-- Overlay -->
              <div class="absolute top-4 right-4 flex gap-2">
                <div
                  class="px-2 py-1 bg-white/90 rounded text-[9px] font-mono text-cyan-600 border border-cyan-200"
                >
                  LIDAR: ACTIVE
                </div>
                <div
                  class="px-2 py-1 bg-white/90 rounded text-[9px] font-mono text-emerald-600 border border-emerald-200"
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
