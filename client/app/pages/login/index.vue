<script setup>
import UiBaseInput from "~/components/ui/BaseInput.vue";
import { useForm } from "vee-validate";
definePageMeta({
  layout: false,
});
const { handleSubmit } = useForm();
const user = useAuthUser();
const { login, me } = useAuthService();
const loading = ref(false);

const submit = handleSubmit(async (values) => {
  console.log("Form values:", values);
  try {
    loading.value = true;

    const res = await login({
      identifier: values.email,
      password: values.password,
    });
    console.log("Login response:", res);
    const userData = await me();
    user.value = userData;
    console.log("Logged user:", userData);
    await navigateTo(`/${userData.hospital.slug}/dashboard`);
  } catch (err) {
    console.log(err);
  } finally {
    loading.value = false;
  }
});
</script>
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
      class="w-full lg:w-1/2 flex items-center justify-center p-12 lg:p-24 bg-white"
    >
      <div class="w-full max-w-lg">
        <NuxtLink
          to="/"
          class="group inline-flex items-center gap-2 mb-10 px-4 py-2 rounded-xl text-sm font-semibold text-slate-600 hover:text-teal-700 hover:bg-teal-50 transition-all duration-300 ease-in-out border border-transparent hover:border-teal-100"
        >
          <div
            class="p-1.5 rounded-lg bg-slate-100 group-hover:bg-teal-100 transition-colors"
          >
            <Icon
              name="bitcoin-icons:arrow-left-filled"
              class="w-5 h-5 text-slate-500 group-hover:text-teal-700 transition-colors font-bold"
            />
          </div>
          Back to Home
        </NuxtLink>
        <div class="mb-16">
          <h2 class="text-4xl font-extrabold text-slate-950 tracking-tight">
            Welcome Back!
          </h2>
          <p class="mt-4 text-base text-slate-600 leading-relaxed max-w-md">
            Sign in to access your dashboard and continue
            <span class="font-medium text-teal-700"
              >managing your hospital fleet.</span
            >
          </p>
        </div>

        <form @submit.prevent="submit" class="w-full space-y-6">
          <div>
            <UiBaseInput
              rules="required|email"
              name="email"
              placeholder="Enter your email"
              leading-icon="mdi:email-outline"
              leadingIconClass="border-r-[1px] border-gray-300 "
              iconLeadingClass="pl-14"
              class="text-gray-600 focus:border-primary duration-200 py-3"
            >
              <template #label>
                <h1
                  class="block text-sm font-medium leading-6 text-gray-900 mb-1 duration-200"
                >
                  Email address <span class="text-red-600">*</span>
                </h1>
              </template>
            </UiBaseInput>
          </div>

          <div>
            <UiBaseInput
              rules="required"
              leading-icon="majesticons:key-line"
              trailingIcon="mdi:eye-outline"
              type="password"
              name="password"
              placeholder="Enter your password"
              leadingIconClass="border-r-[1px] border-gray-300"
              iconLeadingClass="pl-14"
              class="text-gray-600 focus:border-primary duration-200 py-3"
            >
              <template #label>
                <h1
                  class="block text-sm font-medium leading-6 text-gray-900 mb-1 duration-200"
                >
                  Password <span class="text-red-600">*</span>
                </h1>
              </template>
            </UiBaseInput>
          </div>

          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <input
                id="remember-doctorwubshet@example.comme"
                name="remember-me"
                type="checkbox"
                class="h-4 w-4 rounded border-gray-300 text-primary focus:ring-primary duration-200 cursor-pointer"
              />
              <label
                for="remember-me"
                class="ml-3 block text-sm leading-6 text-gray-900 cursor-pointer duration-200"
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
              class="flex gap-x-2 w-full justify-center disabled:bg-gray-400 disabled:cursor-not-allowed rounded-full bg-primary px-3 py-2 font-semibold leading-6 text-white shadow-sm hover:bg-primaryDark focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-gray-200 duration-200 hover:shadow-full"
            >
              <Icon
                name="svg-spinners:ring-resize"
                class="text-2xl text-gray-300"
                v-if="loading"
              />
              Log in
            </button>
          </div>
        </form>

        <div class="relative my-12">
          <div class="absolute inset-0 flex items-center">
            <div class="w-full border-t border-slate-100"></div>
          </div>
          <div class="relative flex justify-center text-sm">
            <span class="px-3 bg-white text-slate-400 font-medium">OR</span>
          </div>
        </div>

        <div class="space-y-4">
          <button
            class="w-full flex items-center justify-center gap-3 py-3 px-4 border border-slate-200 rounded-xl text-md font-medium text-slate-800 hover:bg-slate-50 transition"
          >
            <Icon name="devicon:google" class="h-5 w-5" />
            Continue with Google
          </button>
          <button
            class="w-full flex items-center justify-center gap-3 py-3 px-4 border border-slate-200 rounded-xl text-md font-medium text-slate-800 hover:bg-slate-50 transition"
          >
            <Icon name="ic:outline-apple" class="h-5 w-5" />
            with Apple
          </button>
        </div>

        <div class="flex justify-center text-base mt-10">
          <p
            class="text-gray-600 text-center flex flex-wrap items-center gap-2"
          >
            Don't have an account?
            <a
              href="/signup"
              class="text-primary underline underline-offset-4 flex items-center gap-1"
            >
              Sign Up
              <span>
                <Icon name="i-mdi:arrow-right" class="-rotate-45 ml-1" /> </span
            ></a>
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<!-- <script setup>
definePageMeta({
  layout: false,
});
const form = reactive({
  email: "",
  password: "",
});
const loading = ref(false);

const handleLogin = async () => {
  loading.value = true;
  // Call your Go Backend here!
  console.log("Logging in with:", form.email);
  setTimeout(() => {
    loading.value = false;
  }, 1000);
};
</script> -->
