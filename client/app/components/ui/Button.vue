<script setup>
import { Video } from "lucide-vue-next";

// Simple props for reusability
defineProps({
  label: { type: String, default: "Start New Session" },
  variant: { type: String, default: "primary" }, // You can expand this for 'outline', 'ghost', etc.
});
</script>

<template>
  <button
    class="inline-flex items-center justify-center gap-2 px-4 py-2 rounded-md text-sm font-medium transition-all duration-200 bg-primary text-white hover:bg-slate-500 active:scale-95 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
  >
    <Video class="w-4 h-4" />
    <span>{{ label }}</span>
  </button>
</template>
