<template>
  <footer class="bg-white border-t border-slate-200 pt-24 pb-10">
    <div class="container mx-auto px-8">
      <div class="grid grid-cols-1 md:grid-cols-12 gap-12 mb-20">
        <div class="md:col-span-5">
          <div class="flex items-center gap-3 mb-6">
            <div
              class="w-1.5 h-10 bg-gradient-to-b from-cyan-500 to-emerald-500 rounded-full"
            ></div>
            <h3 class="text-slate-900 font-semibold text-lg tracking-tight">
              IoT-Integrated SLAM Framework
            </h3>
          </div>

          <p class="text-slate-600 text-sm leading-relaxed max-w-md">
            Developed by the Department of Computer Engineering to address
            manual monitoring challenges in Ethiopian hospitals through
            autonomous robotics and real-time telemetry.
          </p>
        </div>
        <div class="md:col-span-4">
          <h4
            class="text-slate-400 text-xs font-semibold uppercase tracking-[0.25em] mb-6"
          >
            Project Developers
          </h4>

          <ul class="space-y-3 text-slate-700 text-sm">
            <li class="flex items-center gap-3 group">
              <span
                class="w-1.5 h-1.5 bg-slate-300 rounded-full group-hover:bg-cyan-500 transition-colors"
              ></span>
              <span class="group-hover:text-slate-900 transition-colors">
                Wubshet Ayellew
              </span>
            </li>
          </ul>
        </div>
        <div class="md:col-span-3">
          <h4
            class="text-slate-400 text-xs font-semibold uppercase tracking-[0.25em] mb-6"
          >
            Institution
          </h4>

          <p class="text-slate-600 text-sm leading-relaxed">
            <span class="text-slate-900 font-medium">
              University of Gondar </span
            ><br />
            Institute of Technology<br />
            Gondar, Ethiopia
          </p>
        </div>
      </div>
      <div
        class="pt-8 border-t border-slate-200 flex flex-col md:flex-row justify-between items-center gap-4"
      >
        <p class="text-slate-500 text-[11px] font-medium tracking-wide">
          © 2026 Health Robotics Framework. All rights reserved.
        </p>

        <div class="flex gap-8 text-sm text-slate-500">
          <span
            class="relative cursor-pointer hover:text-slate-900 transition-colors after:absolute after:-bottom-1 after:left-0 after:h-[1.5px] after:w-0 after:bg-cyan-500 after:transition-all hover:after:w-full"
          >
            Documentation
          </span>

          <span
            class="relative cursor-pointer hover:text-slate-900 transition-colors after:absolute after:-bottom-1 after:left-0 after:h-[1.5px] after:w-0 after:bg-cyan-500 after:transition-all hover:after:w-full"
          >
            Simulation Logs
          </span>
        </div>
      </div>
    </div>
  </footer>
</template>
