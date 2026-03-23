<template>
  <div
    class="min-h-screen bg-slate-950 font-sans text-slate-200 flex items-center justify-center p-6 relative overflow-hidden"
  >
    <!-- Technical Background Decoration -->
    <div
      class="absolute inset-0 opacity-20 bg-[radial-gradient(circle_at_center,_var(--tw-gradient-stops))] from-cyan-900/20 via-transparent to-transparent"
    ></div>
    <div
      class="absolute top-0 left-0 w-full h-1 bg-gradient-to-r from-transparent via-cyan-500 to-transparent"
    ></div>

    <div class="relative z-10 w-full max-w-md">
      <!-- Back to Home Link -->
      <NuxtLink
        to="/"
        class="inline-flex items-center gap-2 text-slate-500 hover:text-cyan-400 mb-8 transition-colors text-xs font-bold uppercase tracking-widest"
      >
        <Icon name="lucide:arrow-left" />
        Back to Project Site
      </NuxtLink>

      <slot />
    </div>
  </div>
</template>
