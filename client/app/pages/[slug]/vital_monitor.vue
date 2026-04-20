<script setup>
definePageMeta({
  layout: "dashboard",
});
import { Heart, Activity, Thermometer, Droplets } from "lucide-vue-next";
const user = useAuthUser();
const hospitalID = user.value?.hospital.id;
const { vitalsMap, status } = useVitalsSocket(hospitalID);
console.log("the value of the vitla map is ", vitalsMap.value);
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
const livePatients = computed(() => {
  return patientData.map((patient) => {
    const live = vitalsMap.value[patient.id];
    console.log("the value of the live is ", live);

    if (!live) return patient;

    return {
      ...patient,
      vitals: patient.vitals.map((v) => {
        if (v.label === "Heart Rate") {
          return {
            ...v,
            value: live.heart_rate ?? v.value,
            trend: live.heart_rate > 100 ? "up" : "down",
          };
        }

        if (v.label === "SpO2") {
          return {
            ...v,
            value: live.spo2 ?? v.value,
            trend: live.spo2 < 95 ? "down" : "up",
          };
        }

        if (v.label === "Temperature") {
          return {
            ...v,
            value: live.temperature ?? v.value,
          };
        }

        return v;
      }),
    };
  });
});
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
      <p>WS Status: {{ status }}</p>
    </div>
    <div
      class="inline-flex items-center rounded-full border font-medium transition-colors focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 border-transparent bg-slate-50 dark:text-secondary-foreground w-fit gap-1.5 px-3 py-1 text-xs"
    >
      <Icon name="lucide:activity" class="h-3.5 w-3.5" />

      12 Active Monitors
    </div>
  </div>
  <div class="grid grid-cols-1 gap-4 md:grid-cols-2 xl:grid-cols-3">
    <UiPatientVitals
      v-for="patient in livePatients"
      :key="patient.name"
      :patient="patient"
    />
  </div>
</template>
