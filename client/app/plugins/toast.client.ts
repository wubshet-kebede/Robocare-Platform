import * as vt from "vue-toastification";

export default defineNuxtPlugin((nuxtApp) => {
  const modulePayload: any = vt.default || vt;

  nuxtApp.vueApp.use(modulePayload);

  const useToastHook = vt.useToast || modulePayload.useToast;

  return {
    provide: {
      toast:
        typeof useToastHook === "function" ? useToastHook() : modulePayload,
    },
  };
});
