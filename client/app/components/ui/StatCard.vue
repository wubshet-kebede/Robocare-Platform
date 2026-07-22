<script setup>
const cards = [
  { title: "Patients Today", value: 48, change: "+12%", color: "--chart-1" },
  { title: "Appointments", value: 24, change: "-3%", color: "--chart-2" },
  { title: "Bed Occupancy", value: "78%", change: "+5%", color: "--chart-3" },
  { title: "Revenue", value: "$42.5K", change: "+8%", color: "--chart-4" },
];
</script>

<template>
  <div class="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-4 gap-4">
    <div
      v-for="(card, index) in cards"
      :key="index"
      class="bg-white rounded-xl border border-slate-200 p-4 shadow-sm hover:shadow-md transition-all"
      :style="{
        borderLeftWidth: '4px',
        borderLeftColor: `var(${card.color})`,
      }"
    >
      <p class="text-sm text-slate-500 mb-2">
        {{ card.title }}
      </p>

      <div class="flex items-center justify-between">
        <span class="text-xl lg:text-2xl font-bold text-slate-900">
          {{ card.value }}
        </span>

        <span
          class="text-xs font-medium px-2 py-1 rounded-md"
          :class="
            card.change.startsWith('+')
              ? 'bg-emerald-100 text-emerald-700'
              : 'bg-red-100 text-red-700'
          "
        >
          {{ card.change }}
        </span>
      </div>
    </div>
  </div>
</template>
