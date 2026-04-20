<script setup>
import { ref, computed, onMounted } from "vue";
import { useForm } from "vee-validate";

const props = defineProps({
  modelValue: Boolean,
});

const emit = defineEmits(["update:modelValue", "success"]);

const { getRoles } = useRoleService();
const { inviteStaff } = useInvitationService();

const isOpen = computed({
  get: () => props.modelValue,
  set: (val) => emit("update:modelValue", val),
});

const loadingRoles = ref(false);
const loadingSubmit = ref(false);

const roleOptions = ref([]);

const formatRole = (role) =>
  role.replace(/_/g, " ").replace(/\b\w/g, (c) => c.toUpperCase());

const { handleSubmit, values, resetForm } = useForm({
  initialValues: {
    email: "",
    role: "",
  },
});
const fetchRoles = async () => {
  try {
    loadingRoles.value = true;

    const res = await getRoles();

    roleOptions.value = res.map((role) => ({
      id: role,
      name: formatRole(role),
    }));
  } finally {
    loadingRoles.value = false;
  }
};

onMounted(fetchRoles);

const submit = handleSubmit(async (formValues) => {
  try {
    loadingSubmit.value = true;

    await inviteStaff({
      email: formValues.email,
      role: formValues.role,
    });
    resetForm();
    isOpen.value = false;

    emit("success");
  } catch (err) {
    console.error("Invite failed:", err);
  } finally {
    loadingSubmit.value = false;
  }
});
</script>

<template>
  <ModalsModal v-model="isOpen" title="Invite Staff" wrapperClass="max-w-lg">
    <template #content>
      <div class="p-8 space-y-8">
        <UiBaseInput v-model="values.email" name="email" rules="required|email">
          <template #label
            ><h1
              class="block text-md font-medium leading-6 text-gray-900 mb-1 duration-200"
            >
              Email address <span class="text-red-600">*</span>
            </h1></template
          >
        </UiBaseInput>

        <UiListSelect
          rules="required"
          name="role"
          :items="roleOptions"
          v-model="values.role"
          class="text-gray-600 focus:border-primary duration-200 py-3"
        >
          <template #label>
            <div class="flex items-center mb-2">
              <Icon
                name="carbon:user-role"
                class="text-xl text-primary dark:text-white mr-2 mb-1"
              />
              <h1
                class="block text-md font-medium leading-6 text-gray-900 mb-1 duration-200"
              >
                Role
              </h1>
            </div>
          </template>
        </UiListSelect>

        <div class="flex justify-end gap-6 mt-10">
          <button
            @click="isOpen = false"
            class="px-4 py-2 hover:bg-gray-200 rounded-full"
            :disabled="loadingSubmit"
          >
            Cancel
          </button>

          <button
            @click="submit"
            class="bg-primary text-white px-4 py-2 rounded-md"
            :disabled="loadingSubmit"
          >
            {{ loadingSubmit ? "Sending..." : "Send Invite" }}
          </button>
        </div>
      </div>
    </template>
  </ModalsModal>
</template>
