// plugins/toast.client.ts
export default defineNuxtPlugin(async (nuxtApp) => {
  const vue3Toastify = await import("vue3-toastify");
  await import("vue3-toastify/dist/index.css");

  nuxtApp.vueApp.use(vue3Toastify.default, {
    autoClose: 3000,
  });

  return {
    provide: {
      toast: vue3Toastify.toast,
    },
  };
});
