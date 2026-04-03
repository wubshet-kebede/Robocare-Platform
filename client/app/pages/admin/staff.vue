<template>
  <div class="space-y-10">
    <!-- Page Header -->
    <div class="flex items-center justify-between">
      <div>
        <h1 class="text-3xl font-black text-white tracking-tighter">
          Staff Management
        </h1>
        <p
          class="text-slate-500 text-sm font-medium mt-1 uppercase tracking-widest"
        >
          Manage permissions & invite medical personnel
        </p>
      </div>

      <!-- Invite Button -->
      <UiBaseButton variant="gradient" @click="showInviteModal = true">
        <template #icon-left><Icon name="lucide:user-plus" /></template>
        Issue New Invitation
      </UiBaseButton>
    </div>

    <!-- Stats Row (Ember Style) -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      <div
        v-for="stat in stats"
        :key="stat.label"
        class="bg-slate-900/40 border border-white/5 p-6 rounded-3xl backdrop-blur-xl"
      >
        <div
          class="text-slate-500 text-[10px] font-black uppercase tracking-widest mb-2"
        >
          {{ stat.label }}
        </div>
        <div class="text-3xl font-black text-white italic">
          {{ stat.value }}
        </div>
      </div>
    </div>

    <!-- Staff Table -->
    <div
      class="bg-slate-900/20 border border-white/5 rounded-[2.5rem] overflow-hidden backdrop-blur-md shadow-2xl"
    >
      <table class="w-full text-left border-collapse">
        <thead class="bg-white/5 border-b border-white/5">
          <tr>
            <th
              class="px-8 py-5 text-[10px] font-black uppercase tracking-[0.2em] text-slate-500"
            >
              Personnel
            </th>
            <th
              class="px-8 py-5 text-[10px] font-black uppercase tracking-[0.2em] text-slate-500"
            >
              Department
            </th>
            <th
              class="px-8 py-5 text-[10px] font-black uppercase tracking-[0.2em] text-slate-500"
            >
              Status
            </th>
            <th
              class="px-8 py-5 text-[10px] font-black uppercase tracking-[0.2em] text-slate-500 text-right"
            >
              Actions
            </th>
          </tr>
        </thead>
        <tbody class="divide-y divide-white/5">
          <tr
            v-for="member in staffList"
            :key="member.email"
            class="group hover:bg-white/5 transition-colors"
          >
            <td class="px-8 py-6">
              <div class="flex items-center gap-4">
                <div
                  class="w-10 h-10 rounded-full bg-slate-800 border border-white/10 flex items-center justify-center text-cyan-500 font-bold"
                >
                  {{ member.name.charAt(0) }}
                </div>
                <div>
                  <div class="text-white font-bold text-sm">
                    {{ member.name }}
                  </div>
                  <div class="text-slate-500 text-xs">{{ member.email }}</div>
                </div>
              </div>
            </td>
            <td class="px-8 py-6">
              <span
                class="text-slate-300 text-xs font-bold uppercase tracking-widest"
                >{{ member.role }}</span
              >
            </td>
            <td class="px-8 py-6">
              <UiStatusBadge
                :status="member.status"
                :pulse="member.status === 'pending'"
              >
                {{ member.status }}
              </UiStatusBadge>
            </td>
            <td class="px-8 py-6 text-right">
              <button
                class="p-2 text-slate-600 hover:text-red-500 transition-colors"
              >
                <Icon name="lucide:trash-2" />
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Invite Modal Component -->
    <DashboardInviteModal
      v-if="showInviteModal"
      @close="showInviteModal = false"
      @send="handleInvite"
    />
  </div>
</template>

<script setup>
definePageMeta({
  layout: "dashboard",
});
const showInviteModal = ref(false);

const stats = [
  { label: "Active Personnel", value: "14" },
  { label: "Pending Invites", value: "03" },
  { label: "Total Access Logs", value: "1,204" },
];

const staffList = ref([
  {
    name: "Nurse Samuel Ayalew",
    email: "samuel@uog.edu.et",
    role: "Chief Medical Officer",
    status: "online",
  },
  {
    name: "Dr. Wubshet",
    email: "wubshet@hospital.com",
    role: "Chief Medical Officer",
    status: "online",
  },
  {
    name: "Tewodros K.",
    email: "teddy@uog.edu.et",
    role: "IT Support",
    status: "stable",
  },
]);

const handleInvite = (newMember) => {
  staffList.value.push({
    ...newMember,
    status: "pending",
  });
  showInviteModal.value = false;
};
</script>
