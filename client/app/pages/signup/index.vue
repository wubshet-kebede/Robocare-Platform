<template>
  <div class="min-h-screen flex bg-slate-50">
    <div class="hidden lg:block w-1/2 relative overflow-hidden bg-slate-950">
      <video
        autoplay
        loop
        muted
        playsinline
        class="absolute inset-0 w-full h-full object-cover opacity-60"
      >
        <!-- <source src="/assets/robot-animation.mp4" type="video/mp4" /> -->
      </video>
      <img
        src="/assets/images/image_copy.png"
        alt="Hospital Image"
        class="absolute inset-0 w-full h-full object-cover opacity-60"
      />
      <div
        class="absolute inset-0 bg-gradient-to-tr from-teal-900/80 via-transparent to-teal-800/20"
      ></div>
      <div
        class="relative z-10 h-full flex flex-col justify-end p-20 text-white"
      >
        <h1
          class="text-7xl font-extrabold tracking-tighter leading-none italic"
        >
          ROBO<span class="text-teal-400">CARE.</span>
        </h1>
        <p
          class="mt-8 text-2xl text-slate-100 font-light max-w-lg leading-relaxed"
        >
          "Autonomous logistics platform for the
          <span class="text-white font-medium">Ethiopian Hospital</span>, built
          to optimize staff workflows."
        </p>

        <div
          class="mt-12 flex gap-8 opacity-50 text-[10px] font-mono uppercase tracking-widest"
        >
          <div>Gondar, ET</div>
          <div>Lat: 12.6000° N</div>
          <div>Long: 37.4667° E</div>
        </div>
      </div>
    </div>

    <div
      class="w-full lg:w-1/2 flex items-center justify-center p-12 lg:p-24 bg-white overflow-y-auto"
    >
      <div class="w-full max-w-lg">
        <NuxtLink
          to="/"
          class="group inline-flex items-center gap-2 mb-8 px-4 py-2 rounded-xl text-sm font-semibold text-slate-600 hover:text-teal-700 hover:bg-teal-50 transition-all"
        >
          <div
            class="p-1.5 rounded-lg bg-slate-100 group-hover:bg-teal-100 transition-colors"
          >
            <Icon
              name="bitcoin-icons:arrow-left-filled"
              class="w-5 h-5 text-slate-500 group-hover:text-teal-700"
            />
          </div>
          Back to Home
        </NuxtLink>

        <div class="mb-10">
          <h2 class="text-4xl font-extrabold text-slate-950 tracking-tight">
            Sign Up
          </h2>
          <p class="mt-4 text-base text-slate-600 leading-relaxed">
            Join the
            <span class="font-medium text-teal-700">Robocare Network</span> and
            start managing hospital operations.
          </p>
        </div>

        <form @submit.prevent="submit" class="w-full space-y-6">
          <div class="flex flex-row items-center gap-4">
            <UiBaseInput
              rules="required"
              name="name"
              placeholder="Full Name"
              leading-icon="mdi:account-outline"
              leadingIconClass="border-r-[1px] border-gray-300 "
              iconLeadingClass="pl-14"
              class="text-gray-600 focus:border-primary duration-200 py-2"
            >
              <template #label>
                <h1
                  class="block text-sm font-medium leading-6 text-gray-900 mb-1 dark:text-gray-300 duration-200"
                >
                  Full Name <span class="text-red-600">*</span>
                </h1>
              </template>
            </UiBaseInput>

            <UiBaseInput
              rules="required"
              name="name"
              placeholder="Username"
              leading-icon="mdi:account-outline"
              leadingIconClass="border-r-[1px] border-gray-300 "
              iconLeadingClass="pl-14"
              class="text-gray-600 focus:border-primary duration-200 py-2"
            >
              <template #label>
                <h1
                  class="block text-sm font-medium leading-6 text-gray-900 mb-1 dark:text-gray-300 duration-200"
                >
                  Username <span class="text-red-600">*</span>
                </h1>
              </template>
            </UiBaseInput>
          </div>
          <div class="flex flex-row items-center gap-4">
            <UiBaseInput
              rules="required"
              name="hospital"
              placeholder="Hospital Name"
              leading-icon="mdi:hospital-building"
              leadingIconClass="border-r-[1px] border-gray-300 "
              iconLeadingClass="pl-14"
              class="text-gray-600 focus:border-primary duration-200 py-2"
            >
              <template #label>
                <h1
                  class="block text-sm font-medium leading-6 text-gray-900 mb-1 dark:text-gray-300 duration-200"
                >
                  Hospital Name <span class="text-red-600">*</span>
                </h1>
              </template>
            </UiBaseInput>
            <UiBaseInput
              rules="required"
              name="phone"
              placeholder="+2519..."
              leading-icon="mdi:phone-outline"
              leadingIconClass="border-r-[1px] border-gray-300 "
              iconLeadingClass="pl-14"
              class="text-gray-600 focus:border-primary duration-200 py-2"
            >
              <template #label>
                <h1
                  class="block text-sm font-medium leading-6 text-gray-900 mb-1 dark:text-gray-300 duration-200"
                >
                  Phone Number <span class="text-red-600">*</span>
                </h1>
              </template>
            </UiBaseInput>
          </div>
          <div class="flex flex-row gap-4">
            <UiBaseInput
              rules="required"
              name="tin"
              placeholder="123456678"
              leading-icon="mdi:card-account-details-outline"
              leadingIconClass="border-r-[1px] border-gray-300 "
              iconLeadingClass="pl-14"
              class="text-gray-600 focus:border-primary duration-200 py-2"
            >
              <template #label>
                <h1
                  class="block text-sm font-medium leading-6 text-gray-900 mb-1 dark:text-gray-300 duration-200"
                >
                  TIN Number <span class="text-red-600">*</span>
                </h1>
              </template>
            </UiBaseInput>
            <UiBaseInput
              rules="required"
              name="adress"
              placeholder="Addis Ababa, ET"
              leading-icon="mdi:map-marker-outline"
              leadingIconClass="border-r-[1px] border-gray-300 "
              iconLeadingClass="pl-14"
              class="text-gray-600 focus:border-primary duration-200 py-2"
            >
              <template #label>
                <h1
                  class="block text-sm font-medium leading-6 text-gray-900 mb-1 dark:text-gray-300 duration-200"
                >
                  Address <span class="text-red-600">*</span>
                </h1>
              </template>
            </UiBaseInput>
          </div>
          <div>
            <UiBaseInput
              rules="required|email"
              name="email"
              placeholder="Email"
              leading-icon="mdi:email-outline"
              leadingIconClass="border-r-[1px] border-gray-300 "
              iconLeadingClass="pl-14"
              class="text-gray-600 focus:border-primary duration-200 py-2"
            >
              <template #label>
                <h1
                  class="block text-sm font-medium leading-6 text-gray-900 mb-1 dark:text-gray-300 duration-200"
                >
                  Email address <span class="text-red-600">*</span>
                </h1>
              </template>
            </UiBaseInput>
          </div>

          <div class="flex flex-row gap-4">
            <UiBaseInput
              rules="required"
              leading-icon="majesticons:key-line"
              trailingIcon="mdi:eye-outline"
              type="password"
              name="password"
              placeholder="Password"
              leadingIconClass="border-r-[1px] border-gray-300"
              iconLeadingClass="pl-14"
              class="text-gray-600 focus:border-primary duration-200 py-2"
            >
              <template #label>
                <h1
                  class="block text-sm font-medium leading-6 text-gray-900 mb-1 dark:text-gray-300 duration-200"
                >
                  Password <span class="text-red-600">*</span>
                </h1>
              </template>
            </UiBaseInput>
            <UiBaseInput
              rules="required|password"
              leading-icon="majesticons:key-line"
              trailingIcon="mdi:eye-outline"
              type="password"
              name="confirm_password"
              placeholder="Confirm Password"
              leadingIconClass="border-r-[1px] border-gray-300"
              iconLeadingClass="pl-14"
              class="text-gray-600 focus:border-primary duration-200 py-2"
            >
              <template #label>
                <h1
                  class="block text-sm font-medium leading-6 text-gray-900 mb-1 dark:text-gray-300 duration-200"
                >
                  Confirm Password <span class="text-red-600">*</span>
                </h1>
              </template>
            </UiBaseInput>
          </div>

          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <input
                id="remember-me"
                name="remember-me"
                type="checkbox"
                class="h-4 w-4 rounded border-gray-300 text-primary focus:ring-primary duration-200 cursor-pointer"
              />
              <label
                for="remember-me"
                class="ml-3 block text-sm leading-6 text-gray-900 dark:text-gray-300 cursor-pointer duration-200"
                >Remember me</label
              >
            </div>

            <div class="text-sm leading-6">
              <a
                href="#"
                class="font-semibold text-primary hover:text-primaryDark duration-200"
                >Forgot password?</a
              >
            </div>
          </div>

          <div>
            <button
              :disabled="loading"
              type="submit"
              class="flex gap-x-2 w-full justify-center disabled:bg-gray-400 disabled:cursor-not-allowed rounded-md bg-primary px-3 py-2 font-semibold leading-6 text-white shadow-sm hover:bg-primaryDark focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-gray-200 duration-200 hover:shadow-full"
            >
              <Icon
                name="svg-spinners:ring-resize"
                class="text-2xl text-gray-300"
                v-if="loading"
              />
              Sign up
            </button>
          </div>
        </form>

        <p class="mt-10 text-center text-sm text-slate-600">
          Already have an Account?
          <NuxtLink
            to="/login"
            class="font-semibold text-teal-700 hover:text-teal-600 transition underline underline-offset-4"
            >Sign In</NuxtLink
          >
        </p>
      </div>
    </div>
  </div>
</template>
<script setup>
import UiBaseInput from "~/components/ui/BaseInput.vue";
const form = reactive({ email: "", password: "" });
const loading = ref(false);
const passwordVisible = ref(false);
definePageMeta({
  layout: false,
});

const handleSignup = () => {
  console.log(form);
};
</script>
