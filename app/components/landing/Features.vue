<template>
  <section id="features" class="py-24 bg-slate-950 border-t border-slate-900">
    <div class="container mx-auto px-6">
      <div class="text-center max-w-3xl mx-auto mb-16">
        <h2
          class="text-cyan-500 font-black text-sm uppercase tracking-[0.3em] mb-4"
        >
          Capabilities
        </h2>
        <h3 class="text-4xl font-bold text-white mb-6">
          Advanced Healthcare Intelligence
        </h3>
        <p class="text-slate-400">
          Our platform integrates industrial-grade sensors with the ROS2
          ecosystem to provide a reliable, autonomous monitoring solution.
        </p>
      </div>

      <!-- Professional Bento Grid -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <!-- Feature -->
        <div
          class="md:col-span-2 group bg-slate-900/50 border border-slate-800 p-8 rounded-3xl hover:border-cyan-500/50 transition-all"
        >
          <div
            class="w-12 h-12 bg-cyan-500/10 rounded-xl flex items-center justify-center mb-6"
          >
            <Icon name="gravity-ui:map-pin" class="text-2xl text-cyan-500" />
          </div>
          <h4 class="text-xl font-bold text-white mb-3">
            Autonomous SLAM Mapping
          </h4>
          <p class="text-slate-400 text-sm leading-relaxed mb-6">
            Using Slam Toolbox and LiDAR, the robot builds a real-time occupancy
            grid of the hospital ward, allowing for precise point-to-point
            navigation without human intervention.
          </p>
          <div class="flex gap-2">
            <span
              class="px-3 py-1 bg-slate-800 rounded-full text-[10px] font-bold text-slate-300"
              >LIDAR</span
            >
            <span
              class="px-3 py-1 bg-slate-800 rounded-full text-[10px] font-bold text-slate-300"
              >ROS2 JAZZY</span
            >
          </div>
        </div>

        <!-- Feature 2: Vital Signs -->
        <div
          class="group bg-slate-900/50 border border-slate-800 p-8 rounded-3xl hover:border-emerald-500/50 transition-all"
        >
          <div
            class="w-12 h-12 bg-emerald-500/10 rounded-xl flex items-center justify-center mb-6"
          >
            <Icon name="octicon:pulse-24" class="text-2xl text-emerald-500" />
          </div>
          <h4 class="text-xl font-bold text-white mb-3">
            Multi-Sensor Telemetry
          </h4>
          <p class="text-slate-400 text-sm leading-relaxed">
            Integrated MAX30102 and MLX90614 sensors for non-contact temperature
            and heart rate monitoring.
          </p>
        </div>

        <!-- Feature 3: Cloud Sync -->
        <div
          class="group bg-slate-900/50 border border-slate-800 p-8 rounded-3xl hover:border-blue-500/50 transition-all"
        >
          <div
            class="w-12 h-12 bg-blue-500/10 rounded-xl flex items-center justify-center mb-6"
          >
            <Icon
              name="iconamoon:cloud-yes-light"
              class="text-2xl text-blue-500"
            />
          </div>
          <h4 class="text-xl font-bold text-white mb-3">
            Firebase Real-time Sync
          </h4>
          <p class="text-slate-400 text-sm leading-relaxed">
            Zero-latency data bridging ensures doctors receive alerts instantly
            on their dashboard.
          </p>
        </div>

        <!-- Feature 4: Telepresence (Spans 2 columns) -->
        <div
          class="md:col-span-2 group bg-slate-900/50 border border-slate-800 p-8 rounded-3xl hover:border-purple-500/50 transition-all"
        >
          <div class="flex flex-col md:flex-row gap-8 items-center">
            <div class="flex-1">
              <div
                class="w-12 h-12 bg-purple-500/10 rounded-xl flex items-center justify-center mb-6"
              >
                <Icon
                  name="pepicons-pencil:camera"
                  class="text-2xl text-purple-500"
                />
              </div>
              <h4 class="text-xl font-bold text-white mb-3">
                Bidirectional Telepresence
              </h4>
              <p class="text-slate-400 text-sm leading-relaxed">
                High-definition video streaming via Raspberry Pi Camera for
                virtual doctor-patient rounds.
              </p>
            </div>
            <div
              class="flex-1 bg-slate-950 rounded-2xl border border-slate-800 p-2 hidden md:block"
            >
              <div
                class="aspect-video bg-slate-900 rounded-lg flex items-center justify-center text-slate-700 text-xs italic"
              >
                [ Camera Stream Preview ]
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
