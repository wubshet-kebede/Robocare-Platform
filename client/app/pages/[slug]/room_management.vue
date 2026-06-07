<script setup>
definePageMeta({
  layout: "dashboard",
});
const metrics = [
  {
    title: "Total Rooms",
    value: "186",
    icon: "iconoir:house-rooms",
    colorTheme: "bg-rose-50 text-rose-500",
  },

  {
    title: "Available",
    value: "64",
    icon: "carbon:hospital-bed",
    colorTheme: "bg-emerald-50 text-emerald-500",
  },

  {
    title: "Occupied",
    value: "8",
    icon: "fontisto:bed-patient",
    colorTheme: "bg-amber-50 text-amber-500",
  },

  {
    title: "Departments",
    value: "12",
    icon: "fxemoji:departmentstore",
    colorTheme: "bg-violet-50 text-violet-500",
  },
];
const isModalOpen = ref(false);
const openRoomModal = () => {
  isModalOpen.value = true;
};
const isDepartmentModalOpen = ref(false);
const openDepartmentModal = () => {
  isDepartmentModalOpen.value = true;
};
const wards = ref([
  {
    name: "Emergency",
    occupied: 14,
    total: 16,
    available: 1,
    maintenance: 1,
  },
  {
    name: "ICU",
    occupied: 11,
    total: 12,
    available: 1,
    maintenance: 0,
  },
  {
    name: "Cardiology",
    occupied: 12,
    total: 14,
    available: 1,
    maintenance: 1,
  },
  {
    name: "Pediatrics",
    occupied: 10,
    total: 16,
    available: 5,
    maintenance: 1,
  },
  {
    name: "Orthopedics",
    occupied: 11,
    total: 14,
    available: 3,
    maintenance: 0,
  },
  {
    name: "General Medicine",
    occupied: 12,
    total: 16,
    available: 3,
    maintenance: 1,
  },
]);

const occupancyPercentage = (ward) =>
  Math.round((ward.occupied / ward.total) * 100);
</script>
<template>
  <ModalsRoomManagement
    v-if="isModalOpen"
    v-model="isModalOpen"
  ></ModalsRoomManagement>
  <ModalsDepartmmentManagement
    v-if="isDepartmentModalOpen"
    v-model="isDepartmentModalOpen"
  ></ModalsDepartmmentManagement>

  <div
    class="mb-6 flex flex-col gap-4 sm:flex-row items-center justify-between"
  >
    <div>
      <h1 class="text-2xl font-bold tracking-tight">Room Management</h1>
      <p class="mt-1 text-sm text-muted-foreground">
        Manage hospital rooms, schedules, and department assignments.
      </p>
    </div>
    <div class="flex items-center gap-3">
      <button
        type="button"
        @click="openDepartmentModal"
        class="inline-flex items-center gap-2 rounded-md bg-primary px-4 py-3 text-sm font-medium text-primary-foreground transition hover:bg-primary/90 focus-visible:ring-2 focus-visible:ring-ring focus-visible:outline-none disabled:opacity-50"
      >
        <Icon name="boxicons:hospital" class="mr-2 h-4 w-4" />
        Add Department
      </button>
      <div class="relative">
        <UiBaseInput
          v-model="search"
          class="w-64 pl-9"
          placeholder="Search rooms..."
          value=""
          leading-icon="lucide:search"
          leadingIconClass="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
        >
        </UiBaseInput>
      </div>
      <button
        type="button"
        @click="openRoomModal"
        class="inline-flex items-center gap-2 rounded-md bg-primary px-4 py-3 text-sm font-medium text-primary-foreground transition hover:bg-primary/90 focus-visible:ring-2 focus-visible:ring-ring focus-visible:outline-none disabled:opacity-50"
      >
        <Icon name="boxicons:hospital" class="mr-2 h-4 w-4" />
        Add Room
      </button>
    </div>
  </div>
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
    <UiMetricCard
      v-for="metric in metrics"
      :key="metric.title"
      :title="metric.title"
      :value="metric.value"
      :icon="metric.icon"
      :colorTheme="metric.colorTheme"
    />
  </div>
  <div class="mt-8">
    <h2 class="mb-4 text-xl font-bold">Ward Overview</h2>

    <div class="grid grid-cols-1 gap-6 lg:grid-cols-3">
      <div
        v-for="ward in wards"
        :key="ward.name"
        class="rounded-2xl border p-5 shadow-sm"
      >
        <!-- Header -->
        <div class="mb-4 flex items-center justify-between">
          <h3 class="font-semibold text-lg">
            {{ ward.name }}
          </h3>

          <span class="rounded-full px-3 py-1 text-xs font-semibold">
            {{ occupancyPercentage(ward) }}%
          </span>
        </div>

        <!-- Progress -->
        <div class="mb-3 h-2 rounded-full">
          <div
            class="h-2 rounded-full bg-red-500"
            :style="{
              width: `${occupancyPercentage(ward)}%`,
            }"
          />
        </div>

        <p class="mb-4 text-sm text-muted-foreground">
          {{ ward.occupied }} / {{ ward.total }} beds occupied
        </p>

        <!-- Stats -->
        <div class="mb-4 flex flex-wrap gap-4 text-sm">
          <div class="flex items-center gap-2">
            <span class="h-2 w-2 rounded-full bg-green-500"></span>
            {{ ward.available }} available
          </div>

          <div class="flex items-center gap-2">
            <span class="h-2 w-2 rounded-full bg-red-500"></span>
            {{ ward.occupied }} occupied
          </div>

          <div v-if="ward.maintenance" class="flex items-center gap-2">
            <span class="h-2 w-2 rounded-full bg-amber-500"></span>
            {{ ward.maintenance }} maint.
          </div>
        </div>

        <!-- Bed Indicators -->
        <div class="flex flex-wrap gap-2 rounded-xl border p-3">
          <template v-for="n in ward.total" :key="n">
            <div
              v-if="n <= ward.occupied"
              class="h-3 w-3 rounded-full bg-red-500"
            />
            <div
              v-else-if="n <= ward.occupied + ward.maintenance"
              class="h-3 w-3 rounded-full bg-amber-500"
            />
            <div v-else class="h-3 w-3 rounded-full bg-green-500" />
          </template>
        </div>
      </div>
    </div>

    <!-- Legend -->
    <div class="mt-4 flex gap-6 text-sm">
      <div class="flex items-center gap-2">
        <span class="h-3 w-3 rounded-full bg-green-500"></span>
        Available
      </div>

      <div class="flex items-center gap-2">
        <span class="h-3 w-3 rounded-full bg-red-500"></span>
        Occupied
      </div>

      <div class="flex items-center gap-2">
        <span class="h-3 w-3 rounded-full bg-amber-500"></span>
        Maintenance
      </div>
    </div>
  </div>
</template>
