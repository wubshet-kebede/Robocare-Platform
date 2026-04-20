<template>
  <header
    class="fixed top-0 w-full z-50 border-b border-slate-200 bg-white/90 backdrop-blur-xl"
  >
    <nav class="container mx-auto px-8 h-20 flex items-center justify-between">
      <NuxtLink to="/" class="flex items-center gap-3 group">
        <div
          class="relative w-11 h-11 rounded-full bg-gradient-to-br from-cyan-500 to-emerald-500 p-[2px] transition-all duration-300 group-hover:scale-105"
        >
          <div
            class="w-full h-full bg-white rounded-full flex items-center justify-center"
          >
            <img
              src="/assets/images/image.png"
              alt="RoboCare Logo"
              class="w-9 h-9 object-contain rounded-full"
            />
          </div>
        </div>

        <div class="flex flex-col leading-tight">
          <span
            class="text-slate-900 font-semibold tracking-tight text-lg group-hover:text-cyan-600 transition-colors"
          >
            RoboCare Platform
          </span>
          <span
            class="text-slate-400 text-[10px] font-medium tracking-[0.25em] uppercase"
          >
            UoG · Ethiopia
          </span>
        </div>
      </NuxtLink>
      <div
        class="hidden lg:flex items-center gap-10 text-sm font-medium text-slate-500"
      >
        <NuxtLink
          to="#about"
          class="relative hover:text-slate-900 transition-colors uppercase tracking-wider after:absolute after:-bottom-1 after:left-0 after:h-[2px] after:w-0 after:bg-cyan-500 after:transition-all hover:after:w-full"
        >
          About
        </NuxtLink>

        <NuxtLink
          to="#features"
          class="relative hover:text-slate-900 transition-colors uppercase tracking-wider after:absolute after:-bottom-1 after:left-0 after:h-[2px] after:w-0 after:bg-cyan-500 after:transition-all hover:after:w-full"
        >
          Features
        </NuxtLink>

        <NuxtLink
          to="#team"
          class="relative hover:text-slate-900 transition-colors uppercase tracking-wider after:absolute after:-bottom-1 after:left-0 after:h-[2px] after:w-0 after:bg-cyan-500 after:transition-all hover:after:w-full"
        >
          Team
        </NuxtLink>

        <NuxtLink
          to="#contact"
          class="relative hover:text-slate-900 transition-colors uppercase tracking-wider after:absolute after:-bottom-1 after:left-0 after:h-[2px] after:w-0 after:bg-cyan-500 after:transition-all hover:after:w-full"
        >
          Contact
        </NuxtLink>
      </div>

      <!-- Auth Actions -->
      <div class="flex items-center gap-4">
        <NuxtLink
          to="/login"
          class="text-sm font-medium text-slate-500 hover:text-slate-900 transition-colors"
        >
          Sign In
        </NuxtLink>

        <NuxtLink
          to="/signup"
          class="px-6 py-2.5 rounded-full text-sm font-semibold text-white bg-gradient-to-r from-cyan-500 to-emerald-500 hover:from-cyan-600 hover:to-emerald-600 transition-all shadow-md hover:shadow-lg"
        >
          Join Project
        </NuxtLink>
      </div>
    </nav>
  </header>
</template>
