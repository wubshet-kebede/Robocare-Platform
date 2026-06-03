<script setup>
import ApexChart from "vue3-apexcharts";

const props = defineProps({
  series: {
    type: Array,
    required: true,
  },
  categories: {
    type: Array,
    required: true,
  },
});

const chartOptions = {
  chart: {
    type: "line",
    height: 350,
    toolbar: {
      show: false,
    },
    zoom: {
      enabled: false,
    },
    background: "transparent",
  },

  stroke: {
    curve: "smooth",
    width: 2.5,
  },

  dataLabels: {
    enabled: false,
  },

  legend: {
    show: false,
  },

  grid: {
    borderColor: "#E5E7EB",
    strokeDashArray: 4,
  },

  colors: [
    "#06B6D4", // cyan
    "#10B981", // green
    "#EF4444", // red
  ],

  xaxis: {
    categories: props.categories,

    axisBorder: {
      show: false,
    },

    axisTicks: {
      show: false,
    },

    labels: {
      style: {
        colors: "#6B7280",
        fontSize: "12px",
      },
    },
  },

  yaxis: [
    {
      min: 50,
      max: 150,

      title: {
        text: "bpm / mmHg",
        style: {
          color: "#6B7280",
        },
      },

      labels: {
        style: {
          colors: "#6B7280",
        },
      },
    },

    {
      opposite: true,
      min: 90,
      max: 100,

      title: {
        text: "SpO₂ %",
        style: {
          color: "#6B7280",
        },
      },

      labels: {
        style: {
          colors: "#6B7280",
        },
      },
    },
  ],

  tooltip: {
    theme: "light",
  },
};
</script>

<template>
  <ApexChart
    type="line"
    height="350"
    :options="chartOptions"
    :series="series"
  />
</template>
