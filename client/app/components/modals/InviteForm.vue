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

        <UiListSelect name="role" :items="roleOptions" v-model="userData.role">
          <template #label>
            <div class="flex items-center mb-1">
              <Icon
                name="mdi:gender-male-female"
                class="text-xl text-primary dark:text-white mr-2"
              />
              <span class="dark:text-white">Role</span>
            </div>
          </template>
        </UiListSelect>

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
