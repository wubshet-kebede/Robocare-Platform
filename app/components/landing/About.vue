<template>
  <section id="about" class="py-24 bg-slate-950">
    <div class="container mx-auto px-6">
      <div class="flex flex-col lg:flex-row items-center gap-16">
        <!-- Left: Professional Text Content -->
        <div class="flex-1 space-y-6">
          <h2
            class="text-cyan-500 font-black text-sm uppercase tracking-[0.3em]"
          >
            The Mission
          </h2>
          <h3
            class="text-4xl lg:text-5xl font-bold text-white tracking-tight leading-tight"
          >
            Reducing Infection Risks in <br />
            <span class="text-slate-500">Ethiopian Hospitals.</span>
          </h3>
          <p class="text-slate-400 text-lg leading-relaxed">
            Manual health monitoring puts medical staff at high risk of
            cross-infection. Our project introduces an IoT-enabled mobile robot
            that navigates hospital wards autonomously using SLAM technology.
          </p>
          <ul class="space-y-4 pt-4">
            <li class="flex items-start gap-3 text-slate-300">
              <div
                class="mt-1 w-5 h-5 rounded-full bg-cyan-500/20 flex items-center justify-center text-cyan-400"
              >
                ✓
              </div>
              <span>Autonomous vitals collection (BPM, SpO2, Temp).</span>
            </li>
            <li class="flex items-start gap-3 text-slate-300">
              <div
                class="mt-1 w-5 h-5 rounded-full bg-cyan-500/20 flex items-center justify-center text-cyan-400"
              >
                ✓
              </div>
              <span>Real-time data syncing via MQTT and Firebase.</span>
            </li>
          </ul>
        </div>

        <!-- Right: Professional "Stats" Card -->
        <div class="flex-1 w-full">
          <div
            class="bg-slate-900/50 border border-slate-800 p-8 rounded-3xl backdrop-blur-sm relative overflow-hidden group"
          >
            <div
              class="absolute top-0 right-0 p-4 opacity-10 group-hover:opacity-20 transition-opacity"
            >
              <Icon
                name="heroicons:shield-check-solid"
                class="text-9xl text-cyan-500"
              />
            </div>
            <div class="relative z-10">
              <div class="text-5xl font-black text-white mb-2 tracking-tighter">
                0%
              </div>
              <div
                class="text-cyan-500 font-bold uppercase text-xs tracking-widest mb-6"
              >
                Direct Physical Contact Required
              </div>
              <p
                class="text-slate-500 text-sm italic border-l-2 border-slate-800 pl-4"
              >
                "Our framework ensures that doctors can monitor critical patient
                data without entering high-risk zones, significantly lowering
                the workload of nurses."
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
