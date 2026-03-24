<template>
  <footer class="bg-slate-950 border-t border-slate-900 pt-20 pb-10">
    <div class="container mx-auto px-6">
      <div class="grid grid-cols-1 md:grid-cols-12 gap-12 mb-16">
        <!-- Project Summary -->
        <div class="md:col-span-5">
          <div class="flex items-center gap-2 mb-6">
            <div class="w-2 h-8 bg-cyan-500 rounded-full"></div>
            <h3 class="text-white font-bold text-lg uppercase tracking-tighter">
              IoT-Integrated SLAM Framework
            </h3>
          </div>
          <p class="text-slate-500 text-sm leading-relaxed max-w-sm">
            Developed by the Department of Computer Engineering to address
            manual monitoring challenges in Ethiopian hospitals through
            autonomous robotics and real-time telemetry.
          </p>
        </div>

        <!-- Team Credits -->
        <div class="md:col-span-4">
          <h4
            class="text-white text-xs font-black uppercase tracking-[0.2em] mb-6 text-cyan-500/80"
          >
            Project Developers
          </h4>
          <ul class="grid grid-cols-1 gap-3 text-slate-400 text-sm font-medium">
            <li class="flex items-center gap-2">
              <span class="w-1 h-1 bg-slate-700 rounded-full"></span> Wubshet
              Ayellew
            </li>
            <!-- <li class="flex items-center gap-2">
              <span class="w-1 h-1 bg-slate-700 rounded-full"></span> Tewodros
              Kassanew
            </li>
            <li class="flex items-center gap-2">
              <span class="w-1 h-1 bg-slate-700 rounded-full"></span> Samuel
              Ayalew
            </li>
            <li class="flex items-center gap-2">
              <span class="w-1 h-1 bg-slate-700 rounded-full"></span> Samuel
              Teshale
            </li> -->
          </ul>
        </div>

        <!-- Institution -->
        <div class="md:col-span-3">
          <h4
            class="text-white text-xs font-black uppercase tracking-[0.2em] mb-6 text-cyan-500/80"
          >
            Institution
          </h4>
          <p class="text-slate-400 text-sm leading-relaxed">
            University of Gondar<br />
            Institute of Technology<br />
            Gondar, Ethiopia
          </p>
        </div>
      </div>

      <!-- Copyright Bar -->
      <div
        class="pt-8 border-t border-slate-900/50 flex flex-col md:flex-row justify-between items-center gap-4"
      >
        <p
          class="text-slate-600 text-[10px] font-bold uppercase tracking-widest"
        >
          © 2026 Health Robotics Framework.
        </p>
        <div class="flex gap-6 text-slate-500 text-xs">
          <span class="hover:text-cyan-500 cursor-pointer transition-colors"
            >Documentation</span
          >
          <span class="hover:text-cyan-500 cursor-pointer transition-colors"
            >Simulation Logs</span
          >
        </div>
      </div>
    </div>
  </footer>
</template>
