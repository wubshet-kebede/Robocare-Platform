<script setup>
import { useField } from "vee-validate";
// import Visible from "@/plugins/visible";
const props = defineProps({
  modelValue: {
    type: [String, Number, Boolean],
    default: undefined,
  },
  name: {
    type: String,
    default: undefined,
    required: true,
  },
  id: {
    type: String,
    default: undefined,
    required: false,
  },
  type: {
    type: String,
    default: "text",
    required: false,
  },
  placeholder: {
    type: String,
    default: undefined,
    required: false,
  },
  label: {
    type: String,
    default: undefined,
    required: false,
  },
  labelClass: {
    type: String,
    default: () => "",
    required: false,
  },
  hideDetail: {
    type: Boolean,
    default: false,
    required: false,
  },
  trailingIcon: {
    type: String,
    default: "",
    required: false,
  },
  leadingIcon: {
    type: [String],
    default: undefined,
    required: false,
  },
  leadingIconClass: {
    type: [String],
  },
  min: String,
  max: String,
  rules: {
    type: [String, Object],
    default: "",
    required: false,
  },
  disabled: {
    type: Boolean,
    default: false,
  },
  class: String,
  iconLeadingClass: String,
  placeholderStyle: String,
  iconBackground: Boolean,
  mainDiv: String,
});
const emit = defineEmits(["update:modelValue"]);

const {
  errorMessage,
  value: inputValue,
  meta,
} = useField(props.name, props.rules, {
  initialValue: props.modelValue,
});

const type = ref(props.type);

const set = (event) => {
  emit("update:modelValue", event.target.value);
};

const togglePassword = (enabled) => {
  if (type.value === "password") {
    type.value = "text";
    return;
  }
  if (type.value === "text") {
    type.value = "password";
  }
};

// watch(
//   () => props.modelValue,
//   (newVal) => {
//     inputValue.value = props.type == "number" ? Number(newVal) : newVal;
//   }
// );

watch(
  () => props.modelValue,
  (newVal) => {
    inputValue.value = newVal;
  },
);

const clear = () => {
  inputValue.value = "";
  emit("update:modelValue", undefined);
};
</script>
<template>
  <div :class="mainDiv" class="space-y-1">
    <!-- Label -->
    <div class="flex items-center justify-between">
      <slot name="label"></slot>
    </div>

    <!-- Input Wrapper -->
    <div
      class="relative rounded-xl border bg-white shadow-sm transition-all duration-200 group"
      :class="[
        errorMessage
          ? 'border-red-500 focus-within:border-red-500 focus-within:ring-2 focus-within:ring-red-100'
          : 'border-gray-300 hover:border-gray-400 focus-within:border-primary focus-within:ring-2 focus-within:ring-primary/20',
        props.disabled ? 'bg-gray-100 cursor-not-allowed' : 'bg-white',
      ]"
    >
      <!-- Leading Icon -->
      <div
        v-if="leadingIcon"
        class="absolute inset-y-0 left-0 flex items-center pl-3"
      >
        <Icon
          :name="leadingIcon"
          class="text-gray-400 group-focus-within:text-primary transition-colors duration-200"
          size="20"
        />
      </div>

      <!-- Input -->
      <input
        v-model="inputValue"
        @input="set($event)"
        @change="set($event)"
        :type="type"
        :name="props.name"
        :id="id"
        :placeholder="props.placeholder"
        :disabled="props.disabled"
        :min="min"
        :max="max"
        step="any"
        class="block w-full rounded-xl bg-transparent text-sm text-gray-700 placeholder-gray-400 focus:outline-none transition-all duration-200"
        :class="[
          leadingIcon ? 'pl-10' : 'pl-4',
          props.trailingIcon || props.type === 'password' ? 'pr-10' : 'pr-4',
          'py-3',
        ]"
      />

      <!-- Trailing Icon (Password Toggle) -->
      <div
        v-if="props.type === 'password'"
        class="absolute inset-y-0 right-0 flex items-center pr-3 cursor-pointer"
        @click="togglePassword()"
      >
        <Icon
          :name="props.trailingIcon"
          size="20"
          class="transition-colors duration-200"
          :class="type === 'password' ? 'text-gray-400' : 'text-primary'"
        />
      </div>

      <!-- Trailing Icon (Normal) -->
      <div
        v-else-if="props.trailingIcon"
        class="absolute inset-y-0 right-0 flex items-center pr-3"
      >
        <Icon :name="props.trailingIcon" size="18" class="text-gray-400" />
      </div>

      <!-- Custom Trailing Slot -->
      <div class="absolute inset-y-0 right-0 flex items-center pr-3">
        <slot name="trailing" />
      </div>
    </div>

    <!-- Error Message -->
    <p
      v-if="!hideDetail || errorMessage"
      class="text-sm text-red-500 transition-all duration-200"
      :class="
        errorMessage ? 'opacity-100 translate-y-0' : 'opacity-0 -translate-y-1'
      "
    >
      {{ errorMessage }}
    </p>
  </div>
</template>
