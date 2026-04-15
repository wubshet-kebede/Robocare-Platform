export const useAuthService = () => {
  const { $api } = useNuxtApp();

  const login = (data) => {
    return $api("/login", {
      method: "POST",
      body: data,
    });
  };

  const signup = (data) => {
    return $api("/signup", {
      method: "POST",
      body: data,
    });
  };

  return { login, signup };
};
