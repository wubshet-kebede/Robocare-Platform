<script setup>
definePageMeta({
  layout: "dashboard",
});

const loading = ref(false);
const isModalOpen = ref(false);
const openInviteModal = () => {
  isModalOpen.value = true;
};
// const submit = handleSubmit(async (values) => {
//   console.log("Form values:", values);
//   try {
//     loading.value = true;

//     const res = await inviteStaff({
//       email: values.email,
//       role: values.role,
//     });
//     console.log("invitation response:", res);
//     // const user = await me();
//     // user.value = user;
//     // console.log("Logged user:", user);
//     // await navigateTo(`/${user.hospital.slug}/dashboard`);
//   } catch (err) {
//     console.log(err);
//   } finally {
//     loading.value = false;
//   }
// });
const activeTab = ref("All Staff");
const tabs = computed(() => [
  { name: "All Staff", count: staffList.value.length },
  {
    name: "Doctors",
    count: staffList.value.filter((s) => s.role === "Doctors").length,
  },
  {
    name: "Nurses",
    count: staffList.value.filter((s) => s.role === "Nurses").length,
  },
  {
    name: "Administration",
    count: staffList.value.filter((s) => s.role === "Administration").length,
  },
]);
const staffList = ref([
  {
    id: 1,
    name: "Dr. Sarah Mitchell",
    role: "Doctors",
    department: "Cardiology",
    phone: "(555) 201-3344",
    email: "s.mitchell@hospital.org",
    status: "On Duty",
    rating: 4.9,
    exp: 14,
    shift: "Morning",
  },
  {
    id: 2,
    name: "Marla Santos",
    role: "Nurses",
    department: "Emergency",
    phone: "(555) 201-1122",
    email: "m.santos@hospital.org",
    status: "On Duty",
    rating: 4.9,
    exp: 12,
    shift: "Morning",
  },
  {
    id: 3,
    name: "David Thompson",
    role: "Administration",
    department: "Hospital Admin",
    phone: "(555) 201-7788",
    email: "d.thompson@hospital.org",
    status: "On Duty",
    rating: 4.5,
    exp: 8,
    shift: "Morning",
  },
]);
const filteredStaff = computed(() => {
  if (activeTab.value === "All Staff") return staffList.value;
  return staffList.value.filter((staff) => staff.role === activeTab.value);
});
</script>
<template>
  <ModalsInviteForm v-if="isModalOpen" v-model="isModalOpen">
  </ModalsInviteForm>
  <div
    class="mb-6 flex flex-col gap-4 sm:flex-row items-center justify-between"
  >
    <div>
      <h1 class="text-2xl font-bold tracking-tight">Staff Management</h1>
      <p class="mt-1 text-sm text-muted-foreground">
        Manage hospital workforce, schedules, and department assignments.
      </p>
    </div>
    <div class="flex items-center gap-3">
      <div class="relative">
        <UiBaseInput
          v-model="search"
          class="w-64 pl-9"
          placeholder="Search staff..."
          value=""
          leading-icon="lucide:search"
          leadingIconClass="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
        >
        </UiBaseInput>
      </div>
      <button
        type="button"
        @click="openInviteModal"
        class="inline-flex items-center gap-2 rounded-md bg-primary px-4 py-3 text-sm font-medium text-primary-foreground transition hover:bg-primary/90 focus-visible:ring-2 focus-visible:ring-ring focus-visible:outline-none disabled:opacity-50"
      >
        <Icon name="fluent-mdl2:chat-invite-friend" class="mr-2 h-4 w-4" />
        invite Staff Member
      </button>
    </div>
  </div>
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
    <UiMetricCard
      title="Total Staff"
      value="186"
      icon="lucide:users"
      colorTheme="bg-rose-50 text-rose-500"
    />

    <UiMetricCard
      title="On Duty"
      value="64"
      icon="lucide:user-check"
      colorTheme="bg-emerald-50 text-emerald-500"
    />

    <UiMetricCard
      title="On Leave"
      value="8"
      icon="lucide:calendar"
      colorTheme="bg-amber-50 text-amber-500"
    />

    <UiMetricCard
      title="Departments"
      value="12"
      icon="lucide:shield"
      colorTheme="bg-violet-50 text-violet-500"
    />
  </div>
  <div class="p-6">
    <div class="flex gap-2 mb-8">
      <button
        v-for="tab in tabs"
        :key="tab.name"
        @click="activeTab = tab.name"
        :class="[
          'flex items-center gap-2 px-6 py-2.5 rounded-xl font-medium transition-all duration-200',
          activeTab === tab.name
            ? 'bg-white text-gray-900 shadow-sm'
            : 'text-gray-500 hover:text-gray-900',
        ]"
      >
        {{ tab.name }}
        <span
          :class="[
            'text-xs px-2 py-0.5 rounded-full',
            activeTab === tab.name ? 'bg-gray-100' : 'bg-transparent',
          ]"
        >
          {{ tab.count }}
        </span>
      </button>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <UiStaffCard
        v-for="person in filteredStaff"
        :key="person.id"
        :staff="person"
      />
    </div>

    <div
      v-if="filteredStaff.length === 0"
      class="text-center py-20 text-gray-500"
    >
      No staff found in this category.
    </div>
  </div>
</template>
