<script setup>
const isMenuOpen = ref(false);
</script>
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
            Healthcare Robotics
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
      <div class="hidden lg:flex items-center gap-4">
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
          Sign Up
        </NuxtLink>
      </div>

      <!-- Mobile Menu Button -->
      <button
        @click="isMenuOpen = !isMenuOpen"
        class="lg:hidden p-2 rounded-xl border border-slate-200 hover:bg-slate-50"
      >
        <Icon :name="isMenuOpen ? 'lucide:x' : 'lucide:menu'" class="w-5 h-5" />
      </button>
    </nav>
    <Transition
      enter-active-class="transition duration-300 ease-out"
      enter-from-class="opacity-0 -translate-y-3"
      enter-to-class="opacity-100 translate-y-0"
      leave-active-class="transition duration-200 ease-in"
      leave-from-class="opacity-100 translate-y-0"
      leave-to-class="opacity-0 -translate-y-3"
    >
      <div
        v-if="isMenuOpen"
        class="lg:hidden border-t border-slate-200 bg-white"
      >
        <div class="container mx-auto px-8 py-6 flex flex-col gap-5">
          <NuxtLink
            to="#about"
            @click="isMenuOpen = false"
            class="text-slate-700 font-medium"
          >
            About
          </NuxtLink>

          <NuxtLink
            to="#features"
            @click="isMenuOpen = false"
            class="text-slate-700 font-medium"
          >
            Features
          </NuxtLink>

          <NuxtLink
            to="#team"
            @click="isMenuOpen = false"
            class="text-slate-700 font-medium"
          >
            Team
          </NuxtLink>

          <NuxtLink
            to="#contact"
            @click="isMenuOpen = false"
            class="text-slate-700 font-medium"
          >
            Contact
          </NuxtLink>

          <div class="pt-4 border-t border-slate-100 flex flex-col gap-3">
            <NuxtLink
              to="/login"
              class="w-full text-center py-3 rounded-xl border border-slate-200"
            >
              Sign In
            </NuxtLink>

            <NuxtLink
              to="/signup"
              class="w-full text-center py-3 rounded-xl bg-gradient-to-r from-cyan-500 to-emerald-500 text-white font-semibold"
            >
              Sign Up
            </NuxtLink>
          </div>
        </div>
      </div>
    </Transition>
  </header>
</template>
