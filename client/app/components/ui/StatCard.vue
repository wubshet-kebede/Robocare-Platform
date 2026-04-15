<script setup>
const cards = [
  { title: "Patients Today", value: 48, change: "+12%", color: "--chart-1" },
  { title: "Appointments", value: 24, change: "-3%", color: "--chart-2" },
  { title: "Bed Occupancy", value: "78%", change: "+5%", color: "--chart-3" },
  { title: "Revenue", value: "$42.5K", change: "+8%", color: "--chart-4" },
];
</script>

<template>
  <div class="mb-6 grid grid-cols-2 gap-3 md:grid-cols-4">
    <div
      v-for="(card, index) in cards"
      :key="index"
      class="flex items-center gap-3 rounded-lg border dark:border-border dark:bg-card py-2 px-4"
      :style="{ borderLeftWidth: '4px', borderLeftColor: `var(${card.color})` }"
    >
      <div class="flex-1 min-w-0">
        <p class="text-xs dark:text-muted-foreground truncate">
          {{ card.title }}
        </p>
        <div class="flex items-center gap-2">
          <span class="text-lg font-bold">{{ card.value }}</span>
          <div
            class="inline-flex items-center rounded-md border font-medium transition-colors focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 border-transparent shadow-sm text-[10px] px-1.5 py-0"
            :class="[
              card.change.startsWith('+')
                ? 'bg-success dark:text-success-foreground'
                : 'bg-destructive dark:text-destructive-foreground',
            ]"
          >
            {{ card.change }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
