<template>
  <div
    class="bg-slate-900/50 border border-slate-800 p-8 md:p-10 rounded-[2.5rem] backdrop-blur-xl shadow-2xl"
  >
    <!-- Dynamic Header based on Slug -->
    <div class="text-center mb-8">
      <div
        class="inline-flex w-16 h-16 bg-cyan-500/10 rounded-2xl items-center justify-center mb-4 border border-cyan-500/20"
      >
        <Icon name="lucide:hospital" class="text-3xl text-cyan-500" />
      </div>
      <h1
        class="text-xl font-black text-white tracking-tighter uppercase italic"
      >
        {{ hospitalName }}
      </h1>
      <p
        class="text-slate-500 text-[10px] font-black uppercase tracking-[0.2em] mt-2"
      >
        Fleet Management Portal
      </p>
    </div>

    <form @submit.prevent="handleLogin" class="space-y-5">
      <!-- Role Selection -->
      <div
        class="grid grid-cols-2 gap-3 p-1 bg-slate-950 border border-slate-800 rounded-xl"
      >
        <button
          type="button"
          @click="role = 'doctor'"
          :class="
            role === 'doctor'
              ? 'bg-cyan-600 text-slate-950 shadow-lg shadow-cyan-500/20'
              : 'text-slate-500'
          "
          class="py-2 text-[10px] font-black uppercase tracking-widest rounded-lg transition-all"
        >
          Medical Doctor
        </button>
        <button
          type="button"
          @click="role = 'staff'"
          :class="
            role === 'staff'
              ? 'bg-cyan-600 text-slate-950 shadow-lg shadow-cyan-500/20'
              : 'text-slate-500'
          "
          class="py-2 text-[10px] font-black uppercase tracking-widest rounded-lg transition-all"
        >
          Tech Staff
        </button>
      </div>

      <!-- Professional Credentials -->
      <div class="space-y-4 pt-2">
        <div class="group">
          <label
            class="text-[10px] font-black uppercase tracking-widest text-slate-600 ml-1 mb-2 block"
            >System ID</label
          >
          <div class="relative">
            <Icon
              name="lucide:fingerprint"
              class="absolute left-4 top-1/2 -translate-y-1/2 text-slate-600 group-focus-within:text-cyan-500 transition-colors"
            />
            <input
              type="text"
              placeholder="EMP-2026-X"
              class="w-full bg-slate-950 border border-slate-800 rounded-xl pl-12 pr-5 py-4 text-white text-sm focus:border-cyan-500 outline-none transition-all"
            />
          </div>
        </div>

        <div class="group">
          <label
            class="text-[10px] font-black uppercase tracking-widest text-slate-600 ml-1 mb-2 block"
            >Access Key</label
          >
          <div class="relative">
            <Icon
              name="lucide:shield-lock"
              class="absolute left-4 top-1/2 -translate-y-1/2 text-slate-600 group-focus-within:text-cyan-500 transition-colors"
            />
            <input
              type="password"
              placeholder="••••••••"
              class="w-full bg-slate-950 border border-slate-800 rounded-xl pl-12 pr-5 py-4 text-white text-sm focus:border-cyan-500 outline-none transition-all"
            />
          </div>
        </div>
      </div>

      <button
        class="w-full py-4 bg-white hover:bg-cyan-50 text-slate-950 font-black rounded-xl transition-all shadow-xl shadow-white/5 mt-4 flex items-center justify-center gap-2 group"
      >
        ENTER DASHBOARD
        <Icon
          name="lucide:unplug"
          class="group-hover:rotate-12 transition-transform"
        />
      </button>
    </form>

    <div class="mt-8 pt-6 border-t border-slate-800/50 text-center">
      <NuxtLink
        to="/login"
        class="text-[10px] font-black text-slate-500 hover:text-cyan-400 uppercase tracking-[0.2em] transition-colors"
      >
        Change Facility
      </NuxtLink>
    </div>
  </div>
</template>

<script setup>
const route = useRoute();
const slug = route.params.slug; // Gets "gondar-university" from URL
const role = ref("doctor");

// Compute name from slug (e.g., gondar-university -> GONDAR UNIVERSITY)
const hospitalName = computed(() => {
  return slug ? slug.toString().replace(/-/g, " ") : "Unknown Facility";
});

definePageMeta({
  layout: "auth",
});

const handleLogin = () => {
  // Logic to redirect to the specific dashboard
  console.log(`Logging into ${slug} as ${role.value}`);
  navigateTo(`/dashboard/${slug}`);
};
</script>
