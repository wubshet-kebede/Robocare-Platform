<script setup>
import ApexChart from "vue3-apexcharts";

const props = defineProps({
  series: { type: Array, required: true },
  categories: { type: Array, required: true },
});

const chartOptions = {
  chart: {
    type: "line",
    toolbar: { show: false },
    zoom: { enabled: false },
    background: "transparent",
  },
  stroke: { curve: "smooth", width: 2 },
  xaxis: {
    categories: props.categories,
    title: { text: "Time" },
    labels: { style: { colors: "#ccc" } },
  },
  yaxis: [
    {
      title: { text: "bpm / mmHg" },
      min: 50,
      max: 150,
    },
    {
      opposite: true,
      title: { text: "SpO2 %" },
      min: 90,
      max: 100,
    },
  ],
  grid: { borderColor: "#444" },
  tooltip: { enabled: true },
  colors: ["#00BFFF", "#00FF00", "#FF6347"], // match your chart colors
};
</script>
<template>
  <div>
    <ApexChart
      type="line"
      :options="chartOptions"
      :series="series"
      height="350"
    />
  </div>
</template>
