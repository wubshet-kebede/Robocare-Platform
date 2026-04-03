<template>
  <div class="space-y-10">
    <!-- Page Header -->

    <div>
      <UiStatCard />
    </div>
    <div>
      <h2 class="text-lg font-bold mb-2">24h Patient Vitals Monitor</h2>
      <p class="text-sm text-muted mb-4">
        Hourly average heart rate, blood pressure, and SpO₂ across all monitored
        patients
      </p>
      <VitalsChart :series="series" :categories="hours" />
    </div>
  </div>
</template>

<script setup>
definePageMeta({
  layout: "dashboard",
});
import VitalsChart from "~/components/VitalsChart.vue";

const hours = [
  "00:00",
  "03:00",
  "06:00",
  "09:00",
  "12:00",
  "15:00",
  "18:00",
  "21:00",
];

const series = [
  {
    name: "Heart Rate",
    type: "line",
    data: [105, 95, 102, 100, 108, 105, 110, 103],
  },
  {
    name: "Blood Pressure",
    type: "line",
    data: [60, 65, 63, 70, 72, 68, 65, 60],
  },
  {
    name: "SpO2",
    type: "line",
    data: [97, 97, 98, 97, 98, 99, 98, 97],
  },
];
</script>
