<script setup>
definePageMeta({
  layout: "dashboard",
});
import { Heart, Activity, Thermometer, Droplets } from "lucide-vue-next";

const patientData = [
  {
    initials: "MC",
    name: "Margaret Chen",
    room: "201",
    status: "Stable",
    vitals: [
      {
        label: "Heart Rate",
        value: 72,
        unit: "bpm",
        trend: "down",
        icon: Heart,
        iconClass: "text-red-500",
      },
      {
        label: "Blood Pressure",
        value: "120/80",
        unit: "mmHg",
        trend: "down",
        icon: Activity,
        iconClass: "text-blue-500",
      },
      {
        label: "Temperature",
        value: 98.6,
        unit: "°F",
        trend: "down",
        icon: Thermometer,
        iconClass: "text-orange-400",
      },
      {
        label: "SpO2",
        value: 98,
        unit: "%",
        trend: "up",
        icon: Droplets,
        iconClass: "text-cyan-500",
      },
    ],
  },
  {
    initials: "RK",
    name: "Robert Kim",
    room: "412",
    status: "Critical",
    vitals: [
      {
        label: "Heart Rate",
        value: 112,
        unit: "bpm",
        trend: "up",
        icon: Heart,
        iconClass: "text-red-500",
      },
      {
        label: "Blood Pressure",
        value: "120/80",
        unit: "mmHg",
        trend: "down",
        icon: Activity,
        iconClass: "text-blue-500",
      },
      {
        label: "Temperature",
        value: 98.6,
        unit: "°F",
        trend: "down",
        icon: Thermometer,
        iconClass: "text-orange-400",
      },
      {
        label: "SpO2",
        value: 98,
        unit: "%",
        trend: "up",
        icon: Droplets,
        iconClass: "text-cyan-500",
      },
    ],
  },
  {
    initials: "RK",
    name: "Robert Kim",
    room: "412",
    status: "Critical",
    vitals: [
      {
        label: "Heart Rate",
        value: 112,
        unit: "bpm",
        trend: "up",
        icon: Heart,
        iconClass: "text-red-500",
      },
      {
        label: "Blood Pressure",
        value: "120/80",
        unit: "mmHg",
        trend: "down",
        icon: Activity,
        iconClass: "text-blue-500",
      },
      {
        label: "Temperature",
        value: 98.6,
        unit: "°F",
        trend: "down",
        icon: Thermometer,
        iconClass: "text-orange-400",
      },
      {
        label: "SpO2",
        value: 98,
        unit: "%",
        trend: "up",
        icon: Droplets,
        iconClass: "text-cyan-500",
      },
    ],
  },
  {
    initials: "YD",
    name: "Robert Kim",
    room: "412",
    status: "Critical",
    vitals: [
      {
        label: "Heart Rate",
        value: 112,
        unit: "bpm",
        trend: "up",
        icon: Heart,
        iconClass: "text-red-500",
      },
      {
        label: "Blood Pressure",
        value: "120/80",
        unit: "mmHg",
        trend: "down",
        icon: Activity,
        iconClass: "text-blue-500",
      },
      {
        label: "Temperature",
        value: 98.6,
        unit: "°F",
        trend: "down",
        icon: Thermometer,
        iconClass: "text-orange-400",
      },
      {
        label: "SpO2",
        value: 98,
        unit: "%",
        trend: "up",
        icon: Droplets,
        iconClass: "text-cyan-500",
      },
    ],
  },
];
</script>

<template>
  <div
    class="mb-6 flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between"
  >
    <div>
      <h1 class="text-2xl font-bold tracking-tight">Vitals Monitor</h1>
      <p class="mt-1 text-sm text-muted-foreground">
        Real-time patient monitoring
      </p>
    </div>
    <div
      class="inline-flex items-center rounded-full border font-medium transition-colors focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 border-transparent bg-gray-400 text-secondary-foreground w-fit gap-1.5 px-3 py-1 text-xs"
    >
      <Icon name="lucide:activity" class="h-3.5 w-3.5" />

      12 Active Monitors
    </div>
  </div>
  <div class="grid grid-cols-1 gap-4 md:grid-cols-2 xl:grid-cols-3">
    <UiPatientVitals
      v-for="patient in patientData"
      :key="patient.name"
      :patient="patient"
    />
  </div>
</template>
