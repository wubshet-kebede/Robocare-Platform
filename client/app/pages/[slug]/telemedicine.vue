<script setup>
definePageMeta({
  layout: "dashboard",
});
</script>
<template>
  <div
    class="mb-6 flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between"
  >
    <div>
      <h1 class="text-2xl font-bold tracking-tight">Telemedicine</h1>
      <p class="mt-1 text-sm text-muted-foreground">
        Manage virtual consultations, video sessions, and scheduling.
      </p>
    </div>
    <UiButton label="Start New Session" />
  </div>
</template>
