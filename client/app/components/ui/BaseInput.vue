<script setup>
defineProps({
  modelValue: [String, Number],
  label: String,
  placeholder: String,
  class: [String, Object, Array],
  type: { type: String, default: "text" },
  icon: String,
  error: String,
});
defineEmits(["update:modelValue"]);
</script>

<template>
  <div class="space-y-2 w-full">
    <label
      v-if="label"
      class="text-[10px] font-black uppercase tracking-[0.2em] text-slate-500 ml-1"
    >
      {{ label }}
    </label>
    <div class="relative group">
      <!-- Icon Support -->
      <div
        v-if="icon"
        class="absolute left-5 top-1/2 -translate-y-1/2 text-slate-500 group-focus-within:text-cyan-500 transition-colors"
      >
        <Icon :name="icon" />
      </div>

      <input
        :value="modelValue"
        :type="type"
        :placeholder="placeholder"
        :class="class"
        @input="$emit('update:modelValue', $event.target.value)"
      />
    </div>
    <p
      v-if="error"
      class="text-[10px] text-red-500 font-bold ml-1 animate-pulse italic"
    >
      * {{ error }}
    </p>
  </div>
</template>
