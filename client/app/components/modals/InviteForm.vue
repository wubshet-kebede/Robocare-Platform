<script setup>
const props = defineProps({
  modelValue: Boolean,
});

const emit = defineEmits(["update:modelValue", "submit"]);

const form = reactive({
  email: "",
  role: "",
});

const isOpen = computed({
  get: () => props.modelValue,
  set: (val) => emit("update:modelValue", val),
});

const submit = () => {
  emit("submit", form);
};
</script>

<template>
  <ModalsModal v-model="isOpen" title="Invite Staff" wrapperClass="max-w-lg">
    <template #content>
      <div class="p-6 space-y-4">
        <UiBaseInput v-model="form.email" name="email" rules="required|email">
          <template #label>Email *</template>
        </UiBaseInput>

        <select v-model="form.role" class="w-full border p-2 rounded">
          <option disabled value="">Select role</option>
          <option value="doctor">Doctor</option>
          <option value="nurse">Nurse</option>
          <option value="admin">Admin</option>
        </select>

        <div class="flex justify-end gap-3 mt-6">
          <button @click="isOpen = false">Cancel</button>

          <button
            @click="submit"
            class="bg-primary text-white px-4 py-2 rounded"
          >
            Send Invite
          </button>
        </div>
      </div>
    </template>
  </ModalsModal>
</template>
