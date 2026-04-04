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
  <div :class="mainDiv">
    <div class="flex gap-x-1">
      <slot name="label"></slot>
      <!-- <span
        class="text-red-600"
        v-if="
          typeof rules == 'object'
            ? Object.hasOwn(rules, 'required')
            : rules.includes('required')
        "
        >*</span
      > -->
    </div>
    <div
      class="relative rounded-md overflow-clip font-body shadow-sm group"
      :class="props.class"
    >
      <slot name="leading" />

      <Icon
        :name="leadingIcon"
        class="absolute px-3"
        size="45"
        v-if="leadingIcon"
        :class="leadingIconClass"
      />
      <input
        v-model="inputValue"
        @input="set($event)"
        @change="set($event)"
        :min="min"
        :max="max"
        :type="type"
        :name="props.name"
        step="any"
        :id="id"
        :class="[
          errorMessage
            ? 'border-red-500 hover:border-red-500 focus:border-red-500 focus:ring-red-500'
            : 'border-[1px]  border-gray-300 group-hover:border-primary focus:border-primary focus:ring-primary disabled:bg-gray-100 disabled:cursor-not-allowed ',
          props.iconLeadingClass ? props.iconLeadingClass : '',
          props.leadingIcon ? 'pl-10' : '',
          props.trailingIcon ? 'pr-0' : '',
          props.class ? props.class : '',
          props.placeholderStyle
            ? props.placeholderStyle
            : ' placeholder-gray-400 text-sm',
          props.trailingIcon ? 'pr-7' : '',
          props.type == 'password' ? 'pr-7' : '',
        ]"
        class="block w-full rounded-md font-body text-base text-gray-500 placeholder-gray-500/50 dark:placeholder-secondary-lite/50 focus:outline-none dark:bg-secondary-hover dark:text-white"
        :placeholder="props.placeholder"
        aria-invalid="true"
        aria-describedby="email-error"
        :disabled="props.disabled"
      />
      <div
        v-if="props.type == 'password'"
        class="absolute inset-y-0 right-0 flex items-center pr-3 hover:cursor-pointer"
        @click="togglePassword()"
      >
        <Icon
          class="h-5 w-5"
          :class="[
            type === 'password'
              ? 'text-gray-400 dark:text-secondary-lite'
              : 'text-primary',
          ]"
          :name="props.trailingIcon"
          size="25"
        ></Icon>
      </div>
      <div
        v-else-if="props.trailingIcon"
        class="absolute inset-y-0 right-0 flex items-center pr-3 hover:cursor-pointer"
      >
        <Icon
          class="text-gray-500 dark:text-secondary-lite"
          :name="props.trailingIcon"
          size="20"
        ></Icon>
      </div>
      <div
        class="absolute inset-y-0 right-0 flex items-center pr-3 hover:cursor-pointer"
      >
        <slot name="trailing" />
      </div>
    </div>
    <p
      v-if="!hideDetail || errorMessage"
      class="col-span-5 mt-1 font-body text-sm text-red-600 duration-300 transition-all"
      :class="errorMessage ? 'h-5' : 'h-0'"
      id="email-error"
      :visible="errorMessage"
    >
      {{ errorMessage }} &nbsp;
    </p>
  </div>
</template>
