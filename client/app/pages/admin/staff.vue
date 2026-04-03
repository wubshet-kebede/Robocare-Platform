<template>
  <div class="space-y-10">
    <!-- Page Header -->

    <div>
      <UiStatCard />
    </div>
  </div>
  <div
    data-slot="card"
    class="rounded-xl border bg-card text-card-foreground shadow-sm transition-shadow duration-200 mb-6"
  >
    <div data-slot="card-header" class="flex flex-col space-y-1.5 p-6 pb-2">
      <div class="tracking-tight text-base font-semibold">
        24h Patient Vitals Monitor
      </div>
      <p class="text-xs text-muted-foreground">
        Hourly average heart rate, blood pressure, and SpO2 across all monitored
        patients
      </p>
    </div>
    <div data-slot="card-content" class="p-6 pt-4">
      <ChartVitalsChart :series="series" :categories="hours" />
    </div>
  </div>
  <div class="mb-6 grid grid-cols-1 gap-4 md:grid-cols-3">
    <div
      class="rounded-xl border bg-card text-card-foreground shadow-sm transition-shadow duration-200"
    >
      <div data-slot="card-header" class="flex flex-col space-y-1.5 p-6 pb-2">
        <div class="tracking-tight text-base font-semibold">Bed Occupancy</div>
        <p class="text-xs text-muted-foreground">Occupancy rate by ward</p>
      </div>
      <div>
        <ChartBedOccupancyChart />
      </div>
    </div>
  </div>
</template>

<script setup>
definePageMeta({
  layout: "dashboard",
});
import VitalsChart from "~/components/chart/VitalsChart.vue";

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
