<script setup>
import {
  Heart,
  Activity,
  Thermometer,
  Droplets,
  TrendingDown,
  TrendingUp,
} from "lucide-vue-next";

const props = defineProps({
  patient: {
    type: Object,
    required: true,
  },
});
const statusConfig = {
  Stable: {
    border: "border-l-green-500",
    bg: "bg-green-500",
    text: "text-white",
  },
  Warning: {
    border: "border-l-amber-500",
    bg: "bg-amber-500",
    text: "text-white",
  },
  Critical: {
    border: "border-l-red-600",
    bg: "bg-red-600",
    text: "text-white",
  },
};

// Map the icon components to names so we can use them dynamically
const iconMap = {
  heart: Heart,
  activity: Activity,
  thermometer: Thermometer,
  droplets: Droplets,
};
const chartOptions = {
  chart: {
    type: "line",
    animations: {
      enabled: true,
      easing: "linear",
      dynamicAnimation: { speed: 1000 },
    },

    toolbar: { show: false },
    sparkline: { enabled: true },

    tooltip: {
      enabled: true,
      theme: "light",
      x: { show: false },
      y: {
        formatter: (val) => `${val} bpm`,
      },
      marker: { show: false },
    },
  },

  markers: { size: 0, hover: { size: 4 } },

  width: "100%",
  height: 60,
  stroke: { curve: "smooth", width: 2 },
  colors: ["#ef4444"],
  //   fill: {
  //     type: "gradient",
  //     gradient: {
  //       shadeIntensity: 1,
  //       opacityFrom: 0.7,
  //       opacityTo: 0.1,
  //       stops: [0, 100],
  //     },
  //   },
  //   stroke: { curve: "smooth", width: 2 },
};

const series = ref([{ data: Array(20).fill(30) }]);

// Logic to simulate real-time sensor data
let interval = null;
onMounted(() => {
  interval = setInterval(() => {
    const newData = [
      ...series.value[0].data.slice(1),
      Math.floor(Math.random() * 50) + 30,
    ];
    series.value = [{ data: newData }];
  }, 1000);
});

onUnmounted(() => clearInterval(interval));
</script>

<template>
  <div
    class="rounded-xl border bg-white shadow-sm transition-all duration-300 hover:shadow-xl hover:-translate-y-1 hover:border-slate-300"
    :class="statusConfig[patient.status].border"
  >
    <div class="flex flex-col space-y-1.5 p-4 pb-2">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div
            class="flex h-9 w-9 shrink-0 items-center justify-center rounded-full bg-slate-100 text-xs font-bold text-slate-600"
          >
            {{ patient.initials }}
          </div>
          <div>
            <div class="text-sm font-semibold tracking-tight text-slate-900">
              {{ patient.name }}
            </div>
            <p class="text-[11px] text-slate-500">Room {{ patient.room }}</p>
          </div>
        </div>
        <div
          :class="[
            statusConfig[patient.status].bg,
            statusConfig[patient.status].text,
          ]"
          class="px-2.5 py-0.5 rounded-full text-[10px] font-bold uppercase tracking-wider"
        >
          {{ patient.status }}
        </div>
      </div>
    </div>

    <div class="px-4 pb-4 pt-0 space-y-1">
      <div class="divide-y divide-slate-100">
        <div
          v-for="vital in patient.vitals"
          :key="vital.label"
          class="flex items-center justify-between py-2"
        >
          <div class="flex items-center gap-2">
            <component
              :is="iconMap[vital.iconName]"
              :class="vital.iconClass"
              class="h-3.5 w-3.5"
            />
            <span class="text-xs text-slate-500">{{ vital.label }}</span>
          </div>
          <div class="flex items-center gap-1.5">
            <span class="text-sm font-bold text-slate-800">
              {{ vital.value
              }}<span class="ml-0.5 text-[10px] font-normal text-slate-400">{{
                vital.unit
              }}</span>
            </span>
            <TrendingDown
              v-if="vital.trend === 'down'"
              class="h-3 w-3 text-green-500/70"
            />
            <TrendingUp v-else class="h-3 w-3 text-red-500/70" />
          </div>
        </div>
      </div>

      <div class="pt-2 h-[60px] w-full overflow-hidden">
        <ClientOnly>
          <apexchart
            v-if="patient"
            :key="patient.name"
            type="line"
            height="60"
            :options="chartOptions"
            :series="series"
            tooltip.shared:
            true
            tooltip.intersect:
            false
          />

          <template #fallback>
            <div
              class="h-[60px] w-full bg-slate-100 animate-pulse rounded"
            ></div>
          </template>
        </ClientOnly>
      </div>
    </div>
  </div>
</template>
