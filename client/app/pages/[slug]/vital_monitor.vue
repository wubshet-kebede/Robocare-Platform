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
];
</script>

<template>
  <div class="p-10 bg-slate-50 min-h-screen">
    <div
      class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 max-w-7xl mx-auto"
    >
      <UiPatientVitals
        v-for="patient in patientData"
        :key="patient.name"
        :patient="patient"
      />
    </div>
  </div>
</template>
