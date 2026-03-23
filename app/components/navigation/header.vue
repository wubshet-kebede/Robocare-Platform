<template>
  <header
    class="fixed top-0 w-full z-50 border-b border-slate-800/50 bg-slate-950/80 backdrop-blur-xl"
  >
    <nav class="container mx-auto px-6 h-20 flex items-center justify-between">
      <!-- Branding -->
      <NuxtLink to="/" class="flex items-center gap-3 group">
        <div
          class="w-10 h-10 bg-gradient-to-br from-cyan-500 to-emerald-500 rounded-xl flex items-center justify-center shadow-lg shadow-cyan-500/20 group-hover:scale-105 transition-transform"
        >
          <span class="text-slate-950 font-black text-xl">H</span>
        </div>
        <div class="flex flex-col">
          <span class="text-white font-bold tracking-tight text-sm uppercase"
            >HealthBot IoT</span
          >
          <span
            class="text-slate-500 text-[10px] font-bold tracking-[0.2em] uppercase"
            >UoG · Ethiopia</span
          >
        </div>
      </NuxtLink>

      <!-- Navigation Links -->
      <div
        class="hidden lg:flex items-center gap-10 text-sm font-semibold text-slate-400"
      >
        <NuxtLink
          to="#about"
          class="hover:text-cyan-400 transition-colors uppercase tracking-widest"
          >About</NuxtLink
        >
        <NuxtLink
          to="#features"
          class="hover:text-cyan-400 transition-colors uppercase tracking-widest"
          >Features</NuxtLink
        >
        <NuxtLink
          to="#team"
          class="hover:text-cyan-400 transition-colors uppercase tracking-widest"
          >Team</NuxtLink
        >
        <NuxtLink
          to="#contact"
          class="hover:text-cyan-400 transition-colors uppercase tracking-widest"
          >Contact</NuxtLink
        >
      </div>

      <!-- Auth Actions -->
      <div class="flex items-center gap-6">
        <NuxtLink
          to="/login"
          class="text-sm font-bold text-slate-300 hover:text-white transition-colors"
        >
          Sign In
        </NuxtLink>
        <NuxtLink
          to="/signup"
          class="px-6 py-2.5 bg-white hover:bg-cyan-50 text-slate-950 font-bold rounded-full text-sm transition-all shadow-xl shadow-white/5 hover:shadow-cyan-500/10"
        >
          Join Project
        </NuxtLink>
      </div>
    </nav>
  </header>
</template>
