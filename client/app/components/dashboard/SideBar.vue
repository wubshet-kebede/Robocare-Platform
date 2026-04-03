<script setup>
// In a real app, you get this from your Auth Store (Pinia/Supabase)
const userRole = ref("admin"); // Try changing this to 'doctor' or 'receptionist'

const menuGroups = [
  {
    title: "Overview",
    roles: ["admin", "doctor", "receptionist"],
    links: [
      {
        name: "Dashboard",
        path: "/dashboard",
        icon: "lucide:layout-dashboard",
      },
    ],
  },
  {
    title: "Clinical",
    roles: ["doctor", "nurse"],
    links: [
      { name: "Robot Command", path: "/robot", icon: "lucide:bot" },
      { name: "Vitals Monitor", path: "/vitals", icon: "lucide:activity" },
      {
        name: "Patient Records",
        path: "/patients",
        icon: "lucide:folder-heart",
      },
    ],
  },
  {
    title: "Administration",
    roles: ["admin"],
    links: [
      { name: "Staff Management", path: "/admin/staff", icon: "lucide:users" },
      {
        name: "Hospital Settings",
        path: "/admin/settings",
        icon: "lucide:building-2",
      },
    ],
  },
  {
    title: "Operations",
    roles: ["receptionist", "admin"],
    links: [{ name: "Bed Management", path: "/beds", icon: "lucide:bed" }],
  },
];

const filteredMenu = computed(() => {
  return menuGroups.filter((group) => group.roles.includes(userRole.value));
});
</script>

<template>
  <aside
    class="w-72 h-screen bg-slate-950 border-r border-white/5 flex flex-col p-6"
  >
    <!-- Brand -->
    <div class="flex items-center gap-3 mb-12">
      <div class="w-8 h-8 bg-cyan-500 rounded-lg"></div>
      <span class="text-white font-black uppercase tracking-widest text-xs"
        >RoboCare</span
      >
    </div>

    <!-- Navigation Groups -->
    <div class="space-y-8 flex-1">
      <div v-for="group in filteredMenu" :key="group.title">
        <h4
          class="text-[10px] font-black uppercase tracking-[0.3em] text-slate-600 mb-4"
        >
          {{ group.title }}
        </h4>
        <ul class="space-y-2">
          <li v-for="link in group.links" :key="link.name">
            <NuxtLink
              :to="link.path"
              class="flex items-center gap-3 px-4 py-3 rounded-xl text-slate-400 hover:bg-white/5 hover:text-white transition-all group"
            >
              <Icon
                :name="link.icon"
                class="text-lg group-hover:text-cyan-500"
              />
              <span class="text-xs font-bold">{{ link.name }}</span>
            </NuxtLink>
          </li>
        </ul>
      </div>
    </div>

    <!-- User Profile & Log Out -->
    <div class="pt-6 border-t border-white/5">
      <button
        class="w-full text-left text-slate-500 hover:text-red-500 text-[10px] font-black uppercase tracking-widest transition-colors"
      >
        Terminate Session
      </button>
    </div>
  </aside>
</template>
