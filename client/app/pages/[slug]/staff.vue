<script setup>
definePageMeta({
  layout: "dashboard",
});
const staffList = ref([
  {
    id: 1,
    name: "Dr. Sarah Mitchell",
    role: "Cardiologist",
    department: "Cardiology",
    phone: "(555) 201-3344",
    email: "s.mitchell@hospital.org",
    status: "On Duty",
    rating: 4.9,
    exp: 14,
    shift: "Morning",
  },
]);
</script>
<template>
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
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mt-5">
    <UiStaffCard v-for="person in staffList" :key="person.id" :staff="person" />
  </div>
</template>
