<template>
  <section id="about" class="py-24 bg-white">
    <div class="container mx-auto px-8">
      <div class="flex flex-col lg:flex-row items-center gap-16">
        <div class="flex-1 space-y-6">
          <h2
            class="text-cyan-600 font-semibold text-sm uppercase tracking-[0.25em]"
          >
            The Mission
          </h2>

          <h3
            class="text-4xl lg:text-5xl font-bold text-slate-900 tracking-tight leading-tight"
          >
            Reducing Infection Risks in <br />
            <span class="text-slate-400"> Ethiopian Hospitals. </span>
          </h3>

          <p class="text-slate-600 text-lg leading-relaxed">
            Manual health monitoring puts medical staff at high risk of
            cross-infection. Our project introduces an IoT-enabled mobile robot
            that navigates hospital wards autonomously using SLAM technology.
          </p>

          <ul class="space-y-4 pt-4">
            <li class="flex items-start gap-3 text-slate-700">
              <div
                class="mt-1 w-5 h-5 rounded-full bg-cyan-100 flex items-center justify-center text-cyan-600 text-xs font-bold"
              >
                ✓
              </div>
              <span> Autonomous vitals collection (BPM, SpO2, Temp). </span>
            </li>

            <li class="flex items-start gap-3 text-slate-700">
              <div
                class="mt-1 w-5 h-5 rounded-full bg-cyan-100 flex items-center justify-center text-cyan-600 text-xs font-bold"
              >
                ✓
              </div>
              <span> Real-time data syncing via MQTT and Firebase. </span>
            </li>
          </ul>
        </div>
        <div class="flex-1 w-full">
          <div
            class="bg-white border border-slate-200 p-10 rounded-3xl relative overflow-hidden shadow-md hover:shadow-lg transition-all group"
          >
            <div
              class="absolute top-0 right-0 p-6 opacity-5 group-hover:opacity-10 transition-opacity"
            >
              <Icon
                name="fluent:shield-task-32-filled"
                class="text-9xl text-cyan-500"
              />
            </div>

            <div class="relative z-10">
              <div
                class="text-5xl font-extrabold text-slate-900 mb-2 tracking-tight"
              >
                0%
              </div>

              <div
                class="text-cyan-600 font-semibold uppercase text-xs tracking-widest mb-6"
              >
                Direct Physical Contact Required
              </div>

              <p
                class="text-slate-500 text-sm italic border-l-2 border-slate-200 pl-4"
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
