<script setup>
import { format } from "date-fns";
import { useForm } from "vee-validate";
import { ref, watch, computed } from "vue";

const { notify } = Notify();

const props = defineProps({
  modelValue: Boolean,
  currentUser: Object,
  editMode: Boolean,
});

const emit = defineEmits([
  "update:modelValue",

  "refetch",
  "update:currentUser",
]);

const isOpen = computed({
  get() {
    return props.modelValue;
  },
  set(value) {
    emit("update:modelValue", value);
  },
});

const title = computed(() => (props.editMode ? "Edit User" : "Create User"));

const { handleSubmit, resetForm, setValues } = useForm();

const form = ref(props.currentUser || {});
const profileId = ref("");
const isEditing = ref(false);
</script>

<template>
  <ModalsModal v-model="isOpen" :title="title" :wrapperClass="'max-w-[900px]'">
    <template #Heading>
      <div>
        <h4 class="text-primary flex gap-x-3 justify-start items-center">
          <Icon v-if="!currentUser" name="gridicons:create"></Icon>
          <Icon v-if="currentUser" name="iconamoon:edit"></Icon>
          <span> {{ title }}</span>
        </h4>
      </div>
    </template>
    <template #content>
      <form @submit.prevent="saveUser" class="space-y-4 p-6">
        <!-- Increased padding -->

        <div class="grid grid-cols-1 md:grid-cols-4 mb-4 gap-x-10">
          <div class="col-span-2 grid grid-cols-1 gap-4">
            <SelectorsAuthorizerUser
              v-if="!editMode"
              rules="required"
              v-model="form.id"
            />
            <HTextfield
              name="first_name"
              rules="required"
              v-model="form.first_name"
            >
              <template #label>
                <h1
                  class="font-medium mb-2 text-gray-600 dark:text-secondary-lite"
                >
                  First Name
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </HTextfield>

            <HTextfield name="middle_name" v-model="form.middle_name">
              <template #label>
                <h1
                  class="font-medium mb-2 text-gray-600 dark:text-secondary-lite"
                >
                  Father's Name
                </h1>
              </template>
            </HTextfield>

            <HTextfield name="last_name" v-model="form.last_name">
              <template #label>
                <h1
                  class="font-medium mb-2 text-gray-600 dark:text-secondary-lite"
                >
                  Grand Father's Name
                </h1>
              </template>
            </HTextfield>
          </div>
          <!-- Other Fields (Two per Row) -->
          <div class="col-span-2 flex flex-col gap-4">
            <HTextfield
              name="email"
              rules="required|email"
              v-model="form.email"
            >
              <template #label>
                <h1
                  class="font-medium mb-2 text-gray-600 dark:text-secondary-lite"
                >
                  Email
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </HTextfield>

            <HTextfield
              name="phone_number"
              rules="required|ethiopian_phone_number"
              v-model="form.phone_number"
            >
              <template #label>
                <h1
                  class="font-medium mb-2 text-gray-600 dark:text-secondary-lite"
                >
                  Phone
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </HTextfield>

            <HListselect
              rules="required"
              :items="[
                { name: 'Female', id: 'female' },
                { name: 'Male', id: 'male' },
              ]"
              name="gender"
              v-model="form.gender"
            >
              <template #label>
                <h1 class="mb-2 dark:text-secondary-lite">Gender</h1>
              </template>
            </HListselect>
            <SelectorsRoleUser
              v-if="!editMode"
              v-model="form.role"
              rules="required"
            >
              <template #label>
                <h1 class="mb-2 dark:text-secondary-lite">Role</h1>
              </template>
            </SelectorsRoleUser>
          </div>
          <div class="col-span-2 flex justify-end mt-3">
            <SelectorsRoleUser
              v-if="editMode"
              v-model="form.role"
              rules="required"
            >
              <template #label>
                <h1 class="mb-2 dark:text-secondary-lite">Role</h1>
              </template>
            </SelectorsRoleUser>
          </div>
          <div class="flex col-span-4 justify-end gap-5 mt-9">
            <button
              type="button"
              @click="isOpen = false"
              class="px-4 py-2 text-sm font-medium text-gray-700 bg-gray-100 rounded-md hover:bg-gray-200 dark:bg-gray-500 dark:text-white"
            >
              Cancel
            </button>
            <button
              type="submit"
              class="px-4 py-2 text-sm font-medium text-white bg-primary rounded-md hover:bg-primary-dark"
            >
              {{ currentUser ? "Update" : "Create" }}
            </button>
          </div>
        </div>
      </form>
    </template>
  </ModalsModal>
</template>
