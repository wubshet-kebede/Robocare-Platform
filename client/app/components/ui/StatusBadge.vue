<script setup>
const props = defineProps({
  status: { type: String, required: true }, // online, critical, stable, pending
  pulse: { type: Boolean, default: false },
});

const statusConfig = {
  online: {
    classes: "bg-emerald-500/10 border-emerald-500/20 text-emerald-400",
    dot: "bg-emerald-500",
  },
  critical: {
    classes: "bg-red-500/10 border-red-500/20 text-red-500",
    dot: "bg-red-500",
  },
  stable: {
    classes: "bg-cyan-500/10 border-cyan-500/20 text-cyan-400",
    dot: "bg-cyan-500",
  },
  pending: {
    classes: "bg-amber-500/10 border-amber-500/20 text-amber-400",
    dot: "bg-amber-500",
  },
};
</script>

<template>
  <div
    :class="[
      'font-black uppercase tracking-widest',
      statusConfig[status]?.classes ||
        'bg-slate-800 border-white/5 text-slate-400',
    ]"
  >
    <span v-if="pulse" class="relative flex h-1.5 w-1.5">
      <span
        :class="[
          'animate-ping absolute inline-flex h-full w-full rounded-full opacity-75',
          statusConfig[status]?.dot,
        ]"
      ></span>
      <span
        :class="[
          'relative inline-flex rounded-full h-1.5 w-1.5',
          statusConfig[status]?.dot,
        ]"
      ></span>
    </span>
    <slot>{{ status }}</slot>
  </div>
</template>
