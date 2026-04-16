<script setup>
const props = defineProps({
  title: {
    type: String,
    required: true,
  },
  value: {
    type: [String, Number],
    required: true,
  },
  icon: {
    type: String,
    required: true,
  },
  bgColor: {
    type: String,
    default: "bg-chart-1/10",
  },
  textColor: {
    type: String,
    default: "text-chart-1",
  },
  colorTheme: {
    type: String,
    default: "bg-gray-100 text-gray-600",
  },
});

console.log("props", props.value);
</script>
<template>
  <div class="rounded-3xl border border-gray-200 bg-white p-6 shadow-md">
    <div class="flex items-center justify-between">
      <div class="space-y-1">
        <p class="text-sm font-medium text-gray-500">
          {{ props.title }}
        </p>
        <p class="text-3xl font-bold tracking-tight text-gray-900">
          {{ props.value }}
        </p>
      </div>

      <div
        :class="[
          'flex h-12 w-12 items-center justify-center rounded-2xl',
          colorTheme,
        ]"
      >
        <Icon v-if="icon" :name="props.icon" class="h-6 w-6" />
      </div>
    </div>
  </div>
</template>
