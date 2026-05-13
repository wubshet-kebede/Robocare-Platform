<script setup>
definePageMeta({
  layout: "dashboard",
});
const metrics = [
  {
    title: "Live Consultations",
    value: "3",
    icon: "lucide:video",
    colorTheme: "bg-rose-50 text-rose-500",
  },

  {
    title: "Patients Waiting",
    value: "5",
    icon: "lucide:users",
    colorTheme: "bg-amber-50 text-amber-500",
  },

  {
    title: "Robots Available",
    value: "2",
    icon: "lucide:bot",
    colorTheme: "bg-emerald-50 text-emerald-500",
  },

  {
    title: "Emergency Requests",
    value: "1",
    icon: "lucide:siren",
    colorTheme: "bg-violet-50 text-violet-500",
  },
];
const isSessionModalOpen = ref(false);

const openSessionModal = () => {
  isSessionModalOpen.value = true;
};
</script>
<template>
  <ModalsTelepresenceSession v-model="isSessionModalOpen" />
  <div
    class="mb-6 flex flex-col gap-4 sm:flex-row items-center justify-between"
  >
    <div>
      <h1 class="text-2xl font-bold tracking-tight">Robot Telepresence</h1>

      <p class="mt-1 text-sm text-muted-foreground">
        Manage remote consultations, live robot sessions, and patient
        interactions.
      </p>
    </div>

    <div class="flex items-center gap-3">
      <div class="relative">
        <UiBaseInput
          v-model="search"
          class="w-64 pl-9"
          placeholder="Search sessions..."
          value=""
          leading-icon="lucide:search"
          leadingIconClass="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
        >
        </UiBaseInput>
      </div>

      <button
        type="button"
        @click="openSessionModal"
        class="inline-flex items-center gap-2 rounded-md bg-primary px-4 py-3 text-sm font-medium text-primary-foreground transition hover:bg-primary/90 focus-visible:ring-2 focus-visible:ring-ring focus-visible:outline-none disabled:opacity-50"
      >
        <Icon name="lucide:video" class="mr-2 h-4 w-4" />

        Start New Session
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
</template>
