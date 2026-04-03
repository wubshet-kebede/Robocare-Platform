<template>
  <div>
    <apexchart
      type="radialBar"
      height="350"
      :options="chartOptions"
      :series="series"
    />
  </div>
</template>

<script setup>
import ApexCharts from "vue3-apexcharts";

const series = [85, 72, 91]; // ICU, General, Emergency

const chartOptions = {
  chart: {
    type: "radialBar",
  },
  plotOptions: {
    radialBar: {
      offsetY: 0,
      startAngle: 0,
      endAngle: 360,
      hollow: {
        margin: 5,
        size: "50%",
        background: "transparent",
        position: "front",
      },
      dataLabels: {
        name: {
          show: false,
        },
        value: {
          show: true,
          fontSize: "22px",
          formatter: function (val) {
            // Show aggregate in center
            const aggregate = Math.round(
              series.reduce((a, b) => a + b, 0) / series.length,
            );
            return aggregate + "%";
          },
        },
      },
    },
  },
  colors: ["#FF6B6B", "#4BC0C0", "#28A745"], // ICU, General, Emergency
  labels: ["ICU", "General", "Emergency"],
  legend: {
    show: true,
    position: "bottom",
  },
};
</script>
