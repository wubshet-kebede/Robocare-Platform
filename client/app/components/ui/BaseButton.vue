<script setup>
const props = defineProps({
  variant: { type: String, default: "primary" }, // primary, secondary, outline, danger
  size: { type: String, default: "md" }, // sm, md, lg
  loading: { type: Boolean, default: false },
  disabled: { type: Boolean, default: false },
  rounded: { type: Boolean, default: false },
});

const variantClasses = {
  primary:
    "bg-white text-slate-950 hover:shadow-[0_0_20px_rgba(255,255,255,0.2)]",
  secondary:
    "bg-slate-900 border border-white/10 text-white hover:bg-slate-800",
  gradient:
    "bg-gradient-to-r from-cyan-500 to-emerald-500 text-slate-950 hover:shadow-[0_10px_30px_rgba(6,182,212,0.3)]",
  danger:
    "bg-red-500/10 border border-red-500/20 text-red-500 hover:bg-red-500 hover:text-white",
};

const sizeClasses = {
  sm: "px-4 py-2 text-[10px]",
  md: "px-8 py-4 text-xs",
  lg: "px-10 py-5 text-sm",
};

defineEmits(["click"]);
</script>

<template>
  <button
    :disabled="loading || disabled"
    :class="[
      'relative inline-flex items-center justify-center font-black uppercase tracking-widest transition-all duration-300 active:scale-95 disabled:opacity-50 disabled:cursor-not-allowed overflow-hidden',
      variantClasses[variant],
      sizeClasses[size],
      rounded ? 'rounded-full' : 'rounded-2xl',
    ]"
    @click="$emit('click')"
  >
    <!-- Loading Spinner -->
    <div
      v-if="loading"
      class="absolute inset-0 flex items-center justify-center bg-inherit"
    >
      <Icon name="svg-spinners:ring-resize" class="text-xl" />
    </div>

    <!-- Button Content -->
    <span :class="{ 'opacity-0': loading }" class="flex items-center gap-2">
      <slot name="icon-left" />
      <slot />
      <slot name="icon-right" />
    </span>
  </button>
</template>
