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
    name: "SpO₂",
    type: "line",
    data: [97, 97, 98, 97, 98, 99, 98, 97],
  },
];
</script>

<template>
  <div class="space-y-6 lg:space-y-8">
    <!-- Stats Overview -->
    <UiStatCard />

    <!-- Vitals Chart -->
    <div
      class="bg-white rounded-2xl border border-slate-200 shadow-sm hover:shadow-md transition-all overflow-hidden"
    >
      <div class="p-4 sm:p-6">
        <h2 class="text-lg sm:text-xl font-semibold text-slate-900">
          24h Patient Vitals Monitor
        </h2>

        <p class="mt-2 text-sm text-slate-500 max-w-2xl">
          Hourly average heart rate, blood pressure, and SpO₂ readings across
          monitored patients.
        </p>
      </div>

      <div class="px-2 sm:px-4 lg:px-6 pb-4 sm:pb-6">
        <ChartVitalsChart :series="series" :categories="hours" />
      </div>
    </div>

    <!-- Analytics Section -->
    <div class="grid grid-cols-1 xl:grid-cols-3 gap-6">
      <!-- Bed Occupancy -->
      <div
        class="xl:col-span-1 rounded-2xl border border-slate-200 bg-white shadow-sm hover:shadow-md transition-all"
      >
        <div class="p-5 sm:p-6">
          <h3 class="text-base font-semibold text-slate-900">Bed Occupancy</h3>

          <p class="mt-1 text-sm text-slate-500">Occupancy rate by ward</p>
        </div>

        <div class="px-2 pb-4 sm:pb-6">
          <ChartBedOccupancyChart />
        </div>
      </div>

      <!-- Future Cards -->
      <div
        class="xl:col-span-2 rounded-2xl border border-dashed border-slate-300 bg-slate-50 flex items-center justify-center min-h-[320px]"
      >
        <div class="text-center px-6">
          <div
            class="inline-flex items-center justify-center w-12 h-12 rounded-xl bg-white border border-slate-200 mb-4"
          >
            <Icon
              name="lucide:layout-dashboard"
              class="w-6 h-6 text-slate-500"
            />
          </div>

          <h3 class="text-lg font-semibold text-slate-800">Analytics Module</h3>

          <p class="mt-2 text-sm text-slate-500 max-w-sm">
            Additional patient analytics, telemetry insights, and robotic
            performance metrics can be displayed here.
          </p>
        </div>
      </div>
    </div>
  </div>
</template>
