<template>
  <div
    class="bg-slate-900/50 border border-slate-800 p-10 rounded-[2.5rem] backdrop-blur-xl shadow-2xl text-center"
  >
    <Icon name="lucide:search-medical" class="text-5xl text-cyan-500 mb-6" />
    <h1 class="text-2xl font-bold text-white mb-2 tracking-tight">
      Find Your Facility
    </h1>
    <p class="text-slate-500 text-sm mb-8">
      Enter your hospital name to access your local robot fleet.
    </p>

    <div class="relative group">
      <Icon
        name="lucide:search"
        class="absolute left-4 top-1/2 -translate-y-1/2 text-slate-600 group-focus-within:text-cyan-500"
      />
      <input
        v-model="searchQuery"
        type="text"
        placeholder="e.g. Gondar University..."
        class="w-full bg-slate-950 border border-slate-800 rounded-xl pl-12 pr-5 py-4 text-white text-sm focus:border-cyan-500 outline-none transition-all"
        @keyup.enter="goToHospital"
      />
    </div>

    <button
      @click="goToHospital"
      class="w-full mt-6 py-4 bg-white text-slate-950 font-black rounded-xl hover:bg-cyan-50 transition-all shadow-xl shadow-white/5"
    >
      CONTINUE TO PORTAL
    </button>
  </div>
</template>

<script setup>
const searchQuery = ref("");

const goToHospital = () => {
  if (!searchQuery.value) return;

  // Convert "Gondar University" to "gondar-university" (the slug)
  const slug = searchQuery.value.toLowerCase().trim().replace(/\s+/g, "-");

  // Redirect to the dynamic page we built: login/[slug].vue
  navigateTo(`/login/${slug}`);
};

definePageMeta({
  layout: "auth",
});
</script>
