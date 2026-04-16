export default defineNuxtPlugin(async () => {
  const auth = useAuthService();
  const user = useAuthUser();

  try {
    const me = await auth.me();
    user.value = me;
  } catch (err) {
    user.value = null;
  }
});
